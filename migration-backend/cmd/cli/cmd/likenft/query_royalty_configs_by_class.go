package likenft

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

var queryRoyaltyConfigsByClassCmd = &cobra.Command{
	Use:   "query-royalty-configs-by-class class-id",
	Short: "Get Royalty Configs By Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Help()
			return
		}

		classId := args[0]

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		likenftClient := cosmos.NewLikeNFTCosmosClient(
			envCfg.CosmosNodeUrl,
			envCfg.CosmosNftEventsIgnoreToList,
		)

		royaltyConfigs, err := likenftClient.QueryRoyaltyConfigsByClassId(
			cosmos.QueryRoyaltyConfigsByClassIdRequest{
				ClassId: classId,
			})

		if err != nil {
			panic(err)
		}

		jsonByte, err := json.Marshal(royaltyConfigs)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", jsonByte)
	},
}

func init() {
	LikeNFTCmd.AddCommand(queryRoyaltyConfigsByClassCmd)
}
