package likecoinapi

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

var GetUserEVMMigrateCmd = &cobra.Command{
	Use:   "get-user-evm-migrate cosmos-address",
	Short: "Get user EVM migrate",
	Long:  `Get user EVM migrate`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		cosmosAddress := args[0]

		likecoinAPI := likecoin_api.NewLikecoinAPI(envCfg.LikecoinAPIUrlBase)
		response, err := likecoinAPI.GetUserEVMMigrate(cosmosAddress)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User EVM migrate: %+v", response)
	},
}

func init() {
	LikecoinAPICmd.AddCommand(GetUserEVMMigrateCmd)
}
