package context

import (
	"context"

	"likenft-indexer/internal/evm"
)

type evmQueryClientKey struct{}

var EvmQueryClientKey = &evmQueryClientKey{}

func WithEvmQueryClient(ctx context.Context, evmQueryClient evm.EVMQueryClient) context.Context {
	return context.WithValue(ctx, EvmQueryClientKey, evmQueryClient)
}

func EvmQueryClientFromContext(ctx context.Context) evm.EVMQueryClient {
	return ctx.Value(EvmQueryClientKey).(evm.EVMQueryClient)
}
