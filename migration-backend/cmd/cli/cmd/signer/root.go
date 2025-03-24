package signer

import (
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/signer/encode"
)

var SignerCmd = &cobra.Command{
	Use:   "signer",
	Short: "CLI for Signer",
}

func init() {
	SignerCmd.AddCommand(encode.EncodeCmd)
}
