// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package likenft_class

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

// ClassConfig is an auto generated low-level Go binding around an user-defined struct.
type ClassConfig struct {
	MaxSupply uint64
}

// ClassInput is an auto generated low-level Go binding around an user-defined struct.
type ClassInput struct {
	Name     string
	Symbol   string
	Metadata string
	Config   ClassConfig
}

// MsgNewClass is an auto generated low-level Go binding around an user-defined struct.
type MsgNewClass struct {
	Creator  common.Address
	Updaters []common.Address
	Minters  []common.Address
	Input    ClassInput
}

// LikenftClassMetaData contains all meta data concerning the LikenftClass contract.
var LikenftClassMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structClassConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structClassInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewClass\",\"name\":\"msgNewClass\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNftNoSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferWithMemo\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"metadataList\",\"type\":\"string[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferWithMemo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structClassConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structClassInput\",\"name\":\"classInput\",\"type\":\"tuple\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161277738038061277783398101604081905261002e9161047d565b8051606082015180516020909101515f610048838261069b565b506001610055828261069b565b5050506001600160a01b03811661008557604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61008e816101e5565b506060810151517f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f2009081906100c3908261069b565b5060608201516020015160018201906100dc908261069b565b5060608201516040015160028201906100f5908261069b565b506060828101510151516003820180546001600160401b0319166001600160401b039092169190911790555f600c8190555b82604001515181101561018a576101817f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68460400151838151811061016e5761016e610755565b602002602001015161023660201b60201c565b50600101610127565b505f5b8260200151518110156101dd576101d47f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab8460200151838151811061016e5761016e610755565b5060010161018d565b505050610769565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f828152600b602090815260408083206001600160a01b038516845290915281205460ff166102da575f838152600b602090815260408083206001600160a01b03861684529091529020805460ff191660011790556102923390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102dd565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b0381118282101715610319576103196102e3565b60405290565b604051602081016001600160401b0381118282101715610319576103196102e3565b604051601f8201601f191681016001600160401b0381118282101715610369576103696102e3565b604052919050565b80516001600160a01b0381168114610387575f5ffd5b919050565b5f82601f83011261039b575f5ffd5b81516001600160401b038111156103b4576103b46102e3565b8060051b6103c460208201610341565b918252602081850181019290810190868411156103df575f5ffd5b6020860192505b83831015610408576103f783610371565b8252602092830192909101906103e6565b9695505050505050565b5f82601f830112610421575f5ffd5b81516001600160401b0381111561043a5761043a6102e3565b61044d601f8201601f1916602001610341565b818152846020838601011115610461575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f6020828403121561048d575f5ffd5b81516001600160401b038111156104a2575f5ffd5b8201608081850312156104b3575f5ffd5b6104bb6102f7565b6104c482610371565b815260208201516001600160401b038111156104de575f5ffd5b6104ea8682850161038c565b60208301525060408201516001600160401b03811115610508575f5ffd5b6105148682850161038c565b60408301525060608201516001600160401b03811115610532575f5ffd5b91909101908185036080811215610547575f5ffd5b61054f6102f7565b83516001600160401b03811115610564575f5ffd5b61057088828701610412565b82525060208401516001600160401b0381111561058b575f5ffd5b61059788828701610412565b60208301525060408401516001600160401b038111156105b5575f5ffd5b6105c188828701610412565b6040830152506020605f19830112156105d8575f5ffd5b6105e061031f565b606094909401519391506001600160401b03841684146105fe575f5ffd5b9281526060838101919091528101919091529392505050565b600181811c9082168061062b57607f821691505b60208210810361064957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561069657805f5260205f20601f840160051c810160208510156106745750805b601f840160051c820191505b81811015610693575f8155600101610680565b50505b505050565b81516001600160401b038111156106b4576106b46102e3565b6106c8816106c28454610617565b8461064f565b6020601f8211600181146106fa575f83156106e35750848201515b5f19600385901b1c1916600184901b178455610693565b5f84815260208120601f198516915b828110156107295787850151825560209485019460019092019101610709565b508482101561074657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b612001806107765f395ff3fe6080604052600436106101d0575f3560e01c8063715018a6116100fd578063c87b56dd11610092578063dcb49c7311610062578063dcb49c731461055c578063e8a3d4851461056f578063e985e9c514610583578063f2fde38b146105a2575f5ffd5b8063c87b56dd146104cc578063d5391393146104eb578063d547741f1461051e578063d90794cf1461053d575f5ffd5b806395d89b41116100cd57806395d89b4114610467578063a217fddf1461047b578063a22cb4651461048e578063b88d4fde146104ad575f5ffd5b8063715018a6146103f8578063765a15bb1461040c5780638da5cb5b1461042b57806391d1485414610448575f5ffd5b80632f2ff15d1161017357806347e633801161014357806347e63380146103685780634f6ccce71461039b5780636352211e146103ba57806370a08231146103d9575f5ffd5b80632f2ff15d146102ec5780632f745c591461030b57806336568abe1461032a57806342842e0e14610349575f5ffd5b8063095ea7b3116101ae578063095ea7b31461026057806318160ddd1461028157806323b872dd1461029f578063248a9ca3146102be575f5ffd5b806301ffc9a7146101d457806306fdde0314610208578063081812fc14610229575b5f5ffd5b3480156101df575f5ffd5b506101f36101ee3660046116c5565b6105c1565b60405190151581526020015b60405180910390f35b348015610213575f5ffd5b5061021c6105d1565b6040516101ff9190611715565b348015610234575f5ffd5b50610248610243366004611727565b610672565b6040516001600160a01b0390911681526020016101ff565b34801561026b575f5ffd5b5061027f61027a366004611759565b610699565b005b34801561028c575f5ffd5b506008545b6040519081526020016101ff565b3480156102aa575f5ffd5b5061027f6102b9366004611781565b6106a8565b3480156102c9575f5ffd5b506102916102d8366004611727565b5f908152600b602052604090206001015490565b3480156102f7575f5ffd5b5061027f6103063660046117bb565b610736565b348015610316575f5ffd5b50610291610325366004611759565b61075a565b348015610335575f5ffd5b5061027f6103443660046117bb565b6107bd565b348015610354575f5ffd5b5061027f610363366004611781565b6107f5565b348015610373575f5ffd5b506102917f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab81565b3480156103a6575f5ffd5b506102916103b5366004611727565b61080f565b3480156103c5575f5ffd5b506102486103d4366004611727565b610864565b3480156103e4575f5ffd5b506102916103f33660046117e5565b61086e565b348015610403575f5ffd5b5061027f6108b3565b348015610417575f5ffd5b5061027f6104263660046118ef565b6108c6565b348015610436575f5ffd5b50600a546001600160a01b0316610248565b348015610453575f5ffd5b506101f36104623660046117bb565b610981565b348015610472575f5ffd5b5061021c6109ab565b348015610486575f5ffd5b506102915f81565b348015610499575f5ffd5b5061027f6104a83660046119eb565b6109e9565b3480156104b8575f5ffd5b5061027f6104c7366004611a24565b6109f4565b3480156104d7575f5ffd5b5061021c6104e6366004611727565b610a0c565b3480156104f6575f5ffd5b506102917f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b348015610529575f5ffd5b5061027f6105383660046117bb565b610a44565b348015610548575f5ffd5b5061027f610557366004611a9a565b610a68565b61027f61056a366004611b19565b610bae565b34801561057a575f5ffd5b5061021c610c0e565b34801561058e575f5ffd5b506101f361059d366004611bac565b610c56565b3480156105ad575f5ffd5b5061027f6105bc3660046117e5565b610c83565b5f6105cb82610cc0565b92915050565b5f516020611fac5f395f51905f5280546060919081906105f090611bd4565b80601f016020809104026020016040519081016040528092919081815260200182805461061c90611bd4565b80156106675780601f1061063e57610100808354040283529160200191610667565b820191905f5260205f20905b81548152906001019060200180831161064a57829003601f168201915b505050505091505090565b5f61067c82610ce4565b505f828152600460205260409020546001600160a01b03166105cb565b6106a4828233610d1c565b5050565b6001600160a01b0382166106d657604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f6106e2838333610d29565b9050836001600160a01b0316816001600160a01b031614610730576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016106cd565b50505050565b5f828152600b602052604090206001015461075081610dfc565b6107308383610e06565b5f6107648361086e565b82106107955760405163295f44f760e21b81526001600160a01b0384166004820152602481018390526044016106cd565b506001600160a01b03919091165f908152600660209081526040808320938352929052205490565b6001600160a01b03811633146107e65760405163334bd91960e11b815260040160405180910390fd5b6107f08282610e97565b505050565b6107f083838360405180602001604052805f8152506109f4565b5f61081960085490565b82106108415760405163295f44f760e21b81525f6004820152602481018390526044016106cd565b6008828154811061085457610854611c0c565b905f5260205f2001549050919050565b5f6105cb82610ce4565b5f6001600160a01b038216610898576040516322718ad960e21b81525f60048201526024016106cd565b506001600160a01b03165f9081526003602052604090205490565b6108bb610f02565b6108c45f610f2f565b565b6108f07f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab32610981565b61090d57604051636609677b60e11b815260040160405180910390fd5b80515f516020611fac5f395f51905f5290819061092a9082611c64565b506020820151600182019061093f9082611c64565b50604082015160028201906109549082611c64565b506040517fa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962905f90a15050565b5f918252600b602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f20180546060915f516020611fac5f395f51905f52916105f090611bd4565b6106a4338383610f80565b6109ff8484846106a8565b610730338585858561101e565b5f818152600d60209081526040918290209151606092610a2e92909101611d1e565b6040516020818303038152906040529050919050565b5f828152600b6020526040902060010154610a5e81610dfc565b6107308383610e97565b610a927f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a632610981565b610aaf57604051636609677b60e11b815260040160405180910390fd5b7f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f203545f516020611fac5f395f51905f52906001600160401b0316828115801590610b145750816001600160401b031681610b0860085490565b610b129190611dd2565b115b15610b3257604051636a29267160e01b815260040160405180910390fd5b5f5b81811015610ba557858582818110610b4e57610b4e611c0c565b9050602002810190610b609190611de5565b600c545f908152600d6020526040902091610b7c919083611e2e565b50610b8987600c54611146565b600c8054905f610b9883611ee7565b9091555050600101610b34565b50505050505050565b610bb98585856106a8565b82846001600160a01b0316866001600160a01b03167fbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b8585604051610bff929190611eff565b60405180910390a45050505050565b604051606090610c42907f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f20290602001611d1e565b604051602081830303815290604052905090565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b610c8b610f02565b6001600160a01b038116610cb457604051631e4fbdf760e01b81525f60048201526024016106cd565b610cbd81610f2f565b50565b5f6001600160e01b03198216637965db0b60e01b14806105cb57506105cb8261115f565b5f818152600260205260408120546001600160a01b0316806105cb57604051637e27328960e01b8152600481018490526024016106cd565b6107f08383836001611183565b5f5f610d36858585611287565b90506001600160a01b038116610d9257610d8d84600880545f838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610db5565b846001600160a01b0316816001600160a01b031614610db557610db58185611379565b6001600160a01b038516610dd157610dcc846113f6565b610df4565b846001600160a01b0316816001600160a01b031614610df457610df4858561149d565b949350505050565b610cbd81336114eb565b5f610e118383610981565b610e90575f838152600b602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610e483390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016105cb565b505f6105cb565b5f610ea28383610981565b15610e90575f838152600b602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016105cb565b600a546001600160a01b031633146108c45760405163118cdaa760e01b81523360048201526024016106cd565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b038216610fb257604051630b61174360e31b81526001600160a01b03831660048201526024016106cd565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b1561113f57604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290611060908890889087908790600401611f2d565b6020604051808303815f875af192505050801561109a575060408051601f3d908101601f1916820190925261109791810190611f69565b60015b611101573d8080156110c7576040519150601f19603f3d011682016040523d82523d5f602084013e6110cc565b606091505b5080515f036110f957604051633250574960e11b81526001600160a01b03851660048201526024016106cd565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461113d57604051633250574960e11b81526001600160a01b03851660048201526024016106cd565b505b5050505050565b6106a4828260405180602001604052805f815250611524565b5f6001600160e01b0319821663780e9d6360e01b14806105cb57506105cb8261153b565b808061119757506001600160a01b03821615155b15611258575f6111a684610ce4565b90506001600160a01b038316158015906111d25750826001600160a01b0316816001600160a01b031614155b80156111e557506111e38184610c56565b155b1561120e5760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016106cd565b81156112565783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f90815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b5f828152600260205260408120546001600160a01b03908116908316156112b3576112b381848661158a565b6001600160a01b038116156112ed576112ce5f855f5f611183565b6001600160a01b0381165f90815260036020526040902080545f190190555b6001600160a01b0385161561131b576001600160a01b0385165f908152600360205260409020805460010190555b5f8481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b5f6113838361086e565b5f838152600760209081526040808320546001600160a01b03881684526006909252909120919250908183146113d8575f83815260208281526040808320548584528184208190558352600790915290208290555b5f938452600760209081526040808620869055938552525081205550565b6008545f9061140790600190611f84565b5f838152600960205260408120546008805493945090928490811061142e5761142e611c0c565b905f5260205f2001549050806008838154811061144d5761144d611c0c565b5f91825260208083209091019290925582815260099091526040808220849055858252812055600880548061148457611484611f97565b600190038181905f5260205f20015f9055905550505050565b5f60016114a98461086e565b6114b39190611f84565b6001600160a01b039093165f908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6114f58282610981565b6106a45760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016106cd565b61152e83836115ee565b6107f0335f85858561101e565b5f6001600160e01b031982166380ac58cd60e01b148061156b57506001600160e01b03198216635b5e139f60e01b145b806105cb57506301ffc9a760e01b6001600160e01b03198316146105cb565b61159583838361164f565b6107f0576001600160a01b0383166115c357604051637e27328960e01b8152600481018290526024016106cd565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016106cd565b6001600160a01b03821661161757604051633250574960e11b81525f60048201526024016106cd565b5f61162383835f610d29565b90506001600160a01b038116156107f0576040516339e3563760e11b81525f60048201526024016106cd565b5f6001600160a01b03831615801590610df45750826001600160a01b0316846001600160a01b0316148061168857506116888484610c56565b80610df45750505f908152600460205260409020546001600160a01b03908116911614919050565b6001600160e01b031981168114610cbd575f5ffd5b5f602082840312156116d5575f5ffd5b81356116e0816116b0565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6116e060208301846116e7565b5f60208284031215611737575f5ffd5b5035919050565b80356001600160a01b0381168114611754575f5ffd5b919050565b5f5f6040838503121561176a575f5ffd5b6117738361173e565b946020939093013593505050565b5f5f5f60608486031215611793575f5ffd5b61179c8461173e565b92506117aa6020850161173e565b929592945050506040919091013590565b5f5f604083850312156117cc575f5ffd5b823591506117dc6020840161173e565b90509250929050565b5f602082840312156117f5575f5ffd5b6116e08261173e565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b0381118282101715611834576118346117fe565b60405290565b604051602081016001600160401b0381118282101715611834576118346117fe565b5f5f6001600160401b03841115611875576118756117fe565b50604051601f19601f85018116603f011681018181106001600160401b03821117156118a3576118a36117fe565b6040528381529050808284018510156118ba575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f8301126118e0575f5ffd5b6116e08383356020850161185c565b5f602082840312156118ff575f5ffd5b81356001600160401b03811115611914575f5ffd5b82018084036080811215611926575f5ffd5b61192e611812565b82356001600160401b03811115611943575f5ffd5b61194f878286016118d1565b82525060208301356001600160401b0381111561196a575f5ffd5b611976878286016118d1565b60208301525060408301356001600160401b03811115611994575f5ffd5b6119a0878286016118d1565b6040830152506020605f19830112156119b7575f5ffd5b6119bf61183a565b9150606083013592506001600160401b03831683146119dc575f5ffd5b91815260608201529392505050565b5f5f604083850312156119fc575f5ffd5b611a058361173e565b915060208301358015158114611a19575f5ffd5b809150509250929050565b5f5f5f5f60808587031215611a37575f5ffd5b611a408561173e565b9350611a4e6020860161173e565b92506040850135915060608501356001600160401b03811115611a6f575f5ffd5b8501601f81018713611a7f575f5ffd5b611a8e8782356020840161185c565b91505092959194509250565b5f5f5f60408486031215611aac575f5ffd5b611ab58461173e565b925060208401356001600160401b03811115611acf575f5ffd5b8401601f81018613611adf575f5ffd5b80356001600160401b03811115611af4575f5ffd5b8660208260051b8401011115611b08575f5ffd5b939660209190910195509293505050565b5f5f5f5f5f60808688031215611b2d575f5ffd5b611b368661173e565b9450611b446020870161173e565b93506040860135925060608601356001600160401b03811115611b65575f5ffd5b8601601f81018813611b75575f5ffd5b80356001600160401b03811115611b8a575f5ffd5b886020828401011115611b9b575f5ffd5b959894975092955050506020019190565b5f5f60408385031215611bbd575f5ffd5b611bc68361173e565b91506117dc6020840161173e565b600181811c90821680611be857607f821691505b602082108103611c0657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b601f8211156107f057805f5260205f20601f840160051c81016020851015611c455750805b601f840160051c820191505b8181101561113f575f8155600101611c51565b81516001600160401b03811115611c7d57611c7d6117fe565b611c9181611c8b8454611bd4565b84611c20565b6020601f821160018114611cc3575f8315611cac5750848201515b5f19600385901b1c1916600184901b17845561113f565b5f84815260208120601f198516915b82811015611cf25787850151825560209485019460019092019101611cd2565b5084821015611d0f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b7f646174613a6170706c69636174696f6e2f6a736f6e3b757466382c000000000081525f5f8354611d4e81611bd4565b600182168015611d655760018114611d8057611db3565b60ff198316601b870152601b82151583028701019350611db3565b865f5260205f205f5b83811015611da8578154888201601b0152600190910190602001611d89565b5050601b8287010193505b509195945050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156105cb576105cb611dbe565b5f5f8335601e19843603018112611dfa575f5ffd5b8301803591506001600160401b03821115611e13575f5ffd5b602001915036819003821315611e27575f5ffd5b9250929050565b6001600160401b03831115611e4557611e456117fe565b611e5983611e538354611bd4565b83611c20565b5f601f841160018114611e8a575f8515611e735750838201355b5f19600387901b1c1916600186901b17835561113f565b5f83815260208120601f198716915b82811015611eb95786850135825560209485019460019092019101611e99565b5086821015611ed5575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f60018201611ef857611ef8611dbe565b5060010190565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611f5f908301846116e7565b9695505050505050565b5f60208284031215611f79575f5ffd5b81516116e0816116b0565b818103818111156105cb576105cb611dbe565b634e487b7160e01b5f52603160045260245ffdfe99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f200a26469706673582212205e4d03c2a522aac54f3d4903f7b9f3bb11eb75fffa87af77f861b0ab9b66078664736f6c634300081c0033",
}

// LikenftClassABI is the input ABI used to generate the binding from.
// Deprecated: Use LikenftClassMetaData.ABI instead.
var LikenftClassABI = LikenftClassMetaData.ABI

// LikenftClassBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LikenftClassMetaData.Bin instead.
var LikenftClassBin = LikenftClassMetaData.Bin

// DeployLikenftClass deploys a new Ethereum contract, binding an instance of LikenftClass to it.
func DeployLikenftClass(auth *bind.TransactOpts, backend bind.ContractBackend, msgNewClass MsgNewClass) (common.Address, *types.Transaction, *LikenftClass, error) {
	parsed, err := LikenftClassMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LikenftClassBin), backend, msgNewClass)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LikenftClass{LikenftClassCaller: LikenftClassCaller{contract: contract}, LikenftClassTransactor: LikenftClassTransactor{contract: contract}, LikenftClassFilterer: LikenftClassFilterer{contract: contract}}, nil
}

// LikenftClass is an auto generated Go binding around an Ethereum contract.
type LikenftClass struct {
	LikenftClassCaller     // Read-only binding to the contract
	LikenftClassTransactor // Write-only binding to the contract
	LikenftClassFilterer   // Log filterer for contract events
}

// LikenftClassCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikenftClassCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikenftClassTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikenftClassFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikenftClassSession struct {
	Contract     *LikenftClass     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LikenftClassCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikenftClassCallerSession struct {
	Contract *LikenftClassCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LikenftClassTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikenftClassTransactorSession struct {
	Contract     *LikenftClassTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LikenftClassRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikenftClassRaw struct {
	Contract *LikenftClass // Generic contract binding to access the raw methods on
}

// LikenftClassCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikenftClassCallerRaw struct {
	Contract *LikenftClassCaller // Generic read-only contract binding to access the raw methods on
}

// LikenftClassTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikenftClassTransactorRaw struct {
	Contract *LikenftClassTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikenftClass creates a new instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClass(address common.Address, backend bind.ContractBackend) (*LikenftClass, error) {
	contract, err := bindLikenftClass(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikenftClass{LikenftClassCaller: LikenftClassCaller{contract: contract}, LikenftClassTransactor: LikenftClassTransactor{contract: contract}, LikenftClassFilterer: LikenftClassFilterer{contract: contract}}, nil
}

// NewLikenftClassCaller creates a new read-only instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassCaller(address common.Address, caller bind.ContractCaller) (*LikenftClassCaller, error) {
	contract, err := bindLikenftClass(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikenftClassCaller{contract: contract}, nil
}

// NewLikenftClassTransactor creates a new write-only instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassTransactor(address common.Address, transactor bind.ContractTransactor) (*LikenftClassTransactor, error) {
	contract, err := bindLikenftClass(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransactor{contract: contract}, nil
}

// NewLikenftClassFilterer creates a new log filterer instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassFilterer(address common.Address, filterer bind.ContractFilterer) (*LikenftClassFilterer, error) {
	contract, err := bindLikenftClass(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikenftClassFilterer{contract: contract}, nil
}

// bindLikenftClass binds a generic wrapper to an already deployed contract.
func bindLikenftClass(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikenftClassMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikenftClass *LikenftClassRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikenftClass.Contract.LikenftClassCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikenftClass *LikenftClassRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.Contract.LikenftClassTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikenftClass *LikenftClassRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikenftClass.Contract.LikenftClassTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikenftClass *LikenftClassCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikenftClass.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikenftClass *LikenftClassTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikenftClass *LikenftClassTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikenftClass.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LikenftClass.Contract.DEFAULTADMINROLE(&_LikenftClass.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LikenftClass.Contract.DEFAULTADMINROLE(&_LikenftClass.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) MINTERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.MINTERROLE(&_LikenftClass.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) MINTERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.MINTERROLE(&_LikenftClass.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) UPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) UPDATERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.UPDATERROLE(&_LikenftClass.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) UPDATERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.UPDATERROLE(&_LikenftClass.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikenftClass.Contract.BalanceOf(&_LikenftClass.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikenftClass.Contract.BalanceOf(&_LikenftClass.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassSession) ContractURI() (string, error) {
	return _LikenftClass.Contract.ContractURI(&_LikenftClass.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) ContractURI() (string, error) {
	return _LikenftClass.Contract.ContractURI(&_LikenftClass.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.GetApproved(&_LikenftClass.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.GetApproved(&_LikenftClass.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LikenftClass.Contract.GetRoleAdmin(&_LikenftClass.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LikenftClass.Contract.GetRoleAdmin(&_LikenftClass.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LikenftClass.Contract.HasRole(&_LikenftClass.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LikenftClass.Contract.HasRole(&_LikenftClass.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikenftClass.Contract.IsApprovedForAll(&_LikenftClass.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikenftClass.Contract.IsApprovedForAll(&_LikenftClass.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassSession) Name() (string, error) {
	return _LikenftClass.Contract.Name(&_LikenftClass.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) Name() (string, error) {
	return _LikenftClass.Contract.Name(&_LikenftClass.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassSession) Owner() (common.Address, error) {
	return _LikenftClass.Contract.Owner(&_LikenftClass.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassCallerSession) Owner() (common.Address, error) {
	return _LikenftClass.Contract.Owner(&_LikenftClass.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.OwnerOf(&_LikenftClass.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.OwnerOf(&_LikenftClass.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikenftClass.Contract.SupportsInterface(&_LikenftClass.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikenftClass.Contract.SupportsInterface(&_LikenftClass.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassSession) Symbol() (string, error) {
	return _LikenftClass.Contract.Symbol(&_LikenftClass.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) Symbol() (string, error) {
	return _LikenftClass.Contract.Symbol(&_LikenftClass.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LikenftClass.Contract.TokenByIndex(&_LikenftClass.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LikenftClass.Contract.TokenByIndex(&_LikenftClass.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LikenftClass.Contract.TokenOfOwnerByIndex(&_LikenftClass.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LikenftClass.Contract.TokenOfOwnerByIndex(&_LikenftClass.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _LikenftClass.Contract.TokenURI(&_LikenftClass.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _LikenftClass.Contract.TokenURI(&_LikenftClass.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikenftClass *LikenftClassCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikenftClass *LikenftClassSession) TotalSupply() (*big.Int, error) {
	return _LikenftClass.Contract.TotalSupply(&_LikenftClass.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) TotalSupply() (*big.Int, error) {
	return _LikenftClass.Contract.TotalSupply(&_LikenftClass.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.Approve(&_LikenftClass.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.Approve(&_LikenftClass.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.GrantRole(&_LikenftClass.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.GrantRole(&_LikenftClass.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadataList) returns()
func (_LikenftClass *LikenftClassTransactor) Mint(opts *bind.TransactOpts, to common.Address, metadataList []string) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "mint", to, metadataList)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadataList) returns()
func (_LikenftClass *LikenftClassSession) Mint(to common.Address, metadataList []string) (*types.Transaction, error) {
	return _LikenftClass.Contract.Mint(&_LikenftClass.TransactOpts, to, metadataList)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadataList) returns()
func (_LikenftClass *LikenftClassTransactorSession) Mint(to common.Address, metadataList []string) (*types.Transaction, error) {
	return _LikenftClass.Contract.Mint(&_LikenftClass.TransactOpts, to, metadataList)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceOwnership(&_LikenftClass.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceOwnership(&_LikenftClass.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceRole(&_LikenftClass.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceRole(&_LikenftClass.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RevokeRole(&_LikenftClass.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RevokeRole(&_LikenftClass.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikenftClass *LikenftClassTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikenftClass *LikenftClassSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom0(&_LikenftClass.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LikenftClass *LikenftClassTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom0(&_LikenftClass.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.Contract.SetApprovalForAll(&_LikenftClass.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.Contract.SetApprovalForAll(&_LikenftClass.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferOwnership(&_LikenftClass.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferOwnership(&_LikenftClass.TransactOpts, newOwner)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassTransactor) TransferWithMemo(opts *bind.TransactOpts, from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferWithMemo", from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferWithMemo(&_LikenftClass.TransactOpts, from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferWithMemo(&_LikenftClass.TransactOpts, from, to, _tokenId, memo)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassTransactor) Update(opts *bind.TransactOpts, classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "update", classInput)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassSession) Update(classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.Contract.Update(&_LikenftClass.TransactOpts, classInput)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassTransactorSession) Update(classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.Contract.Update(&_LikenftClass.TransactOpts, classInput)
}

// LikenftClassApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LikenftClass contract.
type LikenftClassApprovalIterator struct {
	Event *LikenftClassApproval // Event containing the contract specifics and raw log

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
func (it *LikenftClassApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassApproval)
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
		it.Event = new(LikenftClassApproval)
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
func (it *LikenftClassApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassApproval represents a Approval event raised by the LikenftClass contract.
type LikenftClassApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LikenftClassApprovalIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassApprovalIterator{contract: _LikenftClass.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LikenftClassApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassApproval)
				if err := _LikenftClass.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseApproval(log types.Log) (*LikenftClassApproval, error) {
	event := new(LikenftClassApproval)
	if err := _LikenftClass.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LikenftClass contract.
type LikenftClassApprovalForAllIterator struct {
	Event *LikenftClassApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LikenftClassApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassApprovalForAll)
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
		it.Event = new(LikenftClassApprovalForAll)
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
func (it *LikenftClassApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassApprovalForAll represents a ApprovalForAll event raised by the LikenftClass contract.
type LikenftClassApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikenftClass *LikenftClassFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LikenftClassApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassApprovalForAllIterator{contract: _LikenftClass.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikenftClass *LikenftClassFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LikenftClassApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassApprovalForAll)
				if err := _LikenftClass.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseApprovalForAll(log types.Log) (*LikenftClassApprovalForAll, error) {
	event := new(LikenftClassApprovalForAll)
	if err := _LikenftClass.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the LikenftClass contract.
type LikenftClassContractURIUpdatedIterator struct {
	Event *LikenftClassContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *LikenftClassContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassContractURIUpdated)
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
		it.Event = new(LikenftClassContractURIUpdated)
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
func (it *LikenftClassContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassContractURIUpdated represents a ContractURIUpdated event raised by the LikenftClass contract.
type LikenftClassContractURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_LikenftClass *LikenftClassFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*LikenftClassContractURIUpdatedIterator, error) {

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &LikenftClassContractURIUpdatedIterator{contract: _LikenftClass.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_LikenftClass *LikenftClassFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *LikenftClassContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassContractURIUpdated)
				if err := _LikenftClass.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseContractURIUpdated(log types.Log) (*LikenftClassContractURIUpdated, error) {
	event := new(LikenftClassContractURIUpdated)
	if err := _LikenftClass.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LikenftClass contract.
type LikenftClassOwnershipTransferredIterator struct {
	Event *LikenftClassOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LikenftClassOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassOwnershipTransferred)
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
		it.Event = new(LikenftClassOwnershipTransferred)
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
func (it *LikenftClassOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassOwnershipTransferred represents a OwnershipTransferred event raised by the LikenftClass contract.
type LikenftClassOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikenftClass *LikenftClassFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LikenftClassOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassOwnershipTransferredIterator{contract: _LikenftClass.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikenftClass *LikenftClassFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LikenftClassOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassOwnershipTransferred)
				if err := _LikenftClass.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseOwnershipTransferred(log types.Log) (*LikenftClassOwnershipTransferred, error) {
	event := new(LikenftClassOwnershipTransferred)
	if err := _LikenftClass.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LikenftClass contract.
type LikenftClassRoleAdminChangedIterator struct {
	Event *LikenftClassRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LikenftClassRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleAdminChanged)
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
		it.Event = new(LikenftClassRoleAdminChanged)
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
func (it *LikenftClassRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleAdminChanged represents a RoleAdminChanged event raised by the LikenftClass contract.
type LikenftClassRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LikenftClass *LikenftClassFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LikenftClassRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleAdminChangedIterator{contract: _LikenftClass.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LikenftClass *LikenftClassFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleAdminChanged)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseRoleAdminChanged(log types.Log) (*LikenftClassRoleAdminChanged, error) {
	event := new(LikenftClassRoleAdminChanged)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LikenftClass contract.
type LikenftClassRoleGrantedIterator struct {
	Event *LikenftClassRoleGranted // Event containing the contract specifics and raw log

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
func (it *LikenftClassRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleGranted)
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
		it.Event = new(LikenftClassRoleGranted)
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
func (it *LikenftClassRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleGranted represents a RoleGranted event raised by the LikenftClass contract.
type LikenftClassRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LikenftClassRoleGrantedIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleGrantedIterator{contract: _LikenftClass.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleGranted)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseRoleGranted(log types.Log) (*LikenftClassRoleGranted, error) {
	event := new(LikenftClassRoleGranted)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LikenftClass contract.
type LikenftClassRoleRevokedIterator struct {
	Event *LikenftClassRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LikenftClassRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleRevoked)
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
		it.Event = new(LikenftClassRoleRevoked)
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
func (it *LikenftClassRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleRevoked represents a RoleRevoked event raised by the LikenftClass contract.
type LikenftClassRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LikenftClassRoleRevokedIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleRevokedIterator{contract: _LikenftClass.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleRevoked)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseRoleRevoked(log types.Log) (*LikenftClassRoleRevoked, error) {
	event := new(LikenftClassRoleRevoked)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LikenftClass contract.
type LikenftClassTransferIterator struct {
	Event *LikenftClassTransfer // Event containing the contract specifics and raw log

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
func (it *LikenftClassTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassTransfer)
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
		it.Event = new(LikenftClassTransfer)
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
func (it *LikenftClassTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassTransfer represents a Transfer event raised by the LikenftClass contract.
type LikenftClassTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LikenftClassTransferIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransferIterator{contract: _LikenftClass.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LikenftClassTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassTransfer)
				if err := _LikenftClass.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseTransfer(log types.Log) (*LikenftClassTransfer, error) {
	event := new(LikenftClassTransfer)
	if err := _LikenftClass.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassTransferWithMemoIterator is returned from FilterTransferWithMemo and is used to iterate over the raw logs and unpacked data for TransferWithMemo events raised by the LikenftClass contract.
type LikenftClassTransferWithMemoIterator struct {
	Event *LikenftClassTransferWithMemo // Event containing the contract specifics and raw log

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
func (it *LikenftClassTransferWithMemoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassTransferWithMemo)
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
		it.Event = new(LikenftClassTransferWithMemo)
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
func (it *LikenftClassTransferWithMemoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassTransferWithMemoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassTransferWithMemo represents a TransferWithMemo event raised by the LikenftClass contract.
type LikenftClassTransferWithMemo struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferWithMemo is a free log retrieval operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_LikenftClass *LikenftClassFilterer) FilterTransferWithMemo(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LikenftClassTransferWithMemoIterator, error) {

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

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransferWithMemoIterator{contract: _LikenftClass.contract, event: "TransferWithMemo", logs: logs, sub: sub}, nil
}

// WatchTransferWithMemo is a free log subscription operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_LikenftClass *LikenftClassFilterer) WatchTransferWithMemo(opts *bind.WatchOpts, sink chan<- *LikenftClassTransferWithMemo, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassTransferWithMemo)
				if err := _LikenftClass.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
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
func (_LikenftClass *LikenftClassFilterer) ParseTransferWithMemo(log types.Log) (*LikenftClassTransferWithMemo, error) {
	event := new(LikenftClassTransferWithMemo)
	if err := _LikenftClass.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
