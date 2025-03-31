package context

import (
	"context"

	"github.com/hibiken/asynq"
)

type asynqClientContextKey struct{}

var AsynqClientContextKey = &asynqClientContextKey{}

func WithAsynqClientContext(ctx context.Context, client *asynq.Client) context.Context {
	return context.WithValue(ctx, AsynqClientContextKey, client)
}

func AsynqClientFromContext(ctx context.Context) *asynq.Client {
	return ctx.Value(AsynqClientContextKey).(*asynq.Client)
}

func AsynqMiddlewareWithAsynqClientContext(client *asynq.Client) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithAsynqClientContext(ctx, client), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
