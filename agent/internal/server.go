package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gitgdut/bonded-agent/agent/contracts"
)

// ── JSON types (matching frontend types.ts) ──────────────────

type quoteResponse struct {
	InputAmount    string `json:"inputAmount"`
	OutputToken    string `json:"outputToken"`
	ExpectedOutput string `json:"expectedOutput"`
	SimulatedRate  string `json:"simulatedRate"`
	Timestamp      int64  `json:"timestamp"`
}

type createPlanRequest struct {
	InputAmount     string `json:"inputAmount"`
	ExpectedOutput  string `json:"expectedOutput"`
	DeadlineMinutes int    `json:"deadlineMinutes"`
	UserAddress     string `json:"userAddress"`
}

type createPlanResponse struct {
	PlanID              string `json:"planId"`
	GuaranteedOutput    string `json:"guaranteedOutput"`
	MaxCompensation     string `json:"maxCompensation"`
	FailureCompensation string `json:"failureCompensation"`
	Deadline            int64  `json:"deadline"`
	Target              string `json:"target"`
	CalldataHash        string `json:"calldataHash"`
	TxHash              string `json:"txHash"`
}

type planResponse struct {
	PlanID              string   `json:"planId"`
	Status              string   `json:"status"`
	User                string   `json:"user"`
	InputAmount         string   `json:"inputAmount"`
	ExpectedOutput      string   `json:"expectedOutput"`
	GuaranteedOutput    string   `json:"guaranteedOutput"`
	MaxCompensation     string   `json:"maxCompensation"`
	FailureCompensation string   `json:"failureCompensation"`
	Deadline            int64    `json:"deadline"`
	TxHashes            []string `json:"txHashes"`
	ActualOutput        string   `json:"actualOutput,omitempty"`
	ShortfallPaid       string   `json:"shortfallPaid,omitempty"`
	Compensation        string   `json:"compensation,omitempty"`
	Refunded            bool     `json:"refunded"`
	BondReleased        bool     `json:"bondReleased"`
	UpdatedAt           int64    `json:"updatedAt"`
}

type planSettlement struct {
	status         string // "settled_ok" | "settled_shortfall" | "failed"
	actualOutput   string
	shortfallPaid  string
	compensation   string
	refunded       bool
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ── Server ──────────────────────────────────────────────────

// ServeAPI starts the HTTP API on the given address (e.g. ":8787").
func (a *Agent) ServeAPI(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/quote", a.handleQuote)
	mux.HandleFunc("/plans", a.handlePlans)
	mux.HandleFunc("/plans/", a.handlePlanByID)

	log.Printf("API listening on %s", addr)
	return http.ListenAndServe(addr, withCORS(mux))
}

// ── Handlers ────────────────────────────────────────────────

func (a *Agent) handleQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 GET")
		return
	}

	inputStr := r.URL.Query().Get("inputAmount")
	if inputStr == "" {
		writeError(w, http.StatusBadRequest, "MISSING_PARAM", "缺少 inputAmount 参数")
		return
	}

	inputAmount, ok := new(big.Int).SetString(inputStr, 10)
	if !ok {
		writeError(w, http.StatusBadRequest, "INVALID_AMOUNT", "inputAmount 格式错误，需要 wei 字符串")
		return
	}

	expected, err := a.SimulateSwap(inputAmount)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "SIMULATION_FAILED", err.Error())
		return
	}

	rate, _ := a.GetRate()

	writeJSON(w, http.StatusOK, quoteResponse{
		InputAmount:    inputAmount.String(),
		OutputToken:    "tUSDC",
		ExpectedOutput: expected.String(),
		SimulatedRate:  rate.String(),
		Timestamp:      time.Now().UnixMilli(),
	})
}

func (a *Agent) handlePlans(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// GET /plans — return empty list (frontend builds history from events)
		writeJSON(w, http.StatusOK, []planResponse{})
		return
	}

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 POST")
		return
	}

	var req createPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "请求体格式错误")
		return
	}

	inputAmount, ok := new(big.Int).SetString(req.InputAmount, 10)
	if !ok {
		writeError(w, http.StatusBadRequest, "INVALID_AMOUNT", "inputAmount 格式错误")
		return
	}

	// Use the user address from the request frontend, falling back to operator
	user := a.address
	if req.UserAddress != "" {
		user = common.HexToAddress(req.UserAddress)
	}

	planID, txHash, expected, err := a.OpenPlanQuick(user, inputAmount, a.cfg.DefaultGuaranteeRatio)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "CREATE_FAILED", err.Error())
		return
	}

	// Compute guaranteed and deadline for response
	guaranteed := new(big.Int).Set(expected)
	guaranteed.Mul(guaranteed, big.NewInt(int64(a.cfg.DefaultGuaranteeRatio*1e9)))
	guaranteed.Div(guaranteed, big.NewInt(1e9))

	deadline := time.Now().Add(time.Duration(a.cfg.DefaultDeadlineSeconds) * time.Second)

	calldataHash := ""
	if h, err := a.getCalldataHash(); err == nil {
		calldataHash = h
	}

	writeJSON(w, http.StatusOK, createPlanResponse{
		PlanID:              planID,
		GuaranteedOutput:    guaranteed.String(),
		MaxCompensation:     a.cfg.DefaultMaxCompensation.String(),
		FailureCompensation: a.cfg.DefaultFailureComp.String(),
		Deadline:            deadline.UnixMilli(),
		Target:              a.cfg.MockDex.Hex(),
		CalldataHash:        calldataHash,
		TxHash:              txHash,
	})
}

func (a *Agent) handlePlanByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 GET")
		return
	}

	// Extract planId from /plans/<id>
	id := strings.TrimPrefix(r.URL.Path, "/plans/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "MISSING_ID", "缺少 planId")
		return
	}

	planID := common.HexToHash(id)
	plan, err := a.executor.Plans(nil, planID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "QUERY_FAILED", err.Error())
		return
	}

	if plan.Operator == (common.Address{}) {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "计划不存在")
		return
	}

	// Determine status
	status := "open"
	var settlement planSettlement
	if plan.Executed {
		settlement = a.queryPlanSettlement(planID)
		status = settlement.status
	} else if time.Now().Unix() > plan.Deadline.Int64() {
		status = "expired"
	}

	resp := planResponse{
		PlanID:              fmt.Sprintf("0x%x", planID),
		Status:              status,
		User:                plan.User.Hex(),
		InputAmount:         plan.InputAmount.String(),
		ExpectedOutput:      plan.ExpectedOutput.String(),
		GuaranteedOutput:    plan.GuaranteedOutput.String(),
		MaxCompensation:     plan.MaxCompensation.String(),
		FailureCompensation: plan.FailureCompensation.String(),
		Deadline:            plan.Deadline.Int64() * 1000,
		TxHashes:            []string{},
		ActualOutput:        settlement.actualOutput,
		ShortfallPaid:       settlement.shortfallPaid,
		Compensation:        settlement.compensation,
		Refunded:            settlement.refunded,
		BondReleased:        plan.Executed,
		UpdatedAt:           time.Now().UnixMilli(),
	}

	writeJSON(w, http.StatusOK, resp)
}

// ── Settlement query helpers ────────────────────────────────

// queryPlanSettlement looks up on-chain events to determine the
// actual outcome of an executed plan.
func (a *Agent) queryPlanSettlement(planID [32]byte) planSettlement {
	var s planSettlement
	s.status = "settled_ok" // default

	// Check PlanFailed first — if the swap reverted, this event was emitted
	failedIter, err := a.executor.FilterPlanFailed(nil, [][32]byte{planID})
	if err == nil {
		defer failedIter.Close()
		for failedIter.Next() {
			ev := failedIter.Event
			s.status = "failed"
			s.refunded = true
			if ev.CompensationPaid != nil {
				s.compensation = ev.CompensationPaid.String()
			}
		}
	}

	// If already determined as failed, we're done
	if s.status == "failed" {
		return s
	}

	// Check PlanExecuted — always emitted when the swap call succeeds
	execIter, err := a.executor.FilterPlanExecuted(nil, [][32]byte{planID})
	if err == nil {
		defer execIter.Close()
		for execIter.Next() {
			ev := execIter.Event
			if ev.ActualOutput != nil {
				s.actualOutput = ev.ActualOutput.String()
			}
			// If paidToUser > 0, there was a shortfall
			if ev.PaidToUser != nil && ev.PaidToUser.Sign() > 0 {
				s.status = "settled_shortfall"
				s.compensation = ev.PaidToUser.String()
			}
		}
	}

	// If shortfall, also look up the ShortfallPaid event for exact shortfall amount
	if s.status == "settled_shortfall" {
		shortIter, err := a.executor.FilterShortfallPaid(nil, [][32]byte{planID})
		if err == nil {
			defer shortIter.Close()
			for shortIter.Next() {
				ev := shortIter.Event
				if ev.Shortfall != nil {
					s.shortfallPaid = ev.Shortfall.String()
				}
			}
		}
	}

	return s
}

// ── Helpers ─────────────────────────────────────────────────

func (a *Agent) getCalldataHash() (string, error) {
	swapABI, err := contracts.MockDexMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	calldata, err := swapABI.Pack("swap", big.NewInt(0))
	if err != nil {
		return "", err
	}
	hash := crypto.Keccak256Hash(calldata)
	return hash.Hex(), nil
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, apiError{Code: code, Message: message})
}

// withCORS adds CORS headers for frontend development.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
