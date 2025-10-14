package evmeventprocessor

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/logic/stakingstate"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"

	"github.com/ethereum/go-ethereum/common"
)

type stakingEvmEventProcessor struct {
	evmClient             evm.EVMClient
	stakingStateLoader    stakingstateloader.StakingStateLoader
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor

	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
}

func MakeStakingEvmEventProcessor(
	inj *eventProcessorDeps,
) eventProcessor {
	return &stakingEvmEventProcessor{
		inj.evmClient,
		inj.stakingStateLoader,
		inj.stakingStatePersistor,
		inj.likeCollectiveAddress,
		inj.likeStakePositionAddress,
	}
}

func (e *stakingEvmEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	p := stakingstate.MakeStakingEvmEventProcessor(
		e.evmClient,
		e.stakingStateLoader,
		e.stakingStatePersistor,
		e.likeCollectiveAddress,
		e.likeStakePositionAddress,
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
