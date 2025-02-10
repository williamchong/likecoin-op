package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	client *ethclient.Client
}

func NewEvmClient(url string) (*EvmClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EvmClient{
		client: client,
	}, nil
}

func (c *EvmClient) GetNonce(address common.Address) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}
