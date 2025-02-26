package context

import (
	"context"
	"log/slog"

	"github.com/hibiken/asynq"
)

type loggerContextKey struct{}

var LoggerContextKey = &loggerContextKey{}

func WithLoggerContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	return ctx.Value(LoggerContextKey).(*slog.Logger)
}

func AsynqMiddlewareWithLoggerContext(logger *slog.Logger) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithLoggerContext(ctx, logger), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
