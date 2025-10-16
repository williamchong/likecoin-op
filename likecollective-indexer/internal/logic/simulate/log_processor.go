package simulate

import (
	"context"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/logic/stakingstate"

	"github.com/ethereum/go-ethereum/common"
)

type SimulateLogProcessor interface {
	Process(
		ctx context.Context,
		logger *slog.Logger,
		logsWithHeaders []SimulationLogWithHeader,
	) error
}

type simulateLogProcessor struct {
	stakingEvmEventProcessor stakingstate.StakingEvmEventProcessor

	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
}

func MakeSimulateLogProcessor(
	stakingEvmEventProcessor stakingstate.StakingEvmEventProcessor,

	likeCollectiveAddress common.Address,
	likeStakePositionAddress common.Address,
) SimulateLogProcessor {
	return &simulateLogProcessor{
		stakingEvmEventProcessor,
		likeCollectiveAddress,
		likeStakePositionAddress,
	}
}

func (e *simulateLogProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,
	logsWithHeaders []SimulationLogWithHeader,
) error {
	evmEvents := make([]*ent.EVMEvent, 0)
	for _, logWithHeader := range logsWithHeaders {
		var (
			evmEvent *ent.EVMEvent
			err      error
		)
		switch logWithHeader.Log.Address {
		case e.likeCollectiveAddress:
			evmEvent, err = likeCollectiveLogConverter.ConvertLogToEvmEvent(logWithHeader.Log, logWithHeader.Header)
		case e.likeStakePositionAddress:
			evmEvent, err = likeStakePositionLogConverter.ConvertLogToEvmEvent(logWithHeader.Log, logWithHeader.Header)
		}
		if err != nil {
			return err
		}
		evmEvents = append(evmEvents, evmEvent)
	}
	return e.stakingEvmEventProcessor.Process(ctx, logger, evmEvents)
}
