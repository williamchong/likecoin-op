package context

import (
	"context"

	"likecollective-indexer/internal/cli"
)

type configContextKey struct{}

var ConfigContextKey = configContextKey{}

func ConfigFromContext(ctx context.Context) *cli.EnvConfig {
	return ctx.Value(ConfigContextKey).(*cli.EnvConfig)
}

func WithConfig(ctx context.Context, config *cli.EnvConfig) context.Context {
	return context.WithValue(ctx, ConfigContextKey, config)
}
