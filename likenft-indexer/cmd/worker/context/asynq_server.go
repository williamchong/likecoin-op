package context

import (
	"context"

	"github.com/hibiken/asynq"
)

type asynqServerContextKey struct{}

var AsynqServerContextKey = &asynqServerContextKey{}

func WithAsynqServerContext(ctx context.Context, server *asynq.Server) context.Context {
	return context.WithValue(ctx, AsynqServerContextKey, server)
}

func AsynqServerFromContext(ctx context.Context) *asynq.Server {
	return ctx.Value(AsynqServerContextKey).(*asynq.Server)
}
