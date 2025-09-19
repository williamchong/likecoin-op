package stakingstate

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
)

type stakingEventProcessor struct {
	evmEventRepository    database.EVMEventRepository
	accountRepository     database.AccountRepository
	nftClassRepository    database.NFTClassRepository
	stakingRepository     database.StakingRepository
	stakingStatePersistor StakingStatePersistor
}

func MakeStakingEventProcessor(
	evmEventRepository database.EVMEventRepository,
	accountRepository database.AccountRepository,
	nftClassRepository database.NFTClassRepository,
	stakingRepository database.StakingRepository,
	stakingStatePersistor StakingStatePersistor,
) *stakingEventProcessor {
	return &stakingEventProcessor{
		evmEventRepository,
		accountRepository,
		nftClassRepository,
		stakingRepository,
		stakingStatePersistor,
	}
}

func (e *stakingEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,
	evmEvents []*ent.EVMEvent,
) error {
	stakingEvents := make([]*ent.StakingEvent, 0)

	for _, evmEvent := range evmEvents {
		stakingEvent, err := GetStakingEventsFromEvent(evmEvent)
		if err != nil {
			return err
		}
		stakingEvents = append(stakingEvents, stakingEvent...)
	}

	accountKeys := GetAccountKeysFromEvents(stakingEvents)
	bookNFTKeys := GetBookNFTKeysFromEvents(stakingEvents)
	stakingKeys := GetStakingKeysFromEvents(stakingEvents)

	snapshottedAccounts, err := e.accountRepository.QueryAccountsByEvmAddresses(ctx, accountKeys)
	if err != nil {
		return err
	}

	snapshottedNftClasses, err := e.nftClassRepository.QueryNFTClassesByAddresses(ctx, bookNFTKeys)
	if err != nil {
		return err
	}

	snapshottedStakings, err := e.stakingRepository.QueryStakingsByKeys(ctx, stakingKeys)
	if err != nil {
		return err
	}

	stakingState := NewStakingStateFromEnt(
		snapshottedAccounts,
		snapshottedNftClasses,
		snapshottedStakings,
	)

	stakingState, err = stakingState.HandleStakingEvents(stakingEvents)
	if err != nil {
		return err
	}

	err = stakingState.Persist(ctx, stakingEvents, e.stakingStatePersistor)
	if err != nil {
		return err
	}

	return nil
}
