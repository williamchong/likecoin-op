package context

import (
	"context"

	"likenft-indexer/internal/evm"

	"github.com/hibiken/asynq"
)

type evmQueryClientKey struct{}

var EvmQueryClientKey = &evmQueryClientKey{}

func WithEvmQueryClient(ctx context.Context, evmQueryClient evm.EVMQueryClient) context.Context {
	return context.WithValue(ctx, EvmQueryClientKey, evmQueryClient)
}

func EvmQueryClientFromContext(ctx context.Context) evm.EVMQueryClient {
	return ctx.Value(EvmQueryClientKey).(evm.EVMQueryClient)
}

func AsynqMiddlewareWithEvmQueryClientContext(evmQueryClient evm.EVMQueryClient) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithEvmQueryClient(ctx, evmQueryClient), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
