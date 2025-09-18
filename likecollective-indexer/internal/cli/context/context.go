package context

import (
	"context"

	"likecollective-indexer/internal/cli"
)

func WithCliContext(ctx context.Context, config *cli.EnvConfig) context.Context {
	return WithConfig(ctx, config)
}
