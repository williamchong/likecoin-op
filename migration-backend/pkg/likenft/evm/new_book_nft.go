package evm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) NewBookNFT(
	ctx context.Context,
	logger *slog.Logger,

	msgNewBookNFT like_protocol.MsgNewBookNFT,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("NewBookNFT")

	mylogger := logger.WithGroup("NewBookNFT")

	opts, err := l.transactOpts()
	if err != nil {
		return nil, nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err like_protocol.NewLikeProtocol: %v", err)
	}
	tx, err := instance.NewBookNFT(opts, msgNewBookNFT)
	if err != nil {
		mylogger.Error("instance.NewBookNFT", "err", err)
		return nil, nil, fmt.Errorf("err instance.NewBookNFT: %v", err)
	}
	mylogger = mylogger.With("txHash", tx.Hash().Hex()).With("txNonce", tx.Nonce())

	err = l.Client.SendTransaction(opts.Context, tx)
	if err != nil {
		mylogger.Error("l.Client.SendTransaction", "err", err)
		if strings.Contains(err.Error(), "nonce too low") {
			// retry
			return l.NewBookNFT(ctx, logger, msgNewBookNFT)
		}
	}

	txReceipt, err := ethereum.AwaitTx(ctx, mylogger, l.Client, tx)
	if err != nil {
		mylogger.Error("ethereum.AwaitTx", "err", err)
	}

	return tx, txReceipt, err
}

func (l *LikeProtocol) GetClassIdFromNewClassTransaction(txReceipt *types.Receipt) (*common.Address, error) {
	filterer, err := like_protocol.NewLikeProtocolFilterer(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikeProtocolFilterer: %v", err)
	}

	logs := txReceipt.Logs

	for _, vLog := range logs {
		newClassEvent, err := filterer.ParseNewBookNFT(*vLog)
		if err == nil {
			return &newClassEvent.BookNFT, nil
		}
	}
	return nil, errors.New("err finding new book nft event from tx log")
}
