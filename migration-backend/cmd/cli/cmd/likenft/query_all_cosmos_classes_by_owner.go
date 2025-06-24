package likenft

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

// go run ./cmd/cli likenft query-all-cosmos-classes-by-owner like1gss4uegvsncanff7c2me9hf4v762rmfyz7x2rt
var queryAllCosmosClassesByOwnerCmd = &cobra.Command{
	Use:   "query-all-cosmos-classes-by-owner owner",
	Short: "QueryAllCosmosClassesByOwner",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Help()
			return
		}

		iscnOwner := args[0]

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		likeNFTCosmosClient := cosmos.NewLikeNFTCosmosClient(
			envCfg.CosmosNodeUrl,
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
			envCfg.CosmosNftEventsIgnoreToList,
		)

		response, err := likeNFTCosmosClient.QueryAllNFTClassesByOwner(
			ctx,
			cosmos.QueryAllNFTClassesByOwnerRequest{
				ISCNOwner: iscnOwner,
			},
		)

		if err != nil {
			panic(err)
		}

		jsonByte, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonByte))
	},
}

func init() {
	LikeNFTCmd.AddCommand(queryAllCosmosClassesByOwnerCmd)
}
