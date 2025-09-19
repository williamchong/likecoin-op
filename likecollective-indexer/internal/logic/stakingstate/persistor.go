package stakingstate

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
)

type StakingStatePersistor interface {
	Persist(
		ctx context.Context,
		stakingEvents []*ent.StakingEvent,
		accounts []Account,
		nftClasses []NFTClass,
		stakings []Staking,
	) error
}

type latestStakingStatePersistor struct {
	accountRepository      database.AccountRepository
	nftClassRepository     database.NFTClassRepository
	stakingRepository      database.StakingRepository
	stakingEventRepository database.StakingEventRepository
	dbService              database.Service
}

func MakeStakingStatePersistor(
	dbService database.Service,
) StakingStatePersistor {
	accountRepository := database.MakeAccountRepository(dbService)
	nftClassRepository := database.MakeNFTClassRepository(dbService)
	stakingRepository := database.MakeStakingRepository(dbService)
	stakingEventRepository := database.MakeStakingEventRepository(dbService)
	return &latestStakingStatePersistor{
		accountRepository,
		nftClassRepository,
		stakingRepository,
		stakingEventRepository,
		dbService,
	}
}

func (p *latestStakingStatePersistor) Persist(
	ctx context.Context,
	stakingEvents []*ent.StakingEvent,
	accounts []Account,
	nftClasses []NFTClass,
	stakings []Staking,
) error {
	err := database.WithTx(ctx, p.dbService.Client(), func(tx *ent.Tx) error {
		_, err := p.stakingEventRepository.InsertStakingEventsIfNeeded(
			ctx,
			stakingEvents,
		)
		if err != nil {
			return err
		}

		for _, account := range accounts {
			_, err := p.accountRepository.CreateOrUpdateAccount(
				ctx,
				tx,
				account.EVMAddress.String(),
				account.StakedAmount,
				account.PendingRewardAmount,
				account.ClaimedRewardAmount,
			)
			if err != nil {
				return err
			}
		}

		for _, nftClass := range nftClasses {
			_, err := p.nftClassRepository.CreateOrUpdateNFTClass(
				ctx,
				tx,
				nftClass.EVMAddress.String(),
				nftClass.StakedAmount,
			)
			if err != nil {
				return err
			}
		}

		for _, staking := range stakings {
			_, err := p.stakingRepository.CreateOrUpdateStaking(
				ctx,
				tx,
				staking.BookNFTEvmAddress.String(),
				staking.AccountEVMAddress.String(),
				staking.StakedAmount,
				staking.PendingRewardAmount,
				staking.ClaimedRewardAmount,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
