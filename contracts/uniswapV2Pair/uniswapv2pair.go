// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2pair

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

// Uniswapv2pairMetaData contains all meta data concerning the Uniswapv2pair contract.
var Uniswapv2pairMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv2pairABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv2pairMetaData.ABI instead.
var Uniswapv2pairABI = Uniswapv2pairMetaData.ABI

// Uniswapv2pair is an auto generated Go binding around an Ethereum contract.
type Uniswapv2pair struct {
	Uniswapv2pairCaller     // Read-only binding to the contract
	Uniswapv2pairTransactor // Write-only binding to the contract
	Uniswapv2pairFilterer   // Log filterer for contract events
}

// Uniswapv2pairCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2pairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2pairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2pairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2pairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2pairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2pairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2pairSession struct {
	Contract     *Uniswapv2pair    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2pairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2pairCallerSession struct {
	Contract *Uniswapv2pairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Uniswapv2pairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2pairTransactorSession struct {
	Contract     *Uniswapv2pairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Uniswapv2pairRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2pairRaw struct {
	Contract *Uniswapv2pair // Generic contract binding to access the raw methods on
}

// Uniswapv2pairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2pairCallerRaw struct {
	Contract *Uniswapv2pairCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2pairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2pairTransactorRaw struct {
	Contract *Uniswapv2pairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2pair creates a new instance of Uniswapv2pair, bound to a specific deployed contract.
func NewUniswapv2pair(address common.Address, backend bind.ContractBackend) (*Uniswapv2pair, error) {
	contract, err := bindUniswapv2pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2pair{Uniswapv2pairCaller: Uniswapv2pairCaller{contract: contract}, Uniswapv2pairTransactor: Uniswapv2pairTransactor{contract: contract}, Uniswapv2pairFilterer: Uniswapv2pairFilterer{contract: contract}}, nil
}

// NewUniswapv2pairCaller creates a new read-only instance of Uniswapv2pair, bound to a specific deployed contract.
func NewUniswapv2pairCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2pairCaller, error) {
	contract, err := bindUniswapv2pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2pairCaller{contract: contract}, nil
}

// NewUniswapv2pairTransactor creates a new write-only instance of Uniswapv2pair, bound to a specific deployed contract.
func NewUniswapv2pairTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2pairTransactor, error) {
	contract, err := bindUniswapv2pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2pairTransactor{contract: contract}, nil
}

// NewUniswapv2pairFilterer creates a new log filterer instance of Uniswapv2pair, bound to a specific deployed contract.
func NewUniswapv2pairFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2pairFilterer, error) {
	contract, err := bindUniswapv2pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2pairFilterer{contract: contract}, nil
}

// bindUniswapv2pair binds a generic wrapper to an already deployed contract.
func bindUniswapv2pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv2pairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2pair *Uniswapv2pairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2pair.Contract.Uniswapv2pairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2pair *Uniswapv2pairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Uniswapv2pairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2pair *Uniswapv2pairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Uniswapv2pairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2pair *Uniswapv2pairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2pair *Uniswapv2pairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2pair *Uniswapv2pairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairSession) Factory() (common.Address, error) {
	return _Uniswapv2pair.Contract.Factory(&_Uniswapv2pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) Factory() (common.Address, error) {
	return _Uniswapv2pair.Contract.Factory(&_Uniswapv2pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Uniswapv2pair *Uniswapv2pairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Uniswapv2pair *Uniswapv2pairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Uniswapv2pair.Contract.GetReserves(&_Uniswapv2pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Uniswapv2pair.Contract.GetReserves(&_Uniswapv2pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairSession) KLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.KLast(&_Uniswapv2pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) KLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.KLast(&_Uniswapv2pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairSession) Price0CumulativeLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.Price0CumulativeLast(&_Uniswapv2pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.Price0CumulativeLast(&_Uniswapv2pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairSession) Price1CumulativeLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.Price1CumulativeLast(&_Uniswapv2pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Uniswapv2pair.Contract.Price1CumulativeLast(&_Uniswapv2pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairSession) Token0() (common.Address, error) {
	return _Uniswapv2pair.Contract.Token0(&_Uniswapv2pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) Token0() (common.Address, error) {
	return _Uniswapv2pair.Contract.Token0(&_Uniswapv2pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairSession) Token1() (common.Address, error) {
	return _Uniswapv2pair.Contract.Token1(&_Uniswapv2pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswapv2pair *Uniswapv2pairCallerSession) Token1() (common.Address, error) {
	return _Uniswapv2pair.Contract.Token1(&_Uniswapv2pair.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Uniswapv2pair *Uniswapv2pairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Uniswapv2pair *Uniswapv2pairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Burn(&_Uniswapv2pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Uniswapv2pair *Uniswapv2pairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Burn(&_Uniswapv2pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Uniswapv2pair *Uniswapv2pairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Uniswapv2pair *Uniswapv2pairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Mint(&_Uniswapv2pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Uniswapv2pair *Uniswapv2pairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Mint(&_Uniswapv2pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Uniswapv2pair *Uniswapv2pairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Uniswapv2pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Uniswapv2pair *Uniswapv2pairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Swap(&_Uniswapv2pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Uniswapv2pair *Uniswapv2pairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Swap(&_Uniswapv2pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Uniswapv2pair *Uniswapv2pairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Uniswapv2pair *Uniswapv2pairSession) Sync() (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Sync(&_Uniswapv2pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Uniswapv2pair *Uniswapv2pairTransactorSession) Sync() (*types.Transaction, error) {
	return _Uniswapv2pair.Contract.Sync(&_Uniswapv2pair.TransactOpts)
}
