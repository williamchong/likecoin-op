package webhook

import (
	"encoding/json"
	"fmt"
	"log"

	clicontext "likecollective-indexer/internal/cli/context"
	"likecollective-indexer/pkg/alchemy"

	"github.com/spf13/cobra"
)

var CreateWebhookCmd = &cobra.Command{
	Use:     "create <network> <webhook-type> <webhook-url> <graphql-query>",
	Short:   "Create a webhook",
	Long:    `Create a webhook`,
	Example: `cli alchemy webhook create BASE_SEPOLIA GRAPHQL <webhook-url> "$(cat cmd/cli/cmd/alchemy/webhook/create.gql)"`,
	Run: func(cmd *cobra.Command, args []string) {
		network := args[0]
		webhookType := args[1]
		webhookUrl := args[2]
		graphqlQuery := args[3]
		config := clicontext.ConfigFromContext(cmd.Context())
		alchemyClient := alchemy.NewAlchemyClient(config.AlchemyConfig)

		response, err := alchemyClient.CreateWebhook(&alchemy.CreateWebhookRequest{
			Network:      network,
			WebhookType:  webhookType,
			WebhookUrl:   webhookUrl,
			GraphqlQuery: graphqlQuery,
		})
		if err != nil {
			log.Fatalf("Failed to create webhook: %v", err)
		}

		jsonBytes, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Failed to marshal response: %v", err)
		}
		fmt.Println(string(jsonBytes))
	},
}

func init() {
	WebhookCmd.AddCommand(CreateWebhookCmd)
}
