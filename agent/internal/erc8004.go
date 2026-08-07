// Package internal — ERC-8004 Trustless Agents identity integration.
//
// Provides on-chain identity (IdentityRegistry) and reputation (ReputationRegistry)
// for Bonded Agent operators, following the ERC-8004 standard.
package internal

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gitgdut/bonded-agent/agent/contracts/erc8004"
)

// ERC8004Service manages on-chain identity and reputation via ERC-8004 registries.
type ERC8004Service struct {
	client *ethclient.Client

	identity   *erc8004.ERC8004IdentityRegistry
	reputation *erc8004.ERC8004ReputationRegistry
}

// NewERC8004Service creates the service. Contract addresses must be configured.
func NewERC8004Service(
	client *ethclient.Client,
	identityAddr, reputationAddr common.Address,
) (*ERC8004Service, error) {
	if identityAddr == (common.Address{}) || reputationAddr == (common.Address{}) {
		return nil, fmt.Errorf("ERC-8004 contract addresses not configured")
	}

	identity, err := erc8004.NewERC8004IdentityRegistry(identityAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind IdentityRegistry: %w", err)
	}

	reputation, err := erc8004.NewERC8004ReputationRegistry(reputationAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind ReputationRegistry: %w", err)
	}

	return &ERC8004Service{
		client:     client,
		identity:   identity,
		reputation: reputation,
	}, nil
}

// ── Identity Registry ────────────────────────────────────────

// RegisterAgent mints an ERC-8004 identity NFT for the operator.
// Idempotent — skips if the operator already owns an identity.
// Returns the agentId (ERC-721 token ID).
func (s *ERC8004Service) RegisterAgent(auth *bind.TransactOpts, agentURI string) (*big.Int, error) {
	nextID, err := s.identity.NextAgentId(nil)
	if err != nil {
		return nil, fmt.Errorf("query nextAgentId: %w", err)
	}

	operatorAddr := auth.From

	// Scan existing tokens — if operator already owns one, return it
	for id := int64(1); id < nextID.Int64(); id++ {
		owner, err := s.identity.OwnerOf(nil, big.NewInt(id))
		if err != nil {
			continue
		}
		if owner == operatorAddr {
			log.Printf("ERC-8004: operator %s already registered as agent #%d", operatorAddr.Hex(), id)
			return big.NewInt(id), nil
		}
	}

	// Mint new identity (Register0 = single-string overload)
	tx, err := s.identity.Register0(auth, agentURI)
	if err != nil {
		return nil, fmt.Errorf("register agent: %w", err)
	}

	agentID := new(big.Int).Set(nextID)

	log.Printf("ERC-8004: registered operator %s as agent #%d, tx: %s", operatorAddr.Hex(), agentID, tx.Hash().Hex())
	return agentID, nil
}

// GetIdentityURI returns the agent card URI for a given agent.
func (s *ERC8004Service) GetIdentityURI(agentID *big.Int) (string, error) {
	return s.identity.TokenURI(nil, agentID)
}

// GetAgentWallet returns the linked wallet address for an agent.
func (s *ERC8004Service) GetAgentWallet(agentID *big.Int) (common.Address, error) {
	return s.identity.GetAgentWallet(nil, agentID)
}

// ── Reputation Registry ──────────────────────────────────────

// GiveFeedback posts reputation feedback for an operator after a plan execution.
// value should be a signed fixed-point number: e.g., 9777 with 2 decimals = 97.77%.
func (s *ERC8004Service) GiveFeedback(
	auth *bind.TransactOpts,
	agentID *big.Int,
	value *big.Int,
	valueDecimals uint8,
	tag1, tag2, feedbackURI string,
) error {
	tx, err := s.reputation.GiveFeedback(
		auth, agentID, value, valueDecimals,
		tag1, tag2, "", feedbackURI, [32]byte{},
	)
	if err != nil {
		return fmt.Errorf("giveFeedback: %w", err)
	}

	log.Printf("ERC-8004: posted feedback for agent #%d, tx: %s", agentID, tx.Hash().Hex())
	return nil
}

// GetAgentSummary returns aggregated reputation data for an agent.
func (s *ERC8004Service) GetAgentSummary(agentID *big.Int) (*AgentSummary, error) {
	result, err := s.reputation.GetSummary(nil, agentID, nil, "", "")
	if err != nil {
		return nil, fmt.Errorf("getSummary: %w", err)
	}

	return &AgentSummary{
		AgentID:          agentID.String(),
		FeedbackCount:    result.Count,
		SummaryValue:     result.SummaryValue,
		ValueDecimals:    result.SummaryValueDecimals,
	}, nil
}

// ── Composite types ──────────────────────────────────────────

// AgentSummary mirrors the on-chain reputation summary.
type AgentSummary struct {
	AgentID       string   `json:"agentId"`
	FeedbackCount uint64   `json:"feedbackCount"`
	SummaryValue  *big.Int `json:"summaryValue"`
	ValueDecimals uint8    `json:"valueDecimals"`
}

// AgentIdentity combines on-chain ERC-8004 identity with a Bonded Agent operator.
type AgentIdentity struct {
	AgentID         string `json:"agentId"`
	AgentURI        string `json:"agentUri"`
	Wallet          string `json:"wallet,omitempty"`
	OperatorAddress string `json:"operatorAddress"`
	FeedbackCount   uint64 `json:"feedbackCount"`
	ScoreValue      string `json:"scoreValue"` // human-readable from summaryValue+decimals
}

// GetAgentIdentity returns the full ERC-8004 identity profile for a given agent.
func (s *ERC8004Service) GetAgentIdentity(agentID *big.Int, operatorAddr common.Address) (*AgentIdentity, error) {
	uri, _ := s.GetIdentityURI(agentID)

	wallet, _ := s.GetAgentWallet(agentID)

	summary, err := s.GetAgentSummary(agentID)
	if err != nil {
		summary = &AgentSummary{AgentID: agentID.String()}
	}

	identity := &AgentIdentity{
		AgentID:         agentID.String(),
		AgentURI:        uri,
		OperatorAddress: operatorAddr.Hex(),
		FeedbackCount:   summary.FeedbackCount,
	}

	if wallet != (common.Address{}) {
		identity.Wallet = wallet.Hex()
	}

	if summary.FeedbackCount > 0 && summary.SummaryValue != nil {
		identity.ScoreValue = formatFixedPoint(summary.SummaryValue, summary.ValueDecimals)
	}

	return identity, nil
}

// ── Helpers ─────────────────────────────────────────────────

// BuildAgentURI creates a minimal ERC-8004 agent card JSON string.
// In production this should be an IPFS URI — here we use inline JSON for dev.
func BuildAgentURI(name, description, endpoint string, feeBps int64, guaranteeRatio float64) string {
	card := fmt.Sprintf(`{
  "name": "%s",
  "description": "%s",
  "version": "1.0.0",
  "endpoints": {
    "a2a": "%s"
  },
  "capabilities": ["swap", "guaranteed-execution"],
  "trust": {
    "feeBps": %d,
    "guaranteeRatio": %.2f
  }
}`, name, description, endpoint, feeBps, guaranteeRatio)
	return card
}

// formatFixedPoint converts (*big.Int value, decimals) to "12.34".
func formatFixedPoint(value *big.Int, decimals uint8) string {
	if value == nil {
		return "0"
	}
	s := value.String()
	if decimals == 0 || s == "0" {
		return s
	}

	neg := false
	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}

	dec := int(decimals)
	if len(s) <= dec {
		s = strings.Repeat("0", dec-len(s)+1) + s
	}

	intPart := s[:len(s)-dec]
	fracPart := strings.TrimRight(s[len(s)-dec:], "0")
	if fracPart == "" {
		fracPart = "0"
	}

	result := intPart + "." + fracPart
	if neg {
		result = "-" + result
	}
	return result
}
