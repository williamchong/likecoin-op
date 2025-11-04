package likecoin

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate likecoin-migration-id",
	Short: "Migrate likecoin by likecoin-migration-id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		likecoinMigrationIdStr := args[0]
		likecoinMigrationId, err := strconv.ParseUint(likecoinMigrationIdStr, 10, 64)
		if err != nil {
			panic(err)
		}

		logger := slog.New(slog.Default().Handler()).
			WithGroup("migrate").
			With("likecoinMigrationId", likecoinMigrationId)

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)
		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(err)
		}

		cosmosAPI := api.NewCosmosAPI(
			envCfg.CosmosNodeUrl,
			time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
		)

		cosmosLikeCoinNetworkConfigData, err := os.ReadFile(
			envCfg.CosmosLikeCoinNetworkConfigPath,
		)

		if err != nil {
			panic(err)
		}

		cosmosLikeCoinNetworkConfig, err := model.LoadNetworkConfig(
			cosmosLikeCoinNetworkConfigData,
		)

		if err != nil {
			panic(err)
		}

		cosmosLikeCoinClient := cosmos.NewLikeCoin(
			logger,
			cosmosLikeCoinNetworkConfig,
		)

		signer := signer.NewSignerClient(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			envCfg.EthSignerBaseUrl,
			envCfg.EthSignerAPIKey,
		)

		contractAddress := common.HexToAddress(envCfg.EthTokenAddress)
		likeCoinClient, err := evm.NewLikeCoin(
			logger,
			ethClient,
			signer,
			contractAddress,
		)
		if err != nil {
			panic(err)
		}

		migration, err := likecoin.DoMintLikeCoinByLikeCoinMigrationId(
			ctx,
			logger,
			db,
			ethClient,
			cosmosAPI,
			likeCoinClient,
			cosmosLikeCoinClient,
			likecoinMigrationId,
		)

		if err != nil {
			panic(err)
		}

		fmt.Printf("migrate nft completed, evm tx hash: %v", *migration.EvmTxHash)
	},
}

func init() {
	LikeCoinCmd.AddCommand(migrateCmd)
}
