package likenft

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
)

var prepareERC721MetadataCmd = &cobra.Command{
	Use: "prepare-erc721-metadata class-id nft-id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			_ = cmd.Help()
			return
		}

		classId := args[0]
		nftId := args[1]

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		erc721ExternalURLBuilder, err := erc721externalurl.MakeErc721ExternalURLBuilder3ook("https://3ook.com/store")

		if err != nil {
			panic(err)
		}

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

		cosmosNFT, err := likenftClient.QueryNFT(cosmos.QueryNFTRequest{
			ClassId: classId,
			Id:      nftId,
		})

		if err != nil {
			panic(err)
		}

		metadataOverride, err := cosmos.ProcessQueryNFTExternalMetadataErrors(
			likenftClient.QueryNFTExternalMetadata(cosmosNFT.NFT),
		)

		if err != nil {
			panic(err)
		}

		nftIDMatcher := nftidmatcher.MakeNFTIDMatcher()
		_nftId, ok := nftIDMatcher.ExtractSerialID(nftId)

		var (
			metadataBytes []byte
		)

		if ok {
			metadataBytes, err = json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
				erc721ExternalURLBuilder,
				cosmosNFT.NFT,
				cosmosClass.Class,
				iscnDataResponse,
				metadataOverride,
				classId,
				_nftId,
			))
		} else {
			metadataBytes, err = json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNDataArbitrary(
				erc721ExternalURLBuilder,
				cosmosNFT.NFT,
				cosmosClass.Class,
				iscnDataResponse,
				metadataOverride,
				classId,
			))
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(string(metadataBytes))
	},
}

func init() {
	LikeNFTCmd.AddCommand(prepareERC721MetadataCmd)
}
