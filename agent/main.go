package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ── Config (from env vars) ──────────────────────────────────

type Config struct {
	RPC             string
	PrivateKey      string
	BondedExecutor  common.Address
	MockDex         common.Address
	MockUSDC        common.Address
}

func loadConfig() *Config {
	return &Config{
		RPC:            getEnv("MONAD_RPC", "https://testnet-rpc.monad.xyz"),
		PrivateKey:     getEnv("OPERATOR_PRIVATE_KEY", ""),
		BondedExecutor: common.HexToAddress(getEnv("BONDED_EXECUTOR_ADDR", "")),
		MockDex:        common.HexToAddress(getEnv("DEX_ADDR", "")),
		MockUSDC:       common.HexToAddress(getEnv("USDC_ADDR", "")),
	}
}

// ── Agent Service ───────────────────────────────────────────

// Agent handles the lifecycle: simulation → plan creation → event monitoring
type Agent struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	cfg        *Config
}

func NewAgent(cfg *Config) (*Agent, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Agent{
		client:     client,
		privateKey: privateKey,
		address:    address,
		cfg:        cfg,
	}, nil
}

// SimulateSwap queries MockDex rate and returns expected output
func (a *Agent) SimulateSwap(monAmount *big.Int) (*big.Int, error) {
	// TODO: call MockDex.rate() and compute monAmount * rate / 1e18
	_ = context.Background()
	return big.NewInt(0), fmt.Errorf("not implemented")
}

// OpenPlan creates a guaranteed plan on BondedExecutor
func (a *Agent) OpenPlan(
	user common.Address,
	inputAmount *big.Int,
	expectedOutput *big.Int,
	guaranteedOutput *big.Int,
	maxCompensation *big.Int,
	failureCompensation *big.Int,
	target common.Address,
	calldata []byte,
	deadline *big.Int,
	nonce *big.Int,
) (string, error) {
	// TODO: build Plan struct, compute planId, call BondedExecutor.openPlan()
	_ = context.Background()
	return "", fmt.Errorf("not implemented")
}

// ── Main ────────────────────────────────────────────────────

func main() {
	cfg := loadConfig()
	fmt.Printf("Bonded Agent starting...\n")
	fmt.Printf("  RPC: %s\n", cfg.RPC)
	fmt.Printf("  Executor: %s\n", cfg.BondedExecutor.Hex())

	if cfg.PrivateKey == "" {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatalf("generate key: %v", err)
		}
		cfg.PrivateKey = hex.EncodeToString(crypto.FromECDSA(privateKey))
		fmt.Printf("  ⚠️  No private key set — generated read-only key: %s...\n", cfg.PrivateKey[:16])
	}

	_, err := NewAgent(cfg)
	if err != nil {
		log.Fatalf("init agent: %v", err)
	}

	fmt.Println("Agent initialized successfully.")
	fmt.Println("TODO: implement SimulateSwap, OpenPlan, and event monitoring")
}

// ── Helpers ─────────────────────────────────────────────────

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
