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

// CampaignFactoryMetaData contains all meta data concerning the CampaignFactory contract.
var CampaignFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminMultiSig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reputationRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_donationNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_timelockController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"campaignAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creatorWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"milestoneCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"}],\"name\":\"CampaignCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allCampaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creatorWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"goalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"milestoneDescriptions\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"milestoneAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"milestoneDeadlines\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"milestoneEvidence\",\"type\":\"string[]\"}],\"name\":\"createCampaign\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"campaignAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donationNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCampaignCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getCampaigns\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputationRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelockController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CampaignFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignFactoryMetaData.ABI instead.
var CampaignFactoryABI = CampaignFactoryMetaData.ABI

// CampaignFactory is an auto generated Go binding around an Ethereum contract.
type CampaignFactory struct {
	CampaignFactoryCaller     // Read-only binding to the contract
	CampaignFactoryTransactor // Write-only binding to the contract
	CampaignFactoryFilterer   // Log filterer for contract events
}

// CampaignFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignFactorySession struct {
	Contract     *CampaignFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CampaignFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignFactoryCallerSession struct {
	Contract *CampaignFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CampaignFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignFactoryTransactorSession struct {
	Contract     *CampaignFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CampaignFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignFactoryRaw struct {
	Contract *CampaignFactory // Generic contract binding to access the raw methods on
}

// CampaignFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignFactoryCallerRaw struct {
	Contract *CampaignFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignFactoryTransactorRaw struct {
	Contract *CampaignFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaignFactory creates a new instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactory(address common.Address, backend bind.ContractBackend) (*CampaignFactory, error) {
	contract, err := bindCampaignFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CampaignFactory{CampaignFactoryCaller: CampaignFactoryCaller{contract: contract}, CampaignFactoryTransactor: CampaignFactoryTransactor{contract: contract}, CampaignFactoryFilterer: CampaignFactoryFilterer{contract: contract}}, nil
}

// NewCampaignFactoryCaller creates a new read-only instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryCaller(address common.Address, caller bind.ContractCaller) (*CampaignFactoryCaller, error) {
	contract, err := bindCampaignFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryCaller{contract: contract}, nil
}

// NewCampaignFactoryTransactor creates a new write-only instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignFactoryTransactor, error) {
	contract, err := bindCampaignFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryTransactor{contract: contract}, nil
}

// NewCampaignFactoryFilterer creates a new log filterer instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignFactoryFilterer, error) {
	contract, err := bindCampaignFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryFilterer{contract: contract}, nil
}

// bindCampaignFactory binds a generic wrapper to an already deployed contract.
func bindCampaignFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CampaignFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignFactory *CampaignFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignFactory.Contract.CampaignFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignFactory *CampaignFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CampaignFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignFactory *CampaignFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CampaignFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignFactory *CampaignFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignFactory *CampaignFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignFactory *CampaignFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignFactory.Contract.contract.Transact(opts, method, params...)
}

// AllCampaigns is a free data retrieval call binding the contract method 0xa509f4d1.
//
// Solidity: function allCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) AllCampaigns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "allCampaigns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllCampaigns is a free data retrieval call binding the contract method 0xa509f4d1.
//
// Solidity: function allCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactorySession) AllCampaigns(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.AllCampaigns(&_CampaignFactory.CallOpts, arg0)
}

// AllCampaigns is a free data retrieval call binding the contract method 0xa509f4d1.
//
// Solidity: function allCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) AllCampaigns(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.AllCampaigns(&_CampaignFactory.CallOpts, arg0)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) CampaignById(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "campaignById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactorySession) CampaignById(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.CampaignById(&_CampaignFactory.CallOpts, arg0)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) CampaignById(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.CampaignById(&_CampaignFactory.CallOpts, arg0)
}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) DonationNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "donationNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_CampaignFactory *CampaignFactorySession) DonationNFT() (common.Address, error) {
	return _CampaignFactory.Contract.DonationNFT(&_CampaignFactory.CallOpts)
}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) DonationNFT() (common.Address, error) {
	return _CampaignFactory.Contract.DonationNFT(&_CampaignFactory.CallOpts)
}

// GetCampaignCount is a free data retrieval call binding the contract method 0x6caa9218.
//
// Solidity: function getCampaignCount() view returns(uint256)
func (_CampaignFactory *CampaignFactoryCaller) GetCampaignCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "getCampaignCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCampaignCount is a free data retrieval call binding the contract method 0x6caa9218.
//
// Solidity: function getCampaignCount() view returns(uint256)
func (_CampaignFactory *CampaignFactorySession) GetCampaignCount() (*big.Int, error) {
	return _CampaignFactory.Contract.GetCampaignCount(&_CampaignFactory.CallOpts)
}

// GetCampaignCount is a free data retrieval call binding the contract method 0x6caa9218.
//
// Solidity: function getCampaignCount() view returns(uint256)
func (_CampaignFactory *CampaignFactoryCallerSession) GetCampaignCount() (*big.Int, error) {
	return _CampaignFactory.Contract.GetCampaignCount(&_CampaignFactory.CallOpts)
}

// GetCampaigns is a free data retrieval call binding the contract method 0x09051566.
//
// Solidity: function getCampaigns(uint256 offset, uint256 limit) view returns(address[] result)
func (_CampaignFactory *CampaignFactoryCaller) GetCampaigns(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "getCampaigns", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCampaigns is a free data retrieval call binding the contract method 0x09051566.
//
// Solidity: function getCampaigns(uint256 offset, uint256 limit) view returns(address[] result)
func (_CampaignFactory *CampaignFactorySession) GetCampaigns(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _CampaignFactory.Contract.GetCampaigns(&_CampaignFactory.CallOpts, offset, limit)
}

// GetCampaigns is a free data retrieval call binding the contract method 0x09051566.
//
// Solidity: function getCampaigns(uint256 offset, uint256 limit) view returns(address[] result)
func (_CampaignFactory *CampaignFactoryCallerSession) GetCampaigns(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _CampaignFactory.Contract.GetCampaigns(&_CampaignFactory.CallOpts, offset, limit)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignFactory *CampaignFactorySession) Owner() (common.Address, error) {
	return _CampaignFactory.Contract.Owner(&_CampaignFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) Owner() (common.Address, error) {
	return _CampaignFactory.Contract.Owner(&_CampaignFactory.CallOpts)
}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) ReputationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "reputationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_CampaignFactory *CampaignFactorySession) ReputationRegistry() (common.Address, error) {
	return _CampaignFactory.Contract.ReputationRegistry(&_CampaignFactory.CallOpts)
}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) ReputationRegistry() (common.Address, error) {
	return _CampaignFactory.Contract.ReputationRegistry(&_CampaignFactory.CallOpts)
}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) TimelockController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "timelockController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_CampaignFactory *CampaignFactorySession) TimelockController() (common.Address, error) {
	return _CampaignFactory.Contract.TimelockController(&_CampaignFactory.CallOpts)
}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) TimelockController() (common.Address, error) {
	return _CampaignFactory.Contract.TimelockController(&_CampaignFactory.CallOpts)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xecc17f88.
//
// Solidity: function createCampaign(address creatorWallet, uint256 goalAmount, bytes32 documentHash, string[] milestoneDescriptions, uint256[] milestoneAmounts, uint256[] milestoneDeadlines, string[] milestoneEvidence) returns(address campaignAddress)
func (_CampaignFactory *CampaignFactoryTransactor) CreateCampaign(opts *bind.TransactOpts, creatorWallet common.Address, goalAmount *big.Int, documentHash [32]byte, milestoneDescriptions []string, milestoneAmounts []*big.Int, milestoneDeadlines []*big.Int, milestoneEvidence []string) (*types.Transaction, error) {
	return _CampaignFactory.contract.Transact(opts, "createCampaign", creatorWallet, goalAmount, documentHash, milestoneDescriptions, milestoneAmounts, milestoneDeadlines, milestoneEvidence)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xecc17f88.
//
// Solidity: function createCampaign(address creatorWallet, uint256 goalAmount, bytes32 documentHash, string[] milestoneDescriptions, uint256[] milestoneAmounts, uint256[] milestoneDeadlines, string[] milestoneEvidence) returns(address campaignAddress)
func (_CampaignFactory *CampaignFactorySession) CreateCampaign(creatorWallet common.Address, goalAmount *big.Int, documentHash [32]byte, milestoneDescriptions []string, milestoneAmounts []*big.Int, milestoneDeadlines []*big.Int, milestoneEvidence []string) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CreateCampaign(&_CampaignFactory.TransactOpts, creatorWallet, goalAmount, documentHash, milestoneDescriptions, milestoneAmounts, milestoneDeadlines, milestoneEvidence)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xecc17f88.
//
// Solidity: function createCampaign(address creatorWallet, uint256 goalAmount, bytes32 documentHash, string[] milestoneDescriptions, uint256[] milestoneAmounts, uint256[] milestoneDeadlines, string[] milestoneEvidence) returns(address campaignAddress)
func (_CampaignFactory *CampaignFactoryTransactorSession) CreateCampaign(creatorWallet common.Address, goalAmount *big.Int, documentHash [32]byte, milestoneDescriptions []string, milestoneAmounts []*big.Int, milestoneDeadlines []*big.Int, milestoneEvidence []string) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CreateCampaign(&_CampaignFactory.TransactOpts, creatorWallet, goalAmount, documentHash, milestoneDescriptions, milestoneAmounts, milestoneDeadlines, milestoneEvidence)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignFactory *CampaignFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignFactory *CampaignFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CampaignFactory.Contract.RenounceOwnership(&_CampaignFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignFactory *CampaignFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CampaignFactory.Contract.RenounceOwnership(&_CampaignFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignFactory *CampaignFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CampaignFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignFactory *CampaignFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CampaignFactory.Contract.TransferOwnership(&_CampaignFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignFactory *CampaignFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CampaignFactory.Contract.TransferOwnership(&_CampaignFactory.TransactOpts, newOwner)
}

// CampaignFactoryCampaignCreatedIterator is returned from FilterCampaignCreated and is used to iterate over the raw logs and unpacked data for CampaignCreated events raised by the CampaignFactory contract.
type CampaignFactoryCampaignCreatedIterator struct {
	Event *CampaignFactoryCampaignCreated // Event containing the contract specifics and raw log

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
func (it *CampaignFactoryCampaignCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignFactoryCampaignCreated)
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
		it.Event = new(CampaignFactoryCampaignCreated)
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
func (it *CampaignFactoryCampaignCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignFactoryCampaignCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignFactoryCampaignCreated represents a CampaignCreated event raised by the CampaignFactory contract.
type CampaignFactoryCampaignCreated struct {
	CampaignAddress common.Address
	CreatorWallet   common.Address
	CampaignId      *big.Int
	GoalAmount      *big.Int
	MilestoneCount  *big.Int
	DocumentHash    [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreated is a free log retrieval operation binding the contract event 0x3134cbb780ced4ba25bc7e71a70fee127918e70aa46c0e79eb2f9a8d0b56c4cf.
//
// Solidity: event CampaignCreated(address indexed campaignAddress, address indexed creatorWallet, uint256 indexed campaignId, uint256 goalAmount, uint256 milestoneCount, bytes32 documentHash)
func (_CampaignFactory *CampaignFactoryFilterer) FilterCampaignCreated(opts *bind.FilterOpts, campaignAddress []common.Address, creatorWallet []common.Address, campaignId []*big.Int) (*CampaignFactoryCampaignCreatedIterator, error) {

	var campaignAddressRule []interface{}
	for _, campaignAddressItem := range campaignAddress {
		campaignAddressRule = append(campaignAddressRule, campaignAddressItem)
	}
	var creatorWalletRule []interface{}
	for _, creatorWalletItem := range creatorWallet {
		creatorWalletRule = append(creatorWalletRule, creatorWalletItem)
	}
	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _CampaignFactory.contract.FilterLogs(opts, "CampaignCreated", campaignAddressRule, creatorWalletRule, campaignIdRule)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryCampaignCreatedIterator{contract: _CampaignFactory.contract, event: "CampaignCreated", logs: logs, sub: sub}, nil
}

// WatchCampaignCreated is a free log subscription operation binding the contract event 0x3134cbb780ced4ba25bc7e71a70fee127918e70aa46c0e79eb2f9a8d0b56c4cf.
//
// Solidity: event CampaignCreated(address indexed campaignAddress, address indexed creatorWallet, uint256 indexed campaignId, uint256 goalAmount, uint256 milestoneCount, bytes32 documentHash)
func (_CampaignFactory *CampaignFactoryFilterer) WatchCampaignCreated(opts *bind.WatchOpts, sink chan<- *CampaignFactoryCampaignCreated, campaignAddress []common.Address, creatorWallet []common.Address, campaignId []*big.Int) (event.Subscription, error) {

	var campaignAddressRule []interface{}
	for _, campaignAddressItem := range campaignAddress {
		campaignAddressRule = append(campaignAddressRule, campaignAddressItem)
	}
	var creatorWalletRule []interface{}
	for _, creatorWalletItem := range creatorWallet {
		creatorWalletRule = append(creatorWalletRule, creatorWalletItem)
	}
	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _CampaignFactory.contract.WatchLogs(opts, "CampaignCreated", campaignAddressRule, creatorWalletRule, campaignIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignFactoryCampaignCreated)
				if err := _CampaignFactory.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
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

// ParseCampaignCreated is a log parse operation binding the contract event 0x3134cbb780ced4ba25bc7e71a70fee127918e70aa46c0e79eb2f9a8d0b56c4cf.
//
// Solidity: event CampaignCreated(address indexed campaignAddress, address indexed creatorWallet, uint256 indexed campaignId, uint256 goalAmount, uint256 milestoneCount, bytes32 documentHash)
func (_CampaignFactory *CampaignFactoryFilterer) ParseCampaignCreated(log types.Log) (*CampaignFactoryCampaignCreated, error) {
	event := new(CampaignFactoryCampaignCreated)
	if err := _CampaignFactory.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CampaignFactory contract.
type CampaignFactoryOwnershipTransferredIterator struct {
	Event *CampaignFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CampaignFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignFactoryOwnershipTransferred)
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
		it.Event = new(CampaignFactoryOwnershipTransferred)
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
func (it *CampaignFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CampaignFactory contract.
type CampaignFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CampaignFactory *CampaignFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CampaignFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CampaignFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryOwnershipTransferredIterator{contract: _CampaignFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CampaignFactory *CampaignFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CampaignFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CampaignFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignFactoryOwnershipTransferred)
				if err := _CampaignFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CampaignFactory *CampaignFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CampaignFactoryOwnershipTransferred, error) {
	event := new(CampaignFactoryOwnershipTransferred)
	if err := _CampaignFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
