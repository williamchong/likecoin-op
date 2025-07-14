package openapi

import (
	"context"
	"errors"

	"likenft-indexer/cmd/worker/task"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"

	"github.com/hibiken/asynq"
)

func (h *OpenAPIHandler) IndexActionLikeProtocolPost(
	ctx context.Context,
	params api.IndexActionLikeProtocolPostParams,
) (*api.IndexActionLikeProtocolPostOK, error) {
	task, err := task.NewCheckLikeProtocolToLatestBlockNumberTask(h.likeProtocolAddress)
	if err != nil {
		return nil, err
	}
	taskInfo, err := h.asynqClient.Enqueue(
		task,
		asynq.MaxRetry(0),
	)
	if err != nil {
		if errors.Is(err, asynq.ErrDuplicateTask) {
			return &api.IndexActionLikeProtocolPostOK{
				Message: "Already requested",
			}, nil
		}
		return nil, err
	}
	return &api.IndexActionLikeProtocolPostOK{
		Message: "OK",
		TaskID:  model.MakeOptString(&taskInfo.ID),
	}, nil

}
