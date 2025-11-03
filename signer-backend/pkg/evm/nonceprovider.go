package evm

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NonceProvider interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

type nonceProvider struct {
	logger  *slog.Logger
	rpcUrls []string
	index   int
}

func NewNonceProvider(
	logger *slog.Logger,
	rpcUrls []string,
) NonceProvider {
	index := 0
	return &nonceProvider{
		logger,
		rpcUrls,
		index,
	}
}

func (n *nonceProvider) PendingNonceAt(
	ctx context.Context,
	account common.Address,
) (uint64, error) {
	var ethClient *ethclient.Client
	if n.index >= len(n.rpcUrls) {
		n.index = 0
	}
	rpcUrl := n.rpcUrls[n.index]
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		n.logger.Error("failed to dial nonce provider", "error", err)
		return 0, err
	}
	n.logger.Info("using nonce provider", "index", n.index, "rpcUrl", rpcUrl)
	n.index++
	return ethClient.PendingNonceAt(ctx, account)
}
