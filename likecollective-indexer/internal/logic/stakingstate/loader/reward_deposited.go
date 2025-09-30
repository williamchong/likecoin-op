package loader

import (
	"likecollective-indexer/ent"

	"github.com/ethereum/go-ethereum/common"
)

type loadRewardDepositedStateFactory struct {
	stakingEvent *ent.StakingEvent

	accountEVMAddress common.Address
	nftClassAddress   common.Address
}

func MakeLoadRewardDepositedStateFactory(stakingEvent *ent.StakingEvent) LoadStateFactory {
	accountEVMAddress := common.HexToAddress(stakingEvent.AccountEvmAddress)
	nftClassAddress := common.HexToAddress(stakingEvent.NftClassAddress)
	return &loadRewardDepositedStateFactory{
		stakingEvent,
		accountEVMAddress,
		nftClassAddress,
	}
}

func (f *loadRewardDepositedStateFactory) MakeLoadState() *LoadState {
	return &LoadState{
		AccountByAddress:           MakeAddressList(),
		AccountByBookNFT:           MakeAddressList(f.nftClassAddress),
		BookNFTByAddress:           MakeAddressList(f.nftClassAddress),
		StakingByAccountAndBookNFT: MakeStakingKeyList(),
		StakingByBookNFT:           MakeAddressList(f.nftClassAddress),
	}
}
