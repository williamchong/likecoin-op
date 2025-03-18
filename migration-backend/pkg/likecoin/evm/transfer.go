package evm

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm/likecoin"
)

func (l *AuthedLikeCoin) TransferTo(
	ctx context.Context,
	logger *slog.Logger,

	to common.Address,
	value *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("MintNFTs")

	mylogger := logger.WithGroup("MintNFTs")

	opts, err := l.transactOpts()
	if err != nil {
		return nil, nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := likecoin.NewLikecoin(l.LikeCoin.ContractAddress, l.LikeCoin.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err likecoin.NewLikecoin: %v", err)
	}

	tx, err := instance.Transfer(opts, to, value)
	if err != nil {
		mylogger.Error("instance.TransferFrom", "err", err)
		return nil, nil, fmt.Errorf("err instance.TransferFrom: %v", err)
	}
	mylogger = mylogger.With("txHash", tx.Hash().Hex()).With("txNonce", tx.Nonce())

	err = l.LikeCoin.Client.SendTransaction(opts.Context, tx)
	if err != nil {
		mylogger.Error("l.Client.SendTransaction", "err", err)
		if strings.Contains(err.Error(), "nonce too low") {
			// retry
			return l.TransferTo(ctx, logger, to, value)
		}
	}

	txReceipt, err := ethereum.AwaitTx(ctx, mylogger, l.LikeCoin.Client, tx)
	if err != nil {
		mylogger.Error("ethereum.AwaitTx", "err", err)
	}

	return tx, txReceipt, err
}
