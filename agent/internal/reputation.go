// Package internal — Operator reputation engine.
//
// Scans BondedExecutor on-chain events to build per-operator statistics and
// reputation scores. All data is off-chain derived — the BondedExecutor
// contract needs no changes (Plan.operator already supports multi-operator).
package internal

import (
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gitgdut/bonded-agent/agent/contracts"
)

// ── Operator data types ───────────────────────────────────────

// OperatorProfile holds off-chain metadata (loaded from ops.json).
type OperatorProfile struct {
	Address          string  `json:"address"`
	Name             string  `json:"name"`
	FeeBps           int64   `json:"feeBps"`
	Ratio            float64 `json:"ratio"`
	ERC8004AgentID   string  `json:"erc8004AgentId,omitempty"` // ERC-8004 identity NFT token ID
}

// OperatorStats holds on-chain derived statistics and reputation score.
type OperatorStats struct {
	Address         string  `json:"address"`
	Name            string  `json:"name,omitempty"`
	TotalPlans      int64   `json:"totalPlans"`
	SuccessPlans    int64   `json:"successPlans"`
	ShortfallPlans  int64   `json:"shortfallPlans"`
	FailedPlans     int64   `json:"failedPlans"`
	TotalVolume     string  `json:"totalVolume"` // wei string
	ReputationScore float64 `json:"reputationScore"`
	SuccessRate     float64 `json:"successRate"`
	ServiceFeeBps   int64   `json:"serviceFeeBps"`
	GuaranteedRatio float64 `json:"guaranteedRatio"`
	IsDefault       bool    `json:"isDefault"` // true when this operator is the current server
	ERC8004AgentID  string  `json:"erc8004AgentId,omitempty"` // ERC-8004 identity NFT
}

// ── Reputation engine ─────────────────────────────────────────

// ReputationEngine indexes on-chain plan events per operator.
type ReputationEngine struct {
	mu       sync.RWMutex
	stats    map[common.Address]*OperatorStats
	profiles map[common.Address]OperatorProfile

	// planOperator maps planId → operator address (built from PlanOpened events)
	planOperator map[[32]byte]common.Address

	executor *contracts.BondedExecutor
}

// NewReputationEngine creates the engine, loads profiles, and scans history.
func NewReputationEngine(executor *contracts.BondedExecutor, profiles []OperatorProfile) (*ReputationEngine, error) {
	e := &ReputationEngine{
		stats:        make(map[common.Address]*OperatorStats),
		profiles:     make(map[common.Address]OperatorProfile),
		planOperator: make(map[[32]byte]common.Address),
		executor:     executor,
	}

	for _, p := range profiles {
		addr := common.HexToAddress(p.Address)
		e.profiles[addr] = p
		// Seed stats entry for every profile — operators show up even without on-chain history
		e.EnsureStats(addr)
	}

	// Best-effort historical scan — reputation is advisory
	_ = e.scanHistory()

	return e, nil
}

// ── Historical scan ───────────────────────────────────────────

// scanHistory performs a two-pass scan:
//  1. PlanOpened → build planId→operator map + count totals
//  2. PlanExecuted / PlanFailed → attribute outcomes to operators
func (e *ReputationEngine) scanHistory() error {
	// Pass 1 — PlanOpened
	openIter, err := e.executor.FilterPlanOpened(nil, nil, nil, nil)
	if err != nil {
		return err
	}
	defer openIter.Close()

	for openIter.Next() {
		ev := openIter.Event
		e.planOperator[ev.PlanId] = ev.Operator
	}
	if err := openIter.Error(); err != nil {
		return err
	}

	// Pass 2a — PlanExecuted (success / shortfall)
	execIter, err := e.executor.FilterPlanExecuted(nil, nil)
	if err != nil {
		return err
	}
	defer execIter.Close()

	for execIter.Next() {
		ev := execIter.Event
		operator, ok := e.planOperator[ev.PlanId]
		if !ok {
			continue
		}
		s := e.EnsureStats(operator)
		e.mu.Lock()
		s.SuccessPlans++
		if ev.PaidToUser != nil && ev.PaidToUser.Sign() > 0 {
			s.ShortfallPlans++
		}
		e.mu.Unlock()
	}
	// Don't hard-fail on iterator error — we already have PlanOpened data

	// Pass 2b — PlanFailed
	failIter, err := e.executor.FilterPlanFailed(nil, nil)
	if err != nil {
		return err
	}
	defer failIter.Close()

	for failIter.Next() {
		ev := failIter.Event
		operator, ok := e.planOperator[ev.PlanId]
		if !ok {
			continue
		}
		s := e.EnsureStats(operator)
		e.mu.Lock()
		s.FailedPlans++
		e.mu.Unlock()
	}

	// Compute totalPlans from planOperator map
	for planID, op := range e.planOperator {
		_ = planID
		s := e.EnsureStats(op)
		e.mu.Lock()
		s.TotalPlans++
		e.mu.Unlock()
	}

	return nil
}

// EnsureStats creates a stats entry if none exists. Safe to call from outside.
func (e *ReputationEngine) EnsureStats(addr common.Address) *OperatorStats {
	e.mu.Lock()
	defer e.mu.Unlock()

	if s, ok := e.stats[addr]; ok {
		return s
	}

	s := &OperatorStats{
		Address:  addr.Hex(),
		TotalVolume: "0",
	}

	if p, ok := e.profiles[addr]; ok {
		s.Name = p.Name
		s.ServiceFeeBps = p.FeeBps
		s.GuaranteedRatio = p.Ratio
		s.ERC8004AgentID = p.ERC8004AgentID
	}

	e.stats[addr] = s
	return s
}

// ── Public API ────────────────────────────────────────────────

// GetOperators returns all known operators with derived stats and scores.
func (e *ReputationEngine) GetOperators() []OperatorStats {
	e.mu.RLock()
	defer e.mu.RUnlock()

	result := make([]OperatorStats, 0, len(e.stats))
	for addr, s := range e.stats {
		stats := *s
		stats.Address = addr.Hex()
		e.populateDerived(&stats)
		result = append(result, stats)
	}

	return result
}

// GetOperator returns stats for a specific operator, or nil if unknown.
func (e *ReputationEngine) GetOperator(addr common.Address) *OperatorStats {
	e.mu.RLock()
	defer e.mu.RUnlock()

	s, ok := e.stats[addr]
	if !ok {
		return nil
	}

	stats := *s
	stats.Address = addr.Hex()
	e.populateDerived(&stats)
	return &stats
}

// SetERC8004AgentID records the ERC-8004 identity NFT token ID for an operator.
func (e *ReputationEngine) SetERC8004AgentID(addr common.Address, agentID string) {
	s := e.EnsureStats(addr)
	e.mu.Lock()
	s.ERC8004AgentID = agentID
	e.mu.Unlock()
}

// RecordExecution is called after a plan execution to update stats in real time.
func (e *ReputationEngine) RecordExecution(operator common.Address, planID [32]byte, status string, inputAmount *big.Int) {
	e.planOperator[planID] = operator

	s := e.EnsureStats(operator)
	e.mu.Lock()
	defer e.mu.Unlock()

	s.TotalPlans++
	switch status {
	case "settled_ok":
		s.SuccessPlans++
	case "settled_shortfall":
		s.SuccessPlans++
		s.ShortfallPlans++
	case "failed":
		s.FailedPlans++
	}

	if inputAmount != nil {
		cur := new(big.Int)
		cur.SetString(s.TotalVolume, 10)
		s.TotalVolume = new(big.Int).Add(cur, inputAmount).String()
	}
}

// ── Scoring ───────────────────────────────────────────────────

func (e *ReputationEngine) populateDerived(s *OperatorStats) {
	// Load profile metadata (don't overwrite runtime-set values like ERC-8004 agent ID)
	if p, ok := e.profiles[common.HexToAddress(s.Address)]; ok {
		if s.Name == "" {
			s.Name = p.Name
		}
		if s.ServiceFeeBps == 0 {
			s.ServiceFeeBps = p.FeeBps
		}
		if s.GuaranteedRatio == 0 {
			s.GuaranteedRatio = p.Ratio
		}
		if s.ERC8004AgentID == "" {
			s.ERC8004AgentID = p.ERC8004AgentID
		}
	}

	if s.TotalPlans > 0 {
		s.SuccessRate = math.Round(float64(s.SuccessPlans)/float64(s.TotalPlans)*10000) / 100
	}
	s.ReputationScore = computeScore(s)
}

// computeScore: base(50) + successRate*0.4 + noFailureBonus*10
func computeScore(s *OperatorStats) float64 {
	if s.TotalPlans == 0 {
		return 50
	}

	successRate := float64(s.SuccessPlans) / float64(s.TotalPlans)
	failRate := float64(s.FailedPlans) / float64(s.TotalPlans)

	score := 50.0
	score += successRate * 40.0
	score += (1.0 - failRate) * 10.0

	return math.Min(100, math.Max(0, math.Round(score)))
}
