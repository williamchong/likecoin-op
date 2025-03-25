package db

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/likecoin/like-signer-backend/pkg/model"
)

type QueryEvmTransactionRequestFilter struct {
	Id *uint64
}

func makeQueryEvmTransactionRequestWhereClauseFromFilter(
	filter *QueryEvmTransactionRequestFilter,
) (string, []interface{}) {
	wheres := make([]string, 0)
	attributes := make([]interface{}, 0)

	if filter.Id != nil {
		wheres = append(wheres, fmt.Sprintf("id = $%d", len(wheres)+1))
		attributes = append(attributes, filter.Id)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryEvmTransactionRequest(
	db TxLike,
	filter *QueryEvmTransactionRequestFilter,
) (*model.EvmTransactionRequest, error) {

	whereClause, attributes := makeQueryEvmTransactionRequestWhereClauseFromFilter(filter)

	row := db.QueryRow(
		fmt.Sprintf(
			`SELECT 
				id, 
				created_at, 
				status, 
				signer_address, 
				to_address, 
				amount, 
				method,
				params_hex,
				call_data_hex, 
				gas_limit, 
				gas_price, 
				nonce, 
				signed_tx_hash, 
				submitted_at, 
				block_hash, 
				receipt_status, 
				failed_reason
		FROM evm_transaction_request
		%s
		`,
			whereClause,
		),
		attributes...,
	)

	var evmTransactionRequest model.EvmTransactionRequest

	var amount string
	var gasPrice *string

	err := row.Scan(
		&evmTransactionRequest.Id,
		&evmTransactionRequest.CreatedAt,
		&evmTransactionRequest.Status,
		&evmTransactionRequest.SignerAddress,
		&evmTransactionRequest.ToAddress,
		&amount,
		&evmTransactionRequest.Method,
		&evmTransactionRequest.ParamsHex,
		&evmTransactionRequest.CallDataHex,
		&evmTransactionRequest.GasLimit,
		&gasPrice,
		&evmTransactionRequest.Nonce,
		&evmTransactionRequest.SignedTxHash,
		&evmTransactionRequest.SubmittedAt,
		&evmTransactionRequest.BlockHash,
		&evmTransactionRequest.ReceiptStatus,
		&evmTransactionRequest.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	amountBigInt, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amount: %s", amount)
	}
	evmTransactionRequest.Amount = amountBigInt

	if gasPrice != nil {
		gasPriceBigInt, ok := big.NewInt(0).SetString(*gasPrice, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse gas price: %s", *gasPrice)
		}
		evmTransactionRequest.GasPrice = gasPriceBigInt
	}

	return &evmTransactionRequest, nil
}

func InsertEvmTransactionRequest(
	db TxLike,
	evmTransactionRequest *model.EvmTransactionRequest,
) (*model.EvmTransactionRequest, error) {

	columns := []string{
		"status",
		"signer_address",
		"to_address",
		"amount",
		"method",
		"params_hex",
		"call_data_hex",
		"gas_limit",
		"gas_price",
		"nonce",
		"signed_tx_hash",
		"submitted_at",
		"block_hash",
		"receipt_status",
		"failed_reason",
	}

	parametrisedColumns := make([]string, 0)
	for i := range columns {
		parametrisedColumns = append(parametrisedColumns, fmt.Sprintf("$%d", i+1))
	}

	amountString := evmTransactionRequest.Amount.String()
	var gasPriceString *string
	if evmTransactionRequest.GasPrice != nil {
		_gasPriceString := evmTransactionRequest.GasPrice.String()
		gasPriceString = &_gasPriceString
	}

	row := db.QueryRow(
		fmt.Sprintf(`INSERT INTO evm_transaction_request (
			%s
		) VALUES (%s) RETURNING id`,
			strings.Join(columns, ", "),
			strings.Join(parametrisedColumns, ", "),
		),
		evmTransactionRequest.Status,
		evmTransactionRequest.SignerAddress,
		evmTransactionRequest.ToAddress,
		amountString,
		evmTransactionRequest.Method,
		evmTransactionRequest.ParamsHex,
		evmTransactionRequest.CallDataHex,
		evmTransactionRequest.GasLimit,
		gasPriceString,
		evmTransactionRequest.Nonce,
		evmTransactionRequest.SignedTxHash,
		evmTransactionRequest.SubmittedAt,
		evmTransactionRequest.BlockHash,
		evmTransactionRequest.ReceiptStatus,
		evmTransactionRequest.FailedReason,
	)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	evmTransactionRequest.Id = id

	return evmTransactionRequest, nil
}

func UpdateEvmTransactionRequest(
	db TxLike,
	evmTransactionRequest *model.EvmTransactionRequest,
) (*model.EvmTransactionRequest, error) {

	amountString := evmTransactionRequest.Amount.String()

	var gasPriceString *string
	if evmTransactionRequest.GasPrice != nil {
		_gasPriceString := evmTransactionRequest.GasPrice.String()
		gasPriceString = &_gasPriceString
	}

	result, err := db.Exec(
		`UPDATE evm_transaction_request SET
			status = $1,
			signer_address = $2,
			to_address = $3,
			amount = $4,
			method = $5,
			params_hex = $6,
			call_data_hex = $7,
			gas_limit = $8,
			gas_price = $9,
			nonce = $10,
			signed_tx_hash = $11,
			submitted_at = $12,
			block_hash = $13,
			receipt_status = $14,
			failed_reason = $15
		WHERE id = $16`,
		evmTransactionRequest.Status,
		evmTransactionRequest.SignerAddress,
		evmTransactionRequest.ToAddress,
		amountString,
		evmTransactionRequest.Method,
		evmTransactionRequest.ParamsHex,
		evmTransactionRequest.CallDataHex,
		evmTransactionRequest.GasLimit,
		gasPriceString,
		evmTransactionRequest.Nonce,
		evmTransactionRequest.SignedTxHash,
		evmTransactionRequest.SubmittedAt,
		evmTransactionRequest.BlockHash,
		evmTransactionRequest.ReceiptStatus,
		evmTransactionRequest.FailedReason,
		evmTransactionRequest.Id,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("evm transaction request not found")
	}

	return evmTransactionRequest, nil
}
