package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type QueryLikeNFTMigrationActionTransferClassFilter struct {
	EvmClassId  *string
	CosmosOwner *string
	EvmOwner    *string
}

func makeQueryLikeNFTMigrationActionTransferClassWhereClauseFromFilter(f QueryLikeNFTMigrationActionTransferClassFilter) (string, []any) {
	wheres := make([]string, 0)
	attributes := make([]any, 0)

	if f.EvmClassId != nil {
		wheres = append(wheres, fmt.Sprintf("evm_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.EvmClassId)
	}

	if f.CosmosOwner != nil {
		wheres = append(wheres, fmt.Sprintf("cosmos_owner = $%d", len(wheres)+1))
		attributes = append(attributes, f.CosmosOwner)
	}

	if f.EvmOwner != nil {
		wheres = append(wheres, fmt.Sprintf("evm_owner = $%d", len(wheres)+1))
		attributes = append(attributes, f.EvmOwner)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryLikeNFTMigrationActionTransferClass(
	tx TxLike,
	filter QueryLikeNFTMigrationActionTransferClassFilter,
) (*model.LikeNFTMigrationActionTransferClass, error) {
	whereClause, attributes := makeQueryLikeNFTMigrationActionTransferClassWhereClauseFromFilter(filter)
	row := tx.QueryRow(
		fmt.Sprintf(`SELECT
    id,
    created_at,
    evm_class_id,
    cosmos_owner,
    evm_owner,
    status,
    evm_tx_hash,
    failed_reason
FROM likenft_migration_action_transfer_class
%s
`, whereClause),
		attributes...,
	)

	m := &model.LikeNFTMigrationActionTransferClass{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.EvmClassId,
		&m.CosmosOwner,
		&m.EvmOwner,
		&m.Status,
		&m.EvmTxHash,
		&m.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func InsertLikeNFTMigrationActionTransferClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionTransferClass,
) (*model.LikeNFTMigrationActionTransferClass, error) {
	row := tx.QueryRow(
		`INSERT INTO likenft_migration_action_transfer_class (
    evm_class_id,
    cosmos_owner,
    evm_owner,
    status,
    evm_tx_hash,
    failed_reason
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id`,
		a.EvmClassId,
		a.CosmosOwner,
		a.EvmOwner,
		a.Status,
		a.EvmTxHash,
		a.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	a.Id = uint64(lastInsertId)

	return a, nil
}

func UpdateLikeNFTMigrationActionTransferClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionTransferClass,
) error {
	result, err := tx.Exec(
		`UPDATE likenft_migration_action_transfer_class SET
    evm_class_id = $1,
    cosmos_owner = $2,
    evm_owner = $3,
    status = $4,
    evm_tx_hash = $5,
    failed_reason = $6
WHERE id = $7`,
		a.EvmClassId,
		a.CosmosOwner,
		a.EvmOwner,
		a.Status,
		a.EvmTxHash,
		a.FailedReason,
		a.Id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("update affect no rows")
	}

	return nil
}
