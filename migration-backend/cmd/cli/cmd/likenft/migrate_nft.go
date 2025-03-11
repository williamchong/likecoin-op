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

var migrateNFTCmd = &cobra.Command{
	Use:   "migrate-nft likenft-asset-migration-nft-id",
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
			WithGroup("migrateNFTCmd").
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

		likeProtocolClient := evm.NewLikeProtocol(
			logger,
			ethClient,
			privateKey,
			envCfg.EthChainId,
			contractAddress,
		)
		likeNFTClient := evm.NewBookNFT(
			logger,
			ethClient,
			privateKey,
			envCfg.EthChainId,
		)

		mn, err := likenft.MigrateNFTFromAssetMigration(
			logger,
			db,
			likenftClient,
			&likeProtocolClient,
			&likeNFTClient,
			envCfg.InitialNewClassOwner,
			envCfg.InitialNewClassMinter,
			envCfg.InitialNewClassUpdater,
			envCfg.InitialBatchMintNFTsOwner,
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
