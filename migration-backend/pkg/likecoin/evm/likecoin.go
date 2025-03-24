package evm

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type LikeCoin struct {
	Logger          *slog.Logger
	Client          *ethclient.Client
	Signer          *signer.SignerClient
	ContractAddress common.Address
}

func NewLikeCoin(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
	contractAddress common.Address,
) *LikeCoin {
	return &LikeCoin{
		Logger:          logger,
		Client:          client,
		Signer:          signer,
		ContractAddress: contractAddress,
	}
}
