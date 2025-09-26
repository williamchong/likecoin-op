package persistor

import (
	"context"
	"fmt"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate/model"
)

type StakingStatePersistor interface {
	Persist(
		ctx context.Context,
		stakingEvents []*ent.StakingEvent,
		accounts []*model.Account,
		nftClasses []*model.NFTClass,
		stakings []*model.Staking,
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
	accounts []*model.Account,
	nftClasses []*model.NFTClass,
	stakings []*model.Staking,
) error {
	err := database.WithTx(ctx, p.dbService.Client(), func(tx *ent.Tx) error {
		_, err := p.stakingEventRepository.InsertStakingEventsIfNeeded(
			ctx,
			tx,
			stakingEvents,
		)
		if err != nil {
			return fmt.Errorf("failed to insert staking events if needed: %w", err)
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
				return fmt.Errorf("failed to create or update account: %w", err)
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
				return fmt.Errorf("failed to create or update nft class: %w", err)
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
				return fmt.Errorf("failed to create or update staking: %w", err)
			}
		}

		for _, nftClass := range nftClasses {
			_, err := p.nftClassRepository.RecomputeNumberOfStakersByNFTClassAddress(
				ctx,
				tx,
				nftClass.EVMAddress.String(),
			)
			if err != nil {
				return fmt.Errorf("failed to recompute number of stakers by nft class address: %w", err)
			}

			err = p.stakingRepository.RecomputePoolSharesByNFTClassAddress(
				ctx,
				tx,
				nftClass.EVMAddress.String(),
			)
			if err != nil {
				return fmt.Errorf("failed to recompute pool shares by nft class address: %w", err)
			}
		}

		return nil
	})

	return err
}
