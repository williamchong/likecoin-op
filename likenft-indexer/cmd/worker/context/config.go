package context

import (
	"context"

	"likenft-indexer/cmd/worker/config"

	"github.com/hibiken/asynq"
)

type configContextKey struct{}

var ConfigContextKey = &configContextKey{}

func WithConfigContext(ctx context.Context, config *config.EnvConfig) context.Context {
	return context.WithValue(ctx, ConfigContextKey, config)
}

func ConfigFromContext(ctx context.Context) *config.EnvConfig {
	return ctx.Value(ConfigContextKey).(*config.EnvConfig)
}

func AsynqMiddlewareWithConfigContext(config *config.EnvConfig) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithConfigContext(ctx, config), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
