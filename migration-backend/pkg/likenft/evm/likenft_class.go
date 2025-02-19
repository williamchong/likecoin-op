package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

func (l *LikeNFTClass) pubKey() (*ecdsa.PublicKey, error) {
	publicKey := l.PrivateKey.Public()
	pubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	return pubKey, nil
}

func (l *LikeNFTClass) nonce() (*uint64, error) {
	pubKey, err := l.pubKey()
	if err != nil {
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*pubKey)
	nonce, err := l.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	return &nonce, nil
}

func (l *LikeNFTClass) transactOpts() (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, l.ChainID)
	if err != nil {
		return nil, err
	}
	gasPrice, err := l.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := l.nonce()
	if err != nil {
		return nil, err
	}

	txOpts.Nonce = big.NewInt(int64(*nonce))
	txOpts.Value = big.NewInt(0)       // in wei
	txOpts.GasLimit = uint64(30000000) // in units
	txOpts.GasPrice = gasPrice

	return txOpts, nil
}
