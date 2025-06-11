package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/contractevmeventacquirer"

	"github.com/hibiken/asynq"
)

const TypeCheckBookNFTToLatestBlockNumberPayload = "check-book-nft-to-latest-block-number"

type CheckBookNFTToLatestBlockNumberPayload struct {
	ContractAddress string
}

func NewCheckBookNFTToLatestBlockNumberTask(contractAddress string) (*asynq.Task, error) {
	payload, err := json.Marshal(CheckBookNFTToLatestBlockNumberPayload{
		ContractAddress: contractAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCheckBookNFTToLatestBlockNumberPayload, payload), nil
}

func HandleCheckBookNFTToLatestBlockNumber(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	envCfg := appcontext.ConfigFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)
	evmQueryClient := appcontext.EvmQueryClientFromContext(ctx)
	evmClient := appcontext.EvmClientFromContext(ctx)

	dbService := database.New()
	nftClassRepository := database.MakeNFTClassRepository(dbService)
	evmEventRepository := database.MakeEVMEventRepository(dbService)

	mylogger := logger.WithGroup("HandleCheckBookNFTToLatestBlockNumber")

	var p CheckBookNFTToLatestBlockNumberPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal CheckBookNFTToLatestBlockNumberPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	nftClass, err := nftClassRepository.QueryNFTClassByAddress(ctx, p.ContractAddress)
	if err != nil {
		return err
	}

	latestBlockNumber, err := evmClient.BlockNumber(ctx)
	if err != nil {
		return err
	}

	mylogger = mylogger.With("latestBlockNumber", latestBlockNumber)

	blockStarts := make([]uint64, 0)

	for i := uint64(nftClass.LatestEventBlockNumber); i < latestBlockNumber; i = i + envCfg.EvmEventQueryNumberOfBlocksLimit {
		blockStarts = append(blockStarts, i)
	}

	acquirer := contractevmeventacquirer.NewContractEvmEventsAcquirer(
		evmQueryClient,
		evmEventRepository,
		evmQueryClient,
		evmClient,
		contractevmeventacquirer.ContractEvmEventsAcquirerContractTypeBookNFT,
		[]string{p.ContractAddress},
	)

	for i, blockStart := range blockStarts {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		mylogger := mylogger.With(
			"partition",
			fmt.Sprintf("%d/%d", i, len(blockStarts)),
		)
		newBlockNumber, err := acquirer.Acquire(
			ctx,
			mylogger,
			blockStart,
			envCfg.EvmEventQueryNumberOfBlocksLimit,
		)
		if err != nil {
			return err
		}
		err = nftClassRepository.UpdateNFTClassesLatestEventBlockNumber(
			ctx,
			[]string{p.ContractAddress},
			typeutil.Uint64(newBlockNumber),
		)
		if err != nil {
			return err
		}
		task, err := NewCheckReceivedEVMEventsTask()
		if err != nil {
			return err
		}

		_, err = asynqClient.Enqueue(task, asynq.MaxRetry(0))

		if err != nil {
			return err
		}
	}
	return nil
}
