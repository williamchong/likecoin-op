package evm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) MintNFTs(
	ctx context.Context,
	logger *slog.Logger,

	msgMintNFTsFromTokenId *like_protocol.MsgMintNFTsFromTokenId,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("MintNFTs")

	mylogger := logger.WithGroup("MintNFTs")

	opts, err := l.transactOpts()
	if err != nil {
		return nil, nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err like_protocol.NewLikeProtocol: %v", err)
	}
	tx, err := instance.SafeMintNFTsWithTokenId(opts, *msgMintNFTsFromTokenId)
	if err != nil {
		mylogger.Error("instance.SafeMintNFTsWithTokenId", "err", err)
		return nil, nil, fmt.Errorf("err instance.SafeMintNFTsWithTokenId: %v", err)
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
