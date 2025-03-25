package evm

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

func (l *LikeProtocol) MintNFTs(
	ctx context.Context,
	logger *slog.Logger,

	msgMintNFTsFromTokenId *like_protocol.MsgMintNFTsFromTokenId,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("MintNFTs")

	mylogger := logger.WithGroup("MintNFTs")

	r, err := signer.MakeCreateEvmTransactionRequestRequestBody(
		like_protocol.LikeProtocolMetaData, "safeMintNFTsWithTokenId", *msgMintNFTsFromTokenId,
	)(l.ContractAddress.Hex())
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
