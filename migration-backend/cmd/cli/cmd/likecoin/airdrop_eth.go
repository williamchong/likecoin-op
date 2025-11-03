package likecoin

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/signer"

	"github.com/spf13/cobra"
)

var airdropEthCmd = &cobra.Command{
	Use:   "airdrop-eth user-eth-address amount",
	Short: "Airdrop ETH to a user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		logger := slog.New(slog.Default().Handler()).
			WithGroup("airdrop-eth")

		userEthAddress := common.HexToAddress(args[0])
		amountStr := args[1]
		amount, ok := big.NewInt(0).SetString(amountStr, 10)
		if !ok {
			panic(fmt.Errorf("invalid amount: %s", amountStr))
		}

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

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

		ethereumClient := ethereum.NewClient(logger, ethClient, signer)

		tx, _, err := likecoin.DoAirdropEth(
			ctx, logger, ethereumClient, userEthAddress, amount, nil,
		)
		if err != nil {
			panic(err)
		}

		logger.Info("airdrop eth completed", "txHash", tx.Hash().Hex())

		fmt.Println(tx.Hash().Hex())
	},
}

func init() {
	LikeCoinCmd.AddCommand(airdropEthCmd)
}
