package cmd

import (
	"net/http"
	"strconv"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/evmeventprocessor"

	"github.com/spf13/cobra"
)

// make cli acquire-new-evm-events TransferWithMemo '0x14CE6632272552E676b53FE6202edA8F1Be4992c'
var ProcessEVMEventCmd = &cobra.Command{
	Use:   "process-evm-event id",
	Short: "Process evm event",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		dbidStr := args[0]
		dbid, err := strconv.ParseInt(dbidStr, 10, 32)

		if err != nil {
			panic(err)
		}

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

		evmEvent, err := evmEventRepository.GetEvmEventById(ctx, int(dbid))

		if err != nil {
			panic(err)
		}

		err = processor.Process(ctx, logger, evmEvent)

		if err != nil {
			panic(err)
		}
	},
}
