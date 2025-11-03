package ethereum

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type EthereumClient interface {
	GetSignerAddress(
		ctx context.Context,
	) (common.Address, error)

	TransferToken(
		ctx context.Context,
		to common.Address,
		amount *big.Int,
	) (*types.Transaction, *types.Receipt, error)

	BalanceOf(
		ctx context.Context,
		address common.Address,
	) (*big.Int, error)
}

type ethereumClient struct {
	logger *slog.Logger
	client *ethclient.Client
	signer *signer.SignerClient
}

func NewClient(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
) EthereumClient {
	return &ethereumClient{
		logger,
		client,
		signer,
	}
}

func (e *ethereumClient) GetSignerAddress(
	ctx context.Context,
) (common.Address, error) {
	signerAddress, err := e.signer.GetSignerAddress()
	if err != nil {
		return common.Address{}, err
	}
	return common.HexToAddress(*signerAddress), nil
}
