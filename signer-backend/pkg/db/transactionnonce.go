package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/likecoin/like-signer-backend/pkg/model"
)

type QueryTransactionNonceFilter struct {
	EthAddress              *string
	EvmTransactionRequestId *uint64
}

func makeQueryTransactionNonceWhereClauseFromFilter(
	filter *QueryTransactionNonceFilter,
) (string, []interface{}) {
	wheres := make([]string, 0)
	attributes := make([]interface{}, 0)

	if filter.EthAddress != nil {
		wheres = append(wheres, fmt.Sprintf("eth_address = $%d", len(wheres)+1))
		attributes = append(attributes, filter.EthAddress)
	}

	if filter.EvmTransactionRequestId != nil {
		wheres = append(wheres, fmt.Sprintf("evm_transaction_request_id = $%d", len(wheres)+1))
		attributes = append(attributes, filter.EvmTransactionRequestId)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryTransactionNonce(
	db TxLike,
	filter *QueryTransactionNonceFilter,
) (*model.TransactionNonce, error) {
	whereClause, attributes := makeQueryTransactionNonceWhereClauseFromFilter(filter)

	row := db.QueryRow(
		fmt.Sprintf(
			`SELECT 
				id,
				created_at,
				eth_address,
				nonce,
				evm_transaction_request_id
FROM transaction_nonce 
%s`,
			whereClause,
		),
		attributes...,
	)

	var transactionNonce model.TransactionNonce

	err := row.Scan(
		&transactionNonce.Id,
		&transactionNonce.CreatedAt,
		&transactionNonce.EthAddress,
		&transactionNonce.Nonce,
		&transactionNonce.EvmTransactionRequestId,
	)
	if err != nil {
		return nil, err
	}

	return &transactionNonce, nil
}

func QueryLatestNonce(
	db TxLike,
	ethAddress string,
) (nonce uint64, exists bool, err error) {
	row := db.QueryRow(
		`SELECT nonce 
FROM transaction_nonce 
WHERE eth_address = $1 
ORDER BY created_at DESC 
LIMIT 1`,
		ethAddress,
	)

	err = row.Scan(&nonce)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, false, nil
		}
		return 0, false, err
	}

	return nonce, true, nil
}

func InsertTransactionNonce(
	db TxLike,
	transactionNonce *model.TransactionNonce,
) (*model.TransactionNonce, error) {
	row := db.QueryRow(
		`INSERT INTO transaction_nonce (
			eth_address,
			nonce,
			evm_transaction_request_id
		) VALUES ($1, $2, $3) RETURNING id`,
		transactionNonce.EthAddress,
		transactionNonce.Nonce,
		transactionNonce.EvmTransactionRequestId,
	)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	transactionNonce.Id = id

	return transactionNonce, nil
}

func DeleteTransactionNonce(
	db TxLike,
	transactionNonce *model.TransactionNonce,
) error {
	_, err := db.Exec(
		`DELETE FROM transaction_nonce WHERE id = $1`,
		transactionNonce.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
