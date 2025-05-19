package cmd

import (
	"fmt"
	"net/http"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/evmeventprocessor"

	"github.com/spf13/cobra"
)

// go run ./cmd/cli process-all-evm-events received
var ProcessAllEVMEventCmd = &cobra.Command{
	Use:   "process-all-evm-events { received | enqueued | processing | processed | failed }",
	Short: "Process all evm events with status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		eventRaw := args[0]
		event := evmevent.Status(eventRaw)

		ctx := cmd.Context()
		logger := context.LoggerFromContext(ctx)

		httpClient := &http.Client{}
		dbService := database.New()
		evmClient := context.EvmClientFromContext(ctx)

		nftRepository := database.MakeNFTRepository(dbService)
		nftClassRepository := database.MakeNFTClassRepository(dbService)
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

		evmEvents, err := evmEventRepository.GetEVMEventsByStatus(ctx, event)

		if err != nil {
			panic(err)
		}

		for i, e := range evmEvents {
			logger.Info(fmt.Sprintf("Processing event %d/%d: %v", i+1, len(evmEvents), e))
			err = processor.Process(ctx, logger, e)
			if err != nil {
				panic(err)
			}
		}
	},
}
