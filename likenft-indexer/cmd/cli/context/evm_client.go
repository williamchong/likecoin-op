package context

import (
	"context"

	"likenft-indexer/internal/evm"
)

type evmClientKey struct{}

var EvmClientKey = &evmClientKey{}

func WithEvmClient(ctx context.Context, evmClient *evm.EvmClient) context.Context {
	return context.WithValue(ctx, EvmClientKey, evmClient)
}

func EvmClientFromContext(ctx context.Context) *evm.EvmClient {
	return ctx.Value(EvmClientKey).(*evm.EvmClient)
}
