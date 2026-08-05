package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gitgdut/bonded-agent/agent/internal"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg := internal.LoadConfig()

	// Validate addresses for all commands except info
	cmd := os.Args[1]
	if cmd != "info" && cmd != "help" {
		if err := cfg.Validate(); err != nil {
			log.Fatalf("Config error: %v\nSet USDC_ADDR, DEX_ADDR, EXECUTOR_ADDR env vars", err)
		}
	}

	agent, err := internal.NewAgent(cfg)
	if err != nil {
		log.Fatalf("Init agent: %v", err)
	}

	fmt.Printf("✓ Connected to chain %v as operator %s\n\n", agent.ChainID(), agent.Address().Hex())

	switch cmd {
	case "info":
		printInfo(agent)

	case "rate":
		rate, err := agent.GetRate()
		if err != nil {
			log.Fatalf("Get rate: %v", err)
		}
		fmt.Printf("Rate: 1 MON = %s tUSDC\n", formatToken(rate))

	case "simulate":
		if len(os.Args) < 3 {
			log.Fatal("Usage: agent simulate <MON amount>")
		}
		amount := parseEther(os.Args[2])
		output, err := agent.SimulateSwap(amount)
		if err != nil {
			log.Fatalf("Simulate: %v", err)
		}
		fmt.Printf("Input:  %s MON\n", formatEther(amount))
		fmt.Printf("Output: %s tUSDC\n", formatToken(output))

	case "approve":
		if len(os.Args) < 3 {
			log.Fatal("Usage: agent approve <amount>  (or 'max')")
		}
		amount := parseEther(os.Args[2])
		if os.Args[2] == "max" {
			// max uint256
			amount, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		}
		txHash, err := agent.ApproveUSDC(amount)
		if err != nil {
			log.Fatalf("Approve: %v", err)
		}
		fmt.Printf("Approved %s tUSDC → tx: %s\n", formatToken(amount), txHash)

	case "plan":
		if len(os.Args) < 4 {
			log.Fatal("Usage: agent plan <user-address> <MON-amount> [guarantee-ratio]")
		}
		user := common.HexToAddress(os.Args[2])
		amount := parseEther(os.Args[3])
		ratio := 0.90
		if len(os.Args) >= 5 {
			ratio = parseFloat(os.Args[4])
		}

		fmt.Printf("Creating plan for %s...\n", user.Hex())
		planID, txHash, expected, err := agent.OpenPlanQuick(user, amount, ratio)
		if err != nil {
			log.Fatalf("OpenPlan: %v", err)
		}

		guaranteed := new(big.Int).Set(expected)
		guaranteed.Mul(guaranteed, big.NewInt(int64(ratio*1e9)))
		guaranteed.Div(guaranteed, big.NewInt(1e9))

		fmt.Printf("┌─ Plan Created ──────────────────────\n")
		fmt.Printf("│ Plan ID:   %s\n", planID)
		fmt.Printf("│ Tx Hash:   %s\n", txHash)
		fmt.Printf("│ User:      %s\n", user.Hex())
		fmt.Printf("│ Amount:    %s MON\n", formatEther(amount))
		fmt.Printf("│ Expected:  %s tUSDC\n", formatToken(expected))
		fmt.Printf("│ Guaranteed:%s tUSDC (%.0f%%)\n", formatToken(guaranteed), ratio*100)
		fmt.Printf("└─────────────────────────────────────\n")

	case "get-plan":
		if len(os.Args) < 3 {
			log.Fatal("Usage: agent get-plan <planID>")
		}
		planID := [32]byte(common.HexToHash(os.Args[2]))
		plan, err := agent.GetPlan(planID)
		if err != nil {
			log.Fatalf("Get plan: %v", err)
		}
		fmt.Printf("Plan %s:\n", os.Args[2])
		fmt.Printf("  User:         %s\n", plan.User.Hex())
		fmt.Printf("  Operator:     %s\n", plan.Operator.Hex())
		fmt.Printf("  Input:        %s MON\n", formatEther(plan.InputAmount))
		fmt.Printf("  Expected:     %s tUSDC\n", formatToken(plan.ExpectedOutput))
		fmt.Printf("  Guaranteed:   %s tUSDC\n", formatToken(plan.GuaranteedOutput))
		fmt.Printf("  Max Comp:     %s tUSDC\n", formatToken(plan.MaxCompensation))
		fmt.Printf("  Fail Comp:    %s tUSDC\n", formatToken(plan.FailureCompensation))
		fmt.Printf("  Deadline:     %v\n", plan.Deadline)
		fmt.Printf("  Executed:     %v\n", plan.Executed)

	case "balance":
		addr := agent.Address()
		if len(os.Args) >= 3 {
			addr = common.HexToAddress(os.Args[2])
		}
		monBal, err := agent.GetMONBalance(addr)
		if err != nil {
			log.Fatalf("MON Balance: %v", err)
		}
		usdcBal, err := agent.GetUserBalance(addr)
		if err != nil {
			log.Fatalf("tUSDC Balance: %v", err)
		}
		fmt.Printf("%s:\n", addr.Hex())
		fmt.Printf("  MON:   %s\n", formatEther(monBal))
		fmt.Printf("  tUSDC: %s\n", formatToken(usdcBal))

	case "execute":
		if len(os.Args) < 3 {
			log.Fatal("Usage: agent execute <planID>")
		}
		planID := [32]byte(common.HexToHash(os.Args[2]))
		txHash, err := agent.ExecutePlan(planID)
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
		fmt.Printf("Plan executed ✓\nTx: %s\n", txHash)

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

// ── Helpers ─────────────────────────────────────────────────

func printInfo(agent *internal.Agent) {
	fmt.Printf("Chain ID:    %v\n", agent.ChainID())
	fmt.Printf("Operator:    %s\n", agent.Address().Hex())
}

func printUsage() {
	fmt.Println(strings.TrimSpace(`
Bonded Agent CLI — operator tool for the BondedExecutor protocol.

Usage:
  agent info                      Show chain + operator info
  agent rate                      Get current MockDex rate
  agent simulate <MON>            Simulate swap output (e.g. "1.5")
  agent approve <amount|max>      Approve USDC for BondedExecutor
  agent plan <user> <MON> [0.90]  Open a guaranteed plan
  agent execute <planID>          Execute a plan
  agent get-plan <planID>         Read a plan
  agent balance [addr]            Check tUSDC balance

Env vars:
  MONAD_RPC              RPC URL (default: http://127.0.0.1:8545)
  OPERATOR_PRIVATE_KEY   Operator private key (default: anvil #0)
  USDC_ADDR              MockUSDC address
  DEX_ADDR               MockDex address
  EXECUTOR_ADDR          BondedExecutor address
  GUARANTEE_RATIO        Default guarantee ratio (default: 0.90)
	`))
}

// parseEther parses "1.5" → 1.5e18 wei
func parseEther(s string) *big.Int {
	wei := new(big.Float)
	wei.SetString(s)

	ether := new(big.Float).SetFloat64(1e18)
	wei.Mul(wei, ether)

	result := new(big.Int)
	wei.Int(result)
	return result
}

func parseFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

// formatEther formats wei → "1.500000"
func formatEther(wei *big.Int) string {
	f := new(big.Float).SetInt(wei)
	f.Quo(f, big.NewFloat(1e18))
	return fmt.Sprintf("%.6f", f)
}

// formatToken formats 18-decimal token amount → "100.000000"
func formatToken(amount *big.Int) string {
	return formatEther(amount)
}
