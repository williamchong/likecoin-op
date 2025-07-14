package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/contractevmeventacquirer"
	"likenft-indexer/internal/worker/task"

	"github.com/hibiken/asynq"
)

const TypeAcquireLikeProtocolEventsTaskPayload = "acquire-like-protocol-events"

type AcquireLikeProtocolEventsTaskPayload struct {
	ContractAddress string
}

func NewAcquireLikeProtocolEventsTask(contractAddress string) (*asynq.Task, error) {
	payload, err := json.Marshal(AcquireLikeProtocolEventsTaskPayload{
		ContractAddress: contractAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeAcquireLikeProtocolEventsTaskPayload, payload), nil
}

func HandleAcquireLikeProtocolEventsTask(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	cfg := appcontext.ConfigFromContext(ctx)
	evmEventQueryClient := appcontext.EvmQueryClientFromContext(ctx)
	evmClient := appcontext.EvmClientFromContext(ctx)

	mylogger := logger.WithGroup("HandleAcquireLikeProtocolEventsTask")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p AcquireLikeProtocolEventsTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal AcquireNewBookNFTTaskPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()

	likeProtocolRepository := database.MakeLikeProtocolRepository(dbService)
	evmEventRepository := database.MakeEVMEventRepository(dbService)

	likeProtocol, err := likeProtocolRepository.GetLikeProtocol(ctx, p.ContractAddress)

	var fromBlock uint64
	if err != nil {
		if ent.IsNotFound(err) {
			fromBlock = cfg.EvmEventLikeProtocolInitialBlockHeight
		} else {
			return err
		}
	} else {
		fromBlock = uint64(likeProtocol.LatestEventBlockNumber) + 1
	}

	acquirer := contractevmeventacquirer.NewContractEvmEventsAcquirer(
		evmEventQueryClient,
		evmEventRepository,
		evmEventQueryClient,
		evmClient,
		contractevmeventacquirer.ContractEvmEventsAcquirerContractTypeLikeProtocol,
		[]string{p.ContractAddress},
	)

	newBlockHeight, err := acquirer.Acquire(ctx, logger, fromBlock, cfg.EvmEventQueryNumberOfBlocksLimit)
	if err != nil {
		mylogger.Error("acquirer.Acquire", "err", err)
		return err
	}

	likeProtocolRepository.CreateOrUpdateLatestEventBlockHeight(ctx, p.ContractAddress, typeutil.Uint64(newBlockHeight))

	return nil
}

func init() {
	Tasks.Register(task.DefineTask(
		TypeAcquireLikeProtocolEventsTaskPayload,
		HandleAcquireLikeProtocolEventsTask,
	))
}
