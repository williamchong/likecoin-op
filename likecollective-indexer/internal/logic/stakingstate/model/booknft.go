package model

import (
	"likecollective-indexer/ent"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type NFTClass struct {
	EVMAddress   common.Address
	StakedAmount *uint256.Int
}

func NewNFTClass(evmAddress string) *NFTClass {
	return &NFTClass{
		EVMAddress:   common.HexToAddress(evmAddress),
		StakedAmount: uint256.NewInt(0),
	}
}

func NewNFTClassFromEnt(nftClass *ent.NFTClass) *NFTClass {
	accounts := make(map[string]*Account)
	for _, account := range nftClass.Edges.Accounts {
		accounts[account.EvmAddress] = NewAccountFromEnt(account)
	}
	return &NFTClass{
		EVMAddress:   common.HexToAddress(nftClass.Address),
		StakedAmount: nftClass.StakedAmount,
	}
}
