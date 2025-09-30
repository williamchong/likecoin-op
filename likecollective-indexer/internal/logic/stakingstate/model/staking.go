package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Staking struct {
	AccountEVMAddress   common.Address
	BookNFTEvmAddress   common.Address
	StakedAmount        *uint256.Int
	PendingRewardAmount *uint256.Int
	ClaimedRewardAmount *uint256.Int
}

func NewStaking(accountEVMAddress string, bookNFTEvmAddress string) *Staking {
	return &Staking{
		AccountEVMAddress:   common.HexToAddress(accountEVMAddress),
		BookNFTEvmAddress:   common.HexToAddress(bookNFTEvmAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewStakingFromStakingKey(stakingKey database.StakingKey) *Staking {
	return &Staking{
		AccountEVMAddress:   common.HexToAddress(stakingKey.AccountEVMAddress),
		BookNFTEvmAddress:   common.HexToAddress(stakingKey.BookNFTEVMAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewStakingFromEnt(staking *ent.Staking) *Staking {
	return &Staking{
		AccountEVMAddress:   common.HexToAddress(staking.Edges.Account.EvmAddress),
		BookNFTEvmAddress:   common.HexToAddress(staking.Edges.NftClass.Address),
		StakedAmount:        staking.StakedAmount,
		PendingRewardAmount: staking.PendingRewardAmount,
		ClaimedRewardAmount: staking.ClaimedRewardAmount,
	}
}
