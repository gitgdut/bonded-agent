// Minimal Go binding for SimpleAMMPair — only the methods the Agent needs.
// Pattern mirrors mock_dex.go but omits event filtering to keep it simple.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SimpleAMMPairMetaData contains the ABI for SimpleAMMPair.
var SimpleAMMPairMetaData = &bind.MetaData{
	ABI: "[" +
		"{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"payable\"}," +
		"{\"type\":\"receive\",\"stateMutability\":\"payable\"}," +
		"{\"type\":\"function\",\"name\":\"getAmountOut\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"}," +
		"{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}," +
		"{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"}," +
		"{\"type\":\"function\",\"name\":\"reserve0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"}," +
		"{\"type\":\"function\",\"name\":\"reserve1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"}," +
		"{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"minOutput\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\"}," +
		"{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}," +
		"{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"}," +
		"{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false}," +
		"{\"type\":\"event\",\"name\":\"Sync\",\"inputs\":[{\"name\":\"reserve0\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"reserve1\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false}," +
		"{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]}," +
		"{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]}," +
		"{\"type\":\"error\",\"name\":\"InsufficientOutput\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\"},{\"name\":\"minRequired\",\"type\":\"uint256\"}]}," +
		"{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]}" +
		"]",
}

// SimpleAMMPair wraps the deployed SimpleAMMPair contract.
type SimpleAMMPair struct {
	SimpleAMMPairCaller
	SimpleAMMPairTransactor
}

// SimpleAMMPairCaller provides read-only access.
type SimpleAMMPairCaller struct{ contract *bind.BoundContract }

// SimpleAMMPairTransactor provides write access.
type SimpleAMMPairTransactor struct{ contract *bind.BoundContract }

// NewSimpleAMMPair binds to a deployed SimpleAMMPair contract.
func NewSimpleAMMPair(address common.Address, backend bind.ContractBackend) (*SimpleAMMPair, error) {
	c, err := bindSimpleAMMPair(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAMMPair{
		SimpleAMMPairCaller:     SimpleAMMPairCaller{contract: c},
		SimpleAMMPairTransactor: SimpleAMMPairTransactor{contract: c},
	}, nil
}

func bindSimpleAMMPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := SimpleAMMPairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, nil), nil
}

// ── Caller methods ──────────────────────────────────────────

// GetAmountOut computes expected output: (amountIn * 997 * reserve1) / (reserve0 * 1000 + amountIn * 997)
func (c *SimpleAMMPairCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getAmountOut", amountIn)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// Token returns the tUSDC token address.
func (c *SimpleAMMPairCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "token")
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// Reserve0 returns the MON reserve (contract's MON balance).
func (c *SimpleAMMPairCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "reserve0")
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// Initialized returns whether the pool has been initialized.
func (c *SimpleAMMPairCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "initialized")
	if err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

// ── Transactor methods ──────────────────────────────────────

// Swap calls swap(uint256 minOutput). Selector 0x94b918de — same as MockDex.
func (t *SimpleAMMPairTransactor) Swap(opts *bind.TransactOpts, minOutput *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "swap", minOutput)
}

// Initialize records initial reserves. Must be called after MON + tUSDC are sent.
func (t *SimpleAMMPairTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "initialize")
}

// Sync updates reserves to match actual balances.
func (t *SimpleAMMPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "sync")
}

// Transfer sends MON to the contract.
func (t *SimpleAMMPairTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transfer(opts)
}

// Ensure the abigen-style call pattern is registered
var _ = strings.NewReader
var _ = abi.ConvertType
