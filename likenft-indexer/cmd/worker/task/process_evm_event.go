package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/logic/evmeventprocessor"

	"github.com/hibiken/asynq"
)

const TypeProcessEVMEventPayload = "process-evm-event"

type ProcessEVMEventPayload struct {
	EvmEventId int
}

func NewProcessEVMEvent(evmEventId int) (*asynq.Task, error) {
	payload, err := json.Marshal(ProcessEVMEventPayload{
		EvmEventId: evmEventId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeProcessEVMEventPayload, payload), nil
}

func HandleProcessEVMEvent(ctx context.Context, t *asynq.Task) error {
	cfg := appcontext.ConfigFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleProcessEVMEvent")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p ProcessEVMEventPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal ProcessEVMEventPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	httpClient := &http.Client{}
	dbService := database.New()
	evmClient, err := evm.NewEvmClient(cfg.EthNetworkPublicRPCURL)

	if err != nil {
		mylogger.Error("evm.NewEvmClient", "err", err)
		return err
	}

	nftClassRepository := database.MakeNFTClassRepository(dbService)
	nftRepository := database.MakeNFTRepository(dbService)
	evmEventRepository := database.MakeEVMEventRepository(dbService)
	transactionMemoRepository := database.MakeTransactionMemoRepository(dbService)
	accountRepository := database.MakeAccountRepository(dbService)

	processor := evmeventprocessor.MakeEVMEventProcessor(
		httpClient,
		evmClient,
		nftRepository,
		nftClassRepository,
		evmEventRepository,
		transactionMemoRepository,
		accountRepository,
	)

	evmEvent, err := evmEventRepository.GetEvmEventById(ctx, p.EvmEventId)

	if err != nil {
		mylogger.Error("evmEventRepository.GetEvmEventById", "err", err)
		return err
	}

	err = processor.Process(ctx, logger, evmEvent)

	if err != nil {
		mylogger.Error("processor.Proces", "err", err)
		return err
	}

	return nil
}
