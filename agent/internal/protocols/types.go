// Package protocols defines the Moss-style discover→load→action→simulate pipeline
// for DeFi protocol integration. Each protocol is a self-describing adapter that
// exposes Capabilities (unsigned transactions) and Queries (read-only calls).
package protocols

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ── Core types (matching Moss TypeScript types) ─────────────────

// Verb classifies the action a Capability performs.
type Verb string

const (
	VerbSwap     Verb = "swap"
	VerbWrap     Verb = "wrap"
	VerbUnwrap   Verb = "unwrap"
	VerbTransfer Verb = "transfer"
	VerbApprove  Verb = "approve"
	VerbSupply   Verb = "supply"
	VerbWithdraw Verb = "withdraw"
	VerbBorrow   Verb = "borrow"
	VerbRepay    Verb = "repay"
	VerbStake    Verb = "stake"
	VerbUnstake  Verb = "unstake"
	VerbClaim    Verb = "claim"
	VerbMint     Verb = "mint"
	VerbOpen     Verb = "open"
	VerbClose    Verb = "close"
)

// Category classifies a protocol's domain.
type Category string

const (
	CategoryDex      Category = "dex"
	CategoryLending  Category = "lending"
	CategoryStaking  Category = "staking"
	CategoryRewards  Category = "rewards"
	CategoryToken    Category = "token"
	CategoryNFT      Category = "nft"
)

// RiskLabel marks the risk dimensions of a Capability.
type RiskLabel string

const (
	RiskFundOut     RiskLabel = "fundOut"
	RiskApproval    RiskLabel = "approval"
	RiskPriceImpact RiskLabel = "priceImpact"
)

// Coordinate is a lightweight capability/query descriptor returned by discover.
type Coordinate struct {
	Protocol string   `json:"protocol"`
	Method   string   `json:"method"`
	Kind     string   `json:"kind"` // "capability" | "query"
	Verb     Verb     `json:"verb,omitempty"`
	Category Category `json:"category"`
	Tags     []string `json:"tags"`
	Summary  string   `json:"summary"`
}

// ParamMeta describes one parameter's type and purpose.
type ParamMeta struct {
	Type        string `json:"type"`        // e.g. "uint256", "address", "decimal"
	Description string `json:"description"` // method-specific purpose
}

// Stub is the full calling contract for a method (returned by load).
type Stub struct {
	Protocol string              `json:"protocol"`
	Method   string              `json:"method"`
	Kind     string              `json:"kind"`
	Intent   string              `json:"intent"`
	Verb     Verb                `json:"verb,omitempty"`
	Category Category            `json:"category"`
	Risk     []RiskLabel         `json:"risk"`
	Tags     []string            `json:"tags"`
	Params   map[string]ParamMeta `json:"params"`
}

// ── Action / Simulation types ──────────────────────────────────

// TransactionNode is an unsigned transaction built by a Capability.
type TransactionNode struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Data  []byte         `json:"data"`
	Value *big.Int       `json:"value"`
}

// MarshalJSON customizes JSON output for readability.
func (tx TransactionNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Data  string `json:"data"`
		Value string `json:"value"`
	}{
		From:  tx.From.Hex(),
		To:    tx.To.Hex(),
		Data:  "0x" + common.Bytes2Hex(tx.Data),
		Value: tx.Value.String(),
	})
}

// CapabilityNode is a tree of unsigned transactions.
type CapabilityNode struct {
	Kind     string           `json:"kind"` // "capability"
	Protocol string           `json:"protocol"`
	Method   string           `json:"method"`
	Params   map[string]interface{} `json:"params"`
	Tx       *TransactionNode `json:"tx,omitempty"` // the direct transaction
}

// ActionNode is the result of registry.Action() — either a Query or Capability.
type ActionNode struct {
	Kind     string           `json:"kind"` // "query" | "capability"
	Protocol string           `json:"protocol"`
	Method   string           `json:"method"`
	Data     interface{}      `json:"data,omitempty"`     // query result
	Node     *CapabilityNode  `json:"node,omitempty"`     // capability result
}

// ── Simulation types ───────────────────────────────────────────

// SimulationWarning describes a problem found during simulation.
type SimulationWarning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SimulationResult holds the outcome of simulating a single transaction.
type SimulationResult struct {
	Success      bool                 `json:"success"`
	ReturnData   string               `json:"returnData,omitempty"`
	GasUsed      uint64               `json:"gasUsed,omitempty"`
	Warnings     []SimulationWarning  `json:"warnings,omitempty"`
}

// SimulateOutcome is the top-level simulation result.
type SimulateOutcome struct {
	Halted  bool                `json:"halted"`
	Results []SimulationResult  `json:"results"`
}

// ── Protocol interface ─────────────────────────────────────────

// Protocol is the interface each DEX/lending/staking adapter must implement.
// It follows Moss's @Protocol decorator pattern: self-describing, with
// discoverable Capabilities and Queries.
type Protocol interface {
	// Name returns the kebab-case protocol slug (e.g. "pancakeswap-v2").
	Name() string
	// Category returns the protocol category (e.g. "dex").
	Category() Category
	// Description returns human-readable protocol description.
	Description() string
	// Contracts returns name→address map of known contracts (router, factory, etc).
	Contracts() map[string]common.Address
	// Discover returns all available capabilities and queries.
	Discover() []Coordinate
	// Load returns the full calling contract for one method.
	Load(method string) (*Stub, error)
	// Action executes a query or builds a capability (unsigned tx).
	Action(method string, account common.Address, params map[string]interface{}) (*ActionNode, error)
}

// ── Registry ───────────────────────────────────────────────────

// DiscoverFilter restricts discover results.
type DiscoverFilter struct {
	Verb     Verb
	Category Category
	Protocol string
}
