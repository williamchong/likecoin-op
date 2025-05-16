package cmd

import (
	"time"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/evmeventacquirer"

	"github.com/spf13/cobra"
)

// go run ./cmd/cli acquire-booknfts
// go run ./cmd/cli acquire-booknfts '0x14CE6632272552E676b53FE6202edA8F1Be4992c'
var AcquireBookNFTs = &cobra.Command{
	Use:   "acquire-booknfts [likeprotocol-contract-address]",
	Short: "Acquire booknfts from like protocol contract",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := context.ConfigFromContext(ctx)
		logger := context.LoggerFromContext(ctx)
		evmClient := context.EvmQueryClientFromContext(ctx)

		likeprotocolContractAddress := cfg.EthLikeProtocolContractAddress
		if len(args) > 0 {
			likeprotocolContractAddress = args[0]
		}

		logger.Info("Ready to acquire book nfts...", "likeprotocolContractAddress", likeprotocolContractAddress)
		time.Sleep(2 * time.Second)

		dbService := database.New()

		EVMEventProcessedBlockHeightRepository := database.MakeEVMEventProcessedBlockHeightRepository(dbService)
		EVMEventRepository := database.MakeEVMEventRepository(dbService)

		acquirer := evmeventacquirer.MakeEvmEventsAcquirer(
			EVMEventProcessedBlockHeightRepository,
			EVMEventRepository,
			evmClient,
		)

		err := acquirer.Acquire(ctx, logger, likeprotocolContractAddress, evmeventprocessedblockheight.EventNewBookNFT)

		if err != nil {
			panic(err)
		}
	},
}
