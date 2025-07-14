package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeCheckLikeProtocolPayload = "check-like-protocol"

type CheckLikeProtocolPayload struct {
}

func NewCheckLikeProtocolTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CheckLikeProtocolPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeCheckLikeProtocolPayload,
		payload,
		asynq.Queue(TypeCheckLikeProtocolPayload),
	), nil
}

func HandleCheckLikeProtocol(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)
	cfg := appcontext.ConfigFromContext(ctx)

	mylogger := logger.WithGroup("HandleCheckLikeProtocol")

	var p CheckLikeProtocolPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal CheckLikeProtocolPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	err := handleCheckLikeProtocol_enqueueAcquireLikeProtocolEventsTask(
		mylogger,
		asynqClient,
		cfg.EthLikeProtocolContractAddress,
	)

	if err != nil {
		mylogger.Error("enqueueAcquireLikeProtocolEventsTask", "err", err)
		return err
	}

	return nil
}

func handleCheckLikeProtocol_enqueueAcquireLikeProtocolEventsTask(
	logger *slog.Logger,
	asynqClient *asynq.Client,

	contractAddress string,
) error {
	mylogger := logger.WithGroup("enqueueAcquireLikeProtocolEventsTask")

	mylogger.Info("Enqueueing AcquireLikeProtocolEventsTask task...")
	t, err := NewAcquireLikeProtocolEventsTask(contractAddress)
	if err != nil {
		mylogger.Error("Cannot create task", "err", err)
	}
	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		mylogger.Error("Cannot enqueue task", "err", err)
	}
	mylogger.Info("task enqueued", "taskId", taskInfo.ID)

	return nil
}

func init() {
	t := Tasks.Register(task.DefineTask(
		TypeCheckLikeProtocolPayload,
		HandleCheckLikeProtocol,
	))
	PeriodicTasks.Register(task.DefinePeriodicTask(
		t,
		NewCheckLikeProtocolTask,
	))
}
