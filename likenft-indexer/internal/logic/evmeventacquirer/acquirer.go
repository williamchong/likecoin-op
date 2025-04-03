package evmeventacquirer

import (
	"context"
	"log/slog"
	"reflect"
	"runtime"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/util/logconverter"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type EvmEventsAcquirer interface {
	Acquire(
		ctx context.Context,
		logger *slog.Logger,
		contractAddress string,
		event evmeventprocessedblockheight.Event,
	) error
}

type evmEventsAcquirer struct {
	evmEventProcessedBlockHeightRepository database.EVMEventProcessedBlockHeightRepository
	evmEventRepository                     database.EVMEventRepository
	deps                                   *eventAcquirerDeps
}

var _ EvmEventsAcquirer = &evmEventsAcquirer{}

func MakeEvmEventsAcquirer(
	evmEventProcessedBlockHeightRepository database.EVMEventProcessedBlockHeightRepository,
	evmEventRepository database.EVMEventRepository,
	evmClient evm.EVMQueryClient,
) EvmEventsAcquirer {
	return &evmEventsAcquirer{
		evmEventProcessedBlockHeightRepository: evmEventProcessedBlockHeightRepository,
		evmEventRepository:                     evmEventRepository,
		deps: makeEventProcessorDeps(
			evmClient,
		),
	}
}

func (a *evmEventsAcquirer) Acquire(ctx context.Context, logger *slog.Logger, contractAddress string, event evmeventprocessedblockheight.Event) error {
	eventConfigCreator, err := getEventConfig(event)
	if err != nil {
		return err
	}
	eventConfig := eventConfigCreator(a.deps)

	return a.acquire(
		ctx,
		logger,

		eventConfig.ContractType,
		event,
		eventConfig.Abi,
		eventConfig.LogsRetriever,

		contractAddress,
	)
}

func (a *evmEventsAcquirer) acquire(
	ctx context.Context,
	logger *slog.Logger,

	contractType evmeventprocessedblockheight.ContractType,
	event evmeventprocessedblockheight.Event,
	abi *abi.ABI,
	logsRetriever LogsRetriever,

	contractAddress string,
) error {
	mylogger := logger.WithGroup("AcquireNewEvmEvents").
		With("contractAddress", contractAddress).
		With("contractType", contractType).
		With("event", event)

	dbProcessedBlockHeight, err := a.evmEventProcessedBlockHeightRepository.GetEVMEventProcessedBlockHeight(
		ctx,
		contractType,
		contractAddress,
		event,
	)

	processedBlockHeight := uint64(0)
	if err != nil {
		if ent.IsNotFound(err) {
			processedBlockHeight = 0
		} else {
			mylogger.Error("database.GetEVMEventProcessedBlockHeight", "err", err)
			return err
		}
	} else {
		processedBlockHeight = dbProcessedBlockHeight
	}
	startBlock := processedBlockHeight + 1

	mylogger.Info("querying events...", "startBlock", startBlock)
	logs, err := logsRetriever(ctx, mylogger, contractAddress, startBlock)
	if err != nil {
		mylogger.Error("logsRetriever", "fn", runtime.FuncForPC(reflect.ValueOf(logsRetriever).Pointer()).Name(), "err", err)
		return err
	}
	mylogger.Info("events found", "count", len(logs))

	logConverter := logconverter.NewLogConverter(abi)

	evmEvents := make([]*ent.EVMEvent, 0, len(logs))

	highestBlockHeight := processedBlockHeight

	for _, log := range logs {
		evmEvent, err := logConverter.ConvertLogToEvmEvent(log)

		if err != nil {
			mylogger.Error("logConverter.ConvertLogToEvmEvent", "err", err)
			return err
		}

		evmEvents = append(evmEvents, evmEvent)

		highestBlockHeight = max(highestBlockHeight, log.BlockNumber)
	}

	mylogger.Info("inserting evm events...")
	_, err = a.evmEventRepository.InsertEvmEventsIfNeeded(ctx, evmEvents)

	if err != nil {
		mylogger.Error("database.InsertEvmEventsIfNeeded", "err", err)
		return err
	}

	mylogger.Info("updating processed block height...")
	err = a.evmEventProcessedBlockHeightRepository.UpdateEVMEventProcessedBlockHeight(
		ctx,
		contractType,
		contractAddress,
		event,
		highestBlockHeight,
	)

	if err != nil {
		return err
	}

	return nil

}
