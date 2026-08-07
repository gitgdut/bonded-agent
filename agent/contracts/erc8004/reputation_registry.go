// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc8004

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

// ERC8004ReputationRegistryFeedbackEntry is an auto generated low-level Go binding around an user-defined struct.
type ERC8004ReputationRegistryFeedbackEntry struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	FeedbackURI   string
	Timestamp     uint64
	IsRevoked     bool
}

// ERC8004ReputationRegistryMetaData contains all meta data concerning the ERC8004ReputationRegistry contract.
var ERC8004ReputationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClients\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastIndex\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSummary\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tag1\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag2\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"summaryValue\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"summaryValueDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"giveFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"valueDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag2\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feedbackURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readAllFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structERC8004ReputationRegistry.FeedbackEntry[]\",\"components\":[{\"name\":\"value\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"valueDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag2\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feedbackURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRevoked\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"valueDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag2\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feedbackURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"isRevoked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeedbackGiven\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"int128\",\"indexed\":false,\"internalType\":\"int128\"},{\"name\":\"valueDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeedbackRevoked\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ERC8004ReputationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC8004ReputationRegistryMetaData.ABI instead.
var ERC8004ReputationRegistryABI = ERC8004ReputationRegistryMetaData.ABI

// ERC8004ReputationRegistry is an auto generated Go binding around an Ethereum contract.
type ERC8004ReputationRegistry struct {
	ERC8004ReputationRegistryCaller     // Read-only binding to the contract
	ERC8004ReputationRegistryTransactor // Write-only binding to the contract
	ERC8004ReputationRegistryFilterer   // Log filterer for contract events
}

// ERC8004ReputationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC8004ReputationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004ReputationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC8004ReputationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004ReputationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC8004ReputationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004ReputationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC8004ReputationRegistrySession struct {
	Contract     *ERC8004ReputationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC8004ReputationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC8004ReputationRegistryCallerSession struct {
	Contract *ERC8004ReputationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ERC8004ReputationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC8004ReputationRegistryTransactorSession struct {
	Contract     *ERC8004ReputationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ERC8004ReputationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC8004ReputationRegistryRaw struct {
	Contract *ERC8004ReputationRegistry // Generic contract binding to access the raw methods on
}

// ERC8004ReputationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC8004ReputationRegistryCallerRaw struct {
	Contract *ERC8004ReputationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC8004ReputationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC8004ReputationRegistryTransactorRaw struct {
	Contract *ERC8004ReputationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC8004ReputationRegistry creates a new instance of ERC8004ReputationRegistry, bound to a specific deployed contract.
func NewERC8004ReputationRegistry(address common.Address, backend bind.ContractBackend) (*ERC8004ReputationRegistry, error) {
	contract, err := bindERC8004ReputationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistry{ERC8004ReputationRegistryCaller: ERC8004ReputationRegistryCaller{contract: contract}, ERC8004ReputationRegistryTransactor: ERC8004ReputationRegistryTransactor{contract: contract}, ERC8004ReputationRegistryFilterer: ERC8004ReputationRegistryFilterer{contract: contract}}, nil
}

// NewERC8004ReputationRegistryCaller creates a new read-only instance of ERC8004ReputationRegistry, bound to a specific deployed contract.
func NewERC8004ReputationRegistryCaller(address common.Address, caller bind.ContractCaller) (*ERC8004ReputationRegistryCaller, error) {
	contract, err := bindERC8004ReputationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryCaller{contract: contract}, nil
}

// NewERC8004ReputationRegistryTransactor creates a new write-only instance of ERC8004ReputationRegistry, bound to a specific deployed contract.
func NewERC8004ReputationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC8004ReputationRegistryTransactor, error) {
	contract, err := bindERC8004ReputationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryTransactor{contract: contract}, nil
}

// NewERC8004ReputationRegistryFilterer creates a new log filterer instance of ERC8004ReputationRegistry, bound to a specific deployed contract.
func NewERC8004ReputationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC8004ReputationRegistryFilterer, error) {
	contract, err := bindERC8004ReputationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryFilterer{contract: contract}, nil
}

// bindERC8004ReputationRegistry binds a generic wrapper to an already deployed contract.
func bindERC8004ReputationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC8004ReputationRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8004ReputationRegistry.Contract.ERC8004ReputationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.ERC8004ReputationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.ERC8004ReputationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8004ReputationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) GetClients(opts *bind.CallOpts, agentId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "getClients", agentId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.GetClients(&_ERC8004ReputationRegistry.CallOpts, agentId)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.GetClients(&_ERC8004ReputationRegistry.CallOpts, agentId)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address client) view returns(uint64)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) GetLastIndex(opts *bind.CallOpts, agentId *big.Int, client common.Address) (uint64, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "getLastIndex", agentId, client)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address client) view returns(uint64)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) GetLastIndex(agentId *big.Int, client common.Address) (uint64, error) {
	return _ERC8004ReputationRegistry.Contract.GetLastIndex(&_ERC8004ReputationRegistry.CallOpts, agentId, client)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address client) view returns(uint64)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) GetLastIndex(agentId *big.Int, client common.Address) (uint64, error) {
	return _ERC8004ReputationRegistry.Contract.GetLastIndex(&_ERC8004ReputationRegistry.CallOpts, agentId, client)
}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) GetSummary(opts *bind.CallOpts, agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "getSummary", agentId, clientAddresses, tag1, tag2)

	outstruct := new(struct {
		Count                uint64
		SummaryValue         *big.Int
		SummaryValueDecimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SummaryValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SummaryValueDecimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	return _ERC8004ReputationRegistry.Contract.GetSummary(&_ERC8004ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	return _ERC8004ReputationRegistry.Contract.GetSummary(&_ERC8004ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) IdentityRegistry() (common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.IdentityRegistry(&_ERC8004ReputationRegistry.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) IdentityRegistry() (common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.IdentityRegistry(&_ERC8004ReputationRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) Owner() (common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.Owner(&_ERC8004ReputationRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) Owner() (common.Address, error) {
	return _ERC8004ReputationRegistry.Contract.Owner(&_ERC8004ReputationRegistry.CallOpts)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x68bf3dc2.
//
// Solidity: function readAllFeedback(uint256 agentId, address client) view returns((int128,uint8,string,string,string,uint64,bool)[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) ReadAllFeedback(opts *bind.CallOpts, agentId *big.Int, client common.Address) ([]ERC8004ReputationRegistryFeedbackEntry, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "readAllFeedback", agentId, client)

	if err != nil {
		return *new([]ERC8004ReputationRegistryFeedbackEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC8004ReputationRegistryFeedbackEntry)).(*[]ERC8004ReputationRegistryFeedbackEntry)

	return out0, err

}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x68bf3dc2.
//
// Solidity: function readAllFeedback(uint256 agentId, address client) view returns((int128,uint8,string,string,string,uint64,bool)[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) ReadAllFeedback(agentId *big.Int, client common.Address) ([]ERC8004ReputationRegistryFeedbackEntry, error) {
	return _ERC8004ReputationRegistry.Contract.ReadAllFeedback(&_ERC8004ReputationRegistry.CallOpts, agentId, client)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x68bf3dc2.
//
// Solidity: function readAllFeedback(uint256 agentId, address client) view returns((int128,uint8,string,string,string,uint64,bool)[])
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) ReadAllFeedback(agentId *big.Int, client common.Address) ([]ERC8004ReputationRegistryFeedbackEntry, error) {
	return _ERC8004ReputationRegistry.Contract.ReadAllFeedback(&_ERC8004ReputationRegistry.CallOpts, agentId, client)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address client, uint64 index) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, string feedbackURI, bool isRevoked)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCaller) ReadFeedback(opts *bind.CallOpts, agentId *big.Int, client common.Address, index uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	FeedbackURI   string
	IsRevoked     bool
}, error) {
	var out []interface{}
	err := _ERC8004ReputationRegistry.contract.Call(opts, &out, "readFeedback", agentId, client, index)

	outstruct := new(struct {
		Value         *big.Int
		ValueDecimals uint8
		Tag1          string
		Tag2          string
		FeedbackURI   string
		IsRevoked     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueDecimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Tag1 = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Tag2 = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.FeedbackURI = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IsRevoked = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address client, uint64 index) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, string feedbackURI, bool isRevoked)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) ReadFeedback(agentId *big.Int, client common.Address, index uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	FeedbackURI   string
	IsRevoked     bool
}, error) {
	return _ERC8004ReputationRegistry.Contract.ReadFeedback(&_ERC8004ReputationRegistry.CallOpts, agentId, client, index)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address client, uint64 index) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, string feedbackURI, bool isRevoked)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryCallerSession) ReadFeedback(agentId *big.Int, client common.Address, index uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	FeedbackURI   string
	IsRevoked     bool
}, error) {
	return _ERC8004ReputationRegistry.Contract.ReadFeedback(&_ERC8004ReputationRegistry.CallOpts, agentId, client, index)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string , string feedbackURI, bytes32 ) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactor) GiveFeedback(opts *bind.TransactOpts, agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, arg5 string, feedbackURI string, arg7 [32]byte) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.contract.Transact(opts, "giveFeedback", agentId, value, valueDecimals, tag1, tag2, arg5, feedbackURI, arg7)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string , string feedbackURI, bytes32 ) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) GiveFeedback(agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, arg5 string, feedbackURI string, arg7 [32]byte) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.GiveFeedback(&_ERC8004ReputationRegistry.TransactOpts, agentId, value, valueDecimals, tag1, tag2, arg5, feedbackURI, arg7)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string , string feedbackURI, bytes32 ) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorSession) GiveFeedback(agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, arg5 string, feedbackURI string, arg7 [32]byte) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.GiveFeedback(&_ERC8004ReputationRegistry.TransactOpts, agentId, value, valueDecimals, tag1, tag2, arg5, feedbackURI, arg7)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.RenounceOwnership(&_ERC8004ReputationRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.RenounceOwnership(&_ERC8004ReputationRegistry.TransactOpts)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactor) RevokeFeedback(opts *bind.TransactOpts, agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.contract.Transact(opts, "revokeFeedback", agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.RevokeFeedback(&_ERC8004ReputationRegistry.TransactOpts, agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorSession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.RevokeFeedback(&_ERC8004ReputationRegistry.TransactOpts, agentId, feedbackIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.TransferOwnership(&_ERC8004ReputationRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004ReputationRegistry.Contract.TransferOwnership(&_ERC8004ReputationRegistry.TransactOpts, newOwner)
}

// ERC8004ReputationRegistryFeedbackGivenIterator is returned from FilterFeedbackGiven and is used to iterate over the raw logs and unpacked data for FeedbackGiven events raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryFeedbackGivenIterator struct {
	Event *ERC8004ReputationRegistryFeedbackGiven // Event containing the contract specifics and raw log

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
func (it *ERC8004ReputationRegistryFeedbackGivenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004ReputationRegistryFeedbackGiven)
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
		it.Event = new(ERC8004ReputationRegistryFeedbackGiven)
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
func (it *ERC8004ReputationRegistryFeedbackGivenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004ReputationRegistryFeedbackGivenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004ReputationRegistryFeedbackGiven represents a FeedbackGiven event raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryFeedbackGiven struct {
	AgentId       *big.Int
	Client        common.Address
	Index         uint64
	Value         *big.Int
	ValueDecimals uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeedbackGiven is a free log retrieval operation binding the contract event 0xbd84219cf4c28a7de200f887a0fb6d4d8f7e09a3833463acb9f442d069f8c0f4.
//
// Solidity: event FeedbackGiven(uint256 indexed agentId, address indexed client, uint64 index, int128 value, uint8 valueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) FilterFeedbackGiven(opts *bind.FilterOpts, agentId []*big.Int, client []common.Address) (*ERC8004ReputationRegistryFeedbackGivenIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.FilterLogs(opts, "FeedbackGiven", agentIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryFeedbackGivenIterator{contract: _ERC8004ReputationRegistry.contract, event: "FeedbackGiven", logs: logs, sub: sub}, nil
}

// WatchFeedbackGiven is a free log subscription operation binding the contract event 0xbd84219cf4c28a7de200f887a0fb6d4d8f7e09a3833463acb9f442d069f8c0f4.
//
// Solidity: event FeedbackGiven(uint256 indexed agentId, address indexed client, uint64 index, int128 value, uint8 valueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) WatchFeedbackGiven(opts *bind.WatchOpts, sink chan<- *ERC8004ReputationRegistryFeedbackGiven, agentId []*big.Int, client []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.WatchLogs(opts, "FeedbackGiven", agentIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004ReputationRegistryFeedbackGiven)
				if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "FeedbackGiven", log); err != nil {
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

// ParseFeedbackGiven is a log parse operation binding the contract event 0xbd84219cf4c28a7de200f887a0fb6d4d8f7e09a3833463acb9f442d069f8c0f4.
//
// Solidity: event FeedbackGiven(uint256 indexed agentId, address indexed client, uint64 index, int128 value, uint8 valueDecimals)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) ParseFeedbackGiven(log types.Log) (*ERC8004ReputationRegistryFeedbackGiven, error) {
	event := new(ERC8004ReputationRegistryFeedbackGiven)
	if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "FeedbackGiven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004ReputationRegistryFeedbackRevokedIterator is returned from FilterFeedbackRevoked and is used to iterate over the raw logs and unpacked data for FeedbackRevoked events raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryFeedbackRevokedIterator struct {
	Event *ERC8004ReputationRegistryFeedbackRevoked // Event containing the contract specifics and raw log

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
func (it *ERC8004ReputationRegistryFeedbackRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004ReputationRegistryFeedbackRevoked)
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
		it.Event = new(ERC8004ReputationRegistryFeedbackRevoked)
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
func (it *ERC8004ReputationRegistryFeedbackRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004ReputationRegistryFeedbackRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004ReputationRegistryFeedbackRevoked represents a FeedbackRevoked event raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryFeedbackRevoked struct {
	AgentId *big.Int
	Client  common.Address
	Index   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeedbackRevoked is a free log retrieval operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed client, uint64 index)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) FilterFeedbackRevoked(opts *bind.FilterOpts, agentId []*big.Int, client []common.Address) (*ERC8004ReputationRegistryFeedbackRevokedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.FilterLogs(opts, "FeedbackRevoked", agentIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryFeedbackRevokedIterator{contract: _ERC8004ReputationRegistry.contract, event: "FeedbackRevoked", logs: logs, sub: sub}, nil
}

// WatchFeedbackRevoked is a free log subscription operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed client, uint64 index)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) WatchFeedbackRevoked(opts *bind.WatchOpts, sink chan<- *ERC8004ReputationRegistryFeedbackRevoked, agentId []*big.Int, client []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.WatchLogs(opts, "FeedbackRevoked", agentIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004ReputationRegistryFeedbackRevoked)
				if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
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

// ParseFeedbackRevoked is a log parse operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed client, uint64 index)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) ParseFeedbackRevoked(log types.Log) (*ERC8004ReputationRegistryFeedbackRevoked, error) {
	event := new(ERC8004ReputationRegistryFeedbackRevoked)
	if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004ReputationRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryOwnershipTransferredIterator struct {
	Event *ERC8004ReputationRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC8004ReputationRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004ReputationRegistryOwnershipTransferred)
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
		it.Event = new(ERC8004ReputationRegistryOwnershipTransferred)
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
func (it *ERC8004ReputationRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004ReputationRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004ReputationRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ERC8004ReputationRegistry contract.
type ERC8004ReputationRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC8004ReputationRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004ReputationRegistryOwnershipTransferredIterator{contract: _ERC8004ReputationRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC8004ReputationRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC8004ReputationRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004ReputationRegistryOwnershipTransferred)
				if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC8004ReputationRegistry *ERC8004ReputationRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ERC8004ReputationRegistryOwnershipTransferred, error) {
	event := new(ERC8004ReputationRegistryOwnershipTransferred)
	if err := _ERC8004ReputationRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
