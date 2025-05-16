package cmd

import (
	"fmt"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/evmeventacquirer"

	"github.com/spf13/cobra"
)

// go run ./cmd/cli acquire-all-booknft-evm-events
var AcquireAllBookNFTEvmEvents = &cobra.Command{
	Use:   "acquire-all-booknft-evm-events",
	Short: "Acquire evm events for all book nfts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			_ = cmd.Usage()
			return
		}

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

		nftClass, err := dbService.Client().NFTClass.Query().All(ctx)

		if err != nil {
			panic(err)
		}

		for i, n := range nftClass {
			logger.Info(fmt.Sprintf("[%d/%d] Acquiring events from booknft %s", i+1, len(nftClass), n.Address))

			err = acquirer.Acquire(ctx, logger, n.Address, evmeventprocessedblockheight.EventTransferWithMemo)
			if err != nil {
				logger.Error("error acquiring TransferWithMemo event", "err", err)
			}

			err = acquirer.Acquire(ctx, logger, n.Address, evmeventprocessedblockheight.EventTransfer)
			if err != nil {
				logger.Error("error acquiring TransferWithMemo event", "err", err)
			}

			err = acquirer.Acquire(ctx, logger, n.Address, evmeventprocessedblockheight.EventContractURIUpdated)
			if err != nil {
				logger.Error("error acquiring ContractURIUpdated event", "err", err)
			}

			err = acquirer.Acquire(ctx, logger, n.Address, evmeventprocessedblockheight.EventOwnershipTransferred)
			if err != nil {
				logger.Error("error acquiring OwnershipTransferred event", "err", err)
			}
		}
	},
}
