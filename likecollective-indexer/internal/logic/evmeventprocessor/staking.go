package evmeventprocessor

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/logic/stakingstate"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"
)

type stakingEvmEventProcessor struct {
	stakingStateLoader    stakingstateloader.StakingStateLoader
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor
}

func MakeStakingEvmEventProcessor(
	inj *eventProcessorDeps,
) eventProcessor {
	return &stakingEvmEventProcessor{
		inj.stakingStateLoader,
		inj.stakingStatePersistor,
	}
}

func (e *stakingEvmEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	p := stakingstate.MakeStakingEvmEventProcessor(
		e.stakingStateLoader,
		e.stakingStatePersistor,
	)

	return p.Process(ctx, logger, []*ent.EVMEvent{evmEvent})
}

func init() {
	registerEventProcessor(
		"Staked",
		MakeStakingEvmEventProcessor,
	)
	registerEventProcessor(
		"Unstaked",
		MakeStakingEvmEventProcessor,
	)
	registerEventProcessor(
		"RewardClaimed",
		MakeStakingEvmEventProcessor,
	)
	registerEventProcessor(
		"RewardDeposited",
		MakeStakingEvmEventProcessor,
	)
	registerEventProcessor(
		"AllRewardClaimed",
		MakeStakingEvmEventProcessor,
	)
}
