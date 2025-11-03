package logic

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	appdb "github.com/likecoin/like-signer-backend/pkg/db"
	"github.com/likecoin/like-signer-backend/pkg/model"
)

func GetOrCreateTransferTransactionRequest(
	ctx context.Context,
	db *sql.DB,
	signerAddress common.Address,
	toAddress common.Address,
	amount *big.Int,
) (transactionId uint64, err error) {
	return createTransferTransaction(
		ctx,
		db,
		signerAddress,
		toAddress,
		amount,
	)
}

func createTransferTransaction(
	ctx context.Context,
	db *sql.DB,
	signerAddress common.Address,
	toAddress common.Address,
	amount *big.Int,
) (transactionId uint64, err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	abiString := `[
		{
			"inputs": [
				{ "name": "to", "type": "address" },
				{ "name": "amount", "type": "uint256" }
			],
			"name": "transfer",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`
	metadata, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return 0, err
	}

	method, ok := metadata.Methods["transfer"]
	if !ok {
		return 0, fmt.Errorf("method transfer not found")
	}

	paramsBytes, err := method.Inputs.Pack(toAddress, amount)
	if err != nil {
		return 0, err
	}

	transaction := &model.EvmTransactionRequest{
		SignerAddress: signerAddress.Hex(),
		ToAddress:     toAddress.Hex(),
		Amount:        amount,
		Method:        method.Name,
		ParamsHex:     hex.EncodeToString(paramsBytes),
		Status:        model.EvmTransactionRequestStatusInit,
	}

	transaction, err = appdb.InsertEvmTransactionRequest(tx, transaction)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return transaction.Id, nil
}
