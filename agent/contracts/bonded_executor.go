// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BondedExecutorPlan is an auto generated low-level Go binding around an user-defined struct.
type BondedExecutorPlan struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}

// BondedExecutorMetaData contains all meta data concerning the BondedExecutor contract.
var BondedExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_tUSDC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelExpiredPlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelPlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executePlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calldata_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executePlanWithSignature\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calldata_\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"lockedBond\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"openPlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plan\",\"type\":\"tuple\",\"internalType\":\"structBondedExecutor.Plan\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"failureCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calldataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"calldata_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingRefunds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"plans\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"failureCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calldataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"serviceFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setServiceFee\",\"inputs\":[{\"name\":\"_feeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tUSDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockUSDC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLockedBonds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawPendingRefund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BondReleased\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MONRefundStored\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanCancelled\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bondReturned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanExecuted\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"actualOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"userReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shortfallPaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanFailed\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBondedExecutor.FailureReason\"},{\"name\":\"refundedMON\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"compensationPaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanOpened\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bondDeposited\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"coverageFloor\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceFeeCollected\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"swapOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"userReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceFeeUpdated\",\"inputs\":[{\"name\":\"oldFeeBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShortfallPaid\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualUserReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"compensationPaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyExecuted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BelowCoverageFloor\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"floor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BondInvariantViolated\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"locked\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BondTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CalldataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBond\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonceAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPlanUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PlanExpired\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"now_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// BondedExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BondedExecutorMetaData.ABI instead.
var BondedExecutorABI = BondedExecutorMetaData.ABI

// BondedExecutor is an auto generated Go binding around an Ethereum contract.
type BondedExecutor struct {
	BondedExecutorCaller     // Read-only binding to the contract
	BondedExecutorTransactor // Write-only binding to the contract
	BondedExecutorFilterer   // Log filterer for contract events
}

// BondedExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondedExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondedExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondedExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondedExecutorSession struct {
	Contract     *BondedExecutor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondedExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondedExecutorCallerSession struct {
	Contract *BondedExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BondedExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondedExecutorTransactorSession struct {
	Contract     *BondedExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BondedExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondedExecutorRaw struct {
	Contract *BondedExecutor // Generic contract binding to access the raw methods on
}

// BondedExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondedExecutorCallerRaw struct {
	Contract *BondedExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BondedExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondedExecutorTransactorRaw struct {
	Contract *BondedExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondedExecutor creates a new instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutor(address common.Address, backend bind.ContractBackend) (*BondedExecutor, error) {
	contract, err := bindBondedExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondedExecutor{BondedExecutorCaller: BondedExecutorCaller{contract: contract}, BondedExecutorTransactor: BondedExecutorTransactor{contract: contract}, BondedExecutorFilterer: BondedExecutorFilterer{contract: contract}}, nil
}

// NewBondedExecutorCaller creates a new read-only instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorCaller(address common.Address, caller bind.ContractCaller) (*BondedExecutorCaller, error) {
	contract, err := bindBondedExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorCaller{contract: contract}, nil
}

// NewBondedExecutorTransactor creates a new write-only instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BondedExecutorTransactor, error) {
	contract, err := bindBondedExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorTransactor{contract: contract}, nil
}

// NewBondedExecutorFilterer creates a new log filterer instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BondedExecutorFilterer, error) {
	contract, err := bindBondedExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorFilterer{contract: contract}, nil
}

// bindBondedExecutor binds a generic wrapper to an already deployed contract.
func bindBondedExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BondedExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondedExecutor *BondedExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondedExecutor.Contract.BondedExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondedExecutor *BondedExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.Contract.BondedExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondedExecutor *BondedExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondedExecutor.Contract.BondedExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondedExecutor *BondedExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondedExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondedExecutor *BondedExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondedExecutor *BondedExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondedExecutor.Contract.contract.Transact(opts, method, params...)
}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) LockedBond(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "lockedBond", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) LockedBond(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.LockedBond(&_BondedExecutor.CallOpts, arg0)
}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) LockedBond(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.LockedBond(&_BondedExecutor.CallOpts, arg0)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) PendingRefunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "pendingRefunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.PendingRefunds(&_BondedExecutor.CallOpts, arg0)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.PendingRefunds(&_BondedExecutor.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorCaller) Plans(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "plans", arg0)

	outstruct := new(struct {
		User                common.Address
		Operator            common.Address
		InputAmount         *big.Int
		ExpectedOutput      *big.Int
		GuaranteedOutput    *big.Int
		MaxCompensation     *big.Int
		FailureCompensation *big.Int
		Target              common.Address
		CalldataHash        [32]byte
		Deadline            *big.Int
		Nonce               *big.Int
		Executed            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExpectedOutput = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GuaranteedOutput = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxCompensation = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FailureCompensation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Target = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.CalldataHash = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.Deadline = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorSession) Plans(arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	return _BondedExecutor.Contract.Plans(&_BondedExecutor.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorCallerSession) Plans(arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	return _BondedExecutor.Contract.Plans(&_BondedExecutor.CallOpts, arg0)
}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) ServiceFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "serviceFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) ServiceFeeBps() (*big.Int, error) {
	return _BondedExecutor.Contract.ServiceFeeBps(&_BondedExecutor.CallOpts)
}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) ServiceFeeBps() (*big.Int, error) {
	return _BondedExecutor.Contract.ServiceFeeBps(&_BondedExecutor.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorCaller) TUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "tUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorSession) TUSDC() (common.Address, error) {
	return _BondedExecutor.Contract.TUSDC(&_BondedExecutor.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorCallerSession) TUSDC() (common.Address, error) {
	return _BondedExecutor.Contract.TUSDC(&_BondedExecutor.CallOpts)
}

// TotalLockedBonds is a free data retrieval call binding the contract method 0x01d77f71.
//
// Solidity: function totalLockedBonds() view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) TotalLockedBonds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "totalLockedBonds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedBonds is a free data retrieval call binding the contract method 0x01d77f71.
//
// Solidity: function totalLockedBonds() view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) TotalLockedBonds() (*big.Int, error) {
	return _BondedExecutor.Contract.TotalLockedBonds(&_BondedExecutor.CallOpts)
}

// TotalLockedBonds is a free data retrieval call binding the contract method 0x01d77f71.
//
// Solidity: function totalLockedBonds() view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) TotalLockedBonds() (*big.Int, error) {
	return _BondedExecutor.Contract.TotalLockedBonds(&_BondedExecutor.CallOpts)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorCaller) UsedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "usedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorSession) UsedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _BondedExecutor.Contract.UsedNonces(&_BondedExecutor.CallOpts, arg0, arg1)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorCallerSession) UsedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _BondedExecutor.Contract.UsedNonces(&_BondedExecutor.CallOpts, arg0, arg1)
}

// CancelExpiredPlan is a paid mutator transaction binding the contract method 0x4ee1bee4.
//
// Solidity: function cancelExpiredPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactor) CancelExpiredPlan(opts *bind.TransactOpts, planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "cancelExpiredPlan", planId)
}

// CancelExpiredPlan is a paid mutator transaction binding the contract method 0x4ee1bee4.
//
// Solidity: function cancelExpiredPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorSession) CancelExpiredPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelExpiredPlan(&_BondedExecutor.TransactOpts, planId)
}

// CancelExpiredPlan is a paid mutator transaction binding the contract method 0x4ee1bee4.
//
// Solidity: function cancelExpiredPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) CancelExpiredPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelExpiredPlan(&_BondedExecutor.TransactOpts, planId)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactor) CancelPlan(opts *bind.TransactOpts, planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "cancelPlan", planId)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorSession) CancelPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelPlan(&_BondedExecutor.TransactOpts, planId)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) CancelPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelPlan(&_BondedExecutor.TransactOpts, planId)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorTransactor) ExecutePlan(opts *bind.TransactOpts, planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "executePlan", planId, calldata_)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorSession) ExecutePlan(planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlan(&_BondedExecutor.TransactOpts, planId, calldata_)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorTransactorSession) ExecutePlan(planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlan(&_BondedExecutor.TransactOpts, planId, calldata_)
}

// ExecutePlanWithSignature is a paid mutator transaction binding the contract method 0x8a198306.
//
// Solidity: function executePlanWithSignature(bytes32 planId, bytes calldata_, uint256 deadline, bytes signature) payable returns()
func (_BondedExecutor *BondedExecutorTransactor) ExecutePlanWithSignature(opts *bind.TransactOpts, planId [32]byte, calldata_ []byte, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "executePlanWithSignature", planId, calldata_, deadline, signature)
}

// ExecutePlanWithSignature is a paid mutator transaction binding the contract method 0x8a198306.
//
// Solidity: function executePlanWithSignature(bytes32 planId, bytes calldata_, uint256 deadline, bytes signature) payable returns()
func (_BondedExecutor *BondedExecutorSession) ExecutePlanWithSignature(planId [32]byte, calldata_ []byte, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlanWithSignature(&_BondedExecutor.TransactOpts, planId, calldata_, deadline, signature)
}

// ExecutePlanWithSignature is a paid mutator transaction binding the contract method 0x8a198306.
//
// Solidity: function executePlanWithSignature(bytes32 planId, bytes calldata_, uint256 deadline, bytes signature) payable returns()
func (_BondedExecutor *BondedExecutorTransactorSession) ExecutePlanWithSignature(planId [32]byte, calldata_ []byte, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlanWithSignature(&_BondedExecutor.TransactOpts, planId, calldata_, deadline, signature)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorTransactor) OpenPlan(opts *bind.TransactOpts, planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "openPlan", planId, plan, calldata_)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorSession) OpenPlan(planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.OpenPlan(&_BondedExecutor.TransactOpts, planId, plan, calldata_)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) OpenPlan(planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.OpenPlan(&_BondedExecutor.TransactOpts, planId, plan, calldata_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorTransactor) SetServiceFee(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "setServiceFee", _feeBps)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorSession) SetServiceFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.Contract.SetServiceFee(&_BondedExecutor.TransactOpts, _feeBps)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) SetServiceFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.Contract.SetServiceFee(&_BondedExecutor.TransactOpts, _feeBps)
}

// WithdrawPendingRefund is a paid mutator transaction binding the contract method 0x4ca6fd2c.
//
// Solidity: function withdrawPendingRefund() returns()
func (_BondedExecutor *BondedExecutorTransactor) WithdrawPendingRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "withdrawPendingRefund")
}

// WithdrawPendingRefund is a paid mutator transaction binding the contract method 0x4ca6fd2c.
//
// Solidity: function withdrawPendingRefund() returns()
func (_BondedExecutor *BondedExecutorSession) WithdrawPendingRefund() (*types.Transaction, error) {
	return _BondedExecutor.Contract.WithdrawPendingRefund(&_BondedExecutor.TransactOpts)
}

// WithdrawPendingRefund is a paid mutator transaction binding the contract method 0x4ca6fd2c.
//
// Solidity: function withdrawPendingRefund() returns()
func (_BondedExecutor *BondedExecutorTransactorSession) WithdrawPendingRefund() (*types.Transaction, error) {
	return _BondedExecutor.Contract.WithdrawPendingRefund(&_BondedExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorSession) Receive() (*types.Transaction, error) {
	return _BondedExecutor.Contract.Receive(&_BondedExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _BondedExecutor.Contract.Receive(&_BondedExecutor.TransactOpts)
}

// BondedExecutorBondReleasedIterator is returned from FilterBondReleased and is used to iterate over the raw logs and unpacked data for BondReleased events raised by the BondedExecutor contract.
type BondedExecutorBondReleasedIterator struct {
	Event *BondedExecutorBondReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorBondReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorBondReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorBondReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorBondReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorBondReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorBondReleased represents a BondReleased event raised by the BondedExecutor contract.
type BondedExecutorBondReleased struct {
	PlanId   [32]byte
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondReleased is a free log retrieval operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) FilterBondReleased(opts *bind.FilterOpts, planId [][32]byte, operator []common.Address) (*BondedExecutorBondReleasedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "BondReleased", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorBondReleasedIterator{contract: _BondedExecutor.contract, event: "BondReleased", logs: logs, sub: sub}, nil
}

// WatchBondReleased is a free log subscription operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) WatchBondReleased(opts *bind.WatchOpts, sink chan<- *BondedExecutorBondReleased, planId [][32]byte, operator []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "BondReleased", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorBondReleased)
				if err := _BondedExecutor.contract.UnpackLog(event, "BondReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondReleased is a log parse operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) ParseBondReleased(log types.Log) (*BondedExecutorBondReleased, error) {
	event := new(BondedExecutorBondReleased)
	if err := _BondedExecutor.contract.UnpackLog(event, "BondReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorMONRefundStoredIterator is returned from FilterMONRefundStored and is used to iterate over the raw logs and unpacked data for MONRefundStored events raised by the BondedExecutor contract.
type BondedExecutorMONRefundStoredIterator struct {
	Event *BondedExecutorMONRefundStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorMONRefundStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorMONRefundStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorMONRefundStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorMONRefundStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorMONRefundStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorMONRefundStored represents a MONRefundStored event raised by the BondedExecutor contract.
type BondedExecutorMONRefundStored struct {
	PlanId [32]byte
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMONRefundStored is a free log retrieval operation binding the contract event 0xed05e0f3c7e9f464a39484f2ff8dd2de759bb8d7dfcd36978b194b7124d505d1.
//
// Solidity: event MONRefundStored(bytes32 indexed planId, address indexed user, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) FilterMONRefundStored(opts *bind.FilterOpts, planId [][32]byte, user []common.Address) (*BondedExecutorMONRefundStoredIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "MONRefundStored", planIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorMONRefundStoredIterator{contract: _BondedExecutor.contract, event: "MONRefundStored", logs: logs, sub: sub}, nil
}

// WatchMONRefundStored is a free log subscription operation binding the contract event 0xed05e0f3c7e9f464a39484f2ff8dd2de759bb8d7dfcd36978b194b7124d505d1.
//
// Solidity: event MONRefundStored(bytes32 indexed planId, address indexed user, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) WatchMONRefundStored(opts *bind.WatchOpts, sink chan<- *BondedExecutorMONRefundStored, planId [][32]byte, user []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "MONRefundStored", planIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorMONRefundStored)
				if err := _BondedExecutor.contract.UnpackLog(event, "MONRefundStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMONRefundStored is a log parse operation binding the contract event 0xed05e0f3c7e9f464a39484f2ff8dd2de759bb8d7dfcd36978b194b7124d505d1.
//
// Solidity: event MONRefundStored(bytes32 indexed planId, address indexed user, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) ParseMONRefundStored(log types.Log) (*BondedExecutorMONRefundStored, error) {
	event := new(BondedExecutorMONRefundStored)
	if err := _BondedExecutor.contract.UnpackLog(event, "MONRefundStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanCancelledIterator is returned from FilterPlanCancelled and is used to iterate over the raw logs and unpacked data for PlanCancelled events raised by the BondedExecutor contract.
type BondedExecutorPlanCancelledIterator struct {
	Event *BondedExecutorPlanCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorPlanCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorPlanCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorPlanCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanCancelled represents a PlanCancelled event raised by the BondedExecutor contract.
type BondedExecutorPlanCancelled struct {
	PlanId       [32]byte
	Operator     common.Address
	BondReturned *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPlanCancelled is a free log retrieval operation binding the contract event 0x63f923f5e4c8cd463d0b134dd027b513d638239834fa73dbacddc049e1157aa3.
//
// Solidity: event PlanCancelled(bytes32 indexed planId, address indexed operator, uint256 bondReturned)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanCancelled(opts *bind.FilterOpts, planId [][32]byte, operator []common.Address) (*BondedExecutorPlanCancelledIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanCancelled", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanCancelledIterator{contract: _BondedExecutor.contract, event: "PlanCancelled", logs: logs, sub: sub}, nil
}

// WatchPlanCancelled is a free log subscription operation binding the contract event 0x63f923f5e4c8cd463d0b134dd027b513d638239834fa73dbacddc049e1157aa3.
//
// Solidity: event PlanCancelled(bytes32 indexed planId, address indexed operator, uint256 bondReturned)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanCancelled(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanCancelled, planId [][32]byte, operator []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanCancelled", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanCancelled)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlanCancelled is a log parse operation binding the contract event 0x63f923f5e4c8cd463d0b134dd027b513d638239834fa73dbacddc049e1157aa3.
//
// Solidity: event PlanCancelled(bytes32 indexed planId, address indexed operator, uint256 bondReturned)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanCancelled(log types.Log) (*BondedExecutorPlanCancelled, error) {
	event := new(BondedExecutorPlanCancelled)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanExecutedIterator is returned from FilterPlanExecuted and is used to iterate over the raw logs and unpacked data for PlanExecuted events raised by the BondedExecutor contract.
type BondedExecutorPlanExecutedIterator struct {
	Event *BondedExecutorPlanExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorPlanExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorPlanExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorPlanExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanExecuted represents a PlanExecuted event raised by the BondedExecutor contract.
type BondedExecutorPlanExecuted struct {
	PlanId        [32]byte
	ActualOutput  *big.Int
	UserReceived  *big.Int
	ShortfallPaid *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPlanExecuted is a free log retrieval operation binding the contract event 0x58ec66b1c9ca5cce26751cd0b7ff8b84645928bbcbd2f1c0963e8275239182f9.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 userReceived, uint256 shortfallPaid)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanExecuted(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorPlanExecutedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanExecuted", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanExecutedIterator{contract: _BondedExecutor.contract, event: "PlanExecuted", logs: logs, sub: sub}, nil
}

// WatchPlanExecuted is a free log subscription operation binding the contract event 0x58ec66b1c9ca5cce26751cd0b7ff8b84645928bbcbd2f1c0963e8275239182f9.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 userReceived, uint256 shortfallPaid)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanExecuted(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanExecuted, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanExecuted", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanExecuted)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlanExecuted is a log parse operation binding the contract event 0x58ec66b1c9ca5cce26751cd0b7ff8b84645928bbcbd2f1c0963e8275239182f9.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 userReceived, uint256 shortfallPaid)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanExecuted(log types.Log) (*BondedExecutorPlanExecuted, error) {
	event := new(BondedExecutorPlanExecuted)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanFailedIterator is returned from FilterPlanFailed and is used to iterate over the raw logs and unpacked data for PlanFailed events raised by the BondedExecutor contract.
type BondedExecutorPlanFailedIterator struct {
	Event *BondedExecutorPlanFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorPlanFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorPlanFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorPlanFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanFailed represents a PlanFailed event raised by the BondedExecutor contract.
type BondedExecutorPlanFailed struct {
	PlanId           [32]byte
	Reason           uint8
	RefundedMON      *big.Int
	CompensationPaid *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlanFailed is a free log retrieval operation binding the contract event 0xa43ea96c1ea78b701235496f1fc12c1e46f1b266164ca2bfa684aee65e109f75.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint8 reason, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanFailed(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorPlanFailedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanFailed", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanFailedIterator{contract: _BondedExecutor.contract, event: "PlanFailed", logs: logs, sub: sub}, nil
}

// WatchPlanFailed is a free log subscription operation binding the contract event 0xa43ea96c1ea78b701235496f1fc12c1e46f1b266164ca2bfa684aee65e109f75.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint8 reason, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanFailed(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanFailed, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanFailed", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanFailed)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlanFailed is a log parse operation binding the contract event 0xa43ea96c1ea78b701235496f1fc12c1e46f1b266164ca2bfa684aee65e109f75.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint8 reason, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanFailed(log types.Log) (*BondedExecutorPlanFailed, error) {
	event := new(BondedExecutorPlanFailed)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanOpenedIterator is returned from FilterPlanOpened and is used to iterate over the raw logs and unpacked data for PlanOpened events raised by the BondedExecutor contract.
type BondedExecutorPlanOpenedIterator struct {
	Event *BondedExecutorPlanOpened // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorPlanOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanOpened)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorPlanOpened)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorPlanOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanOpened represents a PlanOpened event raised by the BondedExecutor contract.
type BondedExecutorPlanOpened struct {
	PlanId           [32]byte
	User             common.Address
	Operator         common.Address
	GuaranteedOutput *big.Int
	BondDeposited    *big.Int
	CoverageFloor    *big.Int
	Deadline         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlanOpened is a free log retrieval operation binding the contract event 0xabd58b0451918ad74ab3b65620471fa4310f5db59a6748ce749f287f1da77400.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 bondDeposited, uint256 coverageFloor, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanOpened(opts *bind.FilterOpts, planId [][32]byte, user []common.Address, operator []common.Address) (*BondedExecutorPlanOpenedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanOpened", planIdRule, userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanOpenedIterator{contract: _BondedExecutor.contract, event: "PlanOpened", logs: logs, sub: sub}, nil
}

// WatchPlanOpened is a free log subscription operation binding the contract event 0xabd58b0451918ad74ab3b65620471fa4310f5db59a6748ce749f287f1da77400.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 bondDeposited, uint256 coverageFloor, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanOpened(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanOpened, planId [][32]byte, user []common.Address, operator []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanOpened", planIdRule, userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanOpened)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanOpened", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlanOpened is a log parse operation binding the contract event 0xabd58b0451918ad74ab3b65620471fa4310f5db59a6748ce749f287f1da77400.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 bondDeposited, uint256 coverageFloor, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanOpened(log types.Log) (*BondedExecutorPlanOpened, error) {
	event := new(BondedExecutorPlanOpened)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorServiceFeeCollectedIterator is returned from FilterServiceFeeCollected and is used to iterate over the raw logs and unpacked data for ServiceFeeCollected events raised by the BondedExecutor contract.
type BondedExecutorServiceFeeCollectedIterator struct {
	Event *BondedExecutorServiceFeeCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorServiceFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorServiceFeeCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorServiceFeeCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorServiceFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorServiceFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorServiceFeeCollected represents a ServiceFeeCollected event raised by the BondedExecutor contract.
type BondedExecutorServiceFeeCollected struct {
	PlanId       [32]byte
	SwapOutput   *big.Int
	Fee          *big.Int
	UserReceived *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeCollected is a free log retrieval operation binding the contract event 0x585eccf035397876959a1639f73f3dbfe3006a3035e66f795ab38e4ed42658e4.
//
// Solidity: event ServiceFeeCollected(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) FilterServiceFeeCollected(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorServiceFeeCollectedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ServiceFeeCollected", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorServiceFeeCollectedIterator{contract: _BondedExecutor.contract, event: "ServiceFeeCollected", logs: logs, sub: sub}, nil
}

// WatchServiceFeeCollected is a free log subscription operation binding the contract event 0x585eccf035397876959a1639f73f3dbfe3006a3035e66f795ab38e4ed42658e4.
//
// Solidity: event ServiceFeeCollected(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) WatchServiceFeeCollected(opts *bind.WatchOpts, sink chan<- *BondedExecutorServiceFeeCollected, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ServiceFeeCollected", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorServiceFeeCollected)
				if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceFeeCollected is a log parse operation binding the contract event 0x585eccf035397876959a1639f73f3dbfe3006a3035e66f795ab38e4ed42658e4.
//
// Solidity: event ServiceFeeCollected(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) ParseServiceFeeCollected(log types.Log) (*BondedExecutorServiceFeeCollected, error) {
	event := new(BondedExecutorServiceFeeCollected)
	if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorServiceFeeUpdatedIterator is returned from FilterServiceFeeUpdated and is used to iterate over the raw logs and unpacked data for ServiceFeeUpdated events raised by the BondedExecutor contract.
type BondedExecutorServiceFeeUpdatedIterator struct {
	Event *BondedExecutorServiceFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorServiceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorServiceFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorServiceFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorServiceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorServiceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorServiceFeeUpdated represents a ServiceFeeUpdated event raised by the BondedExecutor contract.
type BondedExecutorServiceFeeUpdated struct {
	OldFeeBps *big.Int
	NewFeeBps *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeUpdated is a free log retrieval operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) FilterServiceFeeUpdated(opts *bind.FilterOpts) (*BondedExecutorServiceFeeUpdatedIterator, error) {

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ServiceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BondedExecutorServiceFeeUpdatedIterator{contract: _BondedExecutor.contract, event: "ServiceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceFeeUpdated is a free log subscription operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) WatchServiceFeeUpdated(opts *bind.WatchOpts, sink chan<- *BondedExecutorServiceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ServiceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorServiceFeeUpdated)
				if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceFeeUpdated is a log parse operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) ParseServiceFeeUpdated(log types.Log) (*BondedExecutorServiceFeeUpdated, error) {
	event := new(BondedExecutorServiceFeeUpdated)
	if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorShortfallPaidIterator is returned from FilterShortfallPaid and is used to iterate over the raw logs and unpacked data for ShortfallPaid events raised by the BondedExecutor contract.
type BondedExecutorShortfallPaidIterator struct {
	Event *BondedExecutorShortfallPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondedExecutorShortfallPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorShortfallPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondedExecutorShortfallPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondedExecutorShortfallPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorShortfallPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorShortfallPaid represents a ShortfallPaid event raised by the BondedExecutor contract.
type BondedExecutorShortfallPaid struct {
	PlanId             [32]byte
	GuaranteedOutput   *big.Int
	ActualUserReceived *big.Int
	CompensationPaid   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterShortfallPaid is a free log retrieval operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteedOutput, uint256 actualUserReceived, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) FilterShortfallPaid(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorShortfallPaidIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ShortfallPaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorShortfallPaidIterator{contract: _BondedExecutor.contract, event: "ShortfallPaid", logs: logs, sub: sub}, nil
}

// WatchShortfallPaid is a free log subscription operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteedOutput, uint256 actualUserReceived, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) WatchShortfallPaid(opts *bind.WatchOpts, sink chan<- *BondedExecutorShortfallPaid, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ShortfallPaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorShortfallPaid)
				if err := _BondedExecutor.contract.UnpackLog(event, "ShortfallPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShortfallPaid is a log parse operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteedOutput, uint256 actualUserReceived, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) ParseShortfallPaid(log types.Log) (*BondedExecutorShortfallPaid, error) {
	event := new(BondedExecutorShortfallPaid)
	if err := _BondedExecutor.contract.UnpackLog(event, "ShortfallPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
