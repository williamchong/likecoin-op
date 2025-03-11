package evm

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
)

func (l *LikeProtocol) TransferNFT(
	ctx context.Context,
	logger *slog.Logger,

	evmClassId common.Address,
	from common.Address,
	to common.Address,
	tokenId *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("TransferNFT")

	mylogger := logger.WithGroup("TransferNFT")

	opts, err := l.transactOpts()
	if err != nil {
		return nil, nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := book_nft.NewBookNft(evmClassId, l.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err book_nft.NewLikenftClass: %v", err)
	}
	tx, err := instance.TransferFrom(opts, from, to, tokenId)
	if err != nil {
		mylogger.Error("instance.TransferFrom", "err", err)
		return nil, nil, fmt.Errorf("err instance.TransferFrom: %v", err)
	}
	mylogger = mylogger.With("txHash", tx.Hash().Hex()).With("txNonce", tx.Nonce())

	err = l.Client.SendTransaction(opts.Context, tx)
	if err != nil {
		mylogger.Error("l.Client.SendTransaction", "err", err)
	}

	txReceipt, err := ethereum.AwaitTx(ctx, mylogger, l.Client, tx)
	if err != nil {
		mylogger.Error("ethereum.AwaitTx", "err", err)
	}

	return tx, txReceipt, err
}
