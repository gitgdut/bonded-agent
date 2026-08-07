package internal

import (
	"encoding/hex"
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
	"github.com/gitgdut/bonded-agent/agent/internal/protocols"
)

// ── JSON types (matching frontend types.ts) ──────────────────

type quoteResponse struct {
	InputAmount    string `json:"inputAmount"`
	OutputToken    string `json:"outputToken"`
	ExpectedOutput string `json:"expectedOutput"`
	SimulatedRate  string `json:"simulatedRate"`
	Timestamp      int64  `json:"timestamp"`
	Protocol       string `json:"protocol"`
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

type executeSignedRequest struct {
	Signature string `json:"signature"` // hex-encoded 65-byte EIP-712 signature
	Deadline  int64  `json:"deadline"`  // unix seconds
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

	// Moss-style discover → load → simulate endpoints
	mux.HandleFunc("/discover", a.handleDiscover)
	mux.HandleFunc("/load", a.handleLoad)
	mux.HandleFunc("/simulate", a.handleSimulate)

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
		Protocol:       "simple-amm",
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

	planID, txHash, netExpected, err := a.OpenPlanQuick(user, inputAmount, a.cfg.DefaultGuaranteeRatio)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "CREATE_FAILED", err.Error())
		return
	}

	// guaranteed = netExpected * ratio (already computed in OpenPlanQuick)
	guaranteed := new(big.Int).Set(netExpected)
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
		Target:              a.cfg.DexAddr.Hex(),
		CalldataHash:        calldataHash,
		TxHash:              txHash,
	})
}

func (a *Agent) handlePlanByID(w http.ResponseWriter, r *http.Request) {
	// Extract planId from /plans/<id> or /plans/<id>/execute-signed
	path := strings.TrimPrefix(r.URL.Path, "/plans/")
	path = strings.TrimSuffix(path, "/execute-signed")

	if path == "" {
		writeError(w, http.StatusBadRequest, "MISSING_ID", "缺少 planId")
		return
	}

	// POST /plans/:id/execute-signed — execute with EIP-712 signature
	if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/execute-signed") {
		a.handleExecuteSigned(w, r, path)
		return
	}

	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 GET 或 POST .../execute-signed")
		return
	}

	id := path

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

// handleExecuteSigned handles POST /plans/:id/execute-signed.
// The user signs off-chain, the operator submits on-chain.
func (a *Agent) handleExecuteSigned(w http.ResponseWriter, r *http.Request, id string) {
	var req executeSignedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "请求体格式错误")
		return
	}

	// Decode hex signature
	sigHex := strings.TrimPrefix(req.Signature, "0x")
	sig, err := hex.DecodeString(sigHex)
	if err != nil || len(sig) != 65 {
		writeError(w, http.StatusBadRequest, "INVALID_SIGNATURE", "签名格式错误，需要 65 字节 hex")
		return
	}

	planID := common.HexToHash(id)
	txHash, err := a.ExecutePlanWithSignature(planID, req.Deadline, sig)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "EXECUTE_FAILED", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"txHash": txHash})
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

// ── Moss-style discover → load → simulate handlers ──────────

// handleDiscover returns all available protocol capabilities and queries.
// GET /discover?verb=swap&category=dex&protocol=simple-amm
func (a *Agent) handleDiscover(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 GET")
		return
	}

	q := r.URL.Query()
	filter := protocols.DiscoverFilter{
		Verb:     protocols.Verb(q.Get("verb")),
		Category: protocols.Category(q.Get("category")),
		Protocol: q.Get("protocol"),
	}

	coords := a.registry.Discover(filter)
	writeJSON(w, http.StatusOK, coords)
}

// handleLoad returns full calling contracts for requested methods.
// POST /load — body: {"items": [{"protocol":"simple-amm","method":"quote"}]}
func (a *Agent) handleLoad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 POST")
		return
	}

	var req struct {
		Items []struct {
			Protocol string `json:"protocol"`
			Method   string `json:"method"`
		} `json:"items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "请求体格式错误")
		return
	}

	var items []struct{ Protocol, Method string }
	for _, item := range req.Items {
		items = append(items, struct{ Protocol, Method string }{item.Protocol, item.Method})
	}

	stubs := a.registry.Load(items)
	writeJSON(w, http.StatusOK, stubs)
}

// handleSimulate simulates a capability transaction via eth_call.
// POST /simulate — body: {"tx": {"from":"0x...","to":"0x...","data":"0x...","value":"0"}}
func (a *Agent) handleSimulate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "仅支持 POST")
		return
	}

	var req struct {
		Tx struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Data  string `json:"data"`
			Value string `json:"value"`
		} `json:"tx"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "请求体格式错误")
		return
	}

	value, _ := new(big.Int).SetString(req.Tx.Value, 10)
	if value == nil {
		value = big.NewInt(0)
	}

	tx := protocols.TransactionNode{
		From:  common.HexToAddress(req.Tx.From),
		To:    common.HexToAddress(req.Tx.To),
		Data:  common.FromHex(req.Tx.Data),
		Value: value,
	}

	sim := protocols.NewSimulator(a.client)
	result, err := sim.Simulate(tx)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	returnDataHex := "0x" + common.Bytes2Hex(result.ReturnData)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":    result.Success,
		"returnData": returnDataHex,
		"gasUsed":    fmt.Sprintf("%d", result.GasUsed),
	})
}

// ── Helpers ─────────────────────────────────────────────────

func (a *Agent) getCalldataHash() (string, error) {
	swapABI, err := contracts.SimpleAMMPairMetaData.GetAbi()
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
