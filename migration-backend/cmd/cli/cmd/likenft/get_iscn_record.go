package likenft

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

var getISCNRecordCmd = &cobra.Command{
	Use:   "get-iscn-record <iscn-id-prefix> [<version>]",
	Short: "Get ISCN Record",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			_ = cmd.Help()
			return
		}

		iscnIDPrefix := args[0]

		version := ""
		if len(args) >= 2 {
			version = args[1]
		}

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		likenftClient := cosmos.NewLikeNFTCosmosClient(envCfg.CosmosNodeUrl)

		iscnRecord, err := likenftClient.GetISCNRecord(iscnIDPrefix, version)

		if err != nil {
			panic(err)
		}

		jsonByte, err := json.Marshal(iscnRecord)

		if err != nil {
			panic(err)
		}

		fmt.Printf("ISCN Record: %s\n", jsonByte)
	},
}

func init() {
	LikeNFTCmd.AddCommand(getISCNRecordCmd)
}
