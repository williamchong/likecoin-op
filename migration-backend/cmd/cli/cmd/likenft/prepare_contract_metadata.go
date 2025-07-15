package likenft

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
)

var prepareContractMetadataCmd = &cobra.Command{
	Use: "prepare-contract-metadata class-id",
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
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
			envCfg.CosmosNftEventsIgnoreToList,
		)

		cosmosClass, err := likenftClient.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
			ClassId: classId,
		})

		if err != nil {
			panic(err)
		}

		iscnDataResponse, err := likenftClient.GetISCNRecord(
			cosmosClass.Class.Data.Parent.IscnIdPrefix,
			cosmosClass.Class.Data.Parent.IscnVersionAtMint,
		)

		if err != nil {
			panic(err)
		}

		royaltyConfigsResponse, err := likenftClient.QueryRoyaltyConfigsByClassId(cosmos.QueryRoyaltyConfigsByClassIdRequest{
			ClassId: classId,
		})
		if err != nil {
			panic(err)
		}

		var royaltyConfig *cosmosmodel.RoyaltyConfig = nil
		if royaltyConfigsResponse != nil {
			royaltyConfig = royaltyConfigsResponse.RoyaltyConfig
		}

		metadataBytes, err := json.Marshal(evm.ContractLevelMetadataFromCosmosClassAndISCN(
			cosmosClass.Class,
			iscnDataResponse,
			royaltyConfig,
		))

		if err != nil {
			panic(err)
		}

		fmt.Println(string(metadataBytes))
	},
}

func init() {
	LikeNFTCmd.AddCommand(prepareContractMetadataCmd)
}
