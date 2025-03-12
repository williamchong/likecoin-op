package evm

import (
	"crypto/ecdsa"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LikeProtocol struct {
	Logger          *slog.Logger
	Client          *ethclient.Client
	PrivateKey      *ecdsa.PrivateKey
	ChainID         *big.Int
	ContractAddress common.Address
}

func NewLikeProtocol(
	logger *slog.Logger,
	client *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
	contractAddress common.Address,
) LikeProtocol {
	return LikeProtocol{
		Logger:          logger,
		Client:          client,
		PrivateKey:      privateKey,
		ChainID:         chainID,
		ContractAddress: contractAddress,
	}
}

func (l *LikeProtocol) transactOpts() (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, l.ChainID)
	if err != nil {
		return nil, err
	}
	return txOpts, nil
}
