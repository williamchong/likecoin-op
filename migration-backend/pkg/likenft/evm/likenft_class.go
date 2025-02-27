package evm

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LikeNFTClass struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	ChainID    *big.Int
}

func NewLikeNFTClass(
	client *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
) LikeNFTClass {
	return LikeNFTClass{
		Client:     client,
		PrivateKey: privateKey,
		ChainID:    chainID,
	}
}

func (l *LikeNFTClass) transactOpts() (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, l.ChainID)
	if err != nil {
		return nil, err
	}
	return txOpts, nil
}
