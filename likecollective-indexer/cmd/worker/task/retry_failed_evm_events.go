package task

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	appcontext "likecollective-indexer/cmd/worker/context"
	"likecollective-indexer/ent/evmevent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/model"
	"likecollective-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeRetryFailedEVMEventsPayload = "retry-failed-evm-events"

// retryFailedEVMEventsLimit caps how many failed events one cycle re-drives.
// The query is otherwise unbounded, and an event that can never succeed stays
// at `failed` and is picked up again every cycle, so without a cap a handful of
// poison events would flood the process-evm-event queue indefinitely.
const retryFailedEVMEventsLimit = 200

type RetryFailedEVMEventsPayload struct {
}

func NewRetryFailedEVMEventsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(RetryFailedEVMEventsPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeRetryFailedEVMEventsPayload,
		payload,
		asynq.Queue(TypeRetryFailedEVMEventsPayload),
	), nil
}

// HandleRetryFailedEVMEvents re-drives events that exhausted their asynq
// retries and were parked at `failed`.
//
// check-received-evm-events only ever queries `received`, so before this task
// existed a failed event stayed failed forever. That is not a delayed number:
// account, nft class and staking totals are accumulated from event deltas, so a
// permanently skipped event offsets them permanently, and every later event
// builds on the wrong base.
//
// The status is deliberately left at `failed`. EVMEventProcessor.Process
// accepts `failed` and re-drives it, and leaving it alone means an event whose
// re-drive dies mid-flight is picked up again next cycle rather than stranded
// in `enqueued`.
//
// The cost of that choice is that an event which can never succeed is re-driven
// every cycle forever, so the batch is capped and the cron is slower than
// check-received-evm-events. That bounds the waste but does not end it: with
// more than retryFailedEVMEventsLimit permanently failing events, the oldest
// would crowd out newer ones, since the cap takes them in block order. Ending
// it properly needs a persisted attempt count on evm_events to retire an event
// after N tries; until then a persistently failing event stays visible in the
// logs, which is the behaviour we want from a backstop.
func HandleRetryFailedEVMEvents(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)

	mylogger := logger.WithGroup("HandleRetryFailedEVMEvents")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p RetryFailedEVMEventsPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal RetryFailedEVMEventsPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()

	evmEventRepository := database.MakeEVMEventRepository(dbService)

	failedEvents, err := evmEventRepository.GetEVMEventsByStatusWithLimit(
		ctx,
		evmevent.StatusFailed,
		retryFailedEVMEventsLimit,
	)

	if err != nil {
		return err
	}

	if len(failedEvents) == 0 {
		return nil
	}

	// Deltas are applied in block order, so re-drive in the same order the
	// normal path uses.
	slices.SortFunc(failedEvents, model.EvmEventsProcessingComparator)

	mylogger.Warn(fmt.Sprintf("%d failed events found, re-driving", len(failedEvents)))

	for _, evmEvent := range failedEvents {
		if evmEvent.FailedReason != nil {
			mylogger.Warn("re-driving failed event",
				"evmEventId", evmEvent.ID,
				"failedReason", *evmEvent.FailedReason,
			)
		}
		if err := enqueueProcessEVMEvent(mylogger, asynqClient, evmEvent); err != nil {
			continue
		}
	}

	return nil
}

func init() {
	t := Tasks.Register(task.DefineTask(
		TypeRetryFailedEVMEventsPayload,
		HandleRetryFailedEVMEvents,
	))
	PeriodicTasks.Register(task.DefinePeriodicTask(
		t,
		NewRetryFailedEVMEventsTask,
	))
}
