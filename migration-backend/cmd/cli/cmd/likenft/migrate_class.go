package likenft

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
)

var migrateClassCmd = &cobra.Command{
	Use:   "migrate-class likenft-asset-migration-class-id",
	Short: "Mint NFT Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		idStr := args[0]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			panic(err)
		}

		logger := slog.New(slog.Default().Handler()).
			WithGroup("migrateClassCmd").
			With("id", id)

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)
		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		likenftClient := cosmos.NewLikeNFTCosmosClient(envCfg.CosmosNodeUrl)

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(err)
		}
		privateKey, err := crypto.HexToECDSA(envCfg.EthWalletPrivateKey)
		if err != nil {
			panic(err)
		}
		contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

		evmLikeNFTClient := evm.NewLikeProtocol(
			logger,
			ethClient,
			privateKey,
			envCfg.EthChainId,
			contractAddress,
		)
		evmLikeNFTClassClient := evm.NewBookNFT(
			logger,
			ethClient,
			privateKey,
			envCfg.EthChainId,
		)

		mc, err := likenft.MigrateClassFromAssetMigration(
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
