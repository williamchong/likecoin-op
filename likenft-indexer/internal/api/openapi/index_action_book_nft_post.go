package openapi

import (
	"context"
	"errors"
	"time"

	"likenft-indexer/cmd/worker/task"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/internal/worker/queue"
	"likenft-indexer/openapi/api"

	"github.com/hibiken/asynq"
)

func (h *OpenAPIHandler) IndexActionBookNftBooknftIDPost(
	ctx context.Context,
	params api.IndexActionBookNftBooknftIDPostParams,
) (*api.IndexActionBookNftBooknftIDPostOK, error) {
	nftClass, err := h.nftClassRepository.QueryNFTClassByAddress(ctx, params.BooknftID)
	if err != nil {
		return nil, err
	}

	task, err := task.NewCheckBookNFTToLatestBlockNumberTask(nftClass.Address)
	if err != nil {
		return nil, err
	}

	taskInfo, err := h.asynqClient.Enqueue(
		task,
		asynq.MaxRetry(0),
		asynq.Unique(10*time.Minute),
		asynq.Queue(string(queue.IndexActionQueue)),
	)
	if err != nil {
		if errors.Is(err, asynq.ErrDuplicateTask) {
			return &api.IndexActionBookNftBooknftIDPostOK{
				Message: "Already requested",
			}, nil
		}
		return nil, err
	}

	return &api.IndexActionBookNftBooknftIDPostOK{
		Message: "OK",
		TaskID:  model.MakeOptString(&taskInfo.ID),
	}, nil
}
