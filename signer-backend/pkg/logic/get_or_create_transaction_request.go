package logic

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	appdb "github.com/likecoin/like-signer-backend/pkg/db"
	"github.com/likecoin/like-signer-backend/pkg/model"
)

// idempotent
// Calling on the same contract address, method and params
// will return the same db transaction id
func GetOrCreateTransactionRequest(
	ctx context.Context,
	db *sql.DB,
	signerAddress common.Address,
	contractAddress common.Address,
	method string,
	paramsHex string,
	data []byte,
) (transactionId uint64, err error) {
	contractAddressStr := contractAddress.Hex()
	contractCall, err := appdb.QueryContractCall(db, &appdb.QueryContractCallFilter{
		ContractAddress: &contractAddressStr,
		Method:          &method,
		ParamsHex:       &paramsHex,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createTransaction(
				ctx,
				db,
				signerAddress,
				contractAddress,
				method,
				paramsHex,
				data,
			)
		} else {
			return 0, err
		}
	}

	return contractCall.EvmTransactionRequestId, nil
}

func createTransaction(
	ctx context.Context,
	db *sql.DB,
	signerAddress common.Address,
	contractAddress common.Address,
	method string,
	paramsHex string,
	data []byte,
) (transactionId uint64, err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	transaction := &model.EvmTransactionRequest{
		SignerAddress: signerAddress.Hex(),
		ToAddress:     contractAddress.Hex(),
		Amount:        big.NewInt(0),
		CallDataHex:   hex.EncodeToString(data),
		Status:        model.EvmTransactionRequestStatusInit,
	}

	transaction, err = appdb.InsertEvmTransactionRequest(tx, transaction)
	if err != nil {
		return 0, err
	}

	contractCall := &model.ContractCall{
		ContractAddress:         contractAddress.Hex(),
		Method:                  method,
		ParamsHex:               paramsHex,
		EvmTransactionRequestId: transaction.Id,
	}

	contractCall, err = appdb.InsertContractCall(tx, contractCall)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return transaction.Id, nil
}
