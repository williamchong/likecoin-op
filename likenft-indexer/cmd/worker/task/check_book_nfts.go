package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"

	"github.com/hibiken/asynq"
)

const TypeCheckBookNFTsPayload = "check-book-nfts"

type CheckBookNFTsPayload struct {
}

func NewCheckBookNFTsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CheckBookNFTsPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCheckBookNFTsPayload, payload), nil
}

func HandleCheckBookNFTs(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)

	mylogger := logger.WithGroup("HandleCheckBookNFTs")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p CheckBookNFTsPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal CheckBookNFTsPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()

	nftClassRepository := database.MakeNFTClassRepository(dbService)

	nftClasses, err := nftClassRepository.QueryAllNFTClassesOfLowestEventBlockHeight(ctx, true)

	if err != nil {
		mylogger.Error("nftClassRepository.QueryAllNFTClasses", "err", err)
		return err
	}

	mylogger.Info(fmt.Sprintf("%d nft classes found", len(nftClasses)))

	err = handleCheckBookNFTs_enqueueAcquireBookNFTEvents(mylogger, asynqClient, nftClasses)
	if err != nil {
		mylogger.Error("handleCheckBookNFTs_enqueueAcquireBookNFTEvents", "err", err)
	}

	return nil
}

func handleCheckBookNFTs_enqueueAcquireBookNFTEvents(
	logger *slog.Logger, asynqClient *asynq.Client, nftClasses []*ent.NFTClass,
) error {

	var addresses = make([]string, len(nftClasses))
	for i, nftClass := range nftClasses {
		addresses[i] = nftClass.Address
	}

	myLogger := logger.With("contractAddresses", addresses)
	myLogger.Info("Enqueueing AcquireBookNFTEvents task...")
	for _, address := range addresses {
		mylogger := myLogger.With("address", address)
		t, err := NewAcquireBookNFTEventsTask([]string{address})
		if err != nil {
			mylogger.Error("Cannot create task", "err", err)
			continue
		}
		taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
		if err != nil {
			mylogger.Error("Cannot enqueue task", "err", err)
			continue
		}
		mylogger.Info("task enqueued", "taskId", taskInfo.ID)
	}
	return nil
}
