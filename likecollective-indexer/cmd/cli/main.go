package main

import (
	"context"
	"log"

	"likecollective-indexer/cmd/cli/cmd"
	"likecollective-indexer/internal/cli"
	clicontext "likecollective-indexer/internal/cli/context"
)

func main() {
	envConfig, err := cli.NewEnvConfig()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	ctx := clicontext.WithCliContext(context.Background(), envConfig)

	cmd.Execute(ctx)
}
