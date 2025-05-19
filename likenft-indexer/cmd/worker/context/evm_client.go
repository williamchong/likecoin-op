package context

import (
	"context"

	"likenft-indexer/internal/evm"

	"github.com/hibiken/asynq"
)

type evmClientKey struct{}

var EvmClientKey = &evmClientKey{}

func WithEvmClient(ctx context.Context, evmClient *evm.EvmClient) context.Context {
	return context.WithValue(ctx, EvmClientKey, evmClient)
}

func EvmClientFromContext(ctx context.Context) *evm.EvmClient {
	return ctx.Value(EvmClientKey).(*evm.EvmClient)
}

func AsynqMiddlewareWithEvmClientContext(evmClient *evm.EvmClient) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithEvmClient(ctx, evmClient), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
