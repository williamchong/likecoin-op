package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"likecollective-indexer/ent"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"
)

var UnknownEvent error = errors.New("unknown event")

type eventProcessorDeps struct {
	stakingStateLoader    stakingstateloader.StakingStateLoader
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor
}

func makeEventProcessorDeps(
	stakingStateLoader stakingstateloader.StakingStateLoader,
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor,
) *eventProcessorDeps {
	return &eventProcessorDeps{
		stakingStateLoader,
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
