package cmd

import (
	"fmt"
	"strconv"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/logic/contractevmeventacquirer"

	"github.com/spf13/cobra"
)

// go run ./cmd/cli acquire-contract-events
// go run ./cmd/cli acquire-contract-events like_protocol '0x67bcd74981c33e95e5e306085754dd0a721183f1' 25837814 499
var AcquireContractEvents = &cobra.Command{
	Use:   "acquire-contract-events [contract-type] [likeprotocol-contract-address] [from-block] [number-of-blocks-limit]",
	Short: "Acquire contract-events from contract",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			_ = cmd.Usage()
			return
		}

		contractType := contractevmeventacquirer.ContractEvmEventsAcquirerContractType(args[0])
		contractAddress := args[1]
		fromBlock, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			panic(err)
		}
		numberOfBlocksLimit, err := strconv.ParseUint(args[3], 10, 64)
		if err != nil {
			panic(err)
		}

		ctx := cmd.Context()
		logger := context.LoggerFromContext(ctx)
		evmEventQueryClient := context.EvmQueryClientFromContext(ctx)
		evmClient := context.EvmClientFromContext(ctx)

		dbService := database.New()

		evmEventRepository := database.MakeEVMEventRepository(dbService)

		acquirer := contractevmeventacquirer.NewContractEvmEventsAcquirer(
			evmEventQueryClient,
			evmEventRepository,
			evmEventQueryClient,
			evmClient,
			contractType,
			[]string{contractAddress},
		)

		newBlockHeight, err := acquirer.Acquire(ctx, logger, fromBlock, numberOfBlocksLimit)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%d", newBlockHeight)
	},
}
