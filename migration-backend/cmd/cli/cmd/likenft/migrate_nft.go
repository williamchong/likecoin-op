package likenft

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"strconv"
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
	"github.com/likecoin/like-migration-backend/pkg/likenftindexer"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

var migrateNFTCmd = &cobra.Command{
	Use:   "migrate-nft likenft-asset-migration-nft-id",
	Short: "Mint NFT Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		idStr := args[0]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			panic(err)
		}

		logger := slog.New(slog.Default().Handler()).
			WithGroup("migrateNFTCmd").
			With("id", id)

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		cosmosNFTIdClassifier := cosmosnftidclassifier.MakeCosmosNFTIDClassifier(
			nftidmatcher.MakeNFTIDMatcher(),
		)

		erc721ExternalURLBuilder, err := erc721externalurl.MakeErc721ExternalURLBuilder3ook(
			envCfg.ERC721MetadataExternalURLBase3ook,
		)

		if err != nil {
			panic(err)
		}

		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		likenftClient := cosmos.NewLikeNFTCosmosClient(
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
			panic(err)
		}

		signer := signer.NewSignerClient(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			envCfg.EthSignerBaseUrl,
			envCfg.EthSignerAPIKey,
		)

		likenftIndexer := likenftindexer.NewLikeNFTIndexerClient(
			envCfg.LikeNFTIndexerBaseURL,
			envCfg.LikeNFTIndexerAPIKey,
		)

		contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

		likeProtocolClient := evm.NewLikeProtocol(
			logger,
			ethClient,
			signer,
			contractAddress,
		)
		likeNFTClient, err := evm.NewBookNFT(
			logger,
			ethClient,
			signer,
		)

		if err != nil {
			panic(err)
		}

		mn, err := likenft.MigrateNFTFromAssetMigration(
			ctx,
			logger,
			db,
			likenftClient,
			likecoinAPI,
			&likeProtocolClient,
			&likeNFTClient,
			likenftIndexer,
			cosmosNFTIdClassifier,
			erc721ExternalURLBuilder,
			envCfg.ShouldPremintAllNFTsWhenNewClass,
			envCfg.PremintAllNFTsWhenNewClassShouldPremintArbitraryNFTIDs,
			envCfg.InitialNewClassOwner,
			envCfg.InitialNewClassMinters,
			envCfg.InitialNewClassUpdaters,
			envCfg.InitialBatchMintNFTsOwner,
			new(big.Int).SetUint64(envCfg.DefaultRoyaltyFraction),
			envCfg.BatchMintItemPerPage,
			id,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("migrate nft completed, evm tx hash: %v", *mn.EvmTxHash)
	},
}

func init() {
	LikeNFTCmd.AddCommand(migrateNFTCmd)
}
