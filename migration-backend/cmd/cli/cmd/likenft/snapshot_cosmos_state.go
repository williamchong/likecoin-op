package likenft

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/spf13/cobra"
)

var initMigration = &cobra.Command{
	Use:   "snapshot-cosmos-state cosmos-address",
	Short: "Initialize LikeNFT Migration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			return
		}

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)

		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		cosmosAPI := &api.CosmosAPI{
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
			NodeURL: envCfg.CosmosNodeUrl,
		}

		likenftClient := &cosmos.LikeNFTCosmosClient{
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
			NodeURL: envCfg.CosmosNodeUrl,
		}

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
