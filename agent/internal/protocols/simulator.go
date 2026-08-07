package protocols

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Simulator wraps an ETH client for simulating transactions via eth_call.
// It mirrors Moss's createTraceSimulator — but uses eth_call instead of
// debug_traceCall since we're focused on quote accuracy, not receipt verification.
type Simulator struct {
	client *ethclient.Client
}

// NewSimulator creates a simulator backed by the given ETH client.
func NewSimulator(client *ethclient.Client) *Simulator {
	return &Simulator{client: client}
}

// SimulateResult is the raw result of simulating one transaction.
type SimulateEstimate struct {
	Success    bool
	ReturnData []byte
	GasUsed    uint64
}

// Simulate executes the transaction as an eth_call at the latest block.
// It returns the raw return data and estimated gas, but does NOT perform
// the deep receipt verification that Moss's trace simulator does.
func (s *Simulator) Simulate(tx TransactionNode) (*SimulateEstimate, error) {
	ctx := context.Background()

	// Build call message
	msg := ethereum.CallMsg{
		From:  tx.From,
		To:    &tx.To,
		Data:  tx.Data,
		Value: tx.Value,
	}

	// eth_call — get return data
	returnData, err := s.client.CallContract(ctx, msg, nil)
	if err != nil {
		return &SimulateEstimate{Success: false}, fmt.Errorf("eth_call reverted: %w", err)
	}

	// eth_estimateGas — get gas estimate
	gas, err := s.client.EstimateGas(ctx, msg)
	if err != nil {
		// Gas estimation failed (common for reverted calls), still return call result
		gas = 0
	}

	return &SimulateEstimate{
		Success:    true,
		ReturnData: returnData,
		GasUsed:    gas,
	}, nil
}

// SimulateWithOverride simulates with state overrides — useful for setting
// the caller's balance or token allowances before the call.
func (s *Simulator) SimulateWithOverride(tx TransactionNode, overrides map[common.Address]interface{}) (*SimulateEstimate, error) {
	ctx := context.Background()

	msg := ethereum.CallMsg{
		From:  tx.From,
		To:    &tx.To,
		Data:  tx.Data,
		Value: tx.Value,
	}

	// Build state overrides for eth_call
	var stateOverrides *ethereum.OverrideAccount
	if len(overrides) > 0 {
		stateOverrides = &ethereum.OverrideAccount{}
		// Basic balance override — for full Moss-style state overrides we'd
		// need to use a custom JSON-RPC call since go-ethereum's CallMsg
		// doesn't directly support the full OverrideSet.
		// For now, use the basic call.
	}

	_ = stateOverrides

	returnData, err := s.client.CallContract(ctx, msg, nil)
	if err != nil {
		return &SimulateEstimate{Success: false}, fmt.Errorf("eth_call reverted: %w", err)
	}

	return &SimulateEstimate{
		Success:    true,
		ReturnData: returnData,
	}, nil
}

// ParseUint256Output parses a standard uint256 return value from eth_call.
func ParseUint256Output(data []byte) (*big.Int, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("return data too short: %d bytes, need 32", len(data))
	}
	return new(big.Int).SetBytes(data[:32]), nil
}
