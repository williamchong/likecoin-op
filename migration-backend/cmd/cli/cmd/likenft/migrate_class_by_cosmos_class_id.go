package likenft

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmosnftidclassifier"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

const (
	MigrateClassByCosmosClassIdCmdFlagNamePremintAllNFTs = "premint-all-nfts"
	MigrateClassByCosmosClassIdCmdFlagNameEvmOwner       = "evm-owner"
)

var migrateClassByCosmosClassIdCmd = &cobra.Command{
	Use:   "migrate-class-by-cosmos-class-id cosmos-class-id",
	Short: "Mint NFT Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		premintAllNFTs, err := cmd.Flags().GetBool(MigrateClassByCosmosClassIdCmdFlagNamePremintAllNFTs)
		if err != nil {
			panic(fmt.Errorf("cmd.Flags().GetBool: %v", err))
		}
		evmOwner, err := cmd.Flags().GetString(MigrateClassByCosmosClassIdCmdFlagNameEvmOwner)
		if err != nil {
			panic(fmt.Errorf("cmd.Flags().GetString: %v", err))
		}

		cosmosClassId := args[0]

		logger := slog.New(slog.Default().Handler()).
			WithGroup("migrateClassCmd").
			With("id", cosmosClassId)

		cosmosNFTIdClassifier := cosmosnftidclassifier.MakeCosmosNFTIDClassifier(
			nftidmatcher.MakeNFTIDMatcher(),
		)

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		if evmOwner == "" {
			evmOwner = envCfg.InitialNewClassOwner
		}

		erc721ExternalURLBuilder, err := erc721externalurl.MakeErc721ExternalURLBuilder3ook(
			envCfg.ERC721MetadataExternalURLBase3ook,
		)

		if err != nil {
			panic(fmt.Errorf("erc721externalurl.MakeErc721ExternalURLBuilder3ook: %v", err))
		}

		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(fmt.Errorf("sql.Open: %v", err))
		}

		likenftCosmosClient := cosmos.NewLikeNFTCosmosClient(
			envCfg.CosmosNodeUrl,
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
			envCfg.CosmosNftEventsIgnoreToList,
		)
		likecoinAPI := likecoin_api.NewLikecoinAPI(
			envCfg.LikecoinAPIUrlBase,
			time.Duration(envCfg.LikecoinAPIHTTPTimeoutSeconds),
		)

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(fmt.Errorf("ethclient.Dial: %v", err))
		}

		signer := signer.NewSignerClient(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			envCfg.EthSignerBaseUrl,
			envCfg.EthSignerAPIKey,
		)

		contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

		evmLikeNFTClient := evm.NewLikeProtocol(
			logger,
			ethClient,
			signer,
			contractAddress,
		)
		evmLikeNFTClassClient := evm.NewBookNFT(
			logger,
			ethClient,
			signer,
		)

		cosmosClass, err := likenftCosmosClient.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
			ClassId: cosmosClassId,
		})
		if err != nil {
			panic(fmt.Errorf("likenftCosmosClient.QueryClassByClassId: %v", err))
		}
		iscn, err := likenftCosmosClient.GetISCNRecord(
			cosmosClass.Class.Data.Parent.IscnIdPrefix,
			cosmosClass.Class.Data.Parent.IscnVersionAtMint,
		)
		if err != nil {
			panic(fmt.Errorf("likenftCosmosClient.GetISCNRecord: %v", err))
		}

		lastActionEvmTxHash, err := likenft.MigrateClass(
			ctx,
			logger,
			db,
			likenftCosmosClient,
			likecoinAPI,
			&evmLikeNFTClient,
			&evmLikeNFTClassClient,
			cosmosNFTIdClassifier,
			erc721ExternalURLBuilder,
			premintAllNFTs,
			cosmosClassId,
			envCfg.InitialNewClassOwner,
			envCfg.InitialNewClassMinters,
			envCfg.InitialNewClassUpdater,
			envCfg.InitialBatchMintNFTsOwner,
			envCfg.BatchMintItemPerPage,
			new(big.Int).SetUint64(envCfg.DefaultRoyaltyFraction),
			iscn.Owner,
			evmOwner,
		)
		if err != nil {
			panic(fmt.Errorf("likenft.MigrateClass: %v", err))
		}
		fmt.Printf("migrate class completed, evm tx hash: %v", *lastActionEvmTxHash)
	},
}

func init() {
	_ = migrateClassByCosmosClassIdCmd.Flags().
		Bool(
			MigrateClassByCosmosClassIdCmdFlagNamePremintAllNFTs,
			false,
			"Should Premint Al NFTs When New Class",
		)
	_ = migrateClassByCosmosClassIdCmd.Flags().
		String(
			MigrateClassByCosmosClassIdCmdFlagNameEvmOwner,
			"",
			"The evm owner of the class",
		)
	LikeNFTCmd.AddCommand(migrateClassByCosmosClassIdCmd)
}
