package evm

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

func MakeNewBookNFTWithRoyaltyRequestBody(
	contractAddress string,
	msgNewBookNFT like_protocol.MsgNewBookNFT,
	royaltyFraction *big.Int,
) (*signer.CreateEvmTransactionRequestRequestBody, error) {
	return signer.MakeCreateEvmTransactionRequestRequestBody(
		like_protocol.LikeProtocolMetaData, "newBookNFTWithRoyalty", msgNewBookNFT, royaltyFraction,
	)(contractAddress)
}

func (l *LikeProtocol) NewBookNFTWithRoyalty(
	ctx context.Context,
	logger *slog.Logger,

	msgNewBookNFT like_protocol.MsgNewBookNFT,
	royaltyFraction *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("NewBookNFTWithRoyalty")

	mylogger := logger.WithGroup("NewBookNFTWithRoyalty")

	r, err := MakeNewBookNFTWithRoyaltyRequestBody(
		l.ContractAddress.Hex(), msgNewBookNFT, royaltyFraction,
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
