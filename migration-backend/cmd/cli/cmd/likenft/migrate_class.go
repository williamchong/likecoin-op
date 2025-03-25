package likenft

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

var migrateClassCmd = &cobra.Command{
	Use:   "migrate-class likenft-asset-migration-class-id",
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
			WithGroup("migrateClassCmd").
			With("id", id)

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)
		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		likenftClient := cosmos.NewLikeNFTCosmosClient(envCfg.CosmosNodeUrl)

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

		mc, err := likenft.MigrateClassFromAssetMigration(
			ctx,
			logger,
			db,
			likenftClient,
			&evmLikeNFTClient,
			&evmLikeNFTClassClient,
			envCfg.InitialNewClassOwner,
			envCfg.InitialNewClassMinter,
			envCfg.InitialNewClassUpdater,
			id,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("migrate class completed, evm tx hash: %v", *mc.EvmTxHash)
	},
}

func init() {
	LikeNFTCmd.AddCommand(migrateClassCmd)
}
