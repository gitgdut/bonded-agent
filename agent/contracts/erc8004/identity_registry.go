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

// ERC8004IdentityRegistryMetaData contains all meta data concerning the ERC8004IdentityRegistry contract.
var ERC8004IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAgentWallet\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadata\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextAgentId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"agentURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"agentWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"agentURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAgentURI\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAgentWallet\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadata\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unsetAgentWallet\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AgentRegistered\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agentURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentWalletSet\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentWalletUnset\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdated\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ERC8004IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC8004IdentityRegistryMetaData.ABI instead.
var ERC8004IdentityRegistryABI = ERC8004IdentityRegistryMetaData.ABI

// ERC8004IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type ERC8004IdentityRegistry struct {
	ERC8004IdentityRegistryCaller     // Read-only binding to the contract
	ERC8004IdentityRegistryTransactor // Write-only binding to the contract
	ERC8004IdentityRegistryFilterer   // Log filterer for contract events
}

// ERC8004IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC8004IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC8004IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC8004IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8004IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC8004IdentityRegistrySession struct {
	Contract     *ERC8004IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC8004IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC8004IdentityRegistryCallerSession struct {
	Contract *ERC8004IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC8004IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC8004IdentityRegistryTransactorSession struct {
	Contract     *ERC8004IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC8004IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC8004IdentityRegistryRaw struct {
	Contract *ERC8004IdentityRegistry // Generic contract binding to access the raw methods on
}

// ERC8004IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC8004IdentityRegistryCallerRaw struct {
	Contract *ERC8004IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC8004IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC8004IdentityRegistryTransactorRaw struct {
	Contract *ERC8004IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC8004IdentityRegistry creates a new instance of ERC8004IdentityRegistry, bound to a specific deployed contract.
func NewERC8004IdentityRegistry(address common.Address, backend bind.ContractBackend) (*ERC8004IdentityRegistry, error) {
	contract, err := bindERC8004IdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistry{ERC8004IdentityRegistryCaller: ERC8004IdentityRegistryCaller{contract: contract}, ERC8004IdentityRegistryTransactor: ERC8004IdentityRegistryTransactor{contract: contract}, ERC8004IdentityRegistryFilterer: ERC8004IdentityRegistryFilterer{contract: contract}}, nil
}

// NewERC8004IdentityRegistryCaller creates a new read-only instance of ERC8004IdentityRegistry, bound to a specific deployed contract.
func NewERC8004IdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*ERC8004IdentityRegistryCaller, error) {
	contract, err := bindERC8004IdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryCaller{contract: contract}, nil
}

// NewERC8004IdentityRegistryTransactor creates a new write-only instance of ERC8004IdentityRegistry, bound to a specific deployed contract.
func NewERC8004IdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC8004IdentityRegistryTransactor, error) {
	contract, err := bindERC8004IdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryTransactor{contract: contract}, nil
}

// NewERC8004IdentityRegistryFilterer creates a new log filterer instance of ERC8004IdentityRegistry, bound to a specific deployed contract.
func NewERC8004IdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC8004IdentityRegistryFilterer, error) {
	contract, err := bindERC8004IdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryFilterer{contract: contract}, nil
}

// bindERC8004IdentityRegistry binds a generic wrapper to an already deployed contract.
func bindERC8004IdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC8004IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8004IdentityRegistry.Contract.ERC8004IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.ERC8004IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.ERC8004IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8004IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC8004IdentityRegistry.Contract.BalanceOf(&_ERC8004IdentityRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC8004IdentityRegistry.Contract.BalanceOf(&_ERC8004IdentityRegistry.CallOpts, owner)
}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) GetAgentWallet(opts *bind.CallOpts, agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "getAgentWallet", agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) GetAgentWallet(agentId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.GetAgentWallet(&_ERC8004IdentityRegistry.CallOpts, agentId)
}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) GetAgentWallet(agentId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.GetAgentWallet(&_ERC8004IdentityRegistry.CallOpts, agentId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.GetApproved(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.GetApproved(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) GetMetadata(opts *bind.CallOpts, agentId *big.Int, key string) ([]byte, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "getMetadata", agentId, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) GetMetadata(agentId *big.Int, key string) ([]byte, error) {
	return _ERC8004IdentityRegistry.Contract.GetMetadata(&_ERC8004IdentityRegistry.CallOpts, agentId, key)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) GetMetadata(agentId *big.Int, key string) ([]byte, error) {
	return _ERC8004IdentityRegistry.Contract.GetMetadata(&_ERC8004IdentityRegistry.CallOpts, agentId, key)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC8004IdentityRegistry.Contract.IsApprovedForAll(&_ERC8004IdentityRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC8004IdentityRegistry.Contract.IsApprovedForAll(&_ERC8004IdentityRegistry.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Name() (string, error) {
	return _ERC8004IdentityRegistry.Contract.Name(&_ERC8004IdentityRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) Name() (string, error) {
	return _ERC8004IdentityRegistry.Contract.Name(&_ERC8004IdentityRegistry.CallOpts)
}

// NextAgentId is a free data retrieval call binding the contract method 0x30efc498.
//
// Solidity: function nextAgentId() view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) NextAgentId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "nextAgentId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAgentId is a free data retrieval call binding the contract method 0x30efc498.
//
// Solidity: function nextAgentId() view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) NextAgentId() (*big.Int, error) {
	return _ERC8004IdentityRegistry.Contract.NextAgentId(&_ERC8004IdentityRegistry.CallOpts)
}

// NextAgentId is a free data retrieval call binding the contract method 0x30efc498.
//
// Solidity: function nextAgentId() view returns(uint256)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) NextAgentId() (*big.Int, error) {
	return _ERC8004IdentityRegistry.Contract.NextAgentId(&_ERC8004IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Owner() (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.Owner(&_ERC8004IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.Owner(&_ERC8004IdentityRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.OwnerOf(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC8004IdentityRegistry.Contract.OwnerOf(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC8004IdentityRegistry.Contract.SupportsInterface(&_ERC8004IdentityRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC8004IdentityRegistry.Contract.SupportsInterface(&_ERC8004IdentityRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Symbol() (string, error) {
	return _ERC8004IdentityRegistry.Contract.Symbol(&_ERC8004IdentityRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) Symbol() (string, error) {
	return _ERC8004IdentityRegistry.Contract.Symbol(&_ERC8004IdentityRegistry.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC8004IdentityRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC8004IdentityRegistry.Contract.TokenURI(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC8004IdentityRegistry.Contract.TokenURI(&_ERC8004IdentityRegistry.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Approve(&_ERC8004IdentityRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Approve(&_ERC8004IdentityRegistry.TransactOpts, to, tokenId)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string agentURI, address agentWallet) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) Register(opts *bind.TransactOpts, agentURI string, agentWallet common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "register", agentURI, agentWallet)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string agentURI, address agentWallet) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Register(agentURI string, agentWallet common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Register(&_ERC8004IdentityRegistry.TransactOpts, agentURI, agentWallet)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string agentURI, address agentWallet) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) Register(agentURI string, agentWallet common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Register(&_ERC8004IdentityRegistry.TransactOpts, agentURI, agentWallet)
}

// Register0 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) Register0(opts *bind.TransactOpts, agentURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "register0", agentURI)
}

// Register0 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) Register0(agentURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Register0(&_ERC8004IdentityRegistry.TransactOpts, agentURI)
}

// Register0 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) Register0(agentURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.Register0(&_ERC8004IdentityRegistry.TransactOpts, agentURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.RenounceOwnership(&_ERC8004IdentityRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.RenounceOwnership(&_ERC8004IdentityRegistry.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SafeTransferFrom(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SafeTransferFrom(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SafeTransferFrom0(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SafeTransferFrom0(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId, data)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SetAgentURI(opts *bind.TransactOpts, agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "setAgentURI", agentId, newURI)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SetAgentURI(agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetAgentURI(&_ERC8004IdentityRegistry.TransactOpts, agentId, newURI)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SetAgentURI(agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetAgentURI(&_ERC8004IdentityRegistry.TransactOpts, agentId, newURI)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SetAgentWallet(opts *bind.TransactOpts, agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "setAgentWallet", agentId, newWallet, deadline, signature)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SetAgentWallet(agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetAgentWallet(&_ERC8004IdentityRegistry.TransactOpts, agentId, newWallet, deadline, signature)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SetAgentWallet(agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetAgentWallet(&_ERC8004IdentityRegistry.TransactOpts, agentId, newWallet, deadline, signature)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetApprovalForAll(&_ERC8004IdentityRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetApprovalForAll(&_ERC8004IdentityRegistry.TransactOpts, operator, approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) SetMetadata(opts *bind.TransactOpts, agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "setMetadata", agentId, key, value)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) SetMetadata(agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetMetadata(&_ERC8004IdentityRegistry.TransactOpts, agentId, key, value)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) SetMetadata(agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.SetMetadata(&_ERC8004IdentityRegistry.TransactOpts, agentId, key, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.TransferFrom(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.TransferFrom(&_ERC8004IdentityRegistry.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.TransferOwnership(&_ERC8004IdentityRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.TransferOwnership(&_ERC8004IdentityRegistry.TransactOpts, newOwner)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactor) UnsetAgentWallet(opts *bind.TransactOpts, agentId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.contract.Transact(opts, "unsetAgentWallet", agentId)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistrySession) UnsetAgentWallet(agentId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.UnsetAgentWallet(&_ERC8004IdentityRegistry.TransactOpts, agentId)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryTransactorSession) UnsetAgentWallet(agentId *big.Int) (*types.Transaction, error) {
	return _ERC8004IdentityRegistry.Contract.UnsetAgentWallet(&_ERC8004IdentityRegistry.TransactOpts, agentId)
}

// ERC8004IdentityRegistryAgentRegisteredIterator is returned from FilterAgentRegistered and is used to iterate over the raw logs and unpacked data for AgentRegistered events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentRegisteredIterator struct {
	Event *ERC8004IdentityRegistryAgentRegistered // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryAgentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryAgentRegistered)
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
		it.Event = new(ERC8004IdentityRegistryAgentRegistered)
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
func (it *ERC8004IdentityRegistryAgentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryAgentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryAgentRegistered represents a AgentRegistered event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentRegistered struct {
	AgentId  *big.Int
	Owner    common.Address
	AgentURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentRegistered is a free log retrieval operation binding the contract event 0x0d063c6022bff16d09991a9f91882ffa112f5fb2529136f65eb4c77bbd047e43.
//
// Solidity: event AgentRegistered(uint256 indexed agentId, address indexed owner, string agentURI)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterAgentRegistered(opts *bind.FilterOpts, agentId []*big.Int, owner []common.Address) (*ERC8004IdentityRegistryAgentRegisteredIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "AgentRegistered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryAgentRegisteredIterator{contract: _ERC8004IdentityRegistry.contract, event: "AgentRegistered", logs: logs, sub: sub}, nil
}

// WatchAgentRegistered is a free log subscription operation binding the contract event 0x0d063c6022bff16d09991a9f91882ffa112f5fb2529136f65eb4c77bbd047e43.
//
// Solidity: event AgentRegistered(uint256 indexed agentId, address indexed owner, string agentURI)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchAgentRegistered(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryAgentRegistered, agentId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "AgentRegistered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryAgentRegistered)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
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

// ParseAgentRegistered is a log parse operation binding the contract event 0x0d063c6022bff16d09991a9f91882ffa112f5fb2529136f65eb4c77bbd047e43.
//
// Solidity: event AgentRegistered(uint256 indexed agentId, address indexed owner, string agentURI)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseAgentRegistered(log types.Log) (*ERC8004IdentityRegistryAgentRegistered, error) {
	event := new(ERC8004IdentityRegistryAgentRegistered)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryAgentWalletSetIterator is returned from FilterAgentWalletSet and is used to iterate over the raw logs and unpacked data for AgentWalletSet events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentWalletSetIterator struct {
	Event *ERC8004IdentityRegistryAgentWalletSet // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryAgentWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryAgentWalletSet)
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
		it.Event = new(ERC8004IdentityRegistryAgentWalletSet)
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
func (it *ERC8004IdentityRegistryAgentWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryAgentWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryAgentWalletSet represents a AgentWalletSet event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentWalletSet struct {
	AgentId   *big.Int
	NewWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentWalletSet is a free log retrieval operation binding the contract event 0xc1b62bacf2a2f4363c2190e00f2370f3b563001a0c56c73eae69619e58e9624c.
//
// Solidity: event AgentWalletSet(uint256 indexed agentId, address indexed newWallet)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterAgentWalletSet(opts *bind.FilterOpts, agentId []*big.Int, newWallet []common.Address) (*ERC8004IdentityRegistryAgentWalletSetIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "AgentWalletSet", agentIdRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryAgentWalletSetIterator{contract: _ERC8004IdentityRegistry.contract, event: "AgentWalletSet", logs: logs, sub: sub}, nil
}

// WatchAgentWalletSet is a free log subscription operation binding the contract event 0xc1b62bacf2a2f4363c2190e00f2370f3b563001a0c56c73eae69619e58e9624c.
//
// Solidity: event AgentWalletSet(uint256 indexed agentId, address indexed newWallet)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchAgentWalletSet(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryAgentWalletSet, agentId []*big.Int, newWallet []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "AgentWalletSet", agentIdRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryAgentWalletSet)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentWalletSet", log); err != nil {
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

// ParseAgentWalletSet is a log parse operation binding the contract event 0xc1b62bacf2a2f4363c2190e00f2370f3b563001a0c56c73eae69619e58e9624c.
//
// Solidity: event AgentWalletSet(uint256 indexed agentId, address indexed newWallet)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseAgentWalletSet(log types.Log) (*ERC8004IdentityRegistryAgentWalletSet, error) {
	event := new(ERC8004IdentityRegistryAgentWalletSet)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentWalletSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryAgentWalletUnsetIterator is returned from FilterAgentWalletUnset and is used to iterate over the raw logs and unpacked data for AgentWalletUnset events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentWalletUnsetIterator struct {
	Event *ERC8004IdentityRegistryAgentWalletUnset // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryAgentWalletUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryAgentWalletUnset)
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
		it.Event = new(ERC8004IdentityRegistryAgentWalletUnset)
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
func (it *ERC8004IdentityRegistryAgentWalletUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryAgentWalletUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryAgentWalletUnset represents a AgentWalletUnset event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryAgentWalletUnset struct {
	AgentId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentWalletUnset is a free log retrieval operation binding the contract event 0x71431e78d89cfcd109b26824c9284976654c91c0a617312646c2febbe0c4352a.
//
// Solidity: event AgentWalletUnset(uint256 indexed agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterAgentWalletUnset(opts *bind.FilterOpts, agentId []*big.Int) (*ERC8004IdentityRegistryAgentWalletUnsetIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "AgentWalletUnset", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryAgentWalletUnsetIterator{contract: _ERC8004IdentityRegistry.contract, event: "AgentWalletUnset", logs: logs, sub: sub}, nil
}

// WatchAgentWalletUnset is a free log subscription operation binding the contract event 0x71431e78d89cfcd109b26824c9284976654c91c0a617312646c2febbe0c4352a.
//
// Solidity: event AgentWalletUnset(uint256 indexed agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchAgentWalletUnset(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryAgentWalletUnset, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "AgentWalletUnset", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryAgentWalletUnset)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentWalletUnset", log); err != nil {
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

// ParseAgentWalletUnset is a log parse operation binding the contract event 0x71431e78d89cfcd109b26824c9284976654c91c0a617312646c2febbe0c4352a.
//
// Solidity: event AgentWalletUnset(uint256 indexed agentId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseAgentWalletUnset(log types.Log) (*ERC8004IdentityRegistryAgentWalletUnset, error) {
	event := new(ERC8004IdentityRegistryAgentWalletUnset)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "AgentWalletUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryApprovalIterator struct {
	Event *ERC8004IdentityRegistryApproval // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryApproval)
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
		it.Event = new(ERC8004IdentityRegistryApproval)
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
func (it *ERC8004IdentityRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryApproval represents a Approval event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC8004IdentityRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryApprovalIterator{contract: _ERC8004IdentityRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryApproval)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseApproval(log types.Log) (*ERC8004IdentityRegistryApproval, error) {
	event := new(ERC8004IdentityRegistryApproval)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryApprovalForAllIterator struct {
	Event *ERC8004IdentityRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryApprovalForAll)
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
		it.Event = new(ERC8004IdentityRegistryApprovalForAll)
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
func (it *ERC8004IdentityRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryApprovalForAll represents a ApprovalForAll event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC8004IdentityRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryApprovalForAllIterator{contract: _ERC8004IdentityRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryApprovalForAll)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseApprovalForAll(log types.Log) (*ERC8004IdentityRegistryApprovalForAll, error) {
	event := new(ERC8004IdentityRegistryApprovalForAll)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryBatchMetadataUpdateIterator struct {
	Event *ERC8004IdentityRegistryBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryBatchMetadataUpdate)
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
		it.Event = new(ERC8004IdentityRegistryBatchMetadataUpdate)
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
func (it *ERC8004IdentityRegistryBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ERC8004IdentityRegistryBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryBatchMetadataUpdateIterator{contract: _ERC8004IdentityRegistry.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryBatchMetadataUpdate)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseBatchMetadataUpdate(log types.Log) (*ERC8004IdentityRegistryBatchMetadataUpdate, error) {
	event := new(ERC8004IdentityRegistryBatchMetadataUpdate)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryMetadataUpdateIterator struct {
	Event *ERC8004IdentityRegistryMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryMetadataUpdate)
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
		it.Event = new(ERC8004IdentityRegistryMetadataUpdate)
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
func (it *ERC8004IdentityRegistryMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryMetadataUpdate represents a MetadataUpdate event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ERC8004IdentityRegistryMetadataUpdateIterator, error) {

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryMetadataUpdateIterator{contract: _ERC8004IdentityRegistry.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryMetadataUpdate)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseMetadataUpdate(log types.Log) (*ERC8004IdentityRegistryMetadataUpdate, error) {
	event := new(ERC8004IdentityRegistryMetadataUpdate)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryMetadataUpdatedIterator is returned from FilterMetadataUpdated and is used to iterate over the raw logs and unpacked data for MetadataUpdated events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryMetadataUpdatedIterator struct {
	Event *ERC8004IdentityRegistryMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryMetadataUpdated)
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
		it.Event = new(ERC8004IdentityRegistryMetadataUpdated)
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
func (it *ERC8004IdentityRegistryMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryMetadataUpdated represents a MetadataUpdated event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryMetadataUpdated struct {
	AgentId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdated is a free log retrieval operation binding the contract event 0x48690d4ac759f299a87d6da64344a884585af467bbcc36ac9b63fe4b7858a473.
//
// Solidity: event MetadataUpdated(uint256 indexed agentId, string key, bytes value)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterMetadataUpdated(opts *bind.FilterOpts, agentId []*big.Int) (*ERC8004IdentityRegistryMetadataUpdatedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "MetadataUpdated", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryMetadataUpdatedIterator{contract: _ERC8004IdentityRegistry.contract, event: "MetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdated is a free log subscription operation binding the contract event 0x48690d4ac759f299a87d6da64344a884585af467bbcc36ac9b63fe4b7858a473.
//
// Solidity: event MetadataUpdated(uint256 indexed agentId, string key, bytes value)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryMetadataUpdated, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "MetadataUpdated", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryMetadataUpdated)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
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

// ParseMetadataUpdated is a log parse operation binding the contract event 0x48690d4ac759f299a87d6da64344a884585af467bbcc36ac9b63fe4b7858a473.
//
// Solidity: event MetadataUpdated(uint256 indexed agentId, string key, bytes value)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseMetadataUpdated(log types.Log) (*ERC8004IdentityRegistryMetadataUpdated, error) {
	event := new(ERC8004IdentityRegistryMetadataUpdated)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryOwnershipTransferredIterator struct {
	Event *ERC8004IdentityRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryOwnershipTransferred)
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
		it.Event = new(ERC8004IdentityRegistryOwnershipTransferred)
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
func (it *ERC8004IdentityRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC8004IdentityRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryOwnershipTransferredIterator{contract: _ERC8004IdentityRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryOwnershipTransferred)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ERC8004IdentityRegistryOwnershipTransferred, error) {
	event := new(ERC8004IdentityRegistryOwnershipTransferred)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8004IdentityRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryTransferIterator struct {
	Event *ERC8004IdentityRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ERC8004IdentityRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8004IdentityRegistryTransfer)
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
		it.Event = new(ERC8004IdentityRegistryTransfer)
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
func (it *ERC8004IdentityRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8004IdentityRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8004IdentityRegistryTransfer represents a Transfer event raised by the ERC8004IdentityRegistry contract.
type ERC8004IdentityRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC8004IdentityRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC8004IdentityRegistryTransferIterator{contract: _ERC8004IdentityRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC8004IdentityRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC8004IdentityRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8004IdentityRegistryTransfer)
				if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC8004IdentityRegistry *ERC8004IdentityRegistryFilterer) ParseTransfer(log types.Log) (*ERC8004IdentityRegistryTransfer, error) {
	event := new(ERC8004IdentityRegistryTransfer)
	if err := _ERC8004IdentityRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
