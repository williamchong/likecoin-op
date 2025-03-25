package evm

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type LikeProtocol struct {
	Logger          *slog.Logger
	Client          *ethclient.Client
	Signer          *signer.SignerClient
	ContractAddress common.Address
}

func NewLikeProtocol(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
	contractAddress common.Address,
) LikeProtocol {
	return LikeProtocol{
		Logger:          logger,
		Client:          client,
		Signer:          signer,
		ContractAddress: contractAddress,
	}
}
