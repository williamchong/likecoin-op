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

func MakeMintNFTsRequestBody(
	contractAddress string,
	fromTokenId *big.Int,
	to common.Address,
	metadataList []string,
) (*signer.CreateEvmTransactionRequestRequestBody, error) {
	return signer.MakeCreateEvmTransactionRequestRequestBody(
		book_nft.BookNftMetaData, "safeMintWithTokenId", fromTokenId, to, metadataList,
	)(contractAddress)
}

func (l *BookNFT) MintNFTs(
	ctx context.Context,
	logger *slog.Logger,

	classId common.Address,
	fromTokenId *big.Int,
	to common.Address,
	metadataList []string,
) (*types.Transaction, *types.Receipt, error) {
	logger.Info("MintNFTs")

	mylogger := logger.WithGroup("MintNFTs")

	r, err := MakeMintNFTsRequestBody(
		classId.Hex(), fromTokenId, to, metadataList,
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
