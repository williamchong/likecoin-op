package evmeventprocessor

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate"
)

type stakingEvmEventProcessor struct {
	evmEventRepository    database.EVMEventRepository
	accountRepository     database.AccountRepository
	nftClassRepository    database.NFTClassRepository
	stakingRepository     database.StakingRepository
	stakingStatePersistor stakingstate.StakingStatePersistor
}

func MakeStakingEvmEventProcessor(
	inj *eventProcessorDeps,
) eventProcessor {
	return &stakingEvmEventProcessor{
		inj.evmEventRepository,
		inj.accountRepository,
		inj.nftClassRepository,
		inj.stakingRepository,
		inj.stakingStatePersistor,
	}
}

func (e *stakingEvmEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	p := stakingstate.MakeStakingEventProcessor(
		e.evmEventRepository,
		e.accountRepository,
		e.nftClassRepository,
		e.stakingRepository,
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
