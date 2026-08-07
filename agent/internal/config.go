package internal

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

// Config holds all deployment parameters, loaded from environment variables.
type Config struct {
	RPC        string
	PrivateKey string

	// Contract addresses
	MockUSDC       common.Address
	DexAddr        common.Address // DEX (MockDex or SimpleAMMPair)
	BondedExecutor common.Address

	// Operator parameters
	Operator         common.Address
	OperatorProfiles []OperatorProfile // loaded from ops.json

	// ERC-8004 identity contracts (deployed on Monad Testnet)
	ERC8004IdentityAddr   common.Address
	ERC8004ReputationAddr common.Address

	// Simulation defaults
	DefaultGuaranteeRatio float64 // e.g. 0.90 = guarantee 90% of expected output
	DefaultMaxCompensation *big.Int
	DefaultFailureComp     *big.Int
	DefaultDeadlineSeconds int64
}

// LoadConfig reads configuration from environment variables with sensible defaults.
func LoadConfig() *Config {
	// Auto-load .env file — no need to source manually
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env") // also try project root

	cfg := &Config{
		RPC:        getEnv("MONAD_RPC", "http://127.0.0.1:8545"),
		PrivateKey: getEnv("OPERATOR_PRIVATE_KEY", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),

		MockUSDC:       common.HexToAddress(getEnv("USDC_ADDR", "")),
		DexAddr:        common.HexToAddress(getEnv("DEX_ADDR", "")),
		BondedExecutor: common.HexToAddress(getEnv("EXECUTOR_ADDR", "")),

		DefaultGuaranteeRatio:  parseFloat(getEnv("GUARANTEE_RATIO", "0.90")),
		DefaultMaxCompensation: parseBig(getEnv("MAX_COMPENSATION", "20000000000000000000")),  // 20 tUSDC
		DefaultFailureComp:     parseBig(getEnv("FAILURE_COMPENSATION", "5000000000000000000")), // 5 tUSDC
		DefaultDeadlineSeconds: parseInt(getEnv("DEADLINE_SECONDS", "86400")),                   // 24h

		ERC8004IdentityAddr:   common.HexToAddress(getEnv("ERC8004_IDENTITY_ADDR", "")),
		ERC8004ReputationAddr: common.HexToAddress(getEnv("ERC8004_REPUTATION_ADDR", "")),

		OperatorProfiles: LoadOperatorProfiles("ops.json"),
	}

	return cfg
}

// Validate checks that all required addresses are set.
func (c *Config) Validate() error {
	if c.MockUSDC == (common.Address{}) {
		return fmt.Errorf("USDC_ADDR not set")
	}
	if c.DexAddr == (common.Address{}) {
		return fmt.Errorf("DEX_ADDR not set")
	}
	if c.BondedExecutor == (common.Address{}) {
		return fmt.Errorf("EXECUTOR_ADDR not set")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseBig(s string) *big.Int {
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}

func parseInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// LoadOperatorProfiles reads operator metadata from ops.json.
// Falls back to empty list if the file is missing or malformed.
func LoadOperatorProfiles(path string) []OperatorProfile {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var profiles []OperatorProfile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil
	}
	return profiles
}
