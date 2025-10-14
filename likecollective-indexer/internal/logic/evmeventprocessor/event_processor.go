package evmeventprocessor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/evm"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"

	"github.com/ethereum/go-ethereum/common"
)

var UnknownEvent error = errors.New("unknown event")

type eventProcessorDeps struct {
	evmClient             evm.EVMClient
	stakingStateLoader    stakingstateloader.StakingStateLoader
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor

	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
}

func makeEventProcessorDeps(
	evmClient evm.EVMClient,
	stakingStateLoader stakingstateloader.StakingStateLoader,
	stakingStatePersistor stakingstatepersistor.StakingStatePersistor,
	likeCollectiveAddress common.Address,
	likeStakePositionAddress common.Address,
) *eventProcessorDeps {
	return &eventProcessorDeps{
		evmClient,
		stakingStateLoader,
		stakingStatePersistor,
		likeCollectiveAddress,
		likeStakePositionAddress,
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
