// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package like_stake_position

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

// LikeStakePositionPosition is an auto generated low-level Go binding around an user-defined struct.
type LikeStakePositionPosition struct {
	BookNFT       common.Address
	StakedAmount  *big.Int
	RewardIndex   *big.Int
	InitialStaker common.Address
}

// LikeStakePositionMetaData contains all meta data concerning the LikeStakePosition contract.
var LikeStakePositionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PositionBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"}],\"name\":\"PositionMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"}],\"name\":\"PositionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialStaker\",\"type\":\"address\"}],\"internalType\":\"structLikeStakePosition.Position\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"}],\"name\":\"getUserPositionByBookNFT\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserPositions\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"}],\"name\":\"mintPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bookNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialStaker\",\"type\":\"address\"}],\"internalType\":\"structLikeStakePosition.Position\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newStakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRewardIndex\",\"type\":\"uint256\"}],\"name\":\"updatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRewardIndex\",\"type\":\"uint256\"}],\"name\":\"updatePositionRewardIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LikeStakePositionABI is the input ABI used to generate the binding from.
// Deprecated: Use LikeStakePositionMetaData.ABI instead.
var LikeStakePositionABI = LikeStakePositionMetaData.ABI

// LikeStakePosition is an auto generated Go binding around an Ethereum contract.
type LikeStakePosition struct {
	LikeStakePositionCaller     // Read-only binding to the contract
	LikeStakePositionTransactor // Write-only binding to the contract
	LikeStakePositionFilterer   // Log filterer for contract events
}

// LikeStakePositionCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikeStakePositionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeStakePositionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikeStakePositionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeStakePositionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikeStakePositionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeStakePositionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikeStakePositionSession struct {
	Contract     *LikeStakePosition // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LikeStakePositionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikeStakePositionCallerSession struct {
	Contract *LikeStakePositionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LikeStakePositionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikeStakePositionTransactorSession struct {
	Contract     *LikeStakePositionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LikeStakePositionRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikeStakePositionRaw struct {
	Contract *LikeStakePosition // Generic contract binding to access the raw methods on
}

// LikeStakePositionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikeStakePositionCallerRaw struct {
	Contract *LikeStakePositionCaller // Generic read-only contract binding to access the raw methods on
}

// LikeStakePositionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikeStakePositionTransactorRaw struct {
	Contract *LikeStakePositionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikeStakePosition creates a new instance of LikeStakePosition, bound to a specific deployed contract.
func NewLikeStakePosition(address common.Address, backend bind.ContractBackend) (*LikeStakePosition, error) {
	contract, err := bindLikeStakePosition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikeStakePosition{LikeStakePositionCaller: LikeStakePositionCaller{contract: contract}, LikeStakePositionTransactor: LikeStakePositionTransactor{contract: contract}, LikeStakePositionFilterer: LikeStakePositionFilterer{contract: contract}}, nil
}

// NewLikeStakePositionCaller creates a new read-only instance of LikeStakePosition, bound to a specific deployed contract.
func NewLikeStakePositionCaller(address common.Address, caller bind.ContractCaller) (*LikeStakePositionCaller, error) {
	contract, err := bindLikeStakePosition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionCaller{contract: contract}, nil
}

// NewLikeStakePositionTransactor creates a new write-only instance of LikeStakePosition, bound to a specific deployed contract.
func NewLikeStakePositionTransactor(address common.Address, transactor bind.ContractTransactor) (*LikeStakePositionTransactor, error) {
	contract, err := bindLikeStakePosition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionTransactor{contract: contract}, nil
}

// NewLikeStakePositionFilterer creates a new log filterer instance of LikeStakePosition, bound to a specific deployed contract.
func NewLikeStakePositionFilterer(address common.Address, filterer bind.ContractFilterer) (*LikeStakePositionFilterer, error) {
	contract, err := bindLikeStakePosition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionFilterer{contract: contract}, nil
}

// bindLikeStakePosition binds a generic wrapper to an already deployed contract.
func bindLikeStakePosition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikeStakePositionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeStakePosition *LikeStakePositionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeStakePosition.Contract.LikeStakePositionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeStakePosition *LikeStakePositionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.LikeStakePositionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeStakePosition *LikeStakePositionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.LikeStakePositionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeStakePosition *LikeStakePositionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeStakePosition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeStakePosition *LikeStakePositionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeStakePosition *LikeStakePositionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeStakePosition *LikeStakePositionCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeStakePosition *LikeStakePositionSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeStakePosition.Contract.UPGRADEINTERFACEVERSION(&_LikeStakePosition.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LikeStakePosition *LikeStakePositionCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LikeStakePosition.Contract.UPGRADEINTERFACEVERSION(&_LikeStakePosition.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikeStakePosition.Contract.BalanceOf(&_LikeStakePosition.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikeStakePosition.Contract.BalanceOf(&_LikeStakePosition.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikeStakePosition.Contract.GetApproved(&_LikeStakePosition.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikeStakePosition.Contract.GetApproved(&_LikeStakePosition.CallOpts, tokenId)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCaller) GetNextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "getNextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionSession) GetNextTokenId() (*big.Int, error) {
	return _LikeStakePosition.Contract.GetNextTokenId(&_LikeStakePosition.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCallerSession) GetNextTokenId() (*big.Int, error) {
	return _LikeStakePosition.Contract.GetNextTokenId(&_LikeStakePosition.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionCaller) GetPosition(opts *bind.CallOpts, tokenId *big.Int) (LikeStakePositionPosition, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "getPosition", tokenId)

	if err != nil {
		return *new(LikeStakePositionPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(LikeStakePositionPosition)).(*LikeStakePositionPosition)

	return out0, err

}

// GetPosition is a free data retrieval call binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionSession) GetPosition(tokenId *big.Int) (LikeStakePositionPosition, error) {
	return _LikeStakePosition.Contract.GetPosition(&_LikeStakePosition.CallOpts, tokenId)
}

// GetPosition is a free data retrieval call binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionCallerSession) GetPosition(tokenId *big.Int) (LikeStakePositionPosition, error) {
	return _LikeStakePosition.Contract.GetPosition(&_LikeStakePosition.CallOpts, tokenId)
}

// GetUserPositionByBookNFT is a free data retrieval call binding the contract method 0x0d998e5d.
//
// Solidity: function getUserPositionByBookNFT(address user, address bookNFT) view returns(uint256[])
func (_LikeStakePosition *LikeStakePositionCaller) GetUserPositionByBookNFT(opts *bind.CallOpts, user common.Address, bookNFT common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "getUserPositionByBookNFT", user, bookNFT)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserPositionByBookNFT is a free data retrieval call binding the contract method 0x0d998e5d.
//
// Solidity: function getUserPositionByBookNFT(address user, address bookNFT) view returns(uint256[])
func (_LikeStakePosition *LikeStakePositionSession) GetUserPositionByBookNFT(user common.Address, bookNFT common.Address) ([]*big.Int, error) {
	return _LikeStakePosition.Contract.GetUserPositionByBookNFT(&_LikeStakePosition.CallOpts, user, bookNFT)
}

// GetUserPositionByBookNFT is a free data retrieval call binding the contract method 0x0d998e5d.
//
// Solidity: function getUserPositionByBookNFT(address user, address bookNFT) view returns(uint256[])
func (_LikeStakePosition *LikeStakePositionCallerSession) GetUserPositionByBookNFT(user common.Address, bookNFT common.Address) ([]*big.Int, error) {
	return _LikeStakePosition.Contract.GetUserPositionByBookNFT(&_LikeStakePosition.CallOpts, user, bookNFT)
}

// GetUserPositions is a free data retrieval call binding the contract method 0x2a6bc2dd.
//
// Solidity: function getUserPositions(address user) view returns(uint256[] positions)
func (_LikeStakePosition *LikeStakePositionCaller) GetUserPositions(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "getUserPositions", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserPositions is a free data retrieval call binding the contract method 0x2a6bc2dd.
//
// Solidity: function getUserPositions(address user) view returns(uint256[] positions)
func (_LikeStakePosition *LikeStakePositionSession) GetUserPositions(user common.Address) ([]*big.Int, error) {
	return _LikeStakePosition.Contract.GetUserPositions(&_LikeStakePosition.CallOpts, user)
}

// GetUserPositions is a free data retrieval call binding the contract method 0x2a6bc2dd.
//
// Solidity: function getUserPositions(address user) view returns(uint256[] positions)
func (_LikeStakePosition *LikeStakePositionCallerSession) GetUserPositions(user common.Address) ([]*big.Int, error) {
	return _LikeStakePosition.Contract.GetUserPositions(&_LikeStakePosition.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikeStakePosition *LikeStakePositionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikeStakePosition *LikeStakePositionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikeStakePosition.Contract.IsApprovedForAll(&_LikeStakePosition.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikeStakePosition *LikeStakePositionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikeStakePosition.Contract.IsApprovedForAll(&_LikeStakePosition.CallOpts, owner, operator)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LikeStakePosition *LikeStakePositionCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LikeStakePosition *LikeStakePositionSession) Manager() (common.Address, error) {
	return _LikeStakePosition.Contract.Manager(&_LikeStakePosition.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_LikeStakePosition *LikeStakePositionCallerSession) Manager() (common.Address, error) {
	return _LikeStakePosition.Contract.Manager(&_LikeStakePosition.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeStakePosition *LikeStakePositionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeStakePosition *LikeStakePositionSession) Name() (string, error) {
	return _LikeStakePosition.Contract.Name(&_LikeStakePosition.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeStakePosition *LikeStakePositionCallerSession) Name() (string, error) {
	return _LikeStakePosition.Contract.Name(&_LikeStakePosition.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeStakePosition *LikeStakePositionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeStakePosition *LikeStakePositionSession) Owner() (common.Address, error) {
	return _LikeStakePosition.Contract.Owner(&_LikeStakePosition.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikeStakePosition *LikeStakePositionCallerSession) Owner() (common.Address, error) {
	return _LikeStakePosition.Contract.Owner(&_LikeStakePosition.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikeStakePosition.Contract.OwnerOf(&_LikeStakePosition.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikeStakePosition *LikeStakePositionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikeStakePosition.Contract.OwnerOf(&_LikeStakePosition.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeStakePosition *LikeStakePositionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeStakePosition *LikeStakePositionSession) Paused() (bool, error) {
	return _LikeStakePosition.Contract.Paused(&_LikeStakePosition.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LikeStakePosition *LikeStakePositionCallerSession) Paused() (bool, error) {
	return _LikeStakePosition.Contract.Paused(&_LikeStakePosition.CallOpts)
}

// PositionInfo is a free data retrieval call binding the contract method 0x89097a6a.
//
// Solidity: function positionInfo(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionCaller) PositionInfo(opts *bind.CallOpts, tokenId *big.Int) (LikeStakePositionPosition, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "positionInfo", tokenId)

	if err != nil {
		return *new(LikeStakePositionPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(LikeStakePositionPosition)).(*LikeStakePositionPosition)

	return out0, err

}

// PositionInfo is a free data retrieval call binding the contract method 0x89097a6a.
//
// Solidity: function positionInfo(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionSession) PositionInfo(tokenId *big.Int) (LikeStakePositionPosition, error) {
	return _LikeStakePosition.Contract.PositionInfo(&_LikeStakePosition.CallOpts, tokenId)
}

// PositionInfo is a free data retrieval call binding the contract method 0x89097a6a.
//
// Solidity: function positionInfo(uint256 tokenId) view returns((address,uint256,uint256,address))
func (_LikeStakePosition *LikeStakePositionCallerSession) PositionInfo(tokenId *big.Int) (LikeStakePositionPosition, error) {
	return _LikeStakePosition.Contract.PositionInfo(&_LikeStakePosition.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeStakePosition *LikeStakePositionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeStakePosition *LikeStakePositionSession) ProxiableUUID() ([32]byte, error) {
	return _LikeStakePosition.Contract.ProxiableUUID(&_LikeStakePosition.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LikeStakePosition *LikeStakePositionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LikeStakePosition.Contract.ProxiableUUID(&_LikeStakePosition.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikeStakePosition *LikeStakePositionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikeStakePosition *LikeStakePositionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikeStakePosition.Contract.SupportsInterface(&_LikeStakePosition.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikeStakePosition *LikeStakePositionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikeStakePosition.Contract.SupportsInterface(&_LikeStakePosition.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeStakePosition *LikeStakePositionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeStakePosition *LikeStakePositionSession) Symbol() (string, error) {
	return _LikeStakePosition.Contract.Symbol(&_LikeStakePosition.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeStakePosition *LikeStakePositionCallerSession) Symbol() (string, error) {
	return _LikeStakePosition.Contract.Symbol(&_LikeStakePosition.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LikeStakePosition.Contract.TokenByIndex(&_LikeStakePosition.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LikeStakePosition.Contract.TokenByIndex(&_LikeStakePosition.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LikeStakePosition.Contract.TokenOfOwnerByIndex(&_LikeStakePosition.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LikeStakePosition.Contract.TokenOfOwnerByIndex(&_LikeStakePosition.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LikeStakePosition *LikeStakePositionCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LikeStakePosition *LikeStakePositionSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LikeStakePosition.Contract.TokenURI(&_LikeStakePosition.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LikeStakePosition *LikeStakePositionCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LikeStakePosition.Contract.TokenURI(&_LikeStakePosition.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeStakePosition.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionSession) TotalSupply() (*big.Int, error) {
	return _LikeStakePosition.Contract.TotalSupply(&_LikeStakePosition.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeStakePosition *LikeStakePositionCallerSession) TotalSupply() (*big.Int, error) {
	return _LikeStakePosition.Contract.TotalSupply(&_LikeStakePosition.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Approve(&_LikeStakePosition.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Approve(&_LikeStakePosition.TransactOpts, to, tokenId)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x38ca63bc.
//
// Solidity: function burnPosition(uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) BurnPosition(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "burnPosition", tokenId)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x38ca63bc.
//
// Solidity: function burnPosition(uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionSession) BurnPosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.BurnPosition(&_LikeStakePosition.TransactOpts, tokenId)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x38ca63bc.
//
// Solidity: function burnPosition(uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) BurnPosition(tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.BurnPosition(&_LikeStakePosition.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeStakePosition *LikeStakePositionSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Initialize(&_LikeStakePosition.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Initialize(&_LikeStakePosition.TransactOpts, initialOwner)
}

// MintPosition is a paid mutator transaction binding the contract method 0x09fbb2d7.
//
// Solidity: function mintPosition(address to, address bookNFT, uint256 stakedAmount, uint256 rewardIndex) returns(uint256 tokenId)
func (_LikeStakePosition *LikeStakePositionTransactor) MintPosition(opts *bind.TransactOpts, to common.Address, bookNFT common.Address, stakedAmount *big.Int, rewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "mintPosition", to, bookNFT, stakedAmount, rewardIndex)
}

// MintPosition is a paid mutator transaction binding the contract method 0x09fbb2d7.
//
// Solidity: function mintPosition(address to, address bookNFT, uint256 stakedAmount, uint256 rewardIndex) returns(uint256 tokenId)
func (_LikeStakePosition *LikeStakePositionSession) MintPosition(to common.Address, bookNFT common.Address, stakedAmount *big.Int, rewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.MintPosition(&_LikeStakePosition.TransactOpts, to, bookNFT, stakedAmount, rewardIndex)
}

// MintPosition is a paid mutator transaction binding the contract method 0x09fbb2d7.
//
// Solidity: function mintPosition(address to, address bookNFT, uint256 stakedAmount, uint256 rewardIndex) returns(uint256 tokenId)
func (_LikeStakePosition *LikeStakePositionTransactorSession) MintPosition(to common.Address, bookNFT common.Address, stakedAmount *big.Int, rewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.MintPosition(&_LikeStakePosition.TransactOpts, to, bookNFT, stakedAmount, rewardIndex)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeStakePosition *LikeStakePositionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeStakePosition *LikeStakePositionSession) Pause() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Pause(&_LikeStakePosition.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) Pause() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Pause(&_LikeStakePosition.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeStakePosition *LikeStakePositionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeStakePosition *LikeStakePositionSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.RenounceOwnership(&_LikeStakePosition.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.RenounceOwnership(&_LikeStakePosition.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SafeTransferFrom(&_LikeStakePosition.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SafeTransferFrom(&_LikeStakePosition.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikeStakePosition *LikeStakePositionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SafeTransferFrom0(&_LikeStakePosition.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SafeTransferFrom0(&_LikeStakePosition.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikeStakePosition *LikeStakePositionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetApprovalForAll(&_LikeStakePosition.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetApprovalForAll(&_LikeStakePosition.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_LikeStakePosition *LikeStakePositionSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetBaseURI(&_LikeStakePosition.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetBaseURI(&_LikeStakePosition.TransactOpts, baseURI_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) SetManager(opts *bind.TransactOpts, manager_ common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "setManager", manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_LikeStakePosition *LikeStakePositionSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetManager(&_LikeStakePosition.TransactOpts, manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.SetManager(&_LikeStakePosition.TransactOpts, manager_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.TransferFrom(&_LikeStakePosition.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.TransferFrom(&_LikeStakePosition.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeStakePosition *LikeStakePositionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.TransferOwnership(&_LikeStakePosition.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.TransferOwnership(&_LikeStakePosition.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeStakePosition *LikeStakePositionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeStakePosition *LikeStakePositionSession) Unpause() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Unpause(&_LikeStakePosition.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) Unpause() (*types.Transaction, error) {
	return _LikeStakePosition.Contract.Unpause(&_LikeStakePosition.TransactOpts)
}

// UpdatePosition is a paid mutator transaction binding the contract method 0x0602e13a.
//
// Solidity: function updatePosition(uint256 tokenId, uint256 newStakedAmount, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) UpdatePosition(opts *bind.TransactOpts, tokenId *big.Int, newStakedAmount *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "updatePosition", tokenId, newStakedAmount, newRewardIndex)
}

// UpdatePosition is a paid mutator transaction binding the contract method 0x0602e13a.
//
// Solidity: function updatePosition(uint256 tokenId, uint256 newStakedAmount, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionSession) UpdatePosition(tokenId *big.Int, newStakedAmount *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpdatePosition(&_LikeStakePosition.TransactOpts, tokenId, newStakedAmount, newRewardIndex)
}

// UpdatePosition is a paid mutator transaction binding the contract method 0x0602e13a.
//
// Solidity: function updatePosition(uint256 tokenId, uint256 newStakedAmount, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) UpdatePosition(tokenId *big.Int, newStakedAmount *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpdatePosition(&_LikeStakePosition.TransactOpts, tokenId, newStakedAmount, newRewardIndex)
}

// UpdatePositionRewardIndex is a paid mutator transaction binding the contract method 0x01862992.
//
// Solidity: function updatePositionRewardIndex(uint256 tokenId, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionTransactor) UpdatePositionRewardIndex(opts *bind.TransactOpts, tokenId *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "updatePositionRewardIndex", tokenId, newRewardIndex)
}

// UpdatePositionRewardIndex is a paid mutator transaction binding the contract method 0x01862992.
//
// Solidity: function updatePositionRewardIndex(uint256 tokenId, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionSession) UpdatePositionRewardIndex(tokenId *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpdatePositionRewardIndex(&_LikeStakePosition.TransactOpts, tokenId, newRewardIndex)
}

// UpdatePositionRewardIndex is a paid mutator transaction binding the contract method 0x01862992.
//
// Solidity: function updatePositionRewardIndex(uint256 tokenId, uint256 newRewardIndex) returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) UpdatePositionRewardIndex(tokenId *big.Int, newRewardIndex *big.Int) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpdatePositionRewardIndex(&_LikeStakePosition.TransactOpts, tokenId, newRewardIndex)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeStakePosition *LikeStakePositionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeStakePosition *LikeStakePositionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpgradeToAndCall(&_LikeStakePosition.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LikeStakePosition *LikeStakePositionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LikeStakePosition.Contract.UpgradeToAndCall(&_LikeStakePosition.TransactOpts, newImplementation, data)
}

// LikeStakePositionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LikeStakePosition contract.
type LikeStakePositionApprovalIterator struct {
	Event *LikeStakePositionApproval // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionApproval)
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
		it.Event = new(LikeStakePositionApproval)
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
func (it *LikeStakePositionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionApproval represents a Approval event raised by the LikeStakePosition contract.
type LikeStakePositionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LikeStakePositionApprovalIterator, error) {

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

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionApprovalIterator{contract: _LikeStakePosition.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LikeStakePositionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionApproval)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseApproval(log types.Log) (*LikeStakePositionApproval, error) {
	event := new(LikeStakePositionApproval)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LikeStakePosition contract.
type LikeStakePositionApprovalForAllIterator struct {
	Event *LikeStakePositionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionApprovalForAll)
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
		it.Event = new(LikeStakePositionApprovalForAll)
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
func (it *LikeStakePositionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionApprovalForAll represents a ApprovalForAll event raised by the LikeStakePosition contract.
type LikeStakePositionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LikeStakePositionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionApprovalForAllIterator{contract: _LikeStakePosition.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LikeStakePositionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionApprovalForAll)
				if err := _LikeStakePosition.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseApprovalForAll(log types.Log) (*LikeStakePositionApprovalForAll, error) {
	event := new(LikeStakePositionApprovalForAll)
	if err := _LikeStakePosition.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionBaseURIUpdatedIterator is returned from FilterBaseURIUpdated and is used to iterate over the raw logs and unpacked data for BaseURIUpdated events raised by the LikeStakePosition contract.
type LikeStakePositionBaseURIUpdatedIterator struct {
	Event *LikeStakePositionBaseURIUpdated // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionBaseURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionBaseURIUpdated)
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
		it.Event = new(LikeStakePositionBaseURIUpdated)
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
func (it *LikeStakePositionBaseURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionBaseURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionBaseURIUpdated represents a BaseURIUpdated event raised by the LikeStakePosition contract.
type LikeStakePositionBaseURIUpdated struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIUpdated is a free log retrieval operation binding the contract event 0x6741b2fc379fad678116fe3d4d4b9a1a184ab53ba36b86ad0fa66340b1ab41ad.
//
// Solidity: event BaseURIUpdated(string baseURI)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterBaseURIUpdated(opts *bind.FilterOpts) (*LikeStakePositionBaseURIUpdatedIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "BaseURIUpdated")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionBaseURIUpdatedIterator{contract: _LikeStakePosition.contract, event: "BaseURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseURIUpdated is a free log subscription operation binding the contract event 0x6741b2fc379fad678116fe3d4d4b9a1a184ab53ba36b86ad0fa66340b1ab41ad.
//
// Solidity: event BaseURIUpdated(string baseURI)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchBaseURIUpdated(opts *bind.WatchOpts, sink chan<- *LikeStakePositionBaseURIUpdated) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "BaseURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionBaseURIUpdated)
				if err := _LikeStakePosition.contract.UnpackLog(event, "BaseURIUpdated", log); err != nil {
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

// ParseBaseURIUpdated is a log parse operation binding the contract event 0x6741b2fc379fad678116fe3d4d4b9a1a184ab53ba36b86ad0fa66340b1ab41ad.
//
// Solidity: event BaseURIUpdated(string baseURI)
func (_LikeStakePosition *LikeStakePositionFilterer) ParseBaseURIUpdated(log types.Log) (*LikeStakePositionBaseURIUpdated, error) {
	event := new(LikeStakePositionBaseURIUpdated)
	if err := _LikeStakePosition.contract.UnpackLog(event, "BaseURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the LikeStakePosition contract.
type LikeStakePositionBatchMetadataUpdateIterator struct {
	Event *LikeStakePositionBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionBatchMetadataUpdate)
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
		it.Event = new(LikeStakePositionBatchMetadataUpdate)
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
func (it *LikeStakePositionBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the LikeStakePosition contract.
type LikeStakePositionBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*LikeStakePositionBatchMetadataUpdateIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionBatchMetadataUpdateIterator{contract: _LikeStakePosition.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LikeStakePositionBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionBatchMetadataUpdate)
				if err := _LikeStakePosition.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseBatchMetadataUpdate(log types.Log) (*LikeStakePositionBatchMetadataUpdate, error) {
	event := new(LikeStakePositionBatchMetadataUpdate)
	if err := _LikeStakePosition.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LikeStakePosition contract.
type LikeStakePositionInitializedIterator struct {
	Event *LikeStakePositionInitialized // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionInitialized)
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
		it.Event = new(LikeStakePositionInitialized)
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
func (it *LikeStakePositionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionInitialized represents a Initialized event raised by the LikeStakePosition contract.
type LikeStakePositionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterInitialized(opts *bind.FilterOpts) (*LikeStakePositionInitializedIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionInitializedIterator{contract: _LikeStakePosition.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LikeStakePositionInitialized) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionInitialized)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseInitialized(log types.Log) (*LikeStakePositionInitialized, error) {
	event := new(LikeStakePositionInitialized)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionManagerUpdatedIterator is returned from FilterManagerUpdated and is used to iterate over the raw logs and unpacked data for ManagerUpdated events raised by the LikeStakePosition contract.
type LikeStakePositionManagerUpdatedIterator struct {
	Event *LikeStakePositionManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionManagerUpdated)
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
		it.Event = new(LikeStakePositionManagerUpdated)
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
func (it *LikeStakePositionManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionManagerUpdated represents a ManagerUpdated event raised by the LikeStakePosition contract.
type LikeStakePositionManagerUpdated struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerUpdated is a free log retrieval operation binding the contract event 0x2c1c11af44aa5608f1dca38c00275c30ea091e02417d36e70e9a1538689c433d.
//
// Solidity: event ManagerUpdated(address indexed manager)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterManagerUpdated(opts *bind.FilterOpts, manager []common.Address) (*LikeStakePositionManagerUpdatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionManagerUpdatedIterator{contract: _LikeStakePosition.contract, event: "ManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchManagerUpdated is a free log subscription operation binding the contract event 0x2c1c11af44aa5608f1dca38c00275c30ea091e02417d36e70e9a1538689c433d.
//
// Solidity: event ManagerUpdated(address indexed manager)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchManagerUpdated(opts *bind.WatchOpts, sink chan<- *LikeStakePositionManagerUpdated, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionManagerUpdated)
				if err := _LikeStakePosition.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
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

// ParseManagerUpdated is a log parse operation binding the contract event 0x2c1c11af44aa5608f1dca38c00275c30ea091e02417d36e70e9a1538689c433d.
//
// Solidity: event ManagerUpdated(address indexed manager)
func (_LikeStakePosition *LikeStakePositionFilterer) ParseManagerUpdated(log types.Log) (*LikeStakePositionManagerUpdated, error) {
	event := new(LikeStakePositionManagerUpdated)
	if err := _LikeStakePosition.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the LikeStakePosition contract.
type LikeStakePositionMetadataUpdateIterator struct {
	Event *LikeStakePositionMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionMetadataUpdate)
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
		it.Event = new(LikeStakePositionMetadataUpdate)
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
func (it *LikeStakePositionMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionMetadataUpdate represents a MetadataUpdate event raised by the LikeStakePosition contract.
type LikeStakePositionMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*LikeStakePositionMetadataUpdateIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionMetadataUpdateIterator{contract: _LikeStakePosition.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *LikeStakePositionMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionMetadataUpdate)
				if err := _LikeStakePosition.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseMetadataUpdate(log types.Log) (*LikeStakePositionMetadataUpdate, error) {
	event := new(LikeStakePositionMetadataUpdate)
	if err := _LikeStakePosition.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LikeStakePosition contract.
type LikeStakePositionOwnershipTransferredIterator struct {
	Event *LikeStakePositionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionOwnershipTransferred)
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
		it.Event = new(LikeStakePositionOwnershipTransferred)
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
func (it *LikeStakePositionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionOwnershipTransferred represents a OwnershipTransferred event raised by the LikeStakePosition contract.
type LikeStakePositionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LikeStakePositionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionOwnershipTransferredIterator{contract: _LikeStakePosition.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LikeStakePositionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionOwnershipTransferred)
				if err := _LikeStakePosition.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseOwnershipTransferred(log types.Log) (*LikeStakePositionOwnershipTransferred, error) {
	event := new(LikeStakePositionOwnershipTransferred)
	if err := _LikeStakePosition.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LikeStakePosition contract.
type LikeStakePositionPausedIterator struct {
	Event *LikeStakePositionPaused // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionPaused)
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
		it.Event = new(LikeStakePositionPaused)
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
func (it *LikeStakePositionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionPaused represents a Paused event raised by the LikeStakePosition contract.
type LikeStakePositionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterPaused(opts *bind.FilterOpts) (*LikeStakePositionPausedIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionPausedIterator{contract: _LikeStakePosition.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LikeStakePositionPaused) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionPaused)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParsePaused(log types.Log) (*LikeStakePositionPaused, error) {
	event := new(LikeStakePositionPaused)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionPositionBurnedIterator is returned from FilterPositionBurned and is used to iterate over the raw logs and unpacked data for PositionBurned events raised by the LikeStakePosition contract.
type LikeStakePositionPositionBurnedIterator struct {
	Event *LikeStakePositionPositionBurned // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionPositionBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionPositionBurned)
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
		it.Event = new(LikeStakePositionPositionBurned)
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
func (it *LikeStakePositionPositionBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionPositionBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionPositionBurned represents a PositionBurned event raised by the LikeStakePosition contract.
type LikeStakePositionPositionBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPositionBurned is a free log retrieval operation binding the contract event 0x65d33d8ef62a81711748bbe2a7b67aef94d1a9af04a2690d3a4dfd13d9c1d22b.
//
// Solidity: event PositionBurned(uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterPositionBurned(opts *bind.FilterOpts, tokenId []*big.Int) (*LikeStakePositionPositionBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "PositionBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionPositionBurnedIterator{contract: _LikeStakePosition.contract, event: "PositionBurned", logs: logs, sub: sub}, nil
}

// WatchPositionBurned is a free log subscription operation binding the contract event 0x65d33d8ef62a81711748bbe2a7b67aef94d1a9af04a2690d3a4dfd13d9c1d22b.
//
// Solidity: event PositionBurned(uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchPositionBurned(opts *bind.WatchOpts, sink chan<- *LikeStakePositionPositionBurned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "PositionBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionPositionBurned)
				if err := _LikeStakePosition.contract.UnpackLog(event, "PositionBurned", log); err != nil {
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

// ParsePositionBurned is a log parse operation binding the contract event 0x65d33d8ef62a81711748bbe2a7b67aef94d1a9af04a2690d3a4dfd13d9c1d22b.
//
// Solidity: event PositionBurned(uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) ParsePositionBurned(log types.Log) (*LikeStakePositionPositionBurned, error) {
	event := new(LikeStakePositionPositionBurned)
	if err := _LikeStakePosition.contract.UnpackLog(event, "PositionBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionPositionMintedIterator is returned from FilterPositionMinted and is used to iterate over the raw logs and unpacked data for PositionMinted events raised by the LikeStakePosition contract.
type LikeStakePositionPositionMintedIterator struct {
	Event *LikeStakePositionPositionMinted // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionPositionMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionPositionMinted)
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
		it.Event = new(LikeStakePositionPositionMinted)
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
func (it *LikeStakePositionPositionMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionPositionMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionPositionMinted represents a PositionMinted event raised by the LikeStakePosition contract.
type LikeStakePositionPositionMinted struct {
	TokenId     *big.Int
	To          common.Address
	BookNFT     common.Address
	Amount      *big.Int
	RewardIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionMinted is a free log retrieval operation binding the contract event 0x8cd7f5babbb088bffdd32aee19a1dc179c8819cd6d02cef02c35cb3ebfa38609.
//
// Solidity: event PositionMinted(uint256 indexed tokenId, address indexed to, address indexed bookNFT, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterPositionMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, bookNFT []common.Address) (*LikeStakePositionPositionMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "PositionMinted", tokenIdRule, toRule, bookNFTRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionPositionMintedIterator{contract: _LikeStakePosition.contract, event: "PositionMinted", logs: logs, sub: sub}, nil
}

// WatchPositionMinted is a free log subscription operation binding the contract event 0x8cd7f5babbb088bffdd32aee19a1dc179c8819cd6d02cef02c35cb3ebfa38609.
//
// Solidity: event PositionMinted(uint256 indexed tokenId, address indexed to, address indexed bookNFT, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchPositionMinted(opts *bind.WatchOpts, sink chan<- *LikeStakePositionPositionMinted, tokenId []*big.Int, to []common.Address, bookNFT []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var bookNFTRule []interface{}
	for _, bookNFTItem := range bookNFT {
		bookNFTRule = append(bookNFTRule, bookNFTItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "PositionMinted", tokenIdRule, toRule, bookNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionPositionMinted)
				if err := _LikeStakePosition.contract.UnpackLog(event, "PositionMinted", log); err != nil {
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

// ParsePositionMinted is a log parse operation binding the contract event 0x8cd7f5babbb088bffdd32aee19a1dc179c8819cd6d02cef02c35cb3ebfa38609.
//
// Solidity: event PositionMinted(uint256 indexed tokenId, address indexed to, address indexed bookNFT, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) ParsePositionMinted(log types.Log) (*LikeStakePositionPositionMinted, error) {
	event := new(LikeStakePositionPositionMinted)
	if err := _LikeStakePosition.contract.UnpackLog(event, "PositionMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionPositionUpdatedIterator is returned from FilterPositionUpdated and is used to iterate over the raw logs and unpacked data for PositionUpdated events raised by the LikeStakePosition contract.
type LikeStakePositionPositionUpdatedIterator struct {
	Event *LikeStakePositionPositionUpdated // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionPositionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionPositionUpdated)
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
		it.Event = new(LikeStakePositionPositionUpdated)
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
func (it *LikeStakePositionPositionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionPositionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionPositionUpdated represents a PositionUpdated event raised by the LikeStakePosition contract.
type LikeStakePositionPositionUpdated struct {
	TokenId     *big.Int
	Amount      *big.Int
	RewardIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionUpdated is a free log retrieval operation binding the contract event 0xf89185fe8013045e431c4a2d209bdc8367e67b3aee3359fde091b6ce08e60550.
//
// Solidity: event PositionUpdated(uint256 indexed tokenId, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterPositionUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*LikeStakePositionPositionUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "PositionUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionPositionUpdatedIterator{contract: _LikeStakePosition.contract, event: "PositionUpdated", logs: logs, sub: sub}, nil
}

// WatchPositionUpdated is a free log subscription operation binding the contract event 0xf89185fe8013045e431c4a2d209bdc8367e67b3aee3359fde091b6ce08e60550.
//
// Solidity: event PositionUpdated(uint256 indexed tokenId, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchPositionUpdated(opts *bind.WatchOpts, sink chan<- *LikeStakePositionPositionUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "PositionUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionPositionUpdated)
				if err := _LikeStakePosition.contract.UnpackLog(event, "PositionUpdated", log); err != nil {
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

// ParsePositionUpdated is a log parse operation binding the contract event 0xf89185fe8013045e431c4a2d209bdc8367e67b3aee3359fde091b6ce08e60550.
//
// Solidity: event PositionUpdated(uint256 indexed tokenId, uint256 amount, uint256 rewardIndex)
func (_LikeStakePosition *LikeStakePositionFilterer) ParsePositionUpdated(log types.Log) (*LikeStakePositionPositionUpdated, error) {
	event := new(LikeStakePositionPositionUpdated)
	if err := _LikeStakePosition.contract.UnpackLog(event, "PositionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LikeStakePosition contract.
type LikeStakePositionTransferIterator struct {
	Event *LikeStakePositionTransfer // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionTransfer)
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
		it.Event = new(LikeStakePositionTransfer)
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
func (it *LikeStakePositionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionTransfer represents a Transfer event raised by the LikeStakePosition contract.
type LikeStakePositionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LikeStakePositionTransferIterator, error) {

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

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionTransferIterator{contract: _LikeStakePosition.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LikeStakePositionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionTransfer)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseTransfer(log types.Log) (*LikeStakePositionTransfer, error) {
	event := new(LikeStakePositionTransfer)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LikeStakePosition contract.
type LikeStakePositionUnpausedIterator struct {
	Event *LikeStakePositionUnpaused // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionUnpaused)
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
		it.Event = new(LikeStakePositionUnpaused)
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
func (it *LikeStakePositionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionUnpaused represents a Unpaused event raised by the LikeStakePosition contract.
type LikeStakePositionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LikeStakePositionUnpausedIterator, error) {

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionUnpausedIterator{contract: _LikeStakePosition.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LikeStakePositionUnpaused) (event.Subscription, error) {

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionUnpaused)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseUnpaused(log types.Log) (*LikeStakePositionUnpaused, error) {
	event := new(LikeStakePositionUnpaused)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeStakePositionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LikeStakePosition contract.
type LikeStakePositionUpgradedIterator struct {
	Event *LikeStakePositionUpgraded // Event containing the contract specifics and raw log

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
func (it *LikeStakePositionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeStakePositionUpgraded)
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
		it.Event = new(LikeStakePositionUpgraded)
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
func (it *LikeStakePositionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeStakePositionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeStakePositionUpgraded represents a Upgraded event raised by the LikeStakePosition contract.
type LikeStakePositionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeStakePosition *LikeStakePositionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LikeStakePositionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeStakePosition.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LikeStakePositionUpgradedIterator{contract: _LikeStakePosition.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LikeStakePosition *LikeStakePositionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LikeStakePositionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LikeStakePosition.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeStakePositionUpgraded)
				if err := _LikeStakePosition.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LikeStakePosition *LikeStakePositionFilterer) ParseUpgraded(log types.Log) (*LikeStakePositionUpgraded, error) {
	event := new(LikeStakePositionUpgraded)
	if err := _LikeStakePosition.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
