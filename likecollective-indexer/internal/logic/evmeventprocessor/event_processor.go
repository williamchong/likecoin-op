package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate"
)

var UnknownEvent error = errors.New("unknown event")

type eventProcessorDeps struct {
	evmEventRepository    database.EVMEventRepository
	accountRepository     database.AccountRepository
	nftClassRepository    database.NFTClassRepository
	stakingRepository     database.StakingRepository
	stakingStatePersistor stakingstate.StakingStatePersistor
}

func makeEventProcessorDeps(
	evmEventRepository database.EVMEventRepository,
	accountRepository database.AccountRepository,
	nftClassRepository database.NFTClassRepository,
	stakingRepository database.StakingRepository,
	stakingStatePersistor stakingstate.StakingStatePersistor,
) *eventProcessorDeps {
	return &eventProcessorDeps{
		evmEventRepository,
		accountRepository,
		nftClassRepository,
		stakingRepository,
		stakingStatePersistor,
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
