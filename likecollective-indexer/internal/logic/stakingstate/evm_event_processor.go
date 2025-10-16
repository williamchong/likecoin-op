package stakingstate

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/logic/stakingstate/loader"
	"likecollective-indexer/internal/logic/stakingstate/persistor"

	"github.com/ethereum/go-ethereum/common"
)

type StakingEvmEventProcessor interface {
	Process(
		ctx context.Context,
		logger *slog.Logger,
		evmEvents []*ent.EVMEvent,
	) error
}

type stakingEvmEventProcessor struct {
	evmClient             evm.EVMClient
	stakingStateLoader    loader.StakingStateLoader
	stakingStatePersistor persistor.StakingStatePersistor

	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
}

func MakeStakingEvmEventProcessor(
	evmClient evm.EVMClient,
	stakingStateLoader loader.StakingStateLoader,
	stakingStatePersistor persistor.StakingStatePersistor,
	likeCollectiveAddress common.Address,
	likeStakePositionAddress common.Address,
) StakingEvmEventProcessor {
	return &stakingEvmEventProcessor{
		evmClient,
		stakingStateLoader,
		stakingStatePersistor,
		likeCollectiveAddress,
		likeStakePositionAddress,
	}
}

func (e *stakingEvmEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,
	evmEvents []*ent.EVMEvent,
) error {
	stakingEvents := make([]*ent.StakingEvent, 0)

	for _, evmEvent := range evmEvents {
		var (
			stakingEvent []*ent.StakingEvent
			err          error
		)
		if common.HexToAddress(evmEvent.Address) == e.likeCollectiveAddress {
			stakingEvent, err = GetStakingEventsFromLikeCollectiveEvent(ctx, e.evmClient, evmEvent)
		} else if common.HexToAddress(evmEvent.Address) == e.likeStakePositionAddress {
			stakingEvent, err = GetStakingEventsFromLikeStakePositionEvent(ctx, e.evmClient, evmEvent)
		}
		if err != nil {
			return err
		}
		stakingEvents = append(stakingEvents, stakingEvent...)
	}

	stakingState, err := LoadStakingState(ctx, e.stakingStateLoader, stakingEvents)
	if err != nil {
		return err
	}

	stakingState, processedStakingEvents, err := stakingState.Process(stakingEvents)
	if err != nil {
		return err
	}

	err = stakingState.Persist(ctx, processedStakingEvents, e.stakingStatePersistor)
	if err != nil {
		return err
	}

	return nil
}
