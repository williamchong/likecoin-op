package context

import (
	"context"

	"github.com/hibiken/asynq"
)

type asynqInspectorContextKey struct{}

var AsynqInspectorContextKey = &asynqInspectorContextKey{}

func WithAsynqInspectorContext(
	ctx context.Context,
	inspector *asynq.Inspector,
) context.Context {
	return context.WithValue(
		ctx,
		AsynqInspectorContextKey,
		inspector,
	)
}

func AsynqInspectorFromContext(
	ctx context.Context,
) *asynq.Inspector {
	return ctx.Value(AsynqInspectorContextKey).(*asynq.Inspector)
}

func AsynqMiddlewareWithAsynqInspectorContext(
	inspector *asynq.Inspector,
) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(
			ctx context.Context,
			t *asynq.Task,
		) error {
			err := h.ProcessTask(
				WithAsynqInspectorContext(ctx, inspector),
				t,
			)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
