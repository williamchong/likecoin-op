package cmd

import (
	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/evmeventacquirer"

	"github.com/spf13/cobra"
)

// make cli acquire-new-evm-events TransferWithMemo '0x14CE6632272552E676b53FE6202edA8F1Be4992c'
var AcquireNewEVMEvents = &cobra.Command{
	Use:   "acquire-new-evm-events event class-id",
	Short: "Acquire new events",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			_ = cmd.Usage()
			return
		}

		eventRaw := args[0]
		event := evmeventprocessedblockheight.Event(eventRaw)
		classIdStr := args[1]

		ctx := cmd.Context()
		logger := context.LoggerFromContext(ctx)
		evmClient := context.EvmQueryClientFromContext(ctx)

		dbService := database.New()

		EVMEventProcessedBlockHeightRepository := database.MakeEVMEventProcessedBlockHeightRepository(dbService)
		EVMEventRepository := database.MakeEVMEventRepository(dbService)

		acquirer := evmeventacquirer.MakeEvmEventsAcquirer(
			EVMEventProcessedBlockHeightRepository,
			EVMEventRepository,
			evmClient,
		)

		err := acquirer.Acquire(ctx, logger, classIdStr, event)

		if err != nil {
			panic(err)
		}
	},
}
