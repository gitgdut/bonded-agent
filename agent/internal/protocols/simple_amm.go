package protocols

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gitgdut/bonded-agent/agent/contracts"
)

// SimpleAMMProtocol adapts our SimpleAMMPair to the Moss Protocol interface.
// It exposes two methods:
//   - "quote" (Query):  calls getAmountOut(amountIn) → expected output
//   - "swap"  (Capability): builds calldata for swap(minOutput)
type SimpleAMMProtocol struct {
	pair    *contracts.SimpleAMMPair
	address common.Address

	// Methods metadata
	methods map[string]Coordinate
}

// NewSimpleAMMProtocol creates a protocol adapter for SimpleAMMPair.
func NewSimpleAMMProtocol(pair *contracts.SimpleAMMPair, addr common.Address) *SimpleAMMProtocol {
	return &SimpleAMMProtocol{
		pair:    pair,
		address: addr,
		methods: map[string]Coordinate{
			"quote": {
				Protocol: "simple-amm",
				Method:   "quote",
				Kind:     "query",
				Category: CategoryDex,
				Tags:     []string{"amm", "v2", "quote"},
				Summary:  "Quote expected output for MON→tUSDC swap using constant-product AMM",
			},
			"swap": {
				Protocol: "simple-amm",
				Method:   "swap",
				Kind:     "capability",
				Verb:     VerbSwap,
				Category: CategoryDex,
				Tags:     []string{"amm", "v2"},
				Summary:  "Swap MON for tUSDC on constant-product AMM, with minOutput protection",
			},
		},
	}
}

// Name returns the protocol slug.
func (p *SimpleAMMProtocol) Name() string { return "simple-amm" }

// Category returns the protocol category.
func (p *SimpleAMMProtocol) Category() Category { return CategoryDex }

// Description returns a human-readable description.
func (p *SimpleAMMProtocol) Description() string {
	return "Constant-product AMM (Uniswap V2-style) for MON→tUSDC swaps — deployed by Bonded Agent for testnet"
}

// Contracts returns protocol contract addresses.
func (p *SimpleAMMProtocol) Contracts() map[string]common.Address {
	return map[string]common.Address{
		"pair": p.address,
	}
}

// Discover returns all available methods.
func (p *SimpleAMMProtocol) Discover() []Coordinate {
	return []Coordinate{p.methods["quote"], p.methods["swap"]}
}

// Load returns the full calling contract for one method.
func (p *SimpleAMMProtocol) Load(method string) (*Stub, error) {
	switch method {
	case "quote":
		return &Stub{
			Protocol: "simple-amm",
			Method:   "quote",
			Kind:     "query",
			Intent:   "Quote expected tUSDC output for a given MON input using constant-product formula",
			Category: CategoryDex,
			Tags:     []string{"amm", "v2", "quote"},
			Params: map[string]ParamMeta{
				"amountIn": {
					Type:        "uint256",
					Description: "MON input amount in wei (1 MON = 10^18 wei)",
				},
			},
		}, nil
	case "swap":
		return &Stub{
			Protocol: "simple-amm",
			Method:   "swap",
			Kind:     "capability",
			Verb:     VerbSwap,
			Intent:   "Swap MON for tUSDC with minimum output guarantee, tolerance {slippage} bps",
			Category: CategoryDex,
			Risk:     []RiskLabel{RiskFundOut, RiskPriceImpact},
			Tags:     []string{"amm", "v2"},
			Params: map[string]ParamMeta{
				"amountIn": {
					Type:        "uint256",
					Description: "MON input amount in wei",
				},
				"minOutput": {
					Type:        "uint256",
					Description: "Minimum acceptable tUSDC output (slippage protection). 0 = no protection.",
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("simple-amm has no method %q", method)
	}
}

// Action executes a query or builds a capability.
func (p *SimpleAMMProtocol) Action(method string, account common.Address, params map[string]interface{}) (*ActionNode, error) {
	switch method {
	case "quote":
		return p.actionQuote(params)
	case "swap":
		return p.actionSwap(account, params)
	default:
		return nil, fmt.Errorf("simple-amm has no method %q", method)
	}
}

// actionQuote handles the "quote" Query.
func (p *SimpleAMMProtocol) actionQuote(params map[string]interface{}) (*ActionNode, error) {
	amountIn, ok := params["amountIn"]
	if !ok {
		return nil, fmt.Errorf("quote requires 'amountIn' parameter")
	}

	var amountInBig *big.Int
	switch v := amountIn.(type) {
	case string:
		var ok bool
		amountInBig, ok = new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("invalid amountIn: %q", v)
		}
	case *big.Int:
		amountInBig = new(big.Int).Set(v)
	default:
		return nil, fmt.Errorf("amountIn must be string or *big.Int, got %T", amountIn)
	}

	expectedOut, err := p.pair.GetAmountOut(nil, amountInBig)
	if err != nil {
		return nil, fmt.Errorf("getAmountOut failed: %w", err)
	}

	return &ActionNode{
		Kind:     "query",
		Protocol: "simple-amm",
		Method:   "quote",
		Data: map[string]interface{}{
			"amountIn":       amountInBig.String(),
			"expectedOutput": expectedOut.String(),
			"pairAddress":    p.address.Hex(),
		},
	}, nil
}

// actionSwap builds the swap Capability.
func (p *SimpleAMMProtocol) actionSwap(account common.Address, params map[string]interface{}) (*ActionNode, error) {
	minOutput := big.NewInt(0) // default: no slippage protection at build time

	if moRaw, ok := params["minOutput"]; ok {
		switch v := moRaw.(type) {
		case string:
			var ok bool
			minOutput, ok = new(big.Int).SetString(v, 10)
			if !ok {
				return nil, fmt.Errorf("invalid minOutput: %q", v)
			}
		case *big.Int:
			minOutput = new(big.Int).Set(v)
		}
	}

	// Build calldata: swap(uint256 minOutput)
	swapABI, err := contracts.SimpleAMMPairMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("get ABI: %w", err)
	}

	calldata, err := swapABI.Pack("swap", minOutput)
	if err != nil {
		return nil, fmt.Errorf("pack calldata: %w", err)
	}

	// The account passed here is the BondedExecutor address — it's the msg.sender
	// when BondedExecutor calls target.call{value}(calldata)
	tx := &TransactionNode{
		From:  account, // BondedExecutor address
		To:    p.address,
		Data:  calldata,
		Value: nil, // value is set at execution time by the BondedExecutor
	}

	return &ActionNode{
		Kind:     "capability",
		Protocol: "simple-amm",
		Method:   "swap",
		Node: &CapabilityNode{
			Kind:     "capability",
			Protocol: "simple-amm",
			Method:   "swap",
			Params: map[string]interface{}{
				"minOutput": minOutput.String(),
			},
			Tx: tx,
		},
	}, nil
}

// Pair returns the underlying contract binding (used by the Simulator for direct calls).
func (p *SimpleAMMProtocol) Pair() *contracts.SimpleAMMPair {
	return p.pair
}

// Address returns the SimpleAMMPair on-chain address.
func (p *SimpleAMMProtocol) Address() common.Address {
	return p.address
}
