package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"likenft-indexer/internal/evm"
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

		rpcFlag, err := rootCmd.PersistentFlags().GetString("rpc")
		if err != nil {
			log.Fatal(err)
		}
		evmClient, err := evm.NewEvmClient(rpcFlag)
		if err != nil {
			log.Fatal(err)
		}

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
