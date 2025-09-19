package stakingstate

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Account struct {
	EVMAddress          common.Address
	StakedAmount        *uint256.Int
	PendingRewardAmount *uint256.Int
	ClaimedRewardAmount *uint256.Int
}

func NewAccount(evmAddress string) Account {
	return Account{
		EVMAddress:          common.HexToAddress(evmAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewAccountFromEnt(account *ent.Account) Account {
	return Account{
		EVMAddress:          common.HexToAddress(account.EvmAddress),
		StakedAmount:        account.StakedAmount,
		PendingRewardAmount: account.PendingRewardAmount,
		ClaimedRewardAmount: account.ClaimedRewardAmount,
	}
}

type NFTClass struct {
	EVMAddress   common.Address
	StakedAmount *uint256.Int
}

func NewNFTClass(evmAddress string) NFTClass {
	return NFTClass{
		EVMAddress:   common.HexToAddress(evmAddress),
		StakedAmount: uint256.NewInt(0),
	}
}

func NewNFTClassFromEnt(nftClass *ent.NFTClass) NFTClass {
	accounts := make(map[string]Account)
	for _, account := range nftClass.Edges.Accounts {
		accounts[account.EvmAddress] = NewAccountFromEnt(account)
	}
	return NFTClass{
		EVMAddress:   common.HexToAddress(nftClass.Address),
		StakedAmount: nftClass.StakedAmount,
	}
}

type Staking struct {
	AccountEVMAddress   common.Address
	BookNFTEvmAddress   common.Address
	StakedAmount        *uint256.Int
	PendingRewardAmount *uint256.Int
	ClaimedRewardAmount *uint256.Int
}

func NewStaking(accountEVMAddress string, bookNFTEvmAddress string) Staking {
	return Staking{
		AccountEVMAddress:   common.HexToAddress(accountEVMAddress),
		BookNFTEvmAddress:   common.HexToAddress(bookNFTEvmAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewStakingFromStakingKey(stakingKey database.StakingKey) Staking {
	return Staking{
		AccountEVMAddress:   common.HexToAddress(stakingKey.AccountEVMAddress),
		BookNFTEvmAddress:   common.HexToAddress(stakingKey.BookNFTEVMAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewStakingFromEnt(staking *ent.Staking) Staking {
	return Staking{
		AccountEVMAddress:   common.HexToAddress(staking.Edges.Account.EvmAddress),
		BookNFTEvmAddress:   common.HexToAddress(staking.Edges.NftClass.Address),
		StakedAmount:        staking.StakedAmount,
		PendingRewardAmount: staking.PendingRewardAmount,
		ClaimedRewardAmount: staking.ClaimedRewardAmount,
	}
}
