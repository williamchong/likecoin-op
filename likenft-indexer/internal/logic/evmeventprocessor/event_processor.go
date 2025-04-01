package evmeventprocessor

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
)

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

func registerEventProcessor(event string, creator eventProcessorCreator) error {
	_, hasEvent := eventProcessorMap[event]
	if hasEvent {
		return fmt.Errorf("event already registered")
	}
	eventProcessorMap[event] = creator
	return nil
}

func getEventProcessor(event string) (eventProcessorCreator, error) {
	eventProcessCreator, ok := eventProcessorMap[event]
	if !ok {
		return nil, fmt.Errorf("unknown event %s", event)
	}
	return eventProcessCreator, nil
}
