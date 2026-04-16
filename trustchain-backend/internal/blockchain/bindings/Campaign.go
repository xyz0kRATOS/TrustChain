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

// CampaignMilestone is an auto generated low-level Go binding around an user-defined struct.
type CampaignMilestone struct {
	Description      string
	Amount           *big.Int
	Deadline         *big.Int
	RequiredEvidence string
	Status           uint8
	FundsReleased    bool
}

// CampaignMetaData contains all meta data concerning the Campaign contract.
var CampaignMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminMultiSig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_timelockController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reputationRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_donationNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_goalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_documentHash\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"_milestoneDescriptions\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_milestoneAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_milestoneDeadlines\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_milestoneEvidence\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CampaignCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"frozenBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CampaignFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"updateHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"updateType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CampaignUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DonationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonorRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FundsReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"amendmentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MilestoneAmended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MilestoneApproved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminMultiSig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"}],\"name\":\"approveMilestone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"documentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsMetadataHash\",\"type\":\"string\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donationNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"}],\"name\":\"executeMilestoneRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDonorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getMilestone\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requiredEvidence\",\"type\":\"string\"},{\"internalType\":\"enumCampaign.MilestoneStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"fundsReleased\",\"type\":\"bool\"}],\"internalType\":\"structCampaign.Milestone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMilestoneCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"milestones\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requiredEvidence\",\"type\":\"string\"},{\"internalType\":\"enumCampaign.MilestoneStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"fundsReleased\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"amendmentHash\",\"type\":\"bytes32\"}],\"name\":\"recordAmendment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"updateHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"updateType\",\"type\":\"string\"}],\"name\":\"recordUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputationRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"milestoneIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumCampaign.MilestoneStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setMilestoneStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumCampaign.CampaignStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelockController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CampaignABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignMetaData.ABI instead.
var CampaignABI = CampaignMetaData.ABI

// Campaign is an auto generated Go binding around an Ethereum contract.
type Campaign struct {
	CampaignCaller     // Read-only binding to the contract
	CampaignTransactor // Write-only binding to the contract
	CampaignFilterer   // Log filterer for contract events
}

// CampaignCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignSession struct {
	Contract     *Campaign         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CampaignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignCallerSession struct {
	Contract *CampaignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CampaignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignTransactorSession struct {
	Contract     *CampaignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CampaignRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignRaw struct {
	Contract *Campaign // Generic contract binding to access the raw methods on
}

// CampaignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignCallerRaw struct {
	Contract *CampaignCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignTransactorRaw struct {
	Contract *CampaignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaign creates a new instance of Campaign, bound to a specific deployed contract.
func NewCampaign(address common.Address, backend bind.ContractBackend) (*Campaign, error) {
	contract, err := bindCampaign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Campaign{CampaignCaller: CampaignCaller{contract: contract}, CampaignTransactor: CampaignTransactor{contract: contract}, CampaignFilterer: CampaignFilterer{contract: contract}}, nil
}

// NewCampaignCaller creates a new read-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignCaller(address common.Address, caller bind.ContractCaller) (*CampaignCaller, error) {
	contract, err := bindCampaign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignCaller{contract: contract}, nil
}

// NewCampaignTransactor creates a new write-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignTransactor, error) {
	contract, err := bindCampaign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignTransactor{contract: contract}, nil
}

// NewCampaignFilterer creates a new log filterer instance of Campaign, bound to a specific deployed contract.
func NewCampaignFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignFilterer, error) {
	contract, err := bindCampaign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignFilterer{contract: contract}, nil
}

// bindCampaign binds a generic wrapper to an already deployed contract.
func bindCampaign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CampaignMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.CampaignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transact(opts, method, params...)
}

// AdminMultiSig is a free data retrieval call binding the contract method 0x45c6a874.
//
// Solidity: function adminMultiSig() view returns(address)
func (_Campaign *CampaignCaller) AdminMultiSig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "adminMultiSig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMultiSig is a free data retrieval call binding the contract method 0x45c6a874.
//
// Solidity: function adminMultiSig() view returns(address)
func (_Campaign *CampaignSession) AdminMultiSig() (common.Address, error) {
	return _Campaign.Contract.AdminMultiSig(&_Campaign.CallOpts)
}

// AdminMultiSig is a free data retrieval call binding the contract method 0x45c6a874.
//
// Solidity: function adminMultiSig() view returns(address)
func (_Campaign *CampaignCallerSession) AdminMultiSig() (common.Address, error) {
	return _Campaign.Contract.AdminMultiSig(&_Campaign.CallOpts)
}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_Campaign *CampaignCaller) CampaignId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "campaignId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_Campaign *CampaignSession) CampaignId() (*big.Int, error) {
	return _Campaign.Contract.CampaignId(&_Campaign.CallOpts)
}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_Campaign *CampaignCallerSession) CampaignId() (*big.Int, error) {
	return _Campaign.Contract.CampaignId(&_Campaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignSession) Creator() (common.Address, error) {
	return _Campaign.Contract.Creator(&_Campaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignCallerSession) Creator() (common.Address, error) {
	return _Campaign.Contract.Creator(&_Campaign.CallOpts)
}

// DocumentHash is a free data retrieval call binding the contract method 0x797bec77.
//
// Solidity: function documentHash() view returns(bytes32)
func (_Campaign *CampaignCaller) DocumentHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "documentHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DocumentHash is a free data retrieval call binding the contract method 0x797bec77.
//
// Solidity: function documentHash() view returns(bytes32)
func (_Campaign *CampaignSession) DocumentHash() ([32]byte, error) {
	return _Campaign.Contract.DocumentHash(&_Campaign.CallOpts)
}

// DocumentHash is a free data retrieval call binding the contract method 0x797bec77.
//
// Solidity: function documentHash() view returns(bytes32)
func (_Campaign *CampaignCallerSession) DocumentHash() ([32]byte, error) {
	return _Campaign.Contract.DocumentHash(&_Campaign.CallOpts)
}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_Campaign *CampaignCaller) DonationNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "donationNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_Campaign *CampaignSession) DonationNFT() (common.Address, error) {
	return _Campaign.Contract.DonationNFT(&_Campaign.CallOpts)
}

// DonationNFT is a free data retrieval call binding the contract method 0x23842202.
//
// Solidity: function donationNFT() view returns(address)
func (_Campaign *CampaignCallerSession) DonationNFT() (common.Address, error) {
	return _Campaign.Contract.DonationNFT(&_Campaign.CallOpts)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Campaign *CampaignCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Campaign *CampaignSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _Campaign.Contract.Donations(&_Campaign.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Campaign *CampaignCallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _Campaign.Contract.Donations(&_Campaign.CallOpts, arg0)
}

// DonorList is a free data retrieval call binding the contract method 0x5eed3dcb.
//
// Solidity: function donorList(uint256 ) view returns(address)
func (_Campaign *CampaignCaller) DonorList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "donorList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonorList is a free data retrieval call binding the contract method 0x5eed3dcb.
//
// Solidity: function donorList(uint256 ) view returns(address)
func (_Campaign *CampaignSession) DonorList(arg0 *big.Int) (common.Address, error) {
	return _Campaign.Contract.DonorList(&_Campaign.CallOpts, arg0)
}

// DonorList is a free data retrieval call binding the contract method 0x5eed3dcb.
//
// Solidity: function donorList(uint256 ) view returns(address)
func (_Campaign *CampaignCallerSession) DonorList(arg0 *big.Int) (common.Address, error) {
	return _Campaign.Contract.DonorList(&_Campaign.CallOpts, arg0)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Campaign *CampaignCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Campaign *CampaignSession) GetContractBalance() (*big.Int, error) {
	return _Campaign.Contract.GetContractBalance(&_Campaign.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Campaign *CampaignCallerSession) GetContractBalance() (*big.Int, error) {
	return _Campaign.Contract.GetContractBalance(&_Campaign.CallOpts)
}

// GetDonorCount is a free data retrieval call binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() view returns(uint256)
func (_Campaign *CampaignCaller) GetDonorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "getDonorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonorCount is a free data retrieval call binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() view returns(uint256)
func (_Campaign *CampaignSession) GetDonorCount() (*big.Int, error) {
	return _Campaign.Contract.GetDonorCount(&_Campaign.CallOpts)
}

// GetDonorCount is a free data retrieval call binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() view returns(uint256)
func (_Campaign *CampaignCallerSession) GetDonorCount() (*big.Int, error) {
	return _Campaign.Contract.GetDonorCount(&_Campaign.CallOpts)
}

// GetMilestone is a free data retrieval call binding the contract method 0x2442e1cb.
//
// Solidity: function getMilestone(uint256 index) view returns((string,uint256,uint256,string,uint8,bool))
func (_Campaign *CampaignCaller) GetMilestone(opts *bind.CallOpts, index *big.Int) (CampaignMilestone, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "getMilestone", index)

	if err != nil {
		return *new(CampaignMilestone), err
	}

	out0 := *abi.ConvertType(out[0], new(CampaignMilestone)).(*CampaignMilestone)

	return out0, err

}

// GetMilestone is a free data retrieval call binding the contract method 0x2442e1cb.
//
// Solidity: function getMilestone(uint256 index) view returns((string,uint256,uint256,string,uint8,bool))
func (_Campaign *CampaignSession) GetMilestone(index *big.Int) (CampaignMilestone, error) {
	return _Campaign.Contract.GetMilestone(&_Campaign.CallOpts, index)
}

// GetMilestone is a free data retrieval call binding the contract method 0x2442e1cb.
//
// Solidity: function getMilestone(uint256 index) view returns((string,uint256,uint256,string,uint8,bool))
func (_Campaign *CampaignCallerSession) GetMilestone(index *big.Int) (CampaignMilestone, error) {
	return _Campaign.Contract.GetMilestone(&_Campaign.CallOpts, index)
}

// GetMilestoneCount is a free data retrieval call binding the contract method 0x8f4cf247.
//
// Solidity: function getMilestoneCount() view returns(uint256)
func (_Campaign *CampaignCaller) GetMilestoneCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "getMilestoneCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMilestoneCount is a free data retrieval call binding the contract method 0x8f4cf247.
//
// Solidity: function getMilestoneCount() view returns(uint256)
func (_Campaign *CampaignSession) GetMilestoneCount() (*big.Int, error) {
	return _Campaign.Contract.GetMilestoneCount(&_Campaign.CallOpts)
}

// GetMilestoneCount is a free data retrieval call binding the contract method 0x8f4cf247.
//
// Solidity: function getMilestoneCount() view returns(uint256)
func (_Campaign *CampaignCallerSession) GetMilestoneCount() (*big.Int, error) {
	return _Campaign.Contract.GetMilestoneCount(&_Campaign.CallOpts)
}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_Campaign *CampaignCaller) GoalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "goalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_Campaign *CampaignSession) GoalAmount() (*big.Int, error) {
	return _Campaign.Contract.GoalAmount(&_Campaign.CallOpts)
}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_Campaign *CampaignCallerSession) GoalAmount() (*big.Int, error) {
	return _Campaign.Contract.GoalAmount(&_Campaign.CallOpts)
}

// Milestones is a free data retrieval call binding the contract method 0xe89e4ed6.
//
// Solidity: function milestones(uint256 ) view returns(string description, uint256 amount, uint256 deadline, string requiredEvidence, uint8 status, bool fundsReleased)
func (_Campaign *CampaignCaller) Milestones(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description      string
	Amount           *big.Int
	Deadline         *big.Int
	RequiredEvidence string
	Status           uint8
	FundsReleased    bool
}, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "milestones", arg0)

	outstruct := new(struct {
		Description      string
		Amount           *big.Int
		Deadline         *big.Int
		RequiredEvidence string
		Status           uint8
		FundsReleased    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequiredEvidence = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.FundsReleased = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Milestones is a free data retrieval call binding the contract method 0xe89e4ed6.
//
// Solidity: function milestones(uint256 ) view returns(string description, uint256 amount, uint256 deadline, string requiredEvidence, uint8 status, bool fundsReleased)
func (_Campaign *CampaignSession) Milestones(arg0 *big.Int) (struct {
	Description      string
	Amount           *big.Int
	Deadline         *big.Int
	RequiredEvidence string
	Status           uint8
	FundsReleased    bool
}, error) {
	return _Campaign.Contract.Milestones(&_Campaign.CallOpts, arg0)
}

// Milestones is a free data retrieval call binding the contract method 0xe89e4ed6.
//
// Solidity: function milestones(uint256 ) view returns(string description, uint256 amount, uint256 deadline, string requiredEvidence, uint8 status, bool fundsReleased)
func (_Campaign *CampaignCallerSession) Milestones(arg0 *big.Int) (struct {
	Description      string
	Amount           *big.Int
	Deadline         *big.Int
	RequiredEvidence string
	Status           uint8
	FundsReleased    bool
}, error) {
	return _Campaign.Contract.Milestones(&_Campaign.CallOpts, arg0)
}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_Campaign *CampaignCaller) ReputationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "reputationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_Campaign *CampaignSession) ReputationRegistry() (common.Address, error) {
	return _Campaign.Contract.ReputationRegistry(&_Campaign.CallOpts)
}

// ReputationRegistry is a free data retrieval call binding the contract method 0xc8db44e3.
//
// Solidity: function reputationRegistry() view returns(address)
func (_Campaign *CampaignCallerSession) ReputationRegistry() (common.Address, error) {
	return _Campaign.Contract.ReputationRegistry(&_Campaign.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Campaign *CampaignCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Campaign *CampaignSession) Status() (uint8, error) {
	return _Campaign.Contract.Status(&_Campaign.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Campaign *CampaignCallerSession) Status() (uint8, error) {
	return _Campaign.Contract.Status(&_Campaign.CallOpts)
}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_Campaign *CampaignCaller) TimelockController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "timelockController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_Campaign *CampaignSession) TimelockController() (common.Address, error) {
	return _Campaign.Contract.TimelockController(&_Campaign.CallOpts)
}

// TimelockController is a free data retrieval call binding the contract method 0xf3388e75.
//
// Solidity: function timelockController() view returns(address)
func (_Campaign *CampaignCallerSession) TimelockController() (common.Address, error) {
	return _Campaign.Contract.TimelockController(&_Campaign.CallOpts)
}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_Campaign *CampaignCaller) TotalRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "totalRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_Campaign *CampaignSession) TotalRaised() (*big.Int, error) {
	return _Campaign.Contract.TotalRaised(&_Campaign.CallOpts)
}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_Campaign *CampaignCallerSession) TotalRaised() (*big.Int, error) {
	return _Campaign.Contract.TotalRaised(&_Campaign.CallOpts)
}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Campaign *CampaignCaller) TotalReleased(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "totalReleased")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Campaign *CampaignSession) TotalReleased() (*big.Int, error) {
	return _Campaign.Contract.TotalReleased(&_Campaign.CallOpts)
}

// TotalReleased is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Campaign *CampaignCallerSession) TotalReleased() (*big.Int, error) {
	return _Campaign.Contract.TotalReleased(&_Campaign.CallOpts)
}

// ApproveMilestone is a paid mutator transaction binding the contract method 0xc438b40f.
//
// Solidity: function approveMilestone(uint256 milestoneIndex) returns()
func (_Campaign *CampaignTransactor) ApproveMilestone(opts *bind.TransactOpts, milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "approveMilestone", milestoneIndex)
}

// ApproveMilestone is a paid mutator transaction binding the contract method 0xc438b40f.
//
// Solidity: function approveMilestone(uint256 milestoneIndex) returns()
func (_Campaign *CampaignSession) ApproveMilestone(milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveMilestone(&_Campaign.TransactOpts, milestoneIndex)
}

// ApproveMilestone is a paid mutator transaction binding the contract method 0xc438b40f.
//
// Solidity: function approveMilestone(uint256 milestoneIndex) returns()
func (_Campaign *CampaignTransactorSession) ApproveMilestone(milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveMilestone(&_Campaign.TransactOpts, milestoneIndex)
}

// Donate is a paid mutator transaction binding the contract method 0xb5aebc80.
//
// Solidity: function donate(string ipfsMetadataHash) payable returns()
func (_Campaign *CampaignTransactor) Donate(opts *bind.TransactOpts, ipfsMetadataHash string) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "donate", ipfsMetadataHash)
}

// Donate is a paid mutator transaction binding the contract method 0xb5aebc80.
//
// Solidity: function donate(string ipfsMetadataHash) payable returns()
func (_Campaign *CampaignSession) Donate(ipfsMetadataHash string) (*types.Transaction, error) {
	return _Campaign.Contract.Donate(&_Campaign.TransactOpts, ipfsMetadataHash)
}

// Donate is a paid mutator transaction binding the contract method 0xb5aebc80.
//
// Solidity: function donate(string ipfsMetadataHash) payable returns()
func (_Campaign *CampaignTransactorSession) Donate(ipfsMetadataHash string) (*types.Transaction, error) {
	return _Campaign.Contract.Donate(&_Campaign.TransactOpts, ipfsMetadataHash)
}

// ExecuteMilestoneRelease is a paid mutator transaction binding the contract method 0xfafeb636.
//
// Solidity: function executeMilestoneRelease(uint256 milestoneIndex) returns()
func (_Campaign *CampaignTransactor) ExecuteMilestoneRelease(opts *bind.TransactOpts, milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "executeMilestoneRelease", milestoneIndex)
}

// ExecuteMilestoneRelease is a paid mutator transaction binding the contract method 0xfafeb636.
//
// Solidity: function executeMilestoneRelease(uint256 milestoneIndex) returns()
func (_Campaign *CampaignSession) ExecuteMilestoneRelease(milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ExecuteMilestoneRelease(&_Campaign.TransactOpts, milestoneIndex)
}

// ExecuteMilestoneRelease is a paid mutator transaction binding the contract method 0xfafeb636.
//
// Solidity: function executeMilestoneRelease(uint256 milestoneIndex) returns()
func (_Campaign *CampaignTransactorSession) ExecuteMilestoneRelease(milestoneIndex *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ExecuteMilestoneRelease(&_Campaign.TransactOpts, milestoneIndex)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Campaign *CampaignTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Campaign *CampaignSession) Freeze() (*types.Transaction, error) {
	return _Campaign.Contract.Freeze(&_Campaign.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Campaign *CampaignTransactorSession) Freeze() (*types.Transaction, error) {
	return _Campaign.Contract.Freeze(&_Campaign.TransactOpts)
}

// RecordAmendment is a paid mutator transaction binding the contract method 0x9b8fa6fb.
//
// Solidity: function recordAmendment(uint256 milestoneIndex, bytes32 amendmentHash) returns()
func (_Campaign *CampaignTransactor) RecordAmendment(opts *bind.TransactOpts, milestoneIndex *big.Int, amendmentHash [32]byte) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "recordAmendment", milestoneIndex, amendmentHash)
}

// RecordAmendment is a paid mutator transaction binding the contract method 0x9b8fa6fb.
//
// Solidity: function recordAmendment(uint256 milestoneIndex, bytes32 amendmentHash) returns()
func (_Campaign *CampaignSession) RecordAmendment(milestoneIndex *big.Int, amendmentHash [32]byte) (*types.Transaction, error) {
	return _Campaign.Contract.RecordAmendment(&_Campaign.TransactOpts, milestoneIndex, amendmentHash)
}

// RecordAmendment is a paid mutator transaction binding the contract method 0x9b8fa6fb.
//
// Solidity: function recordAmendment(uint256 milestoneIndex, bytes32 amendmentHash) returns()
func (_Campaign *CampaignTransactorSession) RecordAmendment(milestoneIndex *big.Int, amendmentHash [32]byte) (*types.Transaction, error) {
	return _Campaign.Contract.RecordAmendment(&_Campaign.TransactOpts, milestoneIndex, amendmentHash)
}

// RecordUpdate is a paid mutator transaction binding the contract method 0xbbc533b6.
//
// Solidity: function recordUpdate(bytes32 updateHash, string updateType) returns()
func (_Campaign *CampaignTransactor) RecordUpdate(opts *bind.TransactOpts, updateHash [32]byte, updateType string) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "recordUpdate", updateHash, updateType)
}

// RecordUpdate is a paid mutator transaction binding the contract method 0xbbc533b6.
//
// Solidity: function recordUpdate(bytes32 updateHash, string updateType) returns()
func (_Campaign *CampaignSession) RecordUpdate(updateHash [32]byte, updateType string) (*types.Transaction, error) {
	return _Campaign.Contract.RecordUpdate(&_Campaign.TransactOpts, updateHash, updateType)
}

// RecordUpdate is a paid mutator transaction binding the contract method 0xbbc533b6.
//
// Solidity: function recordUpdate(bytes32 updateHash, string updateType) returns()
func (_Campaign *CampaignTransactorSession) RecordUpdate(updateHash [32]byte, updateType string) (*types.Transaction, error) {
	return _Campaign.Contract.RecordUpdate(&_Campaign.TransactOpts, updateHash, updateType)
}

// SetMilestoneStatus is a paid mutator transaction binding the contract method 0x5d9bf255.
//
// Solidity: function setMilestoneStatus(uint256 milestoneIndex, uint8 newStatus) returns()
func (_Campaign *CampaignTransactor) SetMilestoneStatus(opts *bind.TransactOpts, milestoneIndex *big.Int, newStatus uint8) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "setMilestoneStatus", milestoneIndex, newStatus)
}

// SetMilestoneStatus is a paid mutator transaction binding the contract method 0x5d9bf255.
//
// Solidity: function setMilestoneStatus(uint256 milestoneIndex, uint8 newStatus) returns()
func (_Campaign *CampaignSession) SetMilestoneStatus(milestoneIndex *big.Int, newStatus uint8) (*types.Transaction, error) {
	return _Campaign.Contract.SetMilestoneStatus(&_Campaign.TransactOpts, milestoneIndex, newStatus)
}

// SetMilestoneStatus is a paid mutator transaction binding the contract method 0x5d9bf255.
//
// Solidity: function setMilestoneStatus(uint256 milestoneIndex, uint8 newStatus) returns()
func (_Campaign *CampaignTransactorSession) SetMilestoneStatus(milestoneIndex *big.Int, newStatus uint8) (*types.Transaction, error) {
	return _Campaign.Contract.SetMilestoneStatus(&_Campaign.TransactOpts, milestoneIndex, newStatus)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Campaign *CampaignTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Campaign *CampaignSession) Receive() (*types.Transaction, error) {
	return _Campaign.Contract.Receive(&_Campaign.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Campaign *CampaignTransactorSession) Receive() (*types.Transaction, error) {
	return _Campaign.Contract.Receive(&_Campaign.TransactOpts)
}

// CampaignCampaignCompletedIterator is returned from FilterCampaignCompleted and is used to iterate over the raw logs and unpacked data for CampaignCompleted events raised by the Campaign contract.
type CampaignCampaignCompletedIterator struct {
	Event *CampaignCampaignCompleted // Event containing the contract specifics and raw log

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
func (it *CampaignCampaignCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignCampaignCompleted)
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
		it.Event = new(CampaignCampaignCompleted)
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
func (it *CampaignCampaignCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignCampaignCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignCampaignCompleted represents a CampaignCompleted event raised by the Campaign contract.
type CampaignCampaignCompleted struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCampaignCompleted is a free log retrieval operation binding the contract event 0xbdef6a2e3aa892961a2dab48272f72f2bc1c7121eb1c0ed7e4a58e081c1e4be0.
//
// Solidity: event CampaignCompleted(uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterCampaignCompleted(opts *bind.FilterOpts) (*CampaignCampaignCompletedIterator, error) {

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "CampaignCompleted")
	if err != nil {
		return nil, err
	}
	return &CampaignCampaignCompletedIterator{contract: _Campaign.contract, event: "CampaignCompleted", logs: logs, sub: sub}, nil
}

// WatchCampaignCompleted is a free log subscription operation binding the contract event 0xbdef6a2e3aa892961a2dab48272f72f2bc1c7121eb1c0ed7e4a58e081c1e4be0.
//
// Solidity: event CampaignCompleted(uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchCampaignCompleted(opts *bind.WatchOpts, sink chan<- *CampaignCampaignCompleted) (event.Subscription, error) {

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "CampaignCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignCampaignCompleted)
				if err := _Campaign.contract.UnpackLog(event, "CampaignCompleted", log); err != nil {
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

// ParseCampaignCompleted is a log parse operation binding the contract event 0xbdef6a2e3aa892961a2dab48272f72f2bc1c7121eb1c0ed7e4a58e081c1e4be0.
//
// Solidity: event CampaignCompleted(uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseCampaignCompleted(log types.Log) (*CampaignCampaignCompleted, error) {
	event := new(CampaignCampaignCompleted)
	if err := _Campaign.contract.UnpackLog(event, "CampaignCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignCampaignFrozenIterator is returned from FilterCampaignFrozen and is used to iterate over the raw logs and unpacked data for CampaignFrozen events raised by the Campaign contract.
type CampaignCampaignFrozenIterator struct {
	Event *CampaignCampaignFrozen // Event containing the contract specifics and raw log

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
func (it *CampaignCampaignFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignCampaignFrozen)
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
		it.Event = new(CampaignCampaignFrozen)
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
func (it *CampaignCampaignFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignCampaignFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignCampaignFrozen represents a CampaignFrozen event raised by the Campaign contract.
type CampaignCampaignFrozen struct {
	FrozenBy  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCampaignFrozen is a free log retrieval operation binding the contract event 0xc60a19caf4129aba2174b1b7a37d436ac19d86d92abc679d195808d7cae974b5.
//
// Solidity: event CampaignFrozen(address indexed frozenBy, uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterCampaignFrozen(opts *bind.FilterOpts, frozenBy []common.Address) (*CampaignCampaignFrozenIterator, error) {

	var frozenByRule []interface{}
	for _, frozenByItem := range frozenBy {
		frozenByRule = append(frozenByRule, frozenByItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "CampaignFrozen", frozenByRule)
	if err != nil {
		return nil, err
	}
	return &CampaignCampaignFrozenIterator{contract: _Campaign.contract, event: "CampaignFrozen", logs: logs, sub: sub}, nil
}

// WatchCampaignFrozen is a free log subscription operation binding the contract event 0xc60a19caf4129aba2174b1b7a37d436ac19d86d92abc679d195808d7cae974b5.
//
// Solidity: event CampaignFrozen(address indexed frozenBy, uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchCampaignFrozen(opts *bind.WatchOpts, sink chan<- *CampaignCampaignFrozen, frozenBy []common.Address) (event.Subscription, error) {

	var frozenByRule []interface{}
	for _, frozenByItem := range frozenBy {
		frozenByRule = append(frozenByRule, frozenByItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "CampaignFrozen", frozenByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignCampaignFrozen)
				if err := _Campaign.contract.UnpackLog(event, "CampaignFrozen", log); err != nil {
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

// ParseCampaignFrozen is a log parse operation binding the contract event 0xc60a19caf4129aba2174b1b7a37d436ac19d86d92abc679d195808d7cae974b5.
//
// Solidity: event CampaignFrozen(address indexed frozenBy, uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseCampaignFrozen(log types.Log) (*CampaignCampaignFrozen, error) {
	event := new(CampaignCampaignFrozen)
	if err := _Campaign.contract.UnpackLog(event, "CampaignFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignCampaignUpdatedIterator is returned from FilterCampaignUpdated and is used to iterate over the raw logs and unpacked data for CampaignUpdated events raised by the Campaign contract.
type CampaignCampaignUpdatedIterator struct {
	Event *CampaignCampaignUpdated // Event containing the contract specifics and raw log

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
func (it *CampaignCampaignUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignCampaignUpdated)
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
		it.Event = new(CampaignCampaignUpdated)
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
func (it *CampaignCampaignUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignCampaignUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignCampaignUpdated represents a CampaignUpdated event raised by the Campaign contract.
type CampaignCampaignUpdated struct {
	CampaignId *big.Int
	UpdateHash [32]byte
	UpdateType string
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignUpdated is a free log retrieval operation binding the contract event 0xa7428d52e88916fbc056cf6f32178460239abb6ff292d18fe2683693fc987d1e.
//
// Solidity: event CampaignUpdated(uint256 indexed campaignId, bytes32 updateHash, string updateType, uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterCampaignUpdated(opts *bind.FilterOpts, campaignId []*big.Int) (*CampaignCampaignUpdatedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return &CampaignCampaignUpdatedIterator{contract: _Campaign.contract, event: "CampaignUpdated", logs: logs, sub: sub}, nil
}

// WatchCampaignUpdated is a free log subscription operation binding the contract event 0xa7428d52e88916fbc056cf6f32178460239abb6ff292d18fe2683693fc987d1e.
//
// Solidity: event CampaignUpdated(uint256 indexed campaignId, bytes32 updateHash, string updateType, uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchCampaignUpdated(opts *bind.WatchOpts, sink chan<- *CampaignCampaignUpdated, campaignId []*big.Int) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "CampaignUpdated", campaignIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignCampaignUpdated)
				if err := _Campaign.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
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

// ParseCampaignUpdated is a log parse operation binding the contract event 0xa7428d52e88916fbc056cf6f32178460239abb6ff292d18fe2683693fc987d1e.
//
// Solidity: event CampaignUpdated(uint256 indexed campaignId, bytes32 updateHash, string updateType, uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseCampaignUpdated(log types.Log) (*CampaignCampaignUpdated, error) {
	event := new(CampaignCampaignUpdated)
	if err := _Campaign.contract.UnpackLog(event, "CampaignUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the Campaign contract.
type CampaignDonationReceivedIterator struct {
	Event *CampaignDonationReceived // Event containing the contract specifics and raw log

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
func (it *CampaignDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignDonationReceived)
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
		it.Event = new(CampaignDonationReceived)
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
func (it *CampaignDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignDonationReceived represents a DonationReceived event raised by the Campaign contract.
type CampaignDonationReceived struct {
	Donor      common.Address
	Amount     *big.Int
	CampaignId *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x3168ca6b1e292883c0668008556fb887f7cce92bcd3c1c57c15a34dcdccd5892.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 indexed campaignId, uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterDonationReceived(opts *bind.FilterOpts, donor []common.Address, campaignId []*big.Int) (*CampaignDonationReceivedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "DonationReceived", donorRule, campaignIdRule)
	if err != nil {
		return nil, err
	}
	return &CampaignDonationReceivedIterator{contract: _Campaign.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x3168ca6b1e292883c0668008556fb887f7cce92bcd3c1c57c15a34dcdccd5892.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 indexed campaignId, uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *CampaignDonationReceived, donor []common.Address, campaignId []*big.Int) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "DonationReceived", donorRule, campaignIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignDonationReceived)
				if err := _Campaign.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x3168ca6b1e292883c0668008556fb887f7cce92bcd3c1c57c15a34dcdccd5892.
//
// Solidity: event DonationReceived(address indexed donor, uint256 amount, uint256 indexed campaignId, uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseDonationReceived(log types.Log) (*CampaignDonationReceived, error) {
	event := new(CampaignDonationReceived)
	if err := _Campaign.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignDonorRefundedIterator is returned from FilterDonorRefunded and is used to iterate over the raw logs and unpacked data for DonorRefunded events raised by the Campaign contract.
type CampaignDonorRefundedIterator struct {
	Event *CampaignDonorRefunded // Event containing the contract specifics and raw log

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
func (it *CampaignDonorRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignDonorRefunded)
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
		it.Event = new(CampaignDonorRefunded)
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
func (it *CampaignDonorRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignDonorRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignDonorRefunded represents a DonorRefunded event raised by the Campaign contract.
type CampaignDonorRefunded struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonorRefunded is a free log retrieval operation binding the contract event 0x33b805a5748a6a0b7957bc0cc438cf7f1aba7f7e70588e183b28e8fdf8c1dffe.
//
// Solidity: event DonorRefunded(address indexed donor, uint256 amount)
func (_Campaign *CampaignFilterer) FilterDonorRefunded(opts *bind.FilterOpts, donor []common.Address) (*CampaignDonorRefundedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "DonorRefunded", donorRule)
	if err != nil {
		return nil, err
	}
	return &CampaignDonorRefundedIterator{contract: _Campaign.contract, event: "DonorRefunded", logs: logs, sub: sub}, nil
}

// WatchDonorRefunded is a free log subscription operation binding the contract event 0x33b805a5748a6a0b7957bc0cc438cf7f1aba7f7e70588e183b28e8fdf8c1dffe.
//
// Solidity: event DonorRefunded(address indexed donor, uint256 amount)
func (_Campaign *CampaignFilterer) WatchDonorRefunded(opts *bind.WatchOpts, sink chan<- *CampaignDonorRefunded, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "DonorRefunded", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignDonorRefunded)
				if err := _Campaign.contract.UnpackLog(event, "DonorRefunded", log); err != nil {
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

// ParseDonorRefunded is a log parse operation binding the contract event 0x33b805a5748a6a0b7957bc0cc438cf7f1aba7f7e70588e183b28e8fdf8c1dffe.
//
// Solidity: event DonorRefunded(address indexed donor, uint256 amount)
func (_Campaign *CampaignFilterer) ParseDonorRefunded(log types.Log) (*CampaignDonorRefunded, error) {
	event := new(CampaignDonorRefunded)
	if err := _Campaign.contract.UnpackLog(event, "DonorRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignFundsReleasedIterator is returned from FilterFundsReleased and is used to iterate over the raw logs and unpacked data for FundsReleased events raised by the Campaign contract.
type CampaignFundsReleasedIterator struct {
	Event *CampaignFundsReleased // Event containing the contract specifics and raw log

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
func (it *CampaignFundsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignFundsReleased)
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
		it.Event = new(CampaignFundsReleased)
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
func (it *CampaignFundsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignFundsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignFundsReleased represents a FundsReleased event raised by the Campaign contract.
type CampaignFundsReleased struct {
	MilestoneIndex *big.Int
	Amount         *big.Int
	Recipient      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsReleased is a free log retrieval operation binding the contract event 0x807158a396c8ce26fb6a4a44930e1b49876133ba7806dbaac491ea4713e4515e.
//
// Solidity: event FundsReleased(uint256 indexed milestoneIndex, uint256 amount, address indexed recipient)
func (_Campaign *CampaignFilterer) FilterFundsReleased(opts *bind.FilterOpts, milestoneIndex []*big.Int, recipient []common.Address) (*CampaignFundsReleasedIterator, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "FundsReleased", milestoneIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CampaignFundsReleasedIterator{contract: _Campaign.contract, event: "FundsReleased", logs: logs, sub: sub}, nil
}

// WatchFundsReleased is a free log subscription operation binding the contract event 0x807158a396c8ce26fb6a4a44930e1b49876133ba7806dbaac491ea4713e4515e.
//
// Solidity: event FundsReleased(uint256 indexed milestoneIndex, uint256 amount, address indexed recipient)
func (_Campaign *CampaignFilterer) WatchFundsReleased(opts *bind.WatchOpts, sink chan<- *CampaignFundsReleased, milestoneIndex []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "FundsReleased", milestoneIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignFundsReleased)
				if err := _Campaign.contract.UnpackLog(event, "FundsReleased", log); err != nil {
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

// ParseFundsReleased is a log parse operation binding the contract event 0x807158a396c8ce26fb6a4a44930e1b49876133ba7806dbaac491ea4713e4515e.
//
// Solidity: event FundsReleased(uint256 indexed milestoneIndex, uint256 amount, address indexed recipient)
func (_Campaign *CampaignFilterer) ParseFundsReleased(log types.Log) (*CampaignFundsReleased, error) {
	event := new(CampaignFundsReleased)
	if err := _Campaign.contract.UnpackLog(event, "FundsReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignMilestoneAmendedIterator is returned from FilterMilestoneAmended and is used to iterate over the raw logs and unpacked data for MilestoneAmended events raised by the Campaign contract.
type CampaignMilestoneAmendedIterator struct {
	Event *CampaignMilestoneAmended // Event containing the contract specifics and raw log

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
func (it *CampaignMilestoneAmendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignMilestoneAmended)
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
		it.Event = new(CampaignMilestoneAmended)
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
func (it *CampaignMilestoneAmendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignMilestoneAmendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignMilestoneAmended represents a MilestoneAmended event raised by the Campaign contract.
type CampaignMilestoneAmended struct {
	MilestoneIndex *big.Int
	AmendmentHash  [32]byte
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMilestoneAmended is a free log retrieval operation binding the contract event 0xf9b5853f7d0dc993174e3c4e9c2e16d05af526a8cb108c33d9edd565df408a8f.
//
// Solidity: event MilestoneAmended(uint256 indexed milestoneIndex, bytes32 amendmentHash, uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterMilestoneAmended(opts *bind.FilterOpts, milestoneIndex []*big.Int) (*CampaignMilestoneAmendedIterator, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "MilestoneAmended", milestoneIndexRule)
	if err != nil {
		return nil, err
	}
	return &CampaignMilestoneAmendedIterator{contract: _Campaign.contract, event: "MilestoneAmended", logs: logs, sub: sub}, nil
}

// WatchMilestoneAmended is a free log subscription operation binding the contract event 0xf9b5853f7d0dc993174e3c4e9c2e16d05af526a8cb108c33d9edd565df408a8f.
//
// Solidity: event MilestoneAmended(uint256 indexed milestoneIndex, bytes32 amendmentHash, uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchMilestoneAmended(opts *bind.WatchOpts, sink chan<- *CampaignMilestoneAmended, milestoneIndex []*big.Int) (event.Subscription, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "MilestoneAmended", milestoneIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignMilestoneAmended)
				if err := _Campaign.contract.UnpackLog(event, "MilestoneAmended", log); err != nil {
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

// ParseMilestoneAmended is a log parse operation binding the contract event 0xf9b5853f7d0dc993174e3c4e9c2e16d05af526a8cb108c33d9edd565df408a8f.
//
// Solidity: event MilestoneAmended(uint256 indexed milestoneIndex, bytes32 amendmentHash, uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseMilestoneAmended(log types.Log) (*CampaignMilestoneAmended, error) {
	event := new(CampaignMilestoneAmended)
	if err := _Campaign.contract.UnpackLog(event, "MilestoneAmended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignMilestoneApprovedIterator is returned from FilterMilestoneApproved and is used to iterate over the raw logs and unpacked data for MilestoneApproved events raised by the Campaign contract.
type CampaignMilestoneApprovedIterator struct {
	Event *CampaignMilestoneApproved // Event containing the contract specifics and raw log

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
func (it *CampaignMilestoneApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignMilestoneApproved)
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
		it.Event = new(CampaignMilestoneApproved)
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
func (it *CampaignMilestoneApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignMilestoneApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignMilestoneApproved represents a MilestoneApproved event raised by the Campaign contract.
type CampaignMilestoneApproved struct {
	MilestoneIndex *big.Int
	Approver       common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMilestoneApproved is a free log retrieval operation binding the contract event 0xfe731b8534f38a55c98725a977efe67da793f35fb32ca4d1d947c01d80259bc2.
//
// Solidity: event MilestoneApproved(uint256 indexed milestoneIndex, address indexed approver, uint256 timestamp)
func (_Campaign *CampaignFilterer) FilterMilestoneApproved(opts *bind.FilterOpts, milestoneIndex []*big.Int, approver []common.Address) (*CampaignMilestoneApprovedIterator, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Campaign.contract.FilterLogs(opts, "MilestoneApproved", milestoneIndexRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &CampaignMilestoneApprovedIterator{contract: _Campaign.contract, event: "MilestoneApproved", logs: logs, sub: sub}, nil
}

// WatchMilestoneApproved is a free log subscription operation binding the contract event 0xfe731b8534f38a55c98725a977efe67da793f35fb32ca4d1d947c01d80259bc2.
//
// Solidity: event MilestoneApproved(uint256 indexed milestoneIndex, address indexed approver, uint256 timestamp)
func (_Campaign *CampaignFilterer) WatchMilestoneApproved(opts *bind.WatchOpts, sink chan<- *CampaignMilestoneApproved, milestoneIndex []*big.Int, approver []common.Address) (event.Subscription, error) {

	var milestoneIndexRule []interface{}
	for _, milestoneIndexItem := range milestoneIndex {
		milestoneIndexRule = append(milestoneIndexRule, milestoneIndexItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Campaign.contract.WatchLogs(opts, "MilestoneApproved", milestoneIndexRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignMilestoneApproved)
				if err := _Campaign.contract.UnpackLog(event, "MilestoneApproved", log); err != nil {
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

// ParseMilestoneApproved is a log parse operation binding the contract event 0xfe731b8534f38a55c98725a977efe67da793f35fb32ca4d1d947c01d80259bc2.
//
// Solidity: event MilestoneApproved(uint256 indexed milestoneIndex, address indexed approver, uint256 timestamp)
func (_Campaign *CampaignFilterer) ParseMilestoneApproved(log types.Log) (*CampaignMilestoneApproved, error) {
	event := new(CampaignMilestoneApproved)
	if err := _Campaign.contract.UnpackLog(event, "MilestoneApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
