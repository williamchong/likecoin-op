package evm

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
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

	r, err := signer.MakeCreateEvmTransactionRequestRequestBody(
		book_nft.BookNftMetaData, "transferFrom", from, to, tokenId,
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
