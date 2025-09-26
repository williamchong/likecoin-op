package model

import (
	"likecollective-indexer/ent"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Account struct {
	EVMAddress          common.Address
	StakedAmount        *uint256.Int
	PendingRewardAmount *uint256.Int
	ClaimedRewardAmount *uint256.Int
}

func NewAccount(evmAddress string) *Account {
	return &Account{
		EVMAddress:          common.HexToAddress(evmAddress),
		StakedAmount:        uint256.NewInt(0),
		PendingRewardAmount: uint256.NewInt(0),
		ClaimedRewardAmount: uint256.NewInt(0),
	}
}

func NewAccountFromEnt(account *ent.Account) *Account {
	return &Account{
		EVMAddress:          common.HexToAddress(account.EvmAddress),
		StakedAmount:        account.StakedAmount,
		PendingRewardAmount: account.PendingRewardAmount,
		ClaimedRewardAmount: account.ClaimedRewardAmount,
	}
}
