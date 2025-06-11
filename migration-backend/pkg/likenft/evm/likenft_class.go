package evm

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type BookNFT struct {
	Logger *slog.Logger
	Client *ethclient.Client
	Signer *signer.SignerClient

	abi *abi.ABI
}

func NewBookNFT(
	logger *slog.Logger,
	client *ethclient.Client,
	signer *signer.SignerClient,
) (BookNFT, error) {

	abi, err := book_nft.BookNftMetaData.GetAbi()
	if err != nil {
		return BookNFT{}, err
	}
	return BookNFT{
		Logger: logger,
		Client: client,
		Signer: signer,
		abi:    abi,
	}, nil
}
