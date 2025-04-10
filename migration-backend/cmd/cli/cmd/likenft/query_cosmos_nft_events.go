package likenft

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

// go run ./cmd/cli likenft query-cosmos-nft-events likenft1x95aucwx9q57d6plk6af56fl29vfy5mjvqsje5hcm0rkfhnq3smqru00dz moneyverse-0001
var queryCosmosNFTEventsCmd = &cobra.Command{
	Use:   "query-cosmos-nft-events class-id nft-id",
	Short: "QueryCosmosNFTEvents",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			_ = cmd.Help()
			return
		}

		classId := args[0]
		nftId := args[1]

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		likenftClient := cosmos.NewLikeNFTCosmosClient(
			envCfg.CosmosNodeUrl,
			envCfg.CosmosNftEventsIgnoreToList,
		)

		nftEvents, err := likenftClient.QueryAllNFTEvents(
			likenftClient.MakeQueryNFTEventsRequest(
				classId,
				nftId,
			),
		)

		if err != nil {
			panic(err)
		}

		jsonByte, err := json.Marshal(nftEvents)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", jsonByte)
	},
}

func init() {
	LikeNFTCmd.AddCommand(queryCosmosNFTEventsCmd)
}
