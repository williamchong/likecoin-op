package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/ent/evmeventprocessedblockheight"
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

	nftClasses, err := nftClassRepository.QueryAllNFTClassesOfLowestEventBlockHeight(ctx)

	if err != nil {
		mylogger.Error("nftClassRepository.QueryAllNFTClasses", "err", err)
		return err
	}

	mylogger.Info(fmt.Sprintf("%d nft classes found", len(nftClasses)))

	err = handleCheckBookNFTs_enqueueAcquireBookNFTEvents(mylogger, asynqClient, nftClasses)
	if err != nil {
		mylogger.Error("handleCheckBookNFTs_enqueueAcquireBookNFTEvents", "err", err)
	}

	for _, nftClass := range nftClasses {
		err := handleCheckBookNFTs_enqueueAcquireTransferWithMemoEVMEvent(mylogger, asynqClient, nftClass)
		if err != nil {
			mylogger.Error("handleCheckBookNFTs_enqueueAcquireTransferWithMemo", "err", err)
		}
		err = handleCheckBookNFTs_enqueueAcquireTransferEVMEvent(mylogger, asynqClient, nftClass)
		if err != nil {
			mylogger.Error("handleCheckBookNFTs_enqueueAcquireTransferEVMEvent", "err", err)
		}
		err = handleCheckBookNFTs_enqueueAcquireContractURIUpdatedEVMEvent(mylogger, asynqClient, nftClass)
		if err != nil {
			mylogger.Error("handleCheckBookNFTs_enqueueAcquireContractURIUpdatedEVMEvent", "err", err)
		}
		err = handleCheckBookNFTs_enqueueAcquireOwnershipTransferredEVMEvent(mylogger, asynqClient, nftClass)
		if err != nil {
			mylogger.Error("handleCheckBookNFTs_enqueueAcquireOwnershipTransferredEVMEvent", "err", err)
		}
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

func handleCheckBookNFTs_enqueueAcquireTransferWithMemoEVMEvent(logger *slog.Logger, asynqClient *asynq.Client, nftClass *ent.NFTClass) error {
	myLogger := logger.With("nftClass dbID", nftClass.ID)
	myLogger.Info("Enqueueing AcquireTransferWithMemo task...")
	t, err := NewAcquireEVMEventsTask(nftClass.Address, evmeventprocessedblockheight.EventTransferWithMemo)
	if err != nil {
		myLogger.Error("Cannot create task", "err", err)
		return err
	}
	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		myLogger.Error("Cannot enqueue task", "err", err)
		return err
	}
	myLogger.Info("task enqueued", "taskId", taskInfo.ID)
	return nil
}

func handleCheckBookNFTs_enqueueAcquireTransferEVMEvent(logger *slog.Logger, asynqClient *asynq.Client, nftClass *ent.NFTClass) error {
	myLogger := logger.With("nftClass dbID", nftClass.ID)
	myLogger.Info("Enqueueing AcquireTransfer task...")
	t, err := NewAcquireEVMEventsTask(nftClass.Address, evmeventprocessedblockheight.EventTransfer)
	if err != nil {
		myLogger.Error("Cannot create task", "err", err)
		return err
	}
	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		myLogger.Error("Cannot enqueue task", "err", err)
		return err
	}
	myLogger.Info("task enqueued", "taskId", taskInfo.ID)
	return nil
}

func handleCheckBookNFTs_enqueueAcquireContractURIUpdatedEVMEvent(logger *slog.Logger, asynqClient *asynq.Client, nftClass *ent.NFTClass) error {
	myLogger := logger.With("nftClass dbID", nftClass.ID)
	myLogger.Info("Enqueueing AcquireContractURIUpdatedEVMEvent task...")
	t, err := NewAcquireEVMEventsTask(nftClass.Address, evmeventprocessedblockheight.EventContractURIUpdated)
	if err != nil {
		myLogger.Error("Cannot create task", "err", err)
		return err
	}
	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		myLogger.Error("Cannot enqueue task", "err", err)
		return err
	}
	myLogger.Info("task enqueued", "taskId", taskInfo.ID)
	return nil
}

func handleCheckBookNFTs_enqueueAcquireOwnershipTransferredEVMEvent(logger *slog.Logger, asynqClient *asynq.Client, nftClass *ent.NFTClass) error {
	myLogger := logger.With("nftClass dbID", nftClass.ID)
	myLogger.Info("Enqueueing AcquireOwnershipTransferredEVMEvent task...")
	t, err := NewAcquireEVMEventsTask(nftClass.Address, evmeventprocessedblockheight.EventOwnershipTransferred)
	if err != nil {
		myLogger.Error("Cannot create task", "err", err)
		return err
	}
	taskInfo, err := asynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		myLogger.Error("Cannot enqueue task", "err", err)
		return err
	}
	myLogger.Info("task enqueued", "taskId", taskInfo.ID)
	return nil
}
