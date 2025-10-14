package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/logic/stakingstate"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"

	"github.com/ethereum/go-ethereum/common"
)

type transferEventProcessor struct {
	evmClient             evm.EVMClient
	stakingStateLoader    stakingstateloader.StakingStateLoader
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor

	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
}

func MakeTransferEventProcessor(
	inj *eventProcessorDeps,
) eventProcessor {
	return &transferEventProcessor{
		inj.evmClient,
		inj.stakingStateLoader,
		inj.stakingStatePersistor,
		inj.likeCollectiveAddress,
		inj.likeStakePositionAddress,
	}
}

func (e *transferEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	if common.HexToAddress(evmEvent.Address) == e.likeStakePositionAddress {
		return e.processStakePosition(ctx, logger, evmEvent)
	}

	return errors.Join(
		UnknownEvent,
		fmt.Errorf("no candidate to process transfer event"),
	)
}

func (e *transferEventProcessor) processStakePosition(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	if common.HexToAddress(*evmEvent.Topic1) == common.HexToAddress("0x0") {
		return errors.Join(
			UnknownEvent,
			fmt.Errorf(
				"skip handling initial transfer event",
			),
		)
	}

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
		"Transfer",
		MakeTransferEventProcessor,
	)
}
