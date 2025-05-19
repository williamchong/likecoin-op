package cmd

import (
	"fmt"
	"likenft-indexer/cmd/worker/context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account related commands",
}

var nonceCmd = &cobra.Command{
	Use:   "nonce 0xADDRESS",
	Short: "Get nonce",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide an address")
		}
		addressArg := args[0]
		address := common.HexToAddress(addressArg)

		evmClient := context.EvmClientFromContext(cmd.Context())

		nonce, err := evmClient.GetNonce(address)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(nonce)
	},
}

func init() {
	accountCmd.AddCommand(nonceCmd)

	rootCmd.AddCommand(accountCmd)
}
