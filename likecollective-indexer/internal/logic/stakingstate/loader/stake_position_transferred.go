package loader

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"

	"github.com/ethereum/go-ethereum/common"
)

type loadStakePositionTransferredStateFactory struct {
	stakingEvent *ent.StakingEvent

	accountEVMAddress common.Address
	nftClassAddress   common.Address
}

func MakeLoadStakePositionTransferredStateFactory(stakingEvent *ent.StakingEvent) LoadStateFactory {
	accountEVMAddress := common.HexToAddress(stakingEvent.AccountEvmAddress)
	nftClassAddress := common.HexToAddress(stakingEvent.NftClassAddress)
	return &loadStakePositionTransferredStateFactory{
		stakingEvent,
		accountEVMAddress,
		nftClassAddress,
	}
}

func (f *loadStakePositionTransferredStateFactory) MakeLoadState() *LoadState {
	return &LoadState{
		AccountByAddress: MakeAddressList(f.accountEVMAddress),
		AccountByBookNFT: MakeAddressList(),
		BookNFTByAddress: MakeAddressList(),
		StakingByAccountAndBookNFT: MakeStakingKeyList(database.NewStakingKey(
			f.accountEVMAddress.String(),
			f.nftClassAddress.String(),
		)),
		StakingByBookNFT: MakeAddressList(),
	}
}
