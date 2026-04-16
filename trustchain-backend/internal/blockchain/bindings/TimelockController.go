// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TimelockControllerMetaData contains all meta data concerning the TimelockController contract.
var TimelockControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TimelockControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use TimelockControllerMetaData.ABI instead.
var TimelockControllerABI = TimelockControllerMetaData.ABI

// TimelockController is an auto generated Go binding around an Ethereum contract.
type TimelockController struct {
	TimelockControllerCaller     // Read-only binding to the contract
	TimelockControllerTransactor // Write-only binding to the contract
	TimelockControllerFilterer   // Log filterer for contract events
}

// TimelockControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimelockControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimelockControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimelockControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimelockControllerSession struct {
	Contract     *TimelockController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimelockControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimelockControllerCallerSession struct {
	Contract *TimelockControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TimelockControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimelockControllerTransactorSession struct {
	Contract     *TimelockControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TimelockControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimelockControllerRaw struct {
	Contract *TimelockController // Generic contract binding to access the raw methods on
}

// TimelockControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimelockControllerCallerRaw struct {
	Contract *TimelockControllerCaller // Generic read-only contract binding to access the raw methods on
}

// TimelockControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimelockControllerTransactorRaw struct {
	Contract *TimelockControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimelockController creates a new instance of TimelockController, bound to a specific deployed contract.
func NewTimelockController(address common.Address, backend bind.ContractBackend) (*TimelockController, error) {
	contract, err := bindTimelockController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimelockController{TimelockControllerCaller: TimelockControllerCaller{contract: contract}, TimelockControllerTransactor: TimelockControllerTransactor{contract: contract}, TimelockControllerFilterer: TimelockControllerFilterer{contract: contract}}, nil
}

// NewTimelockControllerCaller creates a new read-only instance of TimelockController, bound to a specific deployed contract.
func NewTimelockControllerCaller(address common.Address, caller bind.ContractCaller) (*TimelockControllerCaller, error) {
	contract, err := bindTimelockController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerCaller{contract: contract}, nil
}

// NewTimelockControllerTransactor creates a new write-only instance of TimelockController, bound to a specific deployed contract.
func NewTimelockControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*TimelockControllerTransactor, error) {
	contract, err := bindTimelockController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerTransactor{contract: contract}, nil
}

// NewTimelockControllerFilterer creates a new log filterer instance of TimelockController, bound to a specific deployed contract.
func NewTimelockControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*TimelockControllerFilterer, error) {
	contract, err := bindTimelockController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerFilterer{contract: contract}, nil
}

// bindTimelockController binds a generic wrapper to an already deployed contract.
func bindTimelockController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimelockControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimelockController *TimelockControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimelockController.Contract.TimelockControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimelockController *TimelockControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimelockController.Contract.TimelockControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimelockController *TimelockControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimelockController.Contract.TimelockControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimelockController *TimelockControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimelockController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimelockController *TimelockControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimelockController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimelockController *TimelockControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimelockController.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockController *TimelockControllerTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockController.contract.Transact(opts, "execute", target, value, data, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockController *TimelockControllerSession) Execute(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockController.Contract.Execute(&_TimelockController.TransactOpts, target, value, data, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockController *TimelockControllerTransactorSession) Execute(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockController.Contract.Execute(&_TimelockController.TransactOpts, target, value, data, predecessor, salt)
}
