package context

import (
	"context"

	"likenft-indexer/cmd/cli/config"
)

type configContextKey struct{}

var ConfigContextKey = &configContextKey{}

func WithConfigContext(ctx context.Context, config *config.EnvConfig) context.Context {
	return context.WithValue(ctx, ConfigContextKey, config)
}

func ConfigFromContext(ctx context.Context) *config.EnvConfig {
	return ctx.Value(ConfigContextKey).(*config.EnvConfig)
}
