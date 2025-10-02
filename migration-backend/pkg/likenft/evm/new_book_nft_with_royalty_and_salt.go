package evm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

func MakeNewBookNFTWithRoyaltyAndSaltRequestBody(
	contractAddress string,
	salt [32]byte,
	msgNewBookNFT like_protocol.MsgNewBookNFT,
	royaltyFraction *big.Int,
) (*signer.CreateEvmTransactionRequestRequestBody, error) {
	return signer.MakeCreateEvmTransactionRequestRequestBody(
		like_protocol.LikeProtocolMetaData,
		"newBookNFTWithRoyaltyAndSalt",
		salt,
		msgNewBookNFT,
		royaltyFraction,
	)(contractAddress)
}

func (l *LikeProtocol) NewBookNFTWithRoyaltyAndSalt(
	ctx context.Context,
	logger *slog.Logger,

	salt [32]byte,
	msgNewBookNFT like_protocol.MsgNewBookNFT,
	royaltyFraction *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("NewBookNFTWithRoyaltyAndSalt")

	mylogger := logger.WithGroup("NewBookNFTWithRoyaltyAndSalt")

	r, err := MakeNewBookNFTWithRoyaltyAndSaltRequestBody(
		l.ContractAddress.Hex(), salt, msgNewBookNFT, royaltyFraction,
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
