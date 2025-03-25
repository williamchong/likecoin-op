package context

import (
	"context"
	"log/slog"
)

type loggerContextKey struct{}

var LoggerContextKey = &loggerContextKey{}

func WithLoggerContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(LoggerContextKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}
