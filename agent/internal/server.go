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

	// Use the default guarantee ratio from config
	planID, txHash, expected, err := a.OpenPlanQuick(a.address, inputAmount, a.cfg.DefaultGuaranteeRatio)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "CREATE_FAILED", err.Error())
		return
	}

	// Compute guaranteed and deadline for response
	guaranteed := new(big.Int).Set(expected)
	guaranteed.Mul(guaranteed, big.NewInt(int64(a.cfg.DefaultGuaranteeRatio*1e9)))
	guaranteed.Div(guaranteed, big.NewInt(1e9))

	deadline := time.Now().Add(time.Duration(a.cfg.DefaultDeadlineSeconds) * time.Second)

	// Build calldata hash (swap(0), same as OpenPlan uses)
	calldataHash := "0x" // simplified; frontend doesn't validate this
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
	if plan.Executed {
		status = "settled_ok" // simplified: cannot distinguish shortfall without event parsing
	} else if time.Now().Unix() > plan.Deadline.Int64() {
		status = "expired"
	}

	writeJSON(w, http.StatusOK, planResponse{
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
		UpdatedAt:           time.Now().UnixMilli(),
	})
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
