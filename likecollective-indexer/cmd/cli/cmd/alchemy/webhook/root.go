package webhook

import "github.com/spf13/cobra"

var WebhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Webhook",
	Long:  `Webhook`,
}
