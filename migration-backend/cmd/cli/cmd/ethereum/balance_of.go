package ethereum

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
)

var balanceOfCmd = &cobra.Command{
	Use:   "balance-of eth-address",
	Short: "Get the balance of an ETH address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		ethAddress := common.HexToAddress(args[0])

		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
		if err != nil {
			panic(errors.Join(fmt.Errorf("err dial eth client"), err))
		}

		balance, err := ethClient.BalanceAt(ctx, ethAddress, nil)
		if err != nil {
			panic(errors.Join(fmt.Errorf("err get balance"), err))
		}

		fmt.Println(balance.String())
	},
}

func init() {
	EthereumCmd.AddCommand(balanceOfCmd)
}
