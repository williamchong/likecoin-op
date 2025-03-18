package likecoin

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate cosmos-address",
	Short: "Migrate likecoin by cosmos-address",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		cosmosAddress := args[0]

		logger := slog.New(slog.Default().Handler()).
			WithGroup("migrate").
			With("cosmosAddress", cosmosAddress)

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)
		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(err)
		}

		cosmosAPI := api.NewCosmosAPI(envCfg.CosmosNodeUrl)

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

		contractAddress := common.HexToAddress(envCfg.EthTokenAddress)
		likeCoinClient := evm.NewLikeCoin(
			logger,
			ethClient,
			envCfg.EthChainId,
			contractAddress,
		)
		authedLikeCoinClient := likeCoinClient.Auth(envCfg.EthWalletPrivateKey)

		migration, err := likecoin.DoMintLikeCoinByCosmosAddress(
			ctx,
			logger,
			db,
			ethClient,
			cosmosAPI,
			authedLikeCoinClient,
			cosmosLikeCoinClient,
			cosmosAddress,
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
