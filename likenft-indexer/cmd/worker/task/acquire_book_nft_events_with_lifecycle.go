package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/nftclassacquirebooknftevent"
	"likenft-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeAcquireBookNFTEventsTaskPayloadWithLifecyclePayload = "acquire-book-nft-events-with-lifecycle"

type AcquireBookNFTEventsTaskPayloadWithLifecyclePayload struct {
	ContractAddress string
}

func NewTypeAcquireBookNFTEventsTaskPayloadWithLifecycle(contractAddress string) (*asynq.Task, error) {
	payload, err := json.Marshal(AcquireBookNFTEventsTaskPayloadWithLifecyclePayload{
		ContractAddress: contractAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeAcquireBookNFTEventsTaskPayloadWithLifecyclePayload,
		payload,
		asynq.Queue(TypeAcquireBookNFTEventsTaskPayloadWithLifecyclePayload),
	), nil
}

func handlerWithLifecycle(
	handler func(ctx context.Context, t *asynq.Task) error,
) func(ctx context.Context, t *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		config := appcontext.ConfigFromContext(ctx)
		logger := appcontext.LoggerFromContext(ctx)
		dbService := database.New()
		nftClassAcquireBookNFTEventsRepository := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		mylogger := logger.WithGroup("HandlerWithLifecycle")

		var p AcquireBookNFTEventsTaskPayloadWithLifecyclePayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal: %v", err)
		}

		task, err := NewAcquireBookNFTEventsTask(
			[]string{p.ContractAddress},
		)
		if err != nil {
			return fmt.Errorf("NewAcquireBookNFTEventsTask: %v", err)
		}
		lifecycle, err := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycleFromAddress(
			ctx,
			nftClassAcquireBookNFTEventsRepository,
			p.ContractAddress,
			nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				config.TaskAcquireBookNFTEventsNextProcessingScoreBlockHeightContribution,
				config.TaskAcquireBookNFTEventsNextProcessingScoreWeight0Constant,
				config.TaskAcquireBookNFTEventsNextProcessingScoreWeight1Constant,
				config.TaskAcquireBookNFTEventsNextProcessingScoreWeightContribution,
			),
			nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(
				config.TaskAcquireBookNFTEventsInProgressTimeoutSeconds,
			),
			nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				config.TaskAcquireBookNFTEventsRetryInitialTimeoutSeconds,
				config.TaskAcquireBookNFTEventsRetryExponentialBackoffCoeff,
				config.TaskAcquireBookNFTEventsRetryMaxTimeoutSeconds,
			),
		)
		if err != nil {
			return fmt.Errorf("p.ToLifecycle: %v", err)
		}

		if err := lifecycle.WithEnqueued(ctx, func(nftClass *ent.NFTClass) error {
			return handler(ctx, task)
		}); err != nil {
			mylogger.Error("lifecycle.WithEnqueued", "err", err)
		}

		return nil
	}
}

func init() {
	Tasks.Register(task.DefineTask(
		TypeAcquireBookNFTEventsTaskPayloadWithLifecyclePayload,
		handlerWithLifecycle(HandleAcquireBookNFTEventsTask),
	))
}
