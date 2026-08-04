// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// MockDexMetaData contains all meta data concerning the MockDex contract.
var MockDexMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_tUSDC\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRate\",\"inputs\":[{\"name\":\"_newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"minOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tUSDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockUSDC\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RateUpdated\",\"inputs\":[{\"name\":\"oldRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swapped\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"monIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdcOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientOutput\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minRequired\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// MockDexABI is the input ABI used to generate the binding from.
// Deprecated: Use MockDexMetaData.ABI instead.
var MockDexABI = MockDexMetaData.ABI

// MockDex is an auto generated Go binding around an Ethereum contract.
type MockDex struct {
	MockDexCaller     // Read-only binding to the contract
	MockDexTransactor // Write-only binding to the contract
	MockDexFilterer   // Log filterer for contract events
}

// MockDexCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockDexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockDexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockDexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockDexSession struct {
	Contract     *MockDex          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockDexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockDexCallerSession struct {
	Contract *MockDexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MockDexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockDexTransactorSession struct {
	Contract     *MockDexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MockDexRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockDexRaw struct {
	Contract *MockDex // Generic contract binding to access the raw methods on
}

// MockDexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockDexCallerRaw struct {
	Contract *MockDexCaller // Generic read-only contract binding to access the raw methods on
}

// MockDexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockDexTransactorRaw struct {
	Contract *MockDexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockDex creates a new instance of MockDex, bound to a specific deployed contract.
func NewMockDex(address common.Address, backend bind.ContractBackend) (*MockDex, error) {
	contract, err := bindMockDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockDex{MockDexCaller: MockDexCaller{contract: contract}, MockDexTransactor: MockDexTransactor{contract: contract}, MockDexFilterer: MockDexFilterer{contract: contract}}, nil
}

// NewMockDexCaller creates a new read-only instance of MockDex, bound to a specific deployed contract.
func NewMockDexCaller(address common.Address, caller bind.ContractCaller) (*MockDexCaller, error) {
	contract, err := bindMockDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexCaller{contract: contract}, nil
}

// NewMockDexTransactor creates a new write-only instance of MockDex, bound to a specific deployed contract.
func NewMockDexTransactor(address common.Address, transactor bind.ContractTransactor) (*MockDexTransactor, error) {
	contract, err := bindMockDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexTransactor{contract: contract}, nil
}

// NewMockDexFilterer creates a new log filterer instance of MockDex, bound to a specific deployed contract.
func NewMockDexFilterer(address common.Address, filterer bind.ContractFilterer) (*MockDexFilterer, error) {
	contract, err := bindMockDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockDexFilterer{contract: contract}, nil
}

// bindMockDex binds a generic wrapper to an already deployed contract.
func bindMockDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockDexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDex *MockDexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDex.Contract.MockDexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDex *MockDexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDex.Contract.MockDexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDex *MockDexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDex.Contract.MockDexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDex *MockDexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDex *MockDexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDex *MockDexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDex.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockDex *MockDexCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockDex.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockDex *MockDexSession) Owner() (common.Address, error) {
	return _MockDex.Contract.Owner(&_MockDex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockDex *MockDexCallerSession) Owner() (common.Address, error) {
	return _MockDex.Contract.Owner(&_MockDex.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_MockDex *MockDexCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDex.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_MockDex *MockDexSession) Rate() (*big.Int, error) {
	return _MockDex.Contract.Rate(&_MockDex.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_MockDex *MockDexCallerSession) Rate() (*big.Int, error) {
	return _MockDex.Contract.Rate(&_MockDex.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_MockDex *MockDexCaller) TUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockDex.contract.Call(opts, &out, "tUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_MockDex *MockDexSession) TUSDC() (common.Address, error) {
	return _MockDex.Contract.TUSDC(&_MockDex.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_MockDex *MockDexCallerSession) TUSDC() (common.Address, error) {
	return _MockDex.Contract.TUSDC(&_MockDex.CallOpts)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MockDex *MockDexTransactor) SetRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _MockDex.contract.Transact(opts, "setRate", _newRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MockDex *MockDexSession) SetRate(_newRate *big.Int) (*types.Transaction, error) {
	return _MockDex.Contract.SetRate(&_MockDex.TransactOpts, _newRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _newRate) returns()
func (_MockDex *MockDexTransactorSession) SetRate(_newRate *big.Int) (*types.Transaction, error) {
	return _MockDex.Contract.SetRate(&_MockDex.TransactOpts, _newRate)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 minOutput) payable returns(uint256)
func (_MockDex *MockDexTransactor) Swap(opts *bind.TransactOpts, minOutput *big.Int) (*types.Transaction, error) {
	return _MockDex.contract.Transact(opts, "swap", minOutput)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 minOutput) payable returns(uint256)
func (_MockDex *MockDexSession) Swap(minOutput *big.Int) (*types.Transaction, error) {
	return _MockDex.Contract.Swap(&_MockDex.TransactOpts, minOutput)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 minOutput) payable returns(uint256)
func (_MockDex *MockDexTransactorSession) Swap(minOutput *big.Int) (*types.Transaction, error) {
	return _MockDex.Contract.Swap(&_MockDex.TransactOpts, minOutput)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockDex *MockDexTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDex.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockDex *MockDexSession) Receive() (*types.Transaction, error) {
	return _MockDex.Contract.Receive(&_MockDex.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockDex *MockDexTransactorSession) Receive() (*types.Transaction, error) {
	return _MockDex.Contract.Receive(&_MockDex.TransactOpts)
}

// MockDexRateUpdatedIterator is returned from FilterRateUpdated and is used to iterate over the raw logs and unpacked data for RateUpdated events raised by the MockDex contract.
type MockDexRateUpdatedIterator struct {
	Event *MockDexRateUpdated // Event containing the contract specifics and raw log

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
func (it *MockDexRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexRateUpdated)
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
		it.Event = new(MockDexRateUpdated)
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
func (it *MockDexRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexRateUpdated represents a RateUpdated event raised by the MockDex contract.
type MockDexRateUpdated struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateUpdated is a free log retrieval operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(uint256 oldRate, uint256 newRate)
func (_MockDex *MockDexFilterer) FilterRateUpdated(opts *bind.FilterOpts) (*MockDexRateUpdatedIterator, error) {

	logs, sub, err := _MockDex.contract.FilterLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return &MockDexRateUpdatedIterator{contract: _MockDex.contract, event: "RateUpdated", logs: logs, sub: sub}, nil
}

// WatchRateUpdated is a free log subscription operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(uint256 oldRate, uint256 newRate)
func (_MockDex *MockDexFilterer) WatchRateUpdated(opts *bind.WatchOpts, sink chan<- *MockDexRateUpdated) (event.Subscription, error) {

	logs, sub, err := _MockDex.contract.WatchLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexRateUpdated)
				if err := _MockDex.contract.UnpackLog(event, "RateUpdated", log); err != nil {
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

// ParseRateUpdated is a log parse operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(uint256 oldRate, uint256 newRate)
func (_MockDex *MockDexFilterer) ParseRateUpdated(log types.Log) (*MockDexRateUpdated, error) {
	event := new(MockDexRateUpdated)
	if err := _MockDex.contract.UnpackLog(event, "RateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockDexSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the MockDex contract.
type MockDexSwappedIterator struct {
	Event *MockDexSwapped // Event containing the contract specifics and raw log

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
func (it *MockDexSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexSwapped)
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
		it.Event = new(MockDexSwapped)
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
func (it *MockDexSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexSwapped represents a Swapped event raised by the MockDex contract.
type MockDexSwapped struct {
	User    common.Address
	MonIn   *big.Int
	UsdcOut *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x3a9a9f34f5831e9c8ecb66ab3aa308b2ff31eaca434615f6c9cadc656a9af71c.
//
// Solidity: event Swapped(address indexed user, uint256 monIn, uint256 usdcOut)
func (_MockDex *MockDexFilterer) FilterSwapped(opts *bind.FilterOpts, user []common.Address) (*MockDexSwappedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MockDex.contract.FilterLogs(opts, "Swapped", userRule)
	if err != nil {
		return nil, err
	}
	return &MockDexSwappedIterator{contract: _MockDex.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x3a9a9f34f5831e9c8ecb66ab3aa308b2ff31eaca434615f6c9cadc656a9af71c.
//
// Solidity: event Swapped(address indexed user, uint256 monIn, uint256 usdcOut)
func (_MockDex *MockDexFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *MockDexSwapped, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MockDex.contract.WatchLogs(opts, "Swapped", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexSwapped)
				if err := _MockDex.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x3a9a9f34f5831e9c8ecb66ab3aa308b2ff31eaca434615f6c9cadc656a9af71c.
//
// Solidity: event Swapped(address indexed user, uint256 monIn, uint256 usdcOut)
func (_MockDex *MockDexFilterer) ParseSwapped(log types.Log) (*MockDexSwapped, error) {
	event := new(MockDexSwapped)
	if err := _MockDex.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
