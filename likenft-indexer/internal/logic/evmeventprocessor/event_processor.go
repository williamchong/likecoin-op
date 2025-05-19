package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
)

var UnknownEvent error = errors.New("unknown event")

type eventProcessorDeps struct {
	httpClient                *http.Client
	evmClient                 *evm.EvmClient
	nftRepository             database.NFTRepository
	nftClassRepository        database.NFTClassRepository
	transactionMemoRepository database.TransactionMemoRepository
	accountRepository         database.AccountRepository
}

func makeEventProcessorDeps(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftRepository database.NFTRepository,
	nftClassRepository database.NFTClassRepository,
	transactionMemoRepository database.TransactionMemoRepository,
	accountRepository database.AccountRepository,
) *eventProcessorDeps {
	return &eventProcessorDeps{
		httpClient:                httpClient,
		evmClient:                 evmClient,
		nftRepository:             nftRepository,
		nftClassRepository:        nftClassRepository,
		transactionMemoRepository: transactionMemoRepository,
		accountRepository:         accountRepository,
	}
}

type eventProcessor interface {
	Process(
		ctx context.Context,
		logger *slog.Logger,

		evmEvent *ent.EVMEvent,
	) error
}

type eventProcessorCreator func(inj *eventProcessorDeps) eventProcessor

var eventProcessorMap = make(map[string]eventProcessorCreator)

func registerEventProcessor(event string, creator eventProcessorCreator) {
	_, hasEvent := eventProcessorMap[event]
	if hasEvent {
		panic(fmt.Errorf("event %s already registered", event))
	}
	eventProcessorMap[event] = creator
}

func getEventProcessor(event string) (eventProcessorCreator, error) {
	eventProcessCreator, ok := eventProcessorMap[event]
	if !ok {
		return nil, errors.Join(UnknownEvent, fmt.Errorf("%s", event))
	}
	return eventProcessCreator, nil
}
