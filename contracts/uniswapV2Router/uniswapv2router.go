// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2router

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

// Uniswapv2routerMetaData contains all meta data concerning the Uniswapv2router contract.
var Uniswapv2routerMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Uniswapv2routerABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv2routerMetaData.ABI instead.
var Uniswapv2routerABI = Uniswapv2routerMetaData.ABI

// Uniswapv2router is an auto generated Go binding around an Ethereum contract.
type Uniswapv2router struct {
	Uniswapv2routerCaller     // Read-only binding to the contract
	Uniswapv2routerTransactor // Write-only binding to the contract
	Uniswapv2routerFilterer   // Log filterer for contract events
}

// Uniswapv2routerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2routerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2routerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2routerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2routerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2routerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2routerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2routerSession struct {
	Contract     *Uniswapv2router  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2routerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2routerCallerSession struct {
	Contract *Uniswapv2routerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Uniswapv2routerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2routerTransactorSession struct {
	Contract     *Uniswapv2routerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Uniswapv2routerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2routerRaw struct {
	Contract *Uniswapv2router // Generic contract binding to access the raw methods on
}

// Uniswapv2routerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2routerCallerRaw struct {
	Contract *Uniswapv2routerCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2routerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2routerTransactorRaw struct {
	Contract *Uniswapv2routerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2router creates a new instance of Uniswapv2router, bound to a specific deployed contract.
func NewUniswapv2router(address common.Address, backend bind.ContractBackend) (*Uniswapv2router, error) {
	contract, err := bindUniswapv2router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2router{Uniswapv2routerCaller: Uniswapv2routerCaller{contract: contract}, Uniswapv2routerTransactor: Uniswapv2routerTransactor{contract: contract}, Uniswapv2routerFilterer: Uniswapv2routerFilterer{contract: contract}}, nil
}

// NewUniswapv2routerCaller creates a new read-only instance of Uniswapv2router, bound to a specific deployed contract.
func NewUniswapv2routerCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2routerCaller, error) {
	contract, err := bindUniswapv2router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2routerCaller{contract: contract}, nil
}

// NewUniswapv2routerTransactor creates a new write-only instance of Uniswapv2router, bound to a specific deployed contract.
func NewUniswapv2routerTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2routerTransactor, error) {
	contract, err := bindUniswapv2router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2routerTransactor{contract: contract}, nil
}

// NewUniswapv2routerFilterer creates a new log filterer instance of Uniswapv2router, bound to a specific deployed contract.
func NewUniswapv2routerFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2routerFilterer, error) {
	contract, err := bindUniswapv2router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2routerFilterer{contract: contract}, nil
}

// bindUniswapv2router binds a generic wrapper to an already deployed contract.
func bindUniswapv2router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv2routerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2router *Uniswapv2routerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2router.Contract.Uniswapv2routerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2router *Uniswapv2routerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.Uniswapv2routerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2router *Uniswapv2routerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.Uniswapv2routerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2router *Uniswapv2routerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2router *Uniswapv2routerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2router *Uniswapv2routerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.contract.Transact(opts, method, params...)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2router.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2router.Contract.GetAmountsOut(&_Uniswapv2router.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2router.Contract.GetAmountsOut(&_Uniswapv2router.CallOpts, amountIn, path)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.SwapExactTokensForTokens(&_Uniswapv2router.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.SwapExactTokensForTokens(&_Uniswapv2router.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.SwapTokensForExactTokens(&_Uniswapv2router.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2router *Uniswapv2routerTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2router.Contract.SwapTokensForExactTokens(&_Uniswapv2router.TransactOpts, amountOut, amountInMax, path, to, deadline)
}
