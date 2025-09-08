package evm

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type LikeCoin struct {
	Logger          *slog.Logger
	Client          *ethclient.Client
	Signer          *signer.SignerClient
	Likecoin        *likecoin.Likecoin
	ContractAddress common.Address
}

func NewLikeCoin(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
	contractAddress common.Address,
) (*LikeCoin, error) {
	likecoin, err := likecoin.NewLikecoin(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &LikeCoin{
		Logger:          logger,
		Client:          client,
		Signer:          signer,
		Likecoin:        likecoin,
		ContractAddress: contractAddress,
	}, nil
}
