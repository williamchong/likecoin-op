package evm

import (
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LikeCoin struct {
	Logger          *slog.Logger
	Client          *ethclient.Client
	ChainID         *big.Int
	ContractAddress common.Address
}

func NewLikeCoin(
	logger *slog.Logger,
	client *ethclient.Client,
	chainID *big.Int,
	contractAddress common.Address,
) LikeCoin {
	return LikeCoin{
		Logger:          logger,
		Client:          client,
		ChainID:         chainID,
		ContractAddress: contractAddress,
	}
}

func (l *LikeCoin) Auth(privateKeyStr string) *AuthedLikeCoin {
	return &AuthedLikeCoin{
		LikeCoin:      l,
		PrivateKeyStr: privateKeyStr,
	}
}

type AuthedLikeCoin struct {
	LikeCoin      *LikeCoin
	PrivateKeyStr string
}

func (l *AuthedLikeCoin) transactOpts() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(l.PrivateKeyStr)
	if err != nil {
		return nil, err
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, l.LikeCoin.ChainID)
	if err != nil {
		return nil, err
	}
	return txOpts, nil
}
