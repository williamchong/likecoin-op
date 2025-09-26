package loader

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"

	"github.com/ethereum/go-ethereum/common"
)

type loadAllRewardClaimedStateFactory struct {
	stakingEvent *ent.StakingEvent

	accountEVMAddress common.Address
	nftClassAddress   common.Address
}

func MakeLoadAllRewardsClaimedStateFactory(stakingEvent *ent.StakingEvent) LoadStateFactory {
	accountEVMAddress := common.HexToAddress(stakingEvent.AccountEvmAddress)
	nftClassAddress := common.HexToAddress(stakingEvent.NftClassAddress)
	return &loadAllRewardClaimedStateFactory{
		stakingEvent,
		accountEVMAddress,
		nftClassAddress,
	}
}

func (f *loadAllRewardClaimedStateFactory) MakeLoadState() *LoadState {
	return &LoadState{
		AccountByAddress: MakeAddressList(f.accountEVMAddress),
		AccountByBookNFT: MakeAddressList(),
		BookNFTByAddress: MakeAddressList(f.nftClassAddress),
		StakingByAccountAndBookNFT: MakeStakingKeyList(database.NewStakingKey(
			f.accountEVMAddress.String(),
			f.nftClassAddress.String(),
		)),
		StakingByBookNFT: MakeAddressList(),
	}
}
