package evm

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type BookNFT struct {
	Logger *slog.Logger
	Client *ethclient.Client
	Signer *signer.SignerClient
}

func NewBookNFT(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
) BookNFT {
	return BookNFT{
		Logger: logger,
		Client: client,
		Signer: signer,
	}
}
