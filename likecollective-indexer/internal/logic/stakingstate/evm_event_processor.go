package stakingstate

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/logic/stakingstate/loader"
	"likecollective-indexer/internal/logic/stakingstate/persistor"
)

type stakingEvmEventProcessor struct {
	stakingStateLoader    loader.StakingStateLoader
	stakingStatePersistor persistor.StakingStatePersistor
}

func MakeStakingEvmEventProcessor(
	stakingStateLoader loader.StakingStateLoader,
	stakingStatePersistor persistor.StakingStatePersistor,
) *stakingEvmEventProcessor {
	return &stakingEvmEventProcessor{
		stakingStateLoader,
		stakingStatePersistor,
	}
}

func (e *stakingEvmEventProcessor) Process(
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
