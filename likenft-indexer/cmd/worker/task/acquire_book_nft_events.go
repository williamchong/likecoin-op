package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/contractevmeventacquirer"
	"likenft-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeAcquireBookNFTEventsTaskPayload = "acquire-book-nft-events"

type AcquireBookNFTEventsTaskPayload struct {
	ContractAddresses []string
}

func NewAcquireBookNFTEventsTask(contractAddresses []string) (*asynq.Task, error) {
	payload, err := json.Marshal(AcquireBookNFTEventsTaskPayload{
		ContractAddresses: contractAddresses,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeAcquireBookNFTEventsTaskPayload, payload), nil
}

func groupByProperty[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}

func HandleAcquireBookNFTEventsTask(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	cfg := appcontext.ConfigFromContext(ctx)
	evmEventQueryClient := appcontext.EvmQueryClientFromContext(ctx)
	evmClient := appcontext.EvmClientFromContext(ctx)

	mylogger := logger.WithGroup("HandleAcquireBookNFTEventsTask")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p AcquireBookNFTEventsTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal AcquireNewBookNFTTaskPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()

	evmEventRepository := database.MakeEVMEventRepository(dbService)
	nftClassRepository := database.MakeNFTClassRepository(dbService)

	nftClasses, err := nftClassRepository.QueryNFTClassesByAddressesExact(
		ctx, p.ContractAddresses,
	)

	if err != nil {
		mylogger.Error("nftClassRepository.QueryNFTClassesByAddressesExact", "err", err)
	}

	groupedByBlockHeight := groupByProperty(nftClasses, func(nftClass *ent.NFTClass) typeutil.Uint64 {
		return nftClass.LatestEventBlockNumber
	})

	for latestEventsBlockHeight, nftClasses := range groupedByBlockHeight {
		addresses := make([]string, len(nftClasses))
		for i, nftClass := range nftClasses {
			addresses[i] = nftClass.Address
		}

		acquirer := contractevmeventacquirer.NewContractEvmEventsAcquirer(
			evmEventQueryClient,
			evmEventRepository,
			evmEventQueryClient,
			evmClient,
			contractevmeventacquirer.ContractEvmEventsAcquirerContractTypeBookNFT,
			addresses,
		)

		newBlockHeight, err := acquirer.Acquire(ctx, logger, uint64(latestEventsBlockHeight), cfg.EvmEventQueryNumberOfBlocksLimit)

		if err != nil {
			mylogger.Error("acquirer.Acquire", "err", err)
			var errCannotConvertLog *contractevmeventacquirer.ErrCannotConvertLog
			if errors.As(err, &errCannotConvertLog) {
				return errors.Join(
					err,
					nftClassRepository.DisableForIndexing(ctx, errCannotConvertLog.Log.Address.Hex(), err.Error()),
				)
			}
			return err
		}

		err = nftClassRepository.UpdateNFTClassesLatestEventBlockNumber(
			ctx,
			addresses,
			typeutil.Uint64(newBlockHeight),
		)

		if err != nil {
			mylogger.Error("nftClassRepository.UpdateNFTClassesLatestEventsBlockNumber", "err", err)
		}
	}

	return nil
}

func init() {
	Tasks.Register(task.DefineTask(
		TypeAcquireBookNFTEventsTaskPayload,
		HandleAcquireBookNFTEventsTask,
	))
}
