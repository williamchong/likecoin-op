// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package book_nft

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

// BookNftMetaData contains all meta data concerning the BookNft contract.
var BookNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrEmptyName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidSymbol\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrMaxSupplyZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNftNoSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nextTokenId\",\"type\":\"uint256\"}],\"name\":\"ErrTokenIdMintFails\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferWithMemo\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"memos\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"metadataList\",\"type\":\"string[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"memos\",\"type\":\"string[]\"}],\"name\":\"batchTransferWithMemo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewBookNFT\",\"name\":\"msgNewBookNFT\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"memos\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"metadataList\",\"type\":\"string[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"memos\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"metadataList\",\"type\":\"string[]\"}],\"name\":\"safeMintWithTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"royaltyFraction\",\"type\":\"uint96\"}],\"name\":\"setRoyaltyFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferWithMemo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structBookConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BookNftABI is the input ABI used to generate the binding from.
// Deprecated: Use BookNftMetaData.ABI instead.
var BookNftABI = BookNftMetaData.ABI

// BookNft is an auto generated Go binding around an Ethereum contract.
type BookNft struct {
	BookNftCaller     // Read-only binding to the contract
	BookNftTransactor // Write-only binding to the contract
	BookNftFilterer   // Log filterer for contract events
}

// BookNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookNftSession struct {
	Contract     *BookNft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookNftCallerSession struct {
	Contract *BookNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BookNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookNftTransactorSession struct {
	Contract     *BookNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BookNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookNftRaw struct {
	Contract *BookNft // Generic contract binding to access the raw methods on
}

// BookNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookNftCallerRaw struct {
	Contract *BookNftCaller // Generic read-only contract binding to access the raw methods on
}

// BookNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookNftTransactorRaw struct {
	Contract *BookNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookNft creates a new instance of BookNft, bound to a specific deployed contract.
func NewBookNft(address common.Address, backend bind.ContractBackend) (*BookNft, error) {
	contract, err := bindBookNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BookNft{BookNftCaller: BookNftCaller{contract: contract}, BookNftTransactor: BookNftTransactor{contract: contract}, BookNftFilterer: BookNftFilterer{contract: contract}}, nil
}

// NewBookNftCaller creates a new read-only instance of BookNft, bound to a specific deployed contract.
func NewBookNftCaller(address common.Address, caller bind.ContractCaller) (*BookNftCaller, error) {
	contract, err := bindBookNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookNftCaller{contract: contract}, nil
}

// NewBookNftTransactor creates a new write-only instance of BookNft, bound to a specific deployed contract.
func NewBookNftTransactor(address common.Address, transactor bind.ContractTransactor) (*BookNftTransactor, error) {
	contract, err := bindBookNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookNftTransactor{contract: contract}, nil
}

// NewBookNftFilterer creates a new log filterer instance of BookNft, bound to a specific deployed contract.
func NewBookNftFilterer(address common.Address, filterer bind.ContractFilterer) (*BookNftFilterer, error) {
	contract, err := bindBookNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookNftFilterer{contract: contract}, nil
}

// bindBookNft binds a generic wrapper to an already deployed contract.
func bindBookNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BookNft *BookNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BookNft.Contract.BookNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BookNft *BookNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BookNft.Contract.BookNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BookNft *BookNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BookNft.Contract.BookNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BookNft *BookNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BookNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BookNft *BookNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BookNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BookNft *BookNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BookNft.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BookNft *BookNftCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BookNft *BookNftSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BookNft.Contract.DEFAULTADMINROLE(&_BookNft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BookNft *BookNftCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BookNft.Contract.DEFAULTADMINROLE(&_BookNft.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BookNft *BookNftCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BookNft *BookNftSession) MINTERROLE() ([32]byte, error) {
	return _BookNft.Contract.MINTERROLE(&_BookNft.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BookNft *BookNftCallerSession) MINTERROLE() ([32]byte, error) {
	return _BookNft.Contract.MINTERROLE(&_BookNft.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_BookNft *BookNftCaller) UPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_BookNft *BookNftSession) UPDATERROLE() ([32]byte, error) {
	return _BookNft.Contract.UPDATERROLE(&_BookNft.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_BookNft *BookNftCallerSession) UPDATERROLE() ([32]byte, error) {
	return _BookNft.Contract.UPDATERROLE(&_BookNft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BookNft *BookNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BookNft *BookNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BookNft.Contract.BalanceOf(&_BookNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BookNft *BookNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BookNft.Contract.BalanceOf(&_BookNft.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_BookNft *BookNftCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_BookNft *BookNftSession) ContractURI() (string, error) {
	return _BookNft.Contract.ContractURI(&_BookNft.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_BookNft *BookNftCallerSession) ContractURI() (string, error) {
	return _BookNft.Contract.ContractURI(&_BookNft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BookNft *BookNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BookNft *BookNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BookNft.Contract.GetApproved(&_BookNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BookNft *BookNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BookNft.Contract.GetApproved(&_BookNft.CallOpts, tokenId)
}

// GetBookConfig is a free data retrieval call binding the contract method 0xdb0b09b9.
//
// Solidity: function getBookConfig() view returns((string,string,string,uint64))
func (_BookNft *BookNftCaller) GetBookConfig(opts *bind.CallOpts) (BookConfig, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "getBookConfig")

	if err != nil {
		return *new(BookConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BookConfig)).(*BookConfig)

	return out0, err

}

// GetBookConfig is a free data retrieval call binding the contract method 0xdb0b09b9.
//
// Solidity: function getBookConfig() view returns((string,string,string,uint64))
func (_BookNft *BookNftSession) GetBookConfig() (BookConfig, error) {
	return _BookNft.Contract.GetBookConfig(&_BookNft.CallOpts)
}

// GetBookConfig is a free data retrieval call binding the contract method 0xdb0b09b9.
//
// Solidity: function getBookConfig() view returns((string,string,string,uint64))
func (_BookNft *BookNftCallerSession) GetBookConfig() (BookConfig, error) {
	return _BookNft.Contract.GetBookConfig(&_BookNft.CallOpts)
}

// GetCurrentIndex is a free data retrieval call binding the contract method 0x0d9005ae.
//
// Solidity: function getCurrentIndex() view returns(uint256)
func (_BookNft *BookNftCaller) GetCurrentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "getCurrentIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentIndex is a free data retrieval call binding the contract method 0x0d9005ae.
//
// Solidity: function getCurrentIndex() view returns(uint256)
func (_BookNft *BookNftSession) GetCurrentIndex() (*big.Int, error) {
	return _BookNft.Contract.GetCurrentIndex(&_BookNft.CallOpts)
}

// GetCurrentIndex is a free data retrieval call binding the contract method 0x0d9005ae.
//
// Solidity: function getCurrentIndex() view returns(uint256)
func (_BookNft *BookNftCallerSession) GetCurrentIndex() (*big.Int, error) {
	return _BookNft.Contract.GetCurrentIndex(&_BookNft.CallOpts)
}

// GetProtocolBeacon is a free data retrieval call binding the contract method 0xf570d174.
//
// Solidity: function getProtocolBeacon() view returns(address)
func (_BookNft *BookNftCaller) GetProtocolBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "getProtocolBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolBeacon is a free data retrieval call binding the contract method 0xf570d174.
//
// Solidity: function getProtocolBeacon() view returns(address)
func (_BookNft *BookNftSession) GetProtocolBeacon() (common.Address, error) {
	return _BookNft.Contract.GetProtocolBeacon(&_BookNft.CallOpts)
}

// GetProtocolBeacon is a free data retrieval call binding the contract method 0xf570d174.
//
// Solidity: function getProtocolBeacon() view returns(address)
func (_BookNft *BookNftCallerSession) GetProtocolBeacon() (common.Address, error) {
	return _BookNft.Contract.GetProtocolBeacon(&_BookNft.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BookNft *BookNftCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BookNft *BookNftSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BookNft.Contract.GetRoleAdmin(&_BookNft.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BookNft *BookNftCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BookNft.Contract.GetRoleAdmin(&_BookNft.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BookNft *BookNftCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BookNft *BookNftSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BookNft.Contract.HasRole(&_BookNft.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BookNft *BookNftCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BookNft.Contract.HasRole(&_BookNft.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BookNft *BookNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BookNft *BookNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BookNft.Contract.IsApprovedForAll(&_BookNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BookNft *BookNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BookNft.Contract.IsApprovedForAll(&_BookNft.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint64)
func (_BookNft *BookNftCaller) MaxSupply(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint64)
func (_BookNft *BookNftSession) MaxSupply() (uint64, error) {
	return _BookNft.Contract.MaxSupply(&_BookNft.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint64)
func (_BookNft *BookNftCallerSession) MaxSupply() (uint64, error) {
	return _BookNft.Contract.MaxSupply(&_BookNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BookNft *BookNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BookNft *BookNftSession) Name() (string, error) {
	return _BookNft.Contract.Name(&_BookNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BookNft *BookNftCallerSession) Name() (string, error) {
	return _BookNft.Contract.Name(&_BookNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BookNft *BookNftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BookNft *BookNftSession) Owner() (common.Address, error) {
	return _BookNft.Contract.Owner(&_BookNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BookNft *BookNftCallerSession) Owner() (common.Address, error) {
	return _BookNft.Contract.Owner(&_BookNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BookNft *BookNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BookNft *BookNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BookNft.Contract.OwnerOf(&_BookNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BookNft *BookNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BookNft.Contract.OwnerOf(&_BookNft.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_BookNft *BookNftCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "royaltyInfo", arg0, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_BookNft *BookNftSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _BookNft.Contract.RoyaltyInfo(&_BookNft.CallOpts, arg0, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_BookNft *BookNftCallerSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _BookNft.Contract.RoyaltyInfo(&_BookNft.CallOpts, arg0, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BookNft *BookNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BookNft *BookNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BookNft.Contract.SupportsInterface(&_BookNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BookNft *BookNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BookNft.Contract.SupportsInterface(&_BookNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BookNft *BookNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BookNft *BookNftSession) Symbol() (string, error) {
	return _BookNft.Contract.Symbol(&_BookNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BookNft *BookNftCallerSession) Symbol() (string, error) {
	return _BookNft.Contract.Symbol(&_BookNft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BookNft *BookNftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BookNft *BookNftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BookNft.Contract.TokenByIndex(&_BookNft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BookNft *BookNftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BookNft.Contract.TokenByIndex(&_BookNft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BookNft *BookNftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BookNft *BookNftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BookNft.Contract.TokenOfOwnerByIndex(&_BookNft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BookNft *BookNftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BookNft.Contract.TokenOfOwnerByIndex(&_BookNft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_BookNft *BookNftCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_BookNft *BookNftSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _BookNft.Contract.TokenURI(&_BookNft.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_BookNft *BookNftCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _BookNft.Contract.TokenURI(&_BookNft.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BookNft *BookNftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BookNft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BookNft *BookNftSession) TotalSupply() (*big.Int, error) {
	return _BookNft.Contract.TotalSupply(&_BookNft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BookNft *BookNftCallerSession) TotalSupply() (*big.Int, error) {
	return _BookNft.Contract.TotalSupply(&_BookNft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BookNft *BookNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.Approve(&_BookNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.Approve(&_BookNft.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0xdab45bbd.
//
// Solidity: function batchMint(address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactor) BatchMint(opts *bind.TransactOpts, tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "batchMint", tos, memos, metadataList)
}

// BatchMint is a paid mutator transaction binding the contract method 0xdab45bbd.
//
// Solidity: function batchMint(address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftSession) BatchMint(tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.BatchMint(&_BookNft.TransactOpts, tos, memos, metadataList)
}

// BatchMint is a paid mutator transaction binding the contract method 0xdab45bbd.
//
// Solidity: function batchMint(address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactorSession) BatchMint(tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.BatchMint(&_BookNft.TransactOpts, tos, memos, metadataList)
}

// BatchTransferWithMemo is a paid mutator transaction binding the contract method 0xa75c72bd.
//
// Solidity: function batchTransferWithMemo(address from, address[] tos, uint256[] tokenIds, string[] memos) payable returns()
func (_BookNft *BookNftTransactor) BatchTransferWithMemo(opts *bind.TransactOpts, from common.Address, tos []common.Address, tokenIds []*big.Int, memos []string) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "batchTransferWithMemo", from, tos, tokenIds, memos)
}

// BatchTransferWithMemo is a paid mutator transaction binding the contract method 0xa75c72bd.
//
// Solidity: function batchTransferWithMemo(address from, address[] tos, uint256[] tokenIds, string[] memos) payable returns()
func (_BookNft *BookNftSession) BatchTransferWithMemo(from common.Address, tos []common.Address, tokenIds []*big.Int, memos []string) (*types.Transaction, error) {
	return _BookNft.Contract.BatchTransferWithMemo(&_BookNft.TransactOpts, from, tos, tokenIds, memos)
}

// BatchTransferWithMemo is a paid mutator transaction binding the contract method 0xa75c72bd.
//
// Solidity: function batchTransferWithMemo(address from, address[] tos, uint256[] tokenIds, string[] memos) payable returns()
func (_BookNft *BookNftTransactorSession) BatchTransferWithMemo(from common.Address, tos []common.Address, tokenIds []*big.Int, memos []string) (*types.Transaction, error) {
	return _BookNft.Contract.BatchTransferWithMemo(&_BookNft.TransactOpts, from, tos, tokenIds, memos)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BookNft *BookNftTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BookNft *BookNftSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.Burn(&_BookNft.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BookNft *BookNftTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.Burn(&_BookNft.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BookNft *BookNftTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BookNft *BookNftSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.GrantRole(&_BookNft.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BookNft *BookNftTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.GrantRole(&_BookNft.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x52c375e2.
//
// Solidity: function initialize((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns()
func (_BookNft *BookNftTransactor) Initialize(opts *bind.TransactOpts, msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "initialize", msgNewBookNFT)
}

// Initialize is a paid mutator transaction binding the contract method 0x52c375e2.
//
// Solidity: function initialize((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns()
func (_BookNft *BookNftSession) Initialize(msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _BookNft.Contract.Initialize(&_BookNft.TransactOpts, msgNewBookNFT)
}

// Initialize is a paid mutator transaction binding the contract method 0x52c375e2.
//
// Solidity: function initialize((address,address[],address[],(string,string,string,uint64)) msgNewBookNFT) returns()
func (_BookNft *BookNftTransactorSession) Initialize(msgNewBookNFT MsgNewBookNFT) (*types.Transaction, error) {
	return _BookNft.Contract.Initialize(&_BookNft.TransactOpts, msgNewBookNFT)
}

// Mint is a paid mutator transaction binding the contract method 0xa5210ff6.
//
// Solidity: function mint(address to, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactor) Mint(opts *bind.TransactOpts, to common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "mint", to, memos, metadataList)
}

// Mint is a paid mutator transaction binding the contract method 0xa5210ff6.
//
// Solidity: function mint(address to, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftSession) Mint(to common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.Mint(&_BookNft.TransactOpts, to, memos, metadataList)
}

// Mint is a paid mutator transaction binding the contract method 0xa5210ff6.
//
// Solidity: function mint(address to, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactorSession) Mint(to common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.Mint(&_BookNft.TransactOpts, to, memos, metadataList)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BookNft *BookNftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BookNft *BookNftSession) RenounceOwnership() (*types.Transaction, error) {
	return _BookNft.Contract.RenounceOwnership(&_BookNft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BookNft *BookNftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BookNft.Contract.RenounceOwnership(&_BookNft.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BookNft *BookNftTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BookNft *BookNftSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.RenounceRole(&_BookNft.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BookNft *BookNftTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.RenounceRole(&_BookNft.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BookNft *BookNftTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BookNft *BookNftSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.RevokeRole(&_BookNft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BookNft *BookNftTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.RevokeRole(&_BookNft.TransactOpts, role, account)
}

// SafeMintWithTokenId is a paid mutator transaction binding the contract method 0xc7c48e70.
//
// Solidity: function safeMintWithTokenId(uint256 fromTokenId, address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactor) SafeMintWithTokenId(opts *bind.TransactOpts, fromTokenId *big.Int, tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "safeMintWithTokenId", fromTokenId, tos, memos, metadataList)
}

// SafeMintWithTokenId is a paid mutator transaction binding the contract method 0xc7c48e70.
//
// Solidity: function safeMintWithTokenId(uint256 fromTokenId, address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftSession) SafeMintWithTokenId(fromTokenId *big.Int, tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.SafeMintWithTokenId(&_BookNft.TransactOpts, fromTokenId, tos, memos, metadataList)
}

// SafeMintWithTokenId is a paid mutator transaction binding the contract method 0xc7c48e70.
//
// Solidity: function safeMintWithTokenId(uint256 fromTokenId, address[] tos, string[] memos, string[] metadataList) returns()
func (_BookNft *BookNftTransactorSession) SafeMintWithTokenId(fromTokenId *big.Int, tos []common.Address, memos []string, metadataList []string) (*types.Transaction, error) {
	return _BookNft.Contract.SafeMintWithTokenId(&_BookNft.TransactOpts, fromTokenId, tos, memos, metadataList)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.SafeTransferFrom(&_BookNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.SafeTransferFrom(&_BookNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BookNft *BookNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BookNft *BookNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BookNft.Contract.SafeTransferFrom0(&_BookNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BookNft *BookNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BookNft.Contract.SafeTransferFrom0(&_BookNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BookNft *BookNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BookNft *BookNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BookNft.Contract.SetApprovalForAll(&_BookNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BookNft *BookNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BookNft.Contract.SetApprovalForAll(&_BookNft.TransactOpts, operator, approved)
}

// SetRoyaltyFraction is a paid mutator transaction binding the contract method 0xbc09e593.
//
// Solidity: function setRoyaltyFraction(uint96 royaltyFraction) returns()
func (_BookNft *BookNftTransactor) SetRoyaltyFraction(opts *bind.TransactOpts, royaltyFraction *big.Int) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "setRoyaltyFraction", royaltyFraction)
}

// SetRoyaltyFraction is a paid mutator transaction binding the contract method 0xbc09e593.
//
// Solidity: function setRoyaltyFraction(uint96 royaltyFraction) returns()
func (_BookNft *BookNftSession) SetRoyaltyFraction(royaltyFraction *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.SetRoyaltyFraction(&_BookNft.TransactOpts, royaltyFraction)
}

// SetRoyaltyFraction is a paid mutator transaction binding the contract method 0xbc09e593.
//
// Solidity: function setRoyaltyFraction(uint96 royaltyFraction) returns()
func (_BookNft *BookNftTransactorSession) SetRoyaltyFraction(royaltyFraction *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.SetRoyaltyFraction(&_BookNft.TransactOpts, royaltyFraction)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.TransferFrom(&_BookNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BookNft *BookNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BookNft.Contract.TransferFrom(&_BookNft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BookNft *BookNftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BookNft *BookNftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.TransferOwnership(&_BookNft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BookNft *BookNftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BookNft.Contract.TransferOwnership(&_BookNft.TransactOpts, newOwner)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_BookNft *BookNftTransactor) TransferWithMemo(opts *bind.TransactOpts, from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "transferWithMemo", from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_BookNft *BookNftSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _BookNft.Contract.TransferWithMemo(&_BookNft.TransactOpts, from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_BookNft *BookNftTransactorSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _BookNft.Contract.TransferWithMemo(&_BookNft.TransactOpts, from, to, _tokenId, memo)
}

// Update is a paid mutator transaction binding the contract method 0x2b588091.
//
// Solidity: function update((string,string,string,uint64) config) returns()
func (_BookNft *BookNftTransactor) Update(opts *bind.TransactOpts, config BookConfig) (*types.Transaction, error) {
	return _BookNft.contract.Transact(opts, "update", config)
}

// Update is a paid mutator transaction binding the contract method 0x2b588091.
//
// Solidity: function update((string,string,string,uint64) config) returns()
func (_BookNft *BookNftSession) Update(config BookConfig) (*types.Transaction, error) {
	return _BookNft.Contract.Update(&_BookNft.TransactOpts, config)
}

// Update is a paid mutator transaction binding the contract method 0x2b588091.
//
// Solidity: function update((string,string,string,uint64) config) returns()
func (_BookNft *BookNftTransactorSession) Update(config BookConfig) (*types.Transaction, error) {
	return _BookNft.Contract.Update(&_BookNft.TransactOpts, config)
}

// BookNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BookNft contract.
type BookNftApprovalIterator struct {
	Event *BookNftApproval // Event containing the contract specifics and raw log

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
func (it *BookNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftApproval)
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
		it.Event = new(BookNftApproval)
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
func (it *BookNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftApproval represents a Approval event raised by the BookNft contract.
type BookNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BookNft *BookNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BookNftApprovalIterator, error) {

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

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookNftApprovalIterator{contract: _BookNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BookNft *BookNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BookNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftApproval)
				if err := _BookNft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BookNft *BookNftFilterer) ParseApproval(log types.Log) (*BookNftApproval, error) {
	event := new(BookNftApproval)
	if err := _BookNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BookNft contract.
type BookNftApprovalForAllIterator struct {
	Event *BookNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BookNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftApprovalForAll)
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
		it.Event = new(BookNftApprovalForAll)
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
func (it *BookNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftApprovalForAll represents a ApprovalForAll event raised by the BookNft contract.
type BookNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BookNft *BookNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BookNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BookNftApprovalForAllIterator{contract: _BookNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BookNft *BookNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BookNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftApprovalForAll)
				if err := _BookNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_BookNft *BookNftFilterer) ParseApprovalForAll(log types.Log) (*BookNftApprovalForAll, error) {
	event := new(BookNftApprovalForAll)
	if err := _BookNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the BookNft contract.
type BookNftContractURIUpdatedIterator struct {
	Event *BookNftContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *BookNftContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftContractURIUpdated)
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
		it.Event = new(BookNftContractURIUpdated)
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
func (it *BookNftContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftContractURIUpdated represents a ContractURIUpdated event raised by the BookNft contract.
type BookNftContractURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_BookNft *BookNftFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*BookNftContractURIUpdatedIterator, error) {

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &BookNftContractURIUpdatedIterator{contract: _BookNft.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_BookNft *BookNftFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *BookNftContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftContractURIUpdated)
				if err := _BookNft.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_BookNft *BookNftFilterer) ParseContractURIUpdated(log types.Log) (*BookNftContractURIUpdated, error) {
	event := new(BookNftContractURIUpdated)
	if err := _BookNft.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BookNft contract.
type BookNftInitializedIterator struct {
	Event *BookNftInitialized // Event containing the contract specifics and raw log

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
func (it *BookNftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftInitialized)
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
		it.Event = new(BookNftInitialized)
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
func (it *BookNftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftInitialized represents a Initialized event raised by the BookNft contract.
type BookNftInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BookNft *BookNftFilterer) FilterInitialized(opts *bind.FilterOpts) (*BookNftInitializedIterator, error) {

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BookNftInitializedIterator{contract: _BookNft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BookNft *BookNftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BookNftInitialized) (event.Subscription, error) {

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftInitialized)
				if err := _BookNft.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BookNft *BookNftFilterer) ParseInitialized(log types.Log) (*BookNftInitialized, error) {
	event := new(BookNftInitialized)
	if err := _BookNft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BookNft contract.
type BookNftOwnershipTransferredIterator struct {
	Event *BookNftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BookNftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftOwnershipTransferred)
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
		it.Event = new(BookNftOwnershipTransferred)
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
func (it *BookNftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftOwnershipTransferred represents a OwnershipTransferred event raised by the BookNft contract.
type BookNftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BookNft *BookNftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BookNftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BookNftOwnershipTransferredIterator{contract: _BookNft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BookNft *BookNftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BookNftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftOwnershipTransferred)
				if err := _BookNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BookNft *BookNftFilterer) ParseOwnershipTransferred(log types.Log) (*BookNftOwnershipTransferred, error) {
	event := new(BookNftOwnershipTransferred)
	if err := _BookNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BookNft contract.
type BookNftRoleAdminChangedIterator struct {
	Event *BookNftRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BookNftRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftRoleAdminChanged)
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
		it.Event = new(BookNftRoleAdminChanged)
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
func (it *BookNftRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftRoleAdminChanged represents a RoleAdminChanged event raised by the BookNft contract.
type BookNftRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BookNft *BookNftFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BookNftRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BookNftRoleAdminChangedIterator{contract: _BookNft.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BookNft *BookNftFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BookNftRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftRoleAdminChanged)
				if err := _BookNft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BookNft *BookNftFilterer) ParseRoleAdminChanged(log types.Log) (*BookNftRoleAdminChanged, error) {
	event := new(BookNftRoleAdminChanged)
	if err := _BookNft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BookNft contract.
type BookNftRoleGrantedIterator struct {
	Event *BookNftRoleGranted // Event containing the contract specifics and raw log

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
func (it *BookNftRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftRoleGranted)
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
		it.Event = new(BookNftRoleGranted)
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
func (it *BookNftRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftRoleGranted represents a RoleGranted event raised by the BookNft contract.
type BookNftRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BookNftRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BookNftRoleGrantedIterator{contract: _BookNft.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BookNftRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftRoleGranted)
				if err := _BookNft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) ParseRoleGranted(log types.Log) (*BookNftRoleGranted, error) {
	event := new(BookNftRoleGranted)
	if err := _BookNft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BookNft contract.
type BookNftRoleRevokedIterator struct {
	Event *BookNftRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BookNftRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftRoleRevoked)
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
		it.Event = new(BookNftRoleRevoked)
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
func (it *BookNftRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftRoleRevoked represents a RoleRevoked event raised by the BookNft contract.
type BookNftRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BookNftRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BookNftRoleRevokedIterator{contract: _BookNft.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BookNftRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftRoleRevoked)
				if err := _BookNft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BookNft *BookNftFilterer) ParseRoleRevoked(log types.Log) (*BookNftRoleRevoked, error) {
	event := new(BookNftRoleRevoked)
	if err := _BookNft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BookNft contract.
type BookNftTransferIterator struct {
	Event *BookNftTransfer // Event containing the contract specifics and raw log

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
func (it *BookNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftTransfer)
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
		it.Event = new(BookNftTransfer)
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
func (it *BookNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftTransfer represents a Transfer event raised by the BookNft contract.
type BookNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BookNft *BookNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BookNftTransferIterator, error) {

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

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookNftTransferIterator{contract: _BookNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BookNft *BookNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BookNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftTransfer)
				if err := _BookNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BookNft *BookNftFilterer) ParseTransfer(log types.Log) (*BookNftTransfer, error) {
	event := new(BookNftTransfer)
	if err := _BookNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookNftTransferWithMemoIterator is returned from FilterTransferWithMemo and is used to iterate over the raw logs and unpacked data for TransferWithMemo events raised by the BookNft contract.
type BookNftTransferWithMemoIterator struct {
	Event *BookNftTransferWithMemo // Event containing the contract specifics and raw log

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
func (it *BookNftTransferWithMemoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookNftTransferWithMemo)
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
		it.Event = new(BookNftTransferWithMemo)
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
func (it *BookNftTransferWithMemoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookNftTransferWithMemoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookNftTransferWithMemo represents a TransferWithMemo event raised by the BookNft contract.
type BookNftTransferWithMemo struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferWithMemo is a free log retrieval operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_BookNft *BookNftFilterer) FilterTransferWithMemo(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BookNftTransferWithMemoIterator, error) {

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

	logs, sub, err := _BookNft.contract.FilterLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookNftTransferWithMemoIterator{contract: _BookNft.contract, event: "TransferWithMemo", logs: logs, sub: sub}, nil
}

// WatchTransferWithMemo is a free log subscription operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_BookNft *BookNftFilterer) WatchTransferWithMemo(opts *bind.WatchOpts, sink chan<- *BookNftTransferWithMemo, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BookNft.contract.WatchLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookNftTransferWithMemo)
				if err := _BookNft.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
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

// ParseTransferWithMemo is a log parse operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_BookNft *BookNftFilterer) ParseTransferWithMemo(log types.Log) (*BookNftTransferWithMemo, error) {
	event := new(BookNftTransferWithMemo)
	if err := _BookNft.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
