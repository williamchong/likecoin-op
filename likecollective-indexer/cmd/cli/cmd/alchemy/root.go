package alchemy

import (
	"likecollective-indexer/cmd/cli/cmd/alchemy/webhook"

	"github.com/spf13/cobra"
)

var AlchemyCmd = &cobra.Command{
	Use:   "alchemy",
	Short: "Alchemy",
	Long:  `Alchemy`,
}

func init() {
	AlchemyCmd.AddCommand(webhook.WebhookCmd)
}
