package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-signer-backend/pkg/model"
)

type QueryContractCallFilter struct {
	Id                      *uint64
	ContractAddress         *string
	Method                  *string
	ParamsHex               *string
	EvmTransactionRequestId *uint64
}

func makeQueryContractCallWhereClauseFromFilter(
	filter *QueryContractCallFilter,
) (string, []interface{}) {
	wheres := make([]string, 0)
	attributes := make([]interface{}, 0)

	if filter.Id != nil {
		wheres = append(wheres, fmt.Sprintf("id = $%d", len(wheres)+1))
		attributes = append(attributes, filter.Id)
	}

	if filter.ContractAddress != nil {
		wheres = append(wheres, fmt.Sprintf("contract_address = $%d", len(wheres)+1))
		attributes = append(attributes, filter.ContractAddress)
	}

	if filter.Method != nil {
		wheres = append(wheres, fmt.Sprintf("method = $%d", len(wheres)+1))
		attributes = append(attributes, filter.Method)
	}

	if filter.ParamsHex != nil {
		wheres = append(wheres, fmt.Sprintf("params_hex = $%d", len(wheres)+1))
		attributes = append(attributes, filter.ParamsHex)
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

func QueryContractCall(
	db TxLike,
	filter *QueryContractCallFilter,
) (*model.ContractCall, error) {
	whereClause, attributes := makeQueryContractCallWhereClauseFromFilter(filter)

	row := db.QueryRow(
		fmt.Sprintf(
			`SELECT
				id,
				created_at,
				contract_address,
				method,
				params_hex,
				evm_transaction_request_id
FROM contract_call
%s`,
			whereClause,
		),
		attributes...,
	)

	var contractCall model.ContractCall

	err := row.Scan(
		&contractCall.Id,
		&contractCall.CreatedAt,
		&contractCall.ContractAddress,
		&contractCall.Method,
		&contractCall.ParamsHex,
		&contractCall.EvmTransactionRequestId,
	)

	if err != nil {
		return nil, err
	}

	return &contractCall, nil
}

func InsertContractCall(
	db TxLike,
	contractCall *model.ContractCall,
) (*model.ContractCall, error) {
	row := db.QueryRow(
		`INSERT INTO contract_call (
			contract_address,
			method,
			params_hex,
			evm_transaction_request_id
		) VALUES ($1, $2, $3, $4) RETURNING id`,
		contractCall.ContractAddress,
		contractCall.Method,
		contractCall.ParamsHex,
		contractCall.EvmTransactionRequestId,
	)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	contractCall.Id = id

	return contractCall, nil
}

func DeleteContractCall(
	db TxLike,
	contractCall *model.ContractCall,
) error {
	_, err := db.Exec(
		`DELETE FROM contract_call WHERE id = $1`,
		contractCall.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
