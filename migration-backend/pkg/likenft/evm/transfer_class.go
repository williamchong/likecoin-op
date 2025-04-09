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

func MakeTransferClassRequestBody(
	contractAddress string,
	newOwner common.Address,
) (*signer.CreateEvmTransactionRequestRequestBody, error) {
	return signer.MakeCreateEvmTransactionRequestRequestBody(
		like_protocol.LikeProtocolMetaData, "transferOwnership", newOwner,
	)(contractAddress)
}

func (l *BookNFT) TransferClass(
	ctx context.Context,
	logger *slog.Logger,

	evmClassId common.Address,
	newOwner common.Address,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("TransferClass")

	mylogger := logger.WithGroup("TransferClass")

	r, err := MakeTransferClassRequestBody(
		evmClassId.Hex(), newOwner,
	)
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
