// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package like_collective

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

// LikeCollectiveRewardData is an auto generated low-level Go binding around an user-defined struct.
type LikeCollectiveRewardData struct {
	BookNFT        common.Address
	RewardedAmount *big.Int
}

// LikeCollectiveMetaData contains all meta data concerning the LikeCollective contract.
var LikeCollectiveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ErrInsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"ErrInvalidBookNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoRewardsToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardedAmount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLikeCollective.RewardData[]\",\"name\":\"rewardedAmount\",\"type\":\"tuple[]\"}],\"name\":\"AllRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardedAmount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardedAmount\",\"type\":\"uint256\"}],\"name\":\"RewardDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACC_REWARD_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimAllRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseStakePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"getPendingRewardsForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"getPendingRewardsPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRewardsOfPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"getStakeForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseStakeToPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"newStakePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeStakePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"restakeRewardPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"likeStakePosition\",\"type\":\"address\"}],\"name\":\"setLikeStakePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"likecoin\",\"type\":\"address\"}],\"name\":\"setLikecoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LikeCollectiveABI is the input ABI used to generate the binding from.
// Deprecated: Use LikeCollectiveMetaData.ABI instead.
var LikeCollectiveABI = LikeCollectiveMetaData.ABI

// LikeCollective is an auto generated Go binding around an Ethereum contract.
type LikeCollective struct {
	LikeCollectiveCaller     // Read-only binding to the contract
	LikeCollectiveTransactor // Write-only binding to the contract
	LikeCollectiveFilterer   // Log filterer for contract events
}

// LikeCollectiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikeCollectiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeCollectiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikeCollectiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeCollectiveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikeCollectiveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeCollectiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikeCollectiveSession struct {
	Contract     *LikeCollective   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LikeCollectiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikeCollectiveCallerSession struct {
	Contract *LikeCollectiveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LikeCollectiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikeCollectiveTransactorSession struct {
	Contract     *LikeCollectiveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LikeCollectiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikeCollectiveRaw struct {
	Contract *LikeCollective // Generic contract binding to access the raw methods on
}

// LikeCollectiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikeCollectiveCallerRaw struct {
	Contract *LikeCollectiveCaller // Generic read-only contract binding to access the raw methods on
}

// LikeCollectiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikeCollectiveTransactorRaw struct {
	Contract *LikeCollectiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikeCollective creates a new instance of LikeCollective, bound to a specific deployed contract.
func NewLikeCollective(address common.Address, backend bind.ContractBackend) (*LikeCollective, error) {
	contract, err := bindLikeCollective(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikeCollective{LikeCollectiveCaller: LikeCollectiveCaller{contract: contract}, LikeCollectiveTransactor: LikeCollectiveTransactor{contract: contract}, LikeCollectiveFilterer: LikeCollectiveFilterer{contract: contract}}, nil
}

// NewLikeCollectiveCaller creates a new read-only instance of LikeCollective, bound to a specific deployed contract.
func NewLikeCollectiveCaller(address common.Address, caller bind.ContractCaller) (*LikeCollectiveCaller, error) {
	contract, err := bindLikeCollective(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveCaller{contract: contract}, nil
}

// NewLikeCollectiveTransactor creates a new write-only instance of LikeCollective, bound to a specific deployed contract.
func NewLikeCollectiveTransactor(address common.Address, transactor bind.ContractTransactor) (*LikeCollectiveTransactor, error) {
	contract, err := bindLikeCollective(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveTransactor{contract: contract}, nil
}

// NewLikeCollectiveFilterer creates a new log filterer instance of LikeCollective, bound to a specific deployed contract.
func NewLikeCollectiveFilterer(address common.Address, filterer bind.ContractFilterer) (*LikeCollectiveFilterer, error) {
	contract, err := bindLikeCollective(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveFilterer{contract: contract}, nil
}

// bindLikeCollective binds a generic wrapper to an already deployed contract.
func bindLikeCollective(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikeCollectiveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeCollective *LikeCollectiveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeCollective.Contract.LikeCollectiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeCollective *LikeCollectiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeCollective.Contract.LikeCollectiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeCollective *LikeCollectiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeCollective.Contract.LikeCollectiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeCollective *LikeCollectiveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeCollective.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeCollective *LikeCollectiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeCollective.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeCollective *LikeCollectiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeCollective.Contract.contract.Transact(opts, method, params...)
}

// ACCREWARDPRECISION is a free data retrieval call binding the contract method 0xd1c6a231.
//
// Solidity: function ACC_REWARD_PRECISION() view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) ACCREWARDPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "ACC_REWARD_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ACCREWARDPRECISION is a free data retrieval call binding the contract method 0xd1c6a231.
//
// Solidity: function ACC_REWARD_PRECISION() view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) ACCREWARDPRECISION() (*big.Int, error) {
	return _LikeCollective.Contract.ACCREWARDPRECISION(&_LikeCollective.CallOpts)
}

// ACCREWARDPRECISION is a free data retrieval call binding the contract method 0xd1c6a231.
//
// Solidity: function ACC_REWARD_PRECISION() view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) ACCREWARDPRECISION() (*big.Int, error) {
	return _LikeCollective.Contract.ACCREWARDPRECISION(&_LikeCollective.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeCollective *LikeCollectiveCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeCollective *LikeCollectiveSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeCollective.Contract.UPGRADEINTERFACEVERSION(&_LikeCollective.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeCollective *LikeCollectiveCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeCollective.Contract.UPGRADEINTERFACEVERSION(&_LikeCollective.CallOpts)
}

// GetPendingRewardsForUser is a free data retrieval call binding the contract method 0xbb112623.
//
// Solidity: function getPendingRewardsForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) GetPendingRewardsForUser(opts *bind.CallOpts, user common.Address, bookNFT common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "getPendingRewardsForUser", user, bookNFT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewardsForUser is a free data retrieval call binding the contract method 0xbb112623.
//
// Solidity: function getPendingRewardsForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) GetPendingRewardsForUser(user common.Address, bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetPendingRewardsForUser(&_LikeCollective.CallOpts, user, bookNFT)
}

// GetPendingRewardsForUser is a free data retrieval call binding the contract method 0xbb112623.
//
// Solidity: function getPendingRewardsForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) GetPendingRewardsForUser(user common.Address, bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetPendingRewardsForUser(&_LikeCollective.CallOpts, user, bookNFT)
}

// GetPendingRewardsPool is a free data retrieval call binding the contract method 0x77eef8f9.
//
// Solidity: function getPendingRewardsPool(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) GetPendingRewardsPool(opts *bind.CallOpts, bookNFT common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "getPendingRewardsPool", bookNFT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewardsPool is a free data retrieval call binding the contract method 0x77eef8f9.
//
// Solidity: function getPendingRewardsPool(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) GetPendingRewardsPool(bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetPendingRewardsPool(&_LikeCollective.CallOpts, bookNFT)
}

// GetPendingRewardsPool is a free data retrieval call binding the contract method 0x77eef8f9.
//
// Solidity: function getPendingRewardsPool(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) GetPendingRewardsPool(bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetPendingRewardsPool(&_LikeCollective.CallOpts, bookNFT)
}

// GetRewardsOfPosition is a free data retrieval call binding the contract method 0xc86ec26e.
//
// Solidity: function getRewardsOfPosition(uint256 tokenId) view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) GetRewardsOfPosition(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "getRewardsOfPosition", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsOfPosition is a free data retrieval call binding the contract method 0xc86ec26e.
//
// Solidity: function getRewardsOfPosition(uint256 tokenId) view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) GetRewardsOfPosition(tokenId *big.Int) (*big.Int, error) {
	return _LikeCollective.Contract.GetRewardsOfPosition(&_LikeCollective.CallOpts, tokenId)
}

// GetRewardsOfPosition is a free data retrieval call binding the contract method 0xc86ec26e.
//
// Solidity: function getRewardsOfPosition(uint256 tokenId) view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) GetRewardsOfPosition(tokenId *big.Int) (*big.Int, error) {
	return _LikeCollective.Contract.GetRewardsOfPosition(&_LikeCollective.CallOpts, tokenId)
}

// GetStakeForUser is a free data retrieval call binding the contract method 0x4d9453c0.
//
// Solidity: function getStakeForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) GetStakeForUser(opts *bind.CallOpts, user common.Address, bookNFT common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "getStakeForUser", user, bookNFT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeForUser is a free data retrieval call binding the contract method 0x4d9453c0.
//
// Solidity: function getStakeForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) GetStakeForUser(user common.Address, bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetStakeForUser(&_LikeCollective.CallOpts, user, bookNFT)
}

// GetStakeForUser is a free data retrieval call binding the contract method 0x4d9453c0.
//
// Solidity: function getStakeForUser(address user, address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) GetStakeForUser(user common.Address, bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetStakeForUser(&_LikeCollective.CallOpts, user, bookNFT)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCaller) GetTotalStake(opts *bind.CallOpts, bookNFT common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "getTotalStake", bookNFT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveSession) GetTotalStake(bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetTotalStake(&_LikeCollective.CallOpts, bookNFT)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address bookNFT) view returns(uint256)
func (_LikeCollective *LikeCollectiveCallerSession) GetTotalStake(bookNFT common.Address) (*big.Int, error) {
	return _LikeCollective.Contract.GetTotalStake(&_LikeCollective.CallOpts, bookNFT)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeCollective *LikeCollectiveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeCollective *LikeCollectiveSession) Owner() (common.Address, error) {
	return _LikeCollective.Contract.Owner(&_LikeCollective.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeCollective *LikeCollectiveCallerSession) Owner() (common.Address, error) {
	return _LikeCollective.Contract.Owner(&_LikeCollective.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeCollective *LikeCollectiveCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeCollective *LikeCollectiveSession) Paused() (bool, error) {
	return _LikeCollective.Contract.Paused(&_LikeCollective.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeCollective *LikeCollectiveCallerSession) Paused() (bool, error) {
	return _LikeCollective.Contract.Paused(&_LikeCollective.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeCollective *LikeCollectiveCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikeCollective.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeCollective *LikeCollectiveSession) ProxiableUUID() ([32]byte, error) {
	return _LikeCollective.Contract.ProxiableUUID(&_LikeCollective.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeCollective *LikeCollectiveCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LikeCollective.Contract.ProxiableUUID(&_LikeCollective.CallOpts)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xe991560f.
//
// Solidity: function claimAllRewards(address user) returns()
func (_LikeCollective *LikeCollectiveTransactor) ClaimAllRewards(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "claimAllRewards", user)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xe991560f.
//
// Solidity: function claimAllRewards(address user) returns()
func (_LikeCollective *LikeCollectiveSession) ClaimAllRewards(user common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.ClaimAllRewards(&_LikeCollective.TransactOpts, user)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xe991560f.
//
// Solidity: function claimAllRewards(address user) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) ClaimAllRewards(user common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.ClaimAllRewards(&_LikeCollective.TransactOpts, user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactor) ClaimRewards(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "claimRewards", tokenId)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveSession) ClaimRewards(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.ClaimRewards(&_LikeCollective.TransactOpts, tokenId)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) ClaimRewards(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.ClaimRewards(&_LikeCollective.TransactOpts, tokenId)
}

// DecreaseStakePosition is a paid mutator transaction binding the contract method 0x0abe6d56.
//
// Solidity: function decreaseStakePosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactor) DecreaseStakePosition(opts *bind.TransactOpts, tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "decreaseStakePosition", tokenID, amount)
}

// DecreaseStakePosition is a paid mutator transaction binding the contract method 0x0abe6d56.
//
// Solidity: function decreaseStakePosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveSession) DecreaseStakePosition(tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.DecreaseStakePosition(&_LikeCollective.TransactOpts, tokenID, amount)
}

// DecreaseStakePosition is a paid mutator transaction binding the contract method 0x0abe6d56.
//
// Solidity: function decreaseStakePosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) DecreaseStakePosition(tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.DecreaseStakePosition(&_LikeCollective.TransactOpts, tokenID, amount)
}

// DepositReward is a paid mutator transaction binding the contract method 0x7db4e28f.
//
// Solidity: function depositReward(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactor) DepositReward(opts *bind.TransactOpts, bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "depositReward", bookNFT, amount)
}

// DepositReward is a paid mutator transaction binding the contract method 0x7db4e28f.
//
// Solidity: function depositReward(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveSession) DepositReward(bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.DepositReward(&_LikeCollective.TransactOpts, bookNFT, amount)
}

// DepositReward is a paid mutator transaction binding the contract method 0x7db4e28f.
//
// Solidity: function depositReward(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) DepositReward(bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.DepositReward(&_LikeCollective.TransactOpts, bookNFT, amount)
}

// IncreaseStakeToPosition is a paid mutator transaction binding the contract method 0x95a24715.
//
// Solidity: function increaseStakeToPosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactor) IncreaseStakeToPosition(opts *bind.TransactOpts, tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "increaseStakeToPosition", tokenID, amount)
}

// IncreaseStakeToPosition is a paid mutator transaction binding the contract method 0x95a24715.
//
// Solidity: function increaseStakeToPosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveSession) IncreaseStakeToPosition(tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.IncreaseStakeToPosition(&_LikeCollective.TransactOpts, tokenID, amount)
}

// IncreaseStakeToPosition is a paid mutator transaction binding the contract method 0x95a24715.
//
// Solidity: function increaseStakeToPosition(uint256 tokenID, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) IncreaseStakeToPosition(tokenID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.IncreaseStakeToPosition(&_LikeCollective.TransactOpts, tokenID, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeCollective *LikeCollectiveTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeCollective *LikeCollectiveSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.Initialize(&_LikeCollective.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.Initialize(&_LikeCollective.TransactOpts, initialOwner)
}

// NewStakePosition is a paid mutator transaction binding the contract method 0x5237c1b8.
//
// Solidity: function newStakePosition(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactor) NewStakePosition(opts *bind.TransactOpts, bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "newStakePosition", bookNFT, amount)
}

// NewStakePosition is a paid mutator transaction binding the contract method 0x5237c1b8.
//
// Solidity: function newStakePosition(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveSession) NewStakePosition(bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.NewStakePosition(&_LikeCollective.TransactOpts, bookNFT, amount)
}

// NewStakePosition is a paid mutator transaction binding the contract method 0x5237c1b8.
//
// Solidity: function newStakePosition(address bookNFT, uint256 amount) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) NewStakePosition(bookNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.NewStakePosition(&_LikeCollective.TransactOpts, bookNFT, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeCollective *LikeCollectiveTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeCollective *LikeCollectiveSession) Pause() (*types.Transaction, error) {
	return _LikeCollective.Contract.Pause(&_LikeCollective.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeCollective *LikeCollectiveTransactorSession) Pause() (*types.Transaction, error) {
	return _LikeCollective.Contract.Pause(&_LikeCollective.TransactOpts)
}

// RemoveStakePosition is a paid mutator transaction binding the contract method 0x2e2c55c6.
//
// Solidity: function removeStakePosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactor) RemoveStakePosition(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "removeStakePosition", tokenId)
}

// RemoveStakePosition is a paid mutator transaction binding the contract method 0x2e2c55c6.
//
// Solidity: function removeStakePosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveSession) RemoveStakePosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.RemoveStakePosition(&_LikeCollective.TransactOpts, tokenId)
}

// RemoveStakePosition is a paid mutator transaction binding the contract method 0x2e2c55c6.
//
// Solidity: function removeStakePosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) RemoveStakePosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.RemoveStakePosition(&_LikeCollective.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeCollective *LikeCollectiveTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeCollective *LikeCollectiveSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeCollective.Contract.RenounceOwnership(&_LikeCollective.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeCollective *LikeCollectiveTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeCollective.Contract.RenounceOwnership(&_LikeCollective.TransactOpts)
}

// RestakeRewardPosition is a paid mutator transaction binding the contract method 0xcd083e62.
//
// Solidity: function restakeRewardPosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactor) RestakeRewardPosition(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "restakeRewardPosition", tokenId)
}

// RestakeRewardPosition is a paid mutator transaction binding the contract method 0xcd083e62.
//
// Solidity: function restakeRewardPosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveSession) RestakeRewardPosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.RestakeRewardPosition(&_LikeCollective.TransactOpts, tokenId)
}

// RestakeRewardPosition is a paid mutator transaction binding the contract method 0xcd083e62.
//
// Solidity: function restakeRewardPosition(uint256 tokenId) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) RestakeRewardPosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeCollective.Contract.RestakeRewardPosition(&_LikeCollective.TransactOpts, tokenId)
}

// SetLikeStakePosition is a paid mutator transaction binding the contract method 0xd214a7ba.
//
// Solidity: function setLikeStakePosition(address likeStakePosition) returns()
func (_LikeCollective *LikeCollectiveTransactor) SetLikeStakePosition(opts *bind.TransactOpts, likeStakePosition common.Address) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "setLikeStakePosition", likeStakePosition)
}

// SetLikeStakePosition is a paid mutator transaction binding the contract method 0xd214a7ba.
//
// Solidity: function setLikeStakePosition(address likeStakePosition) returns()
func (_LikeCollective *LikeCollectiveSession) SetLikeStakePosition(likeStakePosition common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.SetLikeStakePosition(&_LikeCollective.TransactOpts, likeStakePosition)
}

// SetLikeStakePosition is a paid mutator transaction binding the contract method 0xd214a7ba.
//
// Solidity: function setLikeStakePosition(address likeStakePosition) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) SetLikeStakePosition(likeStakePosition common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.SetLikeStakePosition(&_LikeCollective.TransactOpts, likeStakePosition)
}

// SetLikecoin is a paid mutator transaction binding the contract method 0x72a3ae9d.
//
// Solidity: function setLikecoin(address likecoin) returns()
func (_LikeCollective *LikeCollectiveTransactor) SetLikecoin(opts *bind.TransactOpts, likecoin common.Address) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "setLikecoin", likecoin)
}

// SetLikecoin is a paid mutator transaction binding the contract method 0x72a3ae9d.
//
// Solidity: function setLikecoin(address likecoin) returns()
func (_LikeCollective *LikeCollectiveSession) SetLikecoin(likecoin common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.SetLikecoin(&_LikeCollective.TransactOpts, likecoin)
}

// SetLikecoin is a paid mutator transaction binding the contract method 0x72a3ae9d.
//
// Solidity: function setLikecoin(address likecoin) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) SetLikecoin(likecoin common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.SetLikecoin(&_LikeCollective.TransactOpts, likecoin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeCollective *LikeCollectiveTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeCollective *LikeCollectiveSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.TransferOwnership(&_LikeCollective.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeCollective *LikeCollectiveTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeCollective.Contract.TransferOwnership(&_LikeCollective.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeCollective *LikeCollectiveTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeCollective *LikeCollectiveSession) Unpause() (*types.Transaction, error) {
	return _LikeCollective.Contract.Unpause(&_LikeCollective.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeCollective *LikeCollectiveTransactorSession) Unpause() (*types.Transaction, error) {
	return _LikeCollective.Contract.Unpause(&_LikeCollective.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeCollective *LikeCollectiveTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeCollective.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeCollective *LikeCollectiveSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeCollective.Contract.UpgradeToAndCall(&_LikeCollective.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeCollective *LikeCollectiveTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeCollective.Contract.UpgradeToAndCall(&_LikeCollective.TransactOpts, newImplementation, data)
}

// LikeCollectiveAllRewardClaimedIterator is returned from FilterAllRewardClaimed and is used to iterate over the raw logs and unpacked data for AllRewardClaimed events raised by the LikeCollective contract.
type LikeCollectiveAllRewardClaimedIterator struct {
	Event *LikeCollectiveAllRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveAllRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveAllRewardClaimed)
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
		it.Event = new(LikeCollectiveAllRewardClaimed)
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
func (it *LikeCollectiveAllRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveAllRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveAllRewardClaimed represents a AllRewardClaimed event raised by the LikeCollective contract.
type LikeCollectiveAllRewardClaimed struct {
	Account        common.Address
	RewardedAmount []LikeCollectiveRewardData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllRewardClaimed is a free log retrieval operation binding the contract event 0x445449d0ab504db2855634394b8408e0b876b052ce974f14e16baa7cfe2a987c.
//
// Solidity: event AllRewardClaimed(address indexed account, (address,uint256)[] rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) FilterAllRewardClaimed(opts *bind.FilterOpts, account []common.Address) (*LikeCollectiveAllRewardClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "AllRewardClaimed", accountRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveAllRewardClaimedIterator{contract: _LikeCollective.contract, event: "AllRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchAllRewardClaimed is a free log subscription operation binding the contract event 0x445449d0ab504db2855634394b8408e0b876b052ce974f14e16baa7cfe2a987c.
//
// Solidity: event AllRewardClaimed(address indexed account, (address,uint256)[] rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) WatchAllRewardClaimed(opts *bind.WatchOpts, sink chan<- *LikeCollectiveAllRewardClaimed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "AllRewardClaimed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveAllRewardClaimed)
				if err := _LikeCollective.contract.UnpackLog(event, "AllRewardClaimed", log); err != nil {
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

// ParseAllRewardClaimed is a log parse operation binding the contract event 0x445449d0ab504db2855634394b8408e0b876b052ce974f14e16baa7cfe2a987c.
//
// Solidity: event AllRewardClaimed(address indexed account, (address,uint256)[] rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) ParseAllRewardClaimed(log types.Log) (*LikeCollectiveAllRewardClaimed, error) {
	event := new(LikeCollectiveAllRewardClaimed)
	if err := _LikeCollective.contract.UnpackLog(event, "AllRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LikeCollective contract.
type LikeCollectiveInitializedIterator struct {
	Event *LikeCollectiveInitialized // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveInitialized)
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
		it.Event = new(LikeCollectiveInitialized)
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
func (it *LikeCollectiveInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveInitialized represents a Initialized event raised by the LikeCollective contract.
type LikeCollectiveInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeCollective *LikeCollectiveFilterer) FilterInitialized(opts *bind.FilterOpts) (*LikeCollectiveInitializedIterator, error) {

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveInitializedIterator{contract: _LikeCollective.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeCollective *LikeCollectiveFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LikeCollectiveInitialized) (event.Subscription, error) {

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveInitialized)
				if err := _LikeCollective.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LikeCollective *LikeCollectiveFilterer) ParseInitialized(log types.Log) (*LikeCollectiveInitialized, error) {
	event := new(LikeCollectiveInitialized)
	if err := _LikeCollective.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LikeCollective contract.
type LikeCollectiveOwnershipTransferredIterator struct {
	Event *LikeCollectiveOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveOwnershipTransferred)
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
		it.Event = new(LikeCollectiveOwnershipTransferred)
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
func (it *LikeCollectiveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveOwnershipTransferred represents a OwnershipTransferred event raised by the LikeCollective contract.
type LikeCollectiveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeCollective *LikeCollectiveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LikeCollectiveOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveOwnershipTransferredIterator{contract: _LikeCollective.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeCollective *LikeCollectiveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LikeCollectiveOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveOwnershipTransferred)
				if err := _LikeCollective.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LikeCollective *LikeCollectiveFilterer) ParseOwnershipTransferred(log types.Log) (*LikeCollectiveOwnershipTransferred, error) {
	event := new(LikeCollectiveOwnershipTransferred)
	if err := _LikeCollective.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectivePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LikeCollective contract.
type LikeCollectivePausedIterator struct {
	Event *LikeCollectivePaused // Event containing the contract specifics and raw log

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
func (it *LikeCollectivePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectivePaused)
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
		it.Event = new(LikeCollectivePaused)
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
func (it *LikeCollectivePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectivePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectivePaused represents a Paused event raised by the LikeCollective contract.
type LikeCollectivePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeCollective *LikeCollectiveFilterer) FilterPaused(opts *bind.FilterOpts) (*LikeCollectivePausedIterator, error) {

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LikeCollectivePausedIterator{contract: _LikeCollective.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeCollective *LikeCollectiveFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LikeCollectivePaused) (event.Subscription, error) {

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectivePaused)
				if err := _LikeCollective.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LikeCollective *LikeCollectiveFilterer) ParsePaused(log types.Log) (*LikeCollectivePaused, error) {
	event := new(LikeCollectivePaused)
	if err := _LikeCollective.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the LikeCollective contract.
type LikeCollectiveRewardClaimedIterator struct {
	Event *LikeCollectiveRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveRewardClaimed)
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
		it.Event = new(LikeCollectiveRewardClaimed)
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
func (it *LikeCollectiveRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveRewardClaimed represents a RewardClaimed event raised by the LikeCollective contract.
type LikeCollectiveRewardClaimed struct {
	BookNFT        common.Address
	Account        common.Address
	RewardedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) FilterRewardClaimed(opts *bind.FilterOpts, bookNFT []common.Address, account []common.Address) (*LikeCollectiveRewardClaimedIterator, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "RewardClaimed", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveRewardClaimedIterator{contract: _LikeCollective.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *LikeCollectiveRewardClaimed, bookNFT []common.Address, account []common.Address) (event.Subscription, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "RewardClaimed", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveRewardClaimed)
				if err := _LikeCollective.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) ParseRewardClaimed(log types.Log) (*LikeCollectiveRewardClaimed, error) {
	event := new(LikeCollectiveRewardClaimed)
	if err := _LikeCollective.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveRewardDepositedIterator is returned from FilterRewardDeposited and is used to iterate over the raw logs and unpacked data for RewardDeposited events raised by the LikeCollective contract.
type LikeCollectiveRewardDepositedIterator struct {
	Event *LikeCollectiveRewardDeposited // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveRewardDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveRewardDeposited)
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
		it.Event = new(LikeCollectiveRewardDeposited)
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
func (it *LikeCollectiveRewardDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveRewardDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveRewardDeposited represents a RewardDeposited event raised by the LikeCollective contract.
type LikeCollectiveRewardDeposited struct {
	BookNFT        common.Address
	Account        common.Address
	RewardedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardDeposited is a free log retrieval operation binding the contract event 0xecc7fa2c9c08fa68f2c2942d64a9c92b602004694d2a128c2a034cff48990f51.
//
// Solidity: event RewardDeposited(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) FilterRewardDeposited(opts *bind.FilterOpts, bookNFT []common.Address, account []common.Address) (*LikeCollectiveRewardDepositedIterator, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "RewardDeposited", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveRewardDepositedIterator{contract: _LikeCollective.contract, event: "RewardDeposited", logs: logs, sub: sub}, nil
}

// WatchRewardDeposited is a free log subscription operation binding the contract event 0xecc7fa2c9c08fa68f2c2942d64a9c92b602004694d2a128c2a034cff48990f51.
//
// Solidity: event RewardDeposited(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) WatchRewardDeposited(opts *bind.WatchOpts, sink chan<- *LikeCollectiveRewardDeposited, bookNFT []common.Address, account []common.Address) (event.Subscription, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "RewardDeposited", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveRewardDeposited)
				if err := _LikeCollective.contract.UnpackLog(event, "RewardDeposited", log); err != nil {
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

// ParseRewardDeposited is a log parse operation binding the contract event 0xecc7fa2c9c08fa68f2c2942d64a9c92b602004694d2a128c2a034cff48990f51.
//
// Solidity: event RewardDeposited(address indexed bookNFT, address indexed account, uint256 rewardedAmount)
func (_LikeCollective *LikeCollectiveFilterer) ParseRewardDeposited(log types.Log) (*LikeCollectiveRewardDeposited, error) {
	event := new(LikeCollectiveRewardDeposited)
	if err := _LikeCollective.contract.UnpackLog(event, "RewardDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the LikeCollective contract.
type LikeCollectiveStakedIterator struct {
	Event *LikeCollectiveStaked // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveStaked)
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
		it.Event = new(LikeCollectiveStaked)
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
func (it *LikeCollectiveStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveStaked represents a Staked event raised by the LikeCollective contract.
type LikeCollectiveStaked struct {
	BookNFT      common.Address
	Account      common.Address
	StakedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) FilterStaked(opts *bind.FilterOpts, bookNFT []common.Address, account []common.Address) (*LikeCollectiveStakedIterator, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Staked", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveStakedIterator{contract: _LikeCollective.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LikeCollectiveStaked, bookNFT []common.Address, account []common.Address) (event.Subscription, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Staked", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveStaked)
				if err := _LikeCollective.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) ParseStaked(log types.Log) (*LikeCollectiveStaked, error) {
	event := new(LikeCollectiveStaked)
	if err := _LikeCollective.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LikeCollective contract.
type LikeCollectiveUnpausedIterator struct {
	Event *LikeCollectiveUnpaused // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveUnpaused)
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
		it.Event = new(LikeCollectiveUnpaused)
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
func (it *LikeCollectiveUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveUnpaused represents a Unpaused event raised by the LikeCollective contract.
type LikeCollectiveUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeCollective *LikeCollectiveFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LikeCollectiveUnpausedIterator, error) {

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveUnpausedIterator{contract: _LikeCollective.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeCollective *LikeCollectiveFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LikeCollectiveUnpaused) (event.Subscription, error) {

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveUnpaused)
				if err := _LikeCollective.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LikeCollective *LikeCollectiveFilterer) ParseUnpaused(log types.Log) (*LikeCollectiveUnpaused, error) {
	event := new(LikeCollectiveUnpaused)
	if err := _LikeCollective.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the LikeCollective contract.
type LikeCollectiveUnstakedIterator struct {
	Event *LikeCollectiveUnstaked // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveUnstaked)
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
		it.Event = new(LikeCollectiveUnstaked)
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
func (it *LikeCollectiveUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveUnstaked represents a Unstaked event raised by the LikeCollective contract.
type LikeCollectiveUnstaked struct {
	BookNFT      common.Address
	Account      common.Address
	StakedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) FilterUnstaked(opts *bind.FilterOpts, bookNFT []common.Address, account []common.Address) (*LikeCollectiveUnstakedIterator, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Unstaked", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveUnstakedIterator{contract: _LikeCollective.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *LikeCollectiveUnstaked, bookNFT []common.Address, account []common.Address) (event.Subscription, error) {

	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Unstaked", bookNFTRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveUnstaked)
				if err := _LikeCollective.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed bookNFT, address indexed account, uint256 stakedAmount)
func (_LikeCollective *LikeCollectiveFilterer) ParseUnstaked(log types.Log) (*LikeCollectiveUnstaked, error) {
	event := new(LikeCollectiveUnstaked)
	if err := _LikeCollective.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeCollectiveUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LikeCollective contract.
type LikeCollectiveUpgradedIterator struct {
	Event *LikeCollectiveUpgraded // Event containing the contract specifics and raw log

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
func (it *LikeCollectiveUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeCollectiveUpgraded)
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
		it.Event = new(LikeCollectiveUpgraded)
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
func (it *LikeCollectiveUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeCollectiveUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeCollectiveUpgraded represents a Upgraded event raised by the LikeCollective contract.
type LikeCollectiveUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeCollective *LikeCollectiveFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LikeCollectiveUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeCollective.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LikeCollectiveUpgradedIterator{contract: _LikeCollective.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeCollective *LikeCollectiveFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LikeCollectiveUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeCollective.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeCollectiveUpgraded)
				if err := _LikeCollective.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LikeCollective *LikeCollectiveFilterer) ParseUpgraded(log types.Log) (*LikeCollectiveUpgraded, error) {
	event := new(LikeCollectiveUpgraded)
	if err := _LikeCollective.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
