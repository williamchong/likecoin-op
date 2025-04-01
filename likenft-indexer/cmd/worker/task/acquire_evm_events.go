package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/logic/evmeventacquirer"

	"github.com/hibiken/asynq"
)

const TypeAcquireEVMEventsTaskPayload = "acquire-evm-events"

type AcquireEVMEventsTaskPayload struct {
	ContractAddress string
	Event           evmeventprocessedblockheight.Event
}

func NewAcquireEVMEventsTask(contractAddress string, event evmeventprocessedblockheight.Event) (*asynq.Task, error) {
	payload, err := json.Marshal(AcquireEVMEventsTaskPayload{
		ContractAddress: contractAddress,
		Event:           event,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeAcquireEVMEventsTaskPayload, payload), nil
}

func HandleAcquireEVMEventsTask(ctx context.Context, t *asynq.Task) error {
	cfg := appcontext.ConfigFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleAcquireNewBookNFTTask")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p AcquireEVMEventsTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal AcquireNewBookNFTTaskPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbService := database.New()
	evmClient, err := evm.NewEvmClient(cfg.EthNetworkPublicRPCURL)

	if err != nil {
		mylogger.Error("evm.NewEvmClient", "err", err)
		return err
	}

	evmEventProcessedBlockHeightRepository := database.MakeEVMEventProcessedBlockHeightRepository(dbService)
	evmEventRepository := database.MakeEVMEventRepository(dbService)

	acquirer := evmeventacquirer.MakeEvmEventsAcquirer(
		evmEventProcessedBlockHeightRepository,
		evmEventRepository,
		evmClient,
	)

	err = acquirer.Acquire(
		ctx,
		mylogger,
		p.ContractAddress,
		p.Event,
	)

	if err != nil {
		mylogger.Error("acquirer.Acquire", "err", err)
		return err
	}

	return nil
}
