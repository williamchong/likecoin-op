package evm

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

func (l *BookNFT) TransferClass(
	ctx context.Context,
	logger *slog.Logger,

	evmClassId common.Address,
	newOwner common.Address,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("TransferClass")

	mylogger := logger.WithGroup("TransferClass")

	r, err := signer.MakeCreateEvmTransactionRequestRequestBody(
		like_protocol.LikeProtocolMetaData, "transferOwnership", newOwner,
	)(evmClassId.Hex())
	if err != nil {
		return nil, nil, err
	}
	evmTxRequestResp, err := l.Signer.CreateEvmTransactionRequest(r)
	if err != nil {
		return nil, nil, err
	}

	txReceipt, err := ethereum.AwaitTx(ctx, mylogger, l.Client, l.Signer, *evmTxRequestResp.TransactionId)

	if err != nil {
		return nil, nil, err
	}

	tx, _, err := l.Client.TransactionByHash(ctx, txReceipt.TxHash)
	if err != nil {
		return nil, nil, err
	}

	return tx, txReceipt, err
}
