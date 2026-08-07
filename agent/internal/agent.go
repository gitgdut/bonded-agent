package internal

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gitgdut/bonded-agent/agent/contracts"
)

// Agent handles all interactions with Bonded Agent contracts.
type Agent struct {
	client  *ethclient.Client
	chainID *big.Int
	auth    *bind.TransactOpts
	address common.Address

	cfg      *Config // needed for contract addresses
	usdc     *contracts.MockUSDC
	dex      *contracts.SimpleAMMPair
	executor *contracts.BondedExecutor
}

// ── Initialization ──────────────────────────────────────────

// NewAgent connects to the chain and initializes contract bindings.
func NewAgent(cfg *Config) (*Agent, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("dial RPC %s: %w", cfg.RPC, err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get chain ID: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey[2:]) // strip 0x
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	cfg.Operator = address

	// Bind contracts
	usdc, err := contracts.NewMockUSDC(cfg.MockUSDC, client)
	if err != nil {
		return nil, fmt.Errorf("bind MockUSDC: %w", err)
	}

	dex, err := contracts.NewSimpleAMMPair(cfg.DexAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind SimpleAMMPair: %w", err)
	}

	executor, err := contracts.NewBondedExecutor(cfg.BondedExecutor, client)
	if err != nil {
		return nil, fmt.Errorf("bind BondedExecutor: %w", err)
	}

	return &Agent{
		client:    client,
		chainID:   chainID,
		auth:      auth,
		address:   address,
		cfg:       cfg,
		usdc:      usdc,
		dex:       dex,
		executor:  executor,
	}, nil
}

// Address returns the operator's Ethereum address.
func (a *Agent) Address() common.Address {
	return a.address
}

// ChainID returns the connected chain ID.
func (a *Agent) ChainID() *big.Int {
	return a.chainID
}

// ── Queries (free, read-only) ───────────────────────────────

// GetRate returns the estimated exchange rate (tUSDC per 1 MON).
// Uses the AMM formula: queries getAmountOut(1 MON) from the pool.
func (a *Agent) GetRate() (*big.Int, error) {
	oneMON := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return a.dex.GetAmountOut(nil, oneMON)
}

// SimulateSwap computes the expected tUSDC output for a given MON input.
// Uses the AMM constant-product formula (read-only call).
func (a *Agent) SimulateSwap(monAmount *big.Int) (*big.Int, error) {
	return a.dex.GetAmountOut(nil, monAmount)
}

// GetPlan retrieves a plan by its ID.
func (a *Agent) GetPlan(planID [32]byte) (contracts.BondedExecutorPlan, error) {
	return a.executor.Plans(nil, planID)
}

// GetUserBalance returns the tUSDC balance of an address.
func (a *Agent) GetUserBalance(addr common.Address) (*big.Int, error) {
	return a.usdc.BalanceOf(nil, addr)
}

// GetMONBalance returns the native MON balance of an address.
func (a *Agent) GetMONBalance(addr common.Address) (*big.Int, error) {
	return a.client.BalanceAt(context.Background(), addr, nil)
}

// ── Transactions (cost gas) ─────────────────────────────────

// ApproveUSDC approves the BondedExecutor to spend operator's tUSDC.
func (a *Agent) ApproveUSDC(amount *big.Int) (string, error) {
	tx, err := a.usdc.Approve(a.auth, a.cfg.BondedExecutor, amount)
	if err != nil {
		return "", fmt.Errorf("approve: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// OpenPlan creates a guaranteed plan on the BondedExecutor.
// Returns (planID, txHash, error).
func (a *Agent) OpenPlan(
	user common.Address,
	inputAmount *big.Int,
	expectedOutput *big.Int,
	guaranteedOutput *big.Int,
	maxCompensation *big.Int,
	failureCompensation *big.Int,
	deadline *big.Int,
) (string, string, error) {
	// Build calldata: swap(uint256 minOutput) with minOutput=0
	swapABI, err := contracts.SimpleAMMPairMetaData.GetAbi()
	if err != nil {
		return "", "", fmt.Errorf("get ABI: %w", err)
	}
	calldata, err := swapABI.Pack("swap", big.NewInt(0))
	if err != nil {
		return "", "", fmt.Errorf("pack calldata: %w", err)
	}
	calldataHash := crypto.Keccak256Hash(calldata)

	// Compute planID = keccak256(abi.encode(user, operator, nonce))
	nonce := big.NewInt(time.Now().UnixNano())
	planID := computePlanID(user, a.address, nonce)

	// Build plan struct
	plan := contracts.BondedExecutorPlan{
		User:                user,
		Operator:            a.address,
		InputAmount:         inputAmount,
		ExpectedOutput:      expectedOutput,
		GuaranteedOutput:    guaranteedOutput,
		MaxCompensation:     maxCompensation,
		FailureCompensation: failureCompensation,
		Target:              a.cfg.DexAddr,
		CalldataHash:        calldataHash,
		Deadline:            deadline,
		Nonce:               nonce,
		Executed:            false,
	}

	tx, err := a.executor.OpenPlan(a.auth, planID, plan, calldata)
	if err != nil {
		return "", "", fmt.Errorf("openPlan tx: %w", err)
	}

	return fmt.Sprintf("0x%x", planID), tx.Hash().Hex(), nil
}

// OpenPlanQuick is a convenience method that auto-computes expectedOutput and guaranteedOutput.
func (a *Agent) OpenPlanQuick(
	user common.Address,
	inputAmount *big.Int,
	ratio float64,
) (string, string, *big.Int, error) {
	expected, err := a.SimulateSwap(inputAmount)
	if err != nil {
		return "", "", nil, fmt.Errorf("simulate: %w", err)
	}

	// Query current service fee (basis points)
	feeBps, err := a.executor.ServiceFeeBps(nil)
	if err != nil {
		feeBps = big.NewInt(0) // fallback: no fee
	}

	// netExpected = expected * (10000 - feeBps) / 10000
	// This is what the user actually receives after fee deduction
	netExpected := new(big.Int).Set(expected)
	if feeBps.Sign() > 0 {
		feeFree := new(big.Int).Sub(big.NewInt(10000), feeBps)
		netExpected.Mul(netExpected, feeFree)
		netExpected.Div(netExpected, big.NewInt(10000))
	}

	// guaranteed = netExpected * ratio
	guaranteed := new(big.Int).Set(netExpected)
	guaranteed.Mul(guaranteed, big.NewInt(int64(ratio*1e9)))
	guaranteed.Div(guaranteed, big.NewInt(1e9))

	deadline := big.NewInt(time.Now().Unix() + 86400) // 24h

	maxComp, _ := new(big.Int).SetString("20000000000000000000", 10)  // 20 tUSDC
	failComp, _ := new(big.Int).SetString("5000000000000000000", 10)  // 5 tUSDC

	planID, txHash, err := a.OpenPlan(
		user, inputAmount, expected, guaranteed,
		maxComp,  // 20 tUSDC max compensation
		failComp, // 5 tUSDC failure compensation
		deadline,
	)

	return planID, txHash, netExpected, err
}

// ── Helpers ─────────────────────────────────────────────────

// ExecutePlan executes a plan as the user.
// Sends plan.InputAmount MON and the matching calldata.
func (a *Agent) ExecutePlan(planID [32]byte) (string, error) {
	plan, err := a.executor.Plans(nil, planID)
	if err != nil {
		return "", fmt.Errorf("get plan: %w", err)
	}
	if plan.Executed {
		return "", fmt.Errorf("plan already executed")
	}
	if plan.Operator == (common.Address{}) {
		return "", fmt.Errorf("plan not found")
	}

	// Build calldata matching what OpenPlan stored (swap(0))
	swapABI, err := contracts.SimpleAMMPairMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("get ABI: %w", err)
	}
	calldata, err := swapABI.Pack("swap", big.NewInt(0))
	if err != nil {
		return "", fmt.Errorf("pack calldata: %w", err)
	}

	// Create auth with value = plan.InputAmount
	auth := *a.auth // shallow copy
	auth.Value = new(big.Int).Set(plan.InputAmount)

	tx, err := a.executor.ExecutePlan(&auth, planID, calldata)
	if err != nil {
		return "", fmt.Errorf("executePlan tx: %w", err)
	}

	return tx.Hash().Hex(), nil
}

// ExecutePlanWithSignature executes a plan using an EIP-712 signature from the user.
// The operator sends the MON and pays gas; the user only signed off-chain.
func (a *Agent) ExecutePlanWithSignature(planID [32]byte, deadline int64, signature []byte) (string, error) {
	plan, err := a.executor.Plans(nil, planID)
	if err != nil {
		return "", fmt.Errorf("get plan: %w", err)
	}
	if plan.Executed {
		return "", fmt.Errorf("plan already executed")
	}
	if plan.Operator == (common.Address{}) {
		return "", fmt.Errorf("plan not found")
	}

	// Build calldata
	swapABI, err := contracts.SimpleAMMPairMetaData.GetAbi()
	if err != nil {
		return "", fmt.Errorf("get ABI: %w", err)
	}
	calldata, err := swapABI.Pack("swap", big.NewInt(0))
	if err != nil {
		return "", fmt.Errorf("pack calldata: %w", err)
	}

	// Create auth with value = plan.InputAmount (operator pays MON)
	auth := *a.auth
	auth.Value = new(big.Int).Set(plan.InputAmount)

	tx, err := a.executor.ExecutePlanWithSignature(&auth, planID, calldata, big.NewInt(deadline), signature)
	if err != nil {
		return "", fmt.Errorf("executePlanWithSignature tx: %w", err)
	}

	return tx.Hash().Hex(), nil
}

// computePlanID = keccak256(abi.encode(user, operator, nonce))
func computePlanID(user, operator common.Address, nonce *big.Int) [32]byte {
	// abi.encode(address,address,uint256)
	data := make([]byte, 0, 32+32+32)

	// pad addresses to 32 bytes
	paddedUser := common.LeftPadBytes(user.Bytes(), 32)
	paddedOp := common.LeftPadBytes(operator.Bytes(), 32)
	paddedNonce := common.LeftPadBytes(nonce.Bytes(), 32)

	data = append(data, paddedUser...)
	data = append(data, paddedOp...)
	data = append(data, paddedNonce...)

	return crypto.Keccak256Hash(data)
}
