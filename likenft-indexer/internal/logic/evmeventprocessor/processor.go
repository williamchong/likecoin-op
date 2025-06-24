package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"runtime"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
)

var ErrAlreadyProcessing = errors.New("evmevent already processing")

type EVMEventProcessor interface {
	Process(
		ctx context.Context,
		logger *slog.Logger,
		evmEvent *ent.EVMEvent,
	) error
}

type evmEventProcessor struct {
	httpClient                *http.Client
	evmClient                 *evm.EvmClient
	nftRepository             database.NFTRepository
	nftClassRepository        database.NFTClassRepository
	evmEventRepository        database.EVMEventRepository
	transactionMemoRepository database.TransactionMemoRepository
	accountRepository         database.AccountRepository
}

func MakeEVMEventProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftRepository database.NFTRepository,
	nftClassRepository database.NFTClassRepository,
	evmEventRepository database.EVMEventRepository,
	transactionMemoRepository database.TransactionMemoRepository,
	accountRepository database.AccountRepository,
) EVMEventProcessor {
	return &evmEventProcessor{
		httpClient:                httpClient,
		evmClient:                 evmClient,
		nftRepository:             nftRepository,
		nftClassRepository:        nftClassRepository,
		evmEventRepository:        evmEventRepository,
		transactionMemoRepository: transactionMemoRepository,
		accountRepository:         accountRepository,
	}
}

func (e *evmEventProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) (err error) {
	mylogger := logger.WithGroup("Process").
		With("evmEventId", evmEvent.ID)

	switch evmEvent.Status {
	case evmevent.StatusSkipped:
		mylogger.Warn("The event has already skipped")
		return nil
	case evmevent.StatusProcessing:
		mylogger.Warn("The event has already processing")
		return ErrAlreadyProcessing
	case evmevent.StatusProcessed:
		mylogger.Info("The event has already processed")
		return nil
	case evmevent.StatusFailed:
		mylogger.Info("record failed. retrying...")
	case evmevent.StatusReceived:
	case evmevent.StatusEnqueued:
		break
	default:
		return fmt.Errorf("unknown status %v", evmEvent.Status)
	}

	defer func() {
		if err != nil {
			errMsg := err.Error()
			mylogger.Error("something went wrong", "err", err)
			_, _ = e.evmEventRepository.UpdateEvmEventStatus(ctx, evmEvent, evmevent.StatusFailed, &errMsg)
		}
	}()

	evmEvent, err = e.evmEventRepository.UpdateEvmEventStatus(ctx, evmEvent, evmevent.StatusProcessing, nil)

	if err != nil {
		mylogger.Error("e.evmEventRepository.UpdateEvmEventStatus", "err", err)
		return err
	}

	processorCreator, err := e.getProcessorCreator(evmEvent)
	if err != nil {
		if errors.Is(err, UnknownEvent) {
			evmEvent, err = e.evmEventRepository.UpdateEvmEventStatus(ctx, evmEvent, evmevent.StatusSkipped, nil)
			return nil
		}
		mylogger.Error("e.getProcessorCreator", "err", err)
		return err
	}

	mylogger = mylogger.With(
		"processorCreator",
		runtime.FuncForPC(reflect.ValueOf(processorCreator).Pointer()).Name())

	processor := processorCreator(makeEventProcessorDeps(
		e.httpClient,
		e.evmClient,
		e.nftRepository,
		e.nftClassRepository,
		e.transactionMemoRepository,
		e.accountRepository,
	))

	err = processor.Process(ctx, logger, evmEvent)
	if err != nil {
		mylogger.Error("processor.Process", "err", err)
		return err
	}

	evmEvent, err = e.evmEventRepository.UpdateEvmEventStatus(ctx, evmEvent, evmevent.StatusProcessed, nil)
	mylogger.Info("updated evm event status to processed")

	if err != nil {
		mylogger.Error("e.evmEventRepository.UpdateEvmEventStatus", "err", err)
		return err
	}

	return nil
}

func (e *evmEventProcessor) getProcessorCreator(
	evmEvent *ent.EVMEvent,
) (eventProcessorCreator, error) {
	processorCreator, err := getEventProcessor(evmEvent.Topic0)
	if err != nil {
		return nil, err
	}
	return processorCreator, nil
}
