package likenft

import (
	"database/sql"
	"time"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
)

var initMigration = &cobra.Command{
	Use:   "snapshot-cosmos-state cosmos-address",
	Short: "Initialize LikeNFT Migration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)

		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		cosmosAPI := api.NewCosmosAPI(
			envCfg.CosmosNodeUrl,
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
		)
		likenftClient := cosmos.NewLikeNFTCosmosClient(
			envCfg.CosmosNodeUrl,
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
			envCfg.CosmosNftEventsIgnoreToList,
		)

		cosmosAddress := args[0]

		snapshotCosmosState := likenft.SnapshotCosmosStateLogic{
			DB:                  db,
			CosmosAPI:           cosmosAPI,
			LikeNFTCosmosClient: likenftClient,
		}
		err = snapshotCosmosState.Execute(cosmosAddress)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	LikeNFTCmd.AddCommand(initMigration)
}
