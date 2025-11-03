package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/likecoin/like-migration-backend/pkg/signer"
)

func (e *ethereumClient) TransferToken(
	ctx context.Context,
	to common.Address,
	amount *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	r, err := signer.MakeCreateEvmTransferTransactionRequestRequestBody(to, amount)
	if err != nil {
		return nil, nil, err
	}

	evmTxRequestResp, err := e.signer.CreateEvmTransferTransactionRequest(r)
	if err != nil {
		return nil, nil, err
	}

	txReceipt, err := AwaitTx(
		ctx,
		e.logger,
		e.client,
		e.signer,
		*evmTxRequestResp.TransactionId,
	)
	if err != nil {
		return nil, nil, err
	}

	tx, _, err := e.client.TransactionByHash(ctx, txReceipt.TxHash)
	if err != nil {
		return nil, nil, err
	}

	return tx, txReceipt, nil
}
