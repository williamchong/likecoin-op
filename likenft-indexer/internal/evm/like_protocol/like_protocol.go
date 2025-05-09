// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package like_protocol

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

// BookConfig is an auto generated low-level Go binding around an user-defined struct.
type BookConfig struct {
	Name      string
	Symbol    string
	Metadata  string
	MaxSupply uint64
}

// MsgNewBookNFT is an auto generated low-level Go binding around an user-defined struct.
type MsgNewBookNFT struct {
	Creator  common.Address
	Updaters []common.Address
	Minters  []common.Address
	Config   BookConfig
}

// LikeProtocolMetaData contains all meta data concerning the LikeProtocol contract.
var LikeProtocolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"BookNFTInvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"BookNFTImplementationUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"NewBookNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookNFTImplementation\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"classId\",\"type\":\"address\"}],\"name\":\"isBookNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewBookNFT\",\"name\":\"msgNewBookNFT\",\"type\":\"tuple\"}],\"name\":\"newBookNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bookAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewBookNFT\",\"name\":\"msgNewBookNFT\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"royaltyFraction\",\"type\":\"uint96\"}],\"name\":\"newBookNFTWithRoyalty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bookAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewBookNFT[]\",\"name\":\"msgNewBookNFTs\",\"type\":\"tuple[]\"}],\"name\":\"newBookNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LikeProtocolABI is the input ABI used to generate the binding from.
// Deprecated: Use LikeProtocolMetaData.ABI instead.
var LikeProtocolABI = LikeProtocolMetaData.ABI

// LikeProtocol is an auto generated Go binding around an Ethereum contract.
type LikeProtocol struct {
	LikeProtocolCaller     // Read-only binding to the contract
	LikeProtocolTransactor // Write-only binding to the contract
	LikeProtocolFilterer   // Log filterer for contract events
}

// LikeProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikeProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikeProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikeProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikeProtocolSession struct {
	Contract     *LikeProtocol     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LikeProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikeProtocolCallerSession struct {
	Contract *LikeProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LikeProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikeProtocolTransactorSession struct {
	Contract     *LikeProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LikeProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikeProtocolRaw struct {
	Contract *LikeProtocol // Generic contract binding to access the raw methods on
}

// LikeProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikeProtocolCallerRaw struct {
	Contract *LikeProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// LikeProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikeProtocolTransactorRaw struct {
	Contract *LikeProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikeProtocol creates a new instance of LikeProtocol, bound to a specific deployed contract.
func NewLikeProtocol(address common.Address, backend bind.ContractBackend) (*LikeProtocol, error) {
	contract, err := bindLikeProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikeProtocol{LikeProtocolCaller: LikeProtocolCaller{contract: contract}, LikeProtocolTransactor: LikeProtocolTransactor{contract: contract}, LikeProtocolFilterer: LikeProtocolFilterer{contract: contract}}, nil
}

// NewLikeProtocolCaller creates a new read-only instance of LikeProtocol, bound to a specific deployed contract.
func NewLikeProtocolCaller(address common.Address, caller bind.ContractCaller) (*LikeProtocolCaller, error) {
	contract, err := bindLikeProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikeProtocolCaller{contract: contract}, nil
}

// NewLikeProtocolTransactor creates a new write-only instance of LikeProtocol, bound to a specific deployed contract.
func NewLikeProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*LikeProtocolTransactor, error) {
	contract, err := bindLikeProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikeProtocolTransactor{contract: contract}, nil
}

// NewLikeProtocolFilterer creates a new log filterer instance of LikeProtocol, bound to a specific deployed contract.
func NewLikeProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*LikeProtocolFilterer, error) {
	contract, err := bindLikeProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikeProtocolFilterer{contract: contract}, nil
}

// bindLikeProtocol binds a generic wrapper to an already deployed contract.
func bindLikeProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikeProtocolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeProtocol *LikeProtocolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeProtocol.Contract.LikeProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeProtocol *LikeProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeProtocol.Contract.LikeProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeProtocol *LikeProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeProtocol.Contract.LikeProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeProtocol *LikeProtocolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeProtocol *LikeProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeProtocol *LikeProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeProtocol.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeProtocol *LikeProtocolCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeProtocol *LikeProtocolSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeProtocol.Contract.UPGRADEINTERFACEVERSION(&_LikeProtocol.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeProtocol *LikeProtocolCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeProtocol.Contract.UPGRADEINTERFACEVERSION(&_LikeProtocol.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LikeProtocol *LikeProtocolCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LikeProtocol *LikeProtocolSession) Implementation() (common.Address, error) {
	return _LikeProtocol.Contract.Implementation(&_LikeProtocol.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LikeProtocol *LikeProtocolCallerSession) Implementation() (common.Address, error) {
	return _LikeProtocol.Contract.Implementation(&_LikeProtocol.CallOpts)
}

// IsBookNFT is a free data retrieval call binding the contract method 0xd2380f1d.
//
// Solidity: function isBookNFT(address classId) view returns(bool)
func (_LikeProtocol *LikeProtocolCaller) IsBookNFT(opts *bind.CallOpts, classId common.Address) (bool, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "isBookNFT", classId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBookNFT is a free data retrieval call binding the contract method 0xd2380f1d.
//
// Solidity: function isBookNFT(address classId) view returns(bool)
func (_LikeProtocol *LikeProtocolSession) IsBookNFT(classId common.Address) (bool, error) {
	return _LikeProtocol.Contract.IsBookNFT(&_LikeProtocol.CallOpts, classId)
}

// IsBookNFT is a free data retrieval call binding the contract method 0xd2380f1d.
//
// Solidity: function isBookNFT(address classId) view returns(bool)
func (_LikeProtocol *LikeProtocolCallerSession) IsBookNFT(classId common.Address) (bool, error) {
	return _LikeProtocol.Contract.IsBookNFT(&_LikeProtocol.CallOpts, classId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeProtocol *LikeProtocolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeProtocol *LikeProtocolSession) Owner() (common.Address, error) {
	return _LikeProtocol.Contract.Owner(&_LikeProtocol.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeProtocol *LikeProtocolCallerSession) Owner() (common.Address, error) {
	return _LikeProtocol.Contract.Owner(&_LikeProtocol.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeProtocol *LikeProtocolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeProtocol *LikeProtocolSession) Paused() (bool, error) {
	return _LikeProtocol.Contract.Paused(&_LikeProtocol.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeProtocol *LikeProtocolCallerSession) Paused() (bool, error) {
	return _LikeProtocol.Contract.Paused(&_LikeProtocol.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeProtocol *LikeProtocolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikeProtocol.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeProtocol *LikeProtocolSession) ProxiableUUID() ([32]byte, error) {
	return _LikeProtocol.Contract.ProxiableUUID(&_LikeProtocol.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeProtocol *LikeProtocolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LikeProtocol.Contract.ProxiableUUID(&_LikeProtocol.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address bookNFTImplementation) returns()
func (_LikeProtocol *LikeProtocolTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, bookNFTImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "initialize", initialOwner, bookNFTImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address bookNFTImplementation) returns()
func (_LikeProtocol *LikeProtocolSession) Initialize(initialOwner common.Address, bookNFTImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.Initialize(&_LikeProtocol.TransactOpts, initialOwner, bookNFTImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address bookNFTImplementation) returns()
func (_LikeProtocol *LikeProtocolTransactorSession) Initialize(initialOwner common.Address, bookNFTImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.Initialize(&_LikeProtocol.TransactOpts, initialOwner, bookNFTImplementation)
}

// NewBookNFT is a paid mutator transaction binding the contract method 0xa43e8cfa.
//
// Solidity: function newBookNFT((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolTransactor) NewBookNFT(opts *bind.TransactOpts, msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "newBookNFT", msgNewBookNFT)
}

// NewBookNFT is a paid mutator transaction binding the contract method 0xa43e8cfa.
//
// Solidity: function newBookNFT((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolSession) NewBookNFT(msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFT(&_LikeProtocol.TransactOpts, msgNewBookNFT)
}

// NewBookNFT is a paid mutator transaction binding the contract method 0xa43e8cfa.
//
// Solidity: function newBookNFT((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolTransactorSession) NewBookNFT(msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFT(&_LikeProtocol.TransactOpts, msgNewBookNFT)
}

// NewBookNFTWithRoyalty is a paid mutator transaction binding the contract method 0x498a1c33.
//
// Solidity: function newBookNFTWithRoyalty((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT, uint96 royaltyFraction) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolTransactor) NewBookNFTWithRoyalty(opts *bind.TransactOpts, msgNewBookNFT MsgNewBookNFT, royaltyFraction *big.Int) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "newBookNFTWithRoyalty", msgNewBookNFT, royaltyFraction)
}

// NewBookNFTWithRoyalty is a paid mutator transaction binding the contract method 0x498a1c33.
//
// Solidity: function newBookNFTWithRoyalty((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT, uint96 royaltyFraction) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolSession) NewBookNFTWithRoyalty(msgNewBookNFT MsgNewBookNFT, royaltyFraction *big.Int) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFTWithRoyalty(&_LikeProtocol.TransactOpts, msgNewBookNFT, royaltyFraction)
}

// NewBookNFTWithRoyalty is a paid mutator transaction binding the contract method 0x498a1c33.
//
// Solidity: function newBookNFTWithRoyalty((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT, uint96 royaltyFraction) returns(address bookAddress)
func (_LikeProtocol *LikeProtocolTransactorSession) NewBookNFTWithRoyalty(msgNewBookNFT MsgNewBookNFT, royaltyFraction *big.Int) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFTWithRoyalty(&_LikeProtocol.TransactOpts, msgNewBookNFT, royaltyFraction)
}

// NewBookNFTs is a paid mutator transaction binding the contract method 0xd274a43f.
//
// Solidity: function newBookNFTs((address,address[],address[],(string,string,string,uint64))[] msgNewBookNFTs) returns()
func (_LikeProtocol *LikeProtocolTransactor) NewBookNFTs(opts *bind.TransactOpts, msgNewBookNFTs []MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "newBookNFTs", msgNewBookNFTs)
}

// NewBookNFTs is a paid mutator transaction binding the contract method 0xd274a43f.
//
// Solidity: function newBookNFTs((address,address[],address[],(string,string,string,uint64))[] msgNewBookNFTs) returns()
func (_LikeProtocol *LikeProtocolSession) NewBookNFTs(msgNewBookNFTs []MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFTs(&_LikeProtocol.TransactOpts, msgNewBookNFTs)
}

// NewBookNFTs is a paid mutator transaction binding the contract method 0xd274a43f.
//
// Solidity: function newBookNFTs((address,address[],address[],(string,string,string,uint64))[] msgNewBookNFTs) returns()
func (_LikeProtocol *LikeProtocolTransactorSession) NewBookNFTs(msgNewBookNFTs []MsgNewBookNFT) (*types.Transaction, error) {
	return _LikeProtocol.Contract.NewBookNFTs(&_LikeProtocol.TransactOpts, msgNewBookNFTs)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeProtocol *LikeProtocolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeProtocol *LikeProtocolSession) Pause() (*types.Transaction, error) {
	return _LikeProtocol.Contract.Pause(&_LikeProtocol.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeProtocol *LikeProtocolTransactorSession) Pause() (*types.Transaction, error) {
	return _LikeProtocol.Contract.Pause(&_LikeProtocol.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeProtocol *LikeProtocolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeProtocol *LikeProtocolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeProtocol.Contract.RenounceOwnership(&_LikeProtocol.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeProtocol *LikeProtocolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeProtocol.Contract.RenounceOwnership(&_LikeProtocol.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeProtocol *LikeProtocolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeProtocol *LikeProtocolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.TransferOwnership(&_LikeProtocol.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeProtocol *LikeProtocolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.TransferOwnership(&_LikeProtocol.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeProtocol *LikeProtocolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeProtocol *LikeProtocolSession) Unpause() (*types.Transaction, error) {
	return _LikeProtocol.Contract.Unpause(&_LikeProtocol.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeProtocol *LikeProtocolTransactorSession) Unpause() (*types.Transaction, error) {
	return _LikeProtocol.Contract.Unpause(&_LikeProtocol.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LikeProtocol *LikeProtocolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LikeProtocol *LikeProtocolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.UpgradeTo(&_LikeProtocol.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LikeProtocol *LikeProtocolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LikeProtocol.Contract.UpgradeTo(&_LikeProtocol.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeProtocol *LikeProtocolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeProtocol.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeProtocol *LikeProtocolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeProtocol.Contract.UpgradeToAndCall(&_LikeProtocol.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeProtocol *LikeProtocolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeProtocol.Contract.UpgradeToAndCall(&_LikeProtocol.TransactOpts, newImplementation, data)
}

// LikeProtocolBookNFTImplementationUpgradedIterator is returned from FilterBookNFTImplementationUpgraded and is used to iterate over the raw logs and unpacked data for BookNFTImplementationUpgraded events raised by the LikeProtocol contract.
type LikeProtocolBookNFTImplementationUpgradedIterator struct {
	Event *LikeProtocolBookNFTImplementationUpgraded // Event containing the contract specifics and raw log

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
func (it *LikeProtocolBookNFTImplementationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolBookNFTImplementationUpgraded)
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
		it.Event = new(LikeProtocolBookNFTImplementationUpgraded)
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
func (it *LikeProtocolBookNFTImplementationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolBookNFTImplementationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolBookNFTImplementationUpgraded represents a BookNFTImplementationUpgraded event raised by the LikeProtocol contract.
type LikeProtocolBookNFTImplementationUpgraded struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBookNFTImplementationUpgraded is a free log retrieval operation binding the contract event 0x079b6b58f11e5758083fb4e9adbfbee490af256ccca51426658d4349450125f1.
//
// Solidity: event BookNFTImplementationUpgraded(address newImplementation)
func (_LikeProtocol *LikeProtocolFilterer) FilterBookNFTImplementationUpgraded(opts *bind.FilterOpts) (*LikeProtocolBookNFTImplementationUpgradedIterator, error) {

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "BookNFTImplementationUpgraded")
	if err != nil {
		return nil, err
	}
	return &LikeProtocolBookNFTImplementationUpgradedIterator{contract: _LikeProtocol.contract, event: "BookNFTImplementationUpgraded", logs: logs, sub: sub}, nil
}

// WatchBookNFTImplementationUpgraded is a free log subscription operation binding the contract event 0x079b6b58f11e5758083fb4e9adbfbee490af256ccca51426658d4349450125f1.
//
// Solidity: event BookNFTImplementationUpgraded(address newImplementation)
func (_LikeProtocol *LikeProtocolFilterer) WatchBookNFTImplementationUpgraded(opts *bind.WatchOpts, sink chan<- *LikeProtocolBookNFTImplementationUpgraded) (event.Subscription, error) {

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "BookNFTImplementationUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolBookNFTImplementationUpgraded)
				if err := _LikeProtocol.contract.UnpackLog(event, "BookNFTImplementationUpgraded", log); err != nil {
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

// ParseBookNFTImplementationUpgraded is a log parse operation binding the contract event 0x079b6b58f11e5758083fb4e9adbfbee490af256ccca51426658d4349450125f1.
//
// Solidity: event BookNFTImplementationUpgraded(address newImplementation)
func (_LikeProtocol *LikeProtocolFilterer) ParseBookNFTImplementationUpgraded(log types.Log) (*LikeProtocolBookNFTImplementationUpgraded, error) {
	event := new(LikeProtocolBookNFTImplementationUpgraded)
	if err := _LikeProtocol.contract.UnpackLog(event, "BookNFTImplementationUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LikeProtocol contract.
type LikeProtocolInitializedIterator struct {
	Event *LikeProtocolInitialized // Event containing the contract specifics and raw log

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
func (it *LikeProtocolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolInitialized)
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
		it.Event = new(LikeProtocolInitialized)
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
func (it *LikeProtocolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolInitialized represents a Initialized event raised by the LikeProtocol contract.
type LikeProtocolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeProtocol *LikeProtocolFilterer) FilterInitialized(opts *bind.FilterOpts) (*LikeProtocolInitializedIterator, error) {

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LikeProtocolInitializedIterator{contract: _LikeProtocol.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeProtocol *LikeProtocolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LikeProtocolInitialized) (event.Subscription, error) {

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolInitialized)
				if err := _LikeProtocol.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeProtocol *LikeProtocolFilterer) ParseInitialized(log types.Log) (*LikeProtocolInitialized, error) {
	event := new(LikeProtocolInitialized)
	if err := _LikeProtocol.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolNewBookNFTIterator is returned from FilterNewBookNFT and is used to iterate over the raw logs and unpacked data for NewBookNFT events raised by the LikeProtocol contract.
type LikeProtocolNewBookNFTIterator struct {
	Event *LikeProtocolNewBookNFT // Event containing the contract specifics and raw log

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
func (it *LikeProtocolNewBookNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolNewBookNFT)
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
		it.Event = new(LikeProtocolNewBookNFT)
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
func (it *LikeProtocolNewBookNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolNewBookNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolNewBookNFT represents a NewBookNFT event raised by the LikeProtocol contract.
type LikeProtocolNewBookNFT struct {
	BookNFT common.Address
	Config  BookConfig
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBookNFT is a free log retrieval operation binding the contract event 0xac1baa76250109980b8de5e2b9fcb185acd2bd5127c85c3e83cc05fb3aae5df8.
//
// Solidity: event NewBookNFT(address bookNFT, (string,string,string,uint64) config)
func (_LikeProtocol *LikeProtocolFilterer) FilterNewBookNFT(opts *bind.FilterOpts) (*LikeProtocolNewBookNFTIterator, error) {

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "NewBookNFT")
	if err != nil {
		return nil, err
	}
	return &LikeProtocolNewBookNFTIterator{contract: _LikeProtocol.contract, event: "NewBookNFT", logs: logs, sub: sub}, nil
}

// WatchNewBookNFT is a free log subscription operation binding the contract event 0xac1baa76250109980b8de5e2b9fcb185acd2bd5127c85c3e83cc05fb3aae5df8.
//
// Solidity: event NewBookNFT(address bookNFT, (string,string,string,uint64) config)
func (_LikeProtocol *LikeProtocolFilterer) WatchNewBookNFT(opts *bind.WatchOpts, sink chan<- *LikeProtocolNewBookNFT) (event.Subscription, error) {

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "NewBookNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolNewBookNFT)
				if err := _LikeProtocol.contract.UnpackLog(event, "NewBookNFT", log); err != nil {
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

// ParseNewBookNFT is a log parse operation binding the contract event 0xac1baa76250109980b8de5e2b9fcb185acd2bd5127c85c3e83cc05fb3aae5df8.
//
// Solidity: event NewBookNFT(address bookNFT, (string,string,string,uint64) config)
func (_LikeProtocol *LikeProtocolFilterer) ParseNewBookNFT(log types.Log) (*LikeProtocolNewBookNFT, error) {
	event := new(LikeProtocolNewBookNFT)
	if err := _LikeProtocol.contract.UnpackLog(event, "NewBookNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LikeProtocol contract.
type LikeProtocolOwnershipTransferredIterator struct {
	Event *LikeProtocolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LikeProtocolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolOwnershipTransferred)
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
		it.Event = new(LikeProtocolOwnershipTransferred)
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
func (it *LikeProtocolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolOwnershipTransferred represents a OwnershipTransferred event raised by the LikeProtocol contract.
type LikeProtocolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeProtocol *LikeProtocolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LikeProtocolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LikeProtocolOwnershipTransferredIterator{contract: _LikeProtocol.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeProtocol *LikeProtocolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LikeProtocolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolOwnershipTransferred)
				if err := _LikeProtocol.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LikeProtocol *LikeProtocolFilterer) ParseOwnershipTransferred(log types.Log) (*LikeProtocolOwnershipTransferred, error) {
	event := new(LikeProtocolOwnershipTransferred)
	if err := _LikeProtocol.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LikeProtocol contract.
type LikeProtocolPausedIterator struct {
	Event *LikeProtocolPaused // Event containing the contract specifics and raw log

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
func (it *LikeProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolPaused)
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
		it.Event = new(LikeProtocolPaused)
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
func (it *LikeProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolPaused represents a Paused event raised by the LikeProtocol contract.
type LikeProtocolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeProtocol *LikeProtocolFilterer) FilterPaused(opts *bind.FilterOpts) (*LikeProtocolPausedIterator, error) {

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LikeProtocolPausedIterator{contract: _LikeProtocol.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeProtocol *LikeProtocolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LikeProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolPaused)
				if err := _LikeProtocol.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeProtocol *LikeProtocolFilterer) ParsePaused(log types.Log) (*LikeProtocolPaused, error) {
	event := new(LikeProtocolPaused)
	if err := _LikeProtocol.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LikeProtocol contract.
type LikeProtocolUnpausedIterator struct {
	Event *LikeProtocolUnpaused // Event containing the contract specifics and raw log

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
func (it *LikeProtocolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolUnpaused)
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
		it.Event = new(LikeProtocolUnpaused)
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
func (it *LikeProtocolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolUnpaused represents a Unpaused event raised by the LikeProtocol contract.
type LikeProtocolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeProtocol *LikeProtocolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LikeProtocolUnpausedIterator, error) {

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LikeProtocolUnpausedIterator{contract: _LikeProtocol.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeProtocol *LikeProtocolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LikeProtocolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolUnpaused)
				if err := _LikeProtocol.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeProtocol *LikeProtocolFilterer) ParseUnpaused(log types.Log) (*LikeProtocolUnpaused, error) {
	event := new(LikeProtocolUnpaused)
	if err := _LikeProtocol.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeProtocolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LikeProtocol contract.
type LikeProtocolUpgradedIterator struct {
	Event *LikeProtocolUpgraded // Event containing the contract specifics and raw log

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
func (it *LikeProtocolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeProtocolUpgraded)
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
		it.Event = new(LikeProtocolUpgraded)
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
func (it *LikeProtocolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeProtocolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeProtocolUpgraded represents a Upgraded event raised by the LikeProtocol contract.
type LikeProtocolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeProtocol *LikeProtocolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LikeProtocolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeProtocol.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LikeProtocolUpgradedIterator{contract: _LikeProtocol.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeProtocol *LikeProtocolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LikeProtocolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeProtocol.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeProtocolUpgraded)
				if err := _LikeProtocol.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeProtocol *LikeProtocolFilterer) ParseUpgraded(log types.Log) (*LikeProtocolUpgraded, error) {
	event := new(LikeProtocolUpgraded)
	if err := _LikeProtocol.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
