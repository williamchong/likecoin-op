package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likecollective-indexer/cmd/worker/context"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/evmeventprocessor"
	stakingstateloader "likecollective-indexer/internal/logic/stakingstate/loader"
	stakingstatepersistor "likecollective-indexer/internal/logic/stakingstate/persistor"
	"likecollective-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeProcessEVMEventPayload = "process-evm-event"
const TypeIndexActionProcessEVMEventPayload = "index-action-process-evm-event"

type ProcessEVMEventPayload struct {
	EvmEventId int
}

func NewProcessEVMEvent(evmEventId int) (*asynq.Task, error) {
	payload, err := json.Marshal(ProcessEVMEventPayload{
		EvmEventId: evmEventId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeProcessEVMEventPayload,
		payload,
		asynq.Queue(TypeProcessEVMEventPayload),
	), nil
}

func NewIndexActionProcessEVMEvent(evmEventId int) (*asynq.Task, error) {
	payload, err := json.Marshal(ProcessEVMEventPayload{
		EvmEventId: evmEventId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeIndexActionProcessEVMEventPayload,
		payload,
		asynq.Queue(TypeIndexActionProcessEVMEventPayload),
	), nil
}

func HandleProcessEVMEvent(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleProcessEVMEvent")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p ProcessEVMEventPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal ProcessEVMEventPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()

	evmEventRepository := database.MakeEVMEventRepository(dbService)
	accountRepository := database.MakeAccountRepository(dbService)
	nftClassRepository := database.MakeNFTClassRepository(dbService)
	stakingRepository := database.MakeStakingRepository(dbService)

	stakingStateLoader := stakingstateloader.MakeStakingStateLoader(
		accountRepository,
		nftClassRepository,
		stakingRepository,
	)
	stakingStatePersistor := stakingstatepersistor.MakeStakingStatePersistor(dbService)

	processor := evmeventprocessor.MakeEVMEventProcessor(
		evmEventRepository,
		stakingStateLoader,
		stakingStatePersistor,
	)

	evmEvent, err := evmEventRepository.GetEvmEventById(ctx, p.EvmEventId)

	if err != nil {
		mylogger.Error("evmEventRepository.GetEvmEventById", "err", err)
		return err
	}

	err = processor.Process(ctx, logger, evmEvent)

	if err != nil {
		mylogger.Error("processor.Proces", "err", err)
		return err
	}

	return nil
}

func init() {
	Tasks.Register(task.DefineTask(
		TypeProcessEVMEventPayload,
		HandleProcessEVMEvent,
	))
	Tasks.Register(task.DefineTask(
		TypeIndexActionProcessEVMEventPayload,
		HandleProcessEVMEvent,
	))
}
