package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type QueryLikeNFTMigrationActionNewClassFilter struct {
	CosmosClassId *string
	EvmClassId    *string
}

func makeQueryLikeNFTMigrationActionNewClassWhereClauseFromFilter(f QueryLikeNFTMigrationActionNewClassFilter) (string, []any) {
	wheres := make([]string, 0)
	attributes := make([]any, 0)

	if f.CosmosClassId != nil {
		wheres = append(wheres, fmt.Sprintf("cosmos_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.CosmosClassId)
	}

	if f.EvmClassId != nil {
		wheres = append(wheres, fmt.Sprintf("evm_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.EvmClassId)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryLikeNFTMigrationActionNewClass(
	tx TxLike,
	filter QueryLikeNFTMigrationActionNewClassFilter,
) (*model.LikeNFTMigrationActionNewClass, error) {
	whereClause, attributes := makeQueryLikeNFTMigrationActionNewClassWhereClauseFromFilter(filter)
	row := tx.QueryRow(
		fmt.Sprintf(`SELECT
    id,
    created_at,
    cosmos_class_id,
    initial_owner,
    initial_minter,
    initial_updater,
    status,
    evm_class_id,
    evm_tx_hash,
    failed_reason
FROM likenft_migration_action_new_class
%s
`, whereClause),
		attributes...,
	)

	m := &model.LikeNFTMigrationActionNewClass{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.CosmosClassId,
		&m.InitialOwner,
		&m.InitialMintersStr,
		&m.InitialUpdater,
		&m.Status,
		&m.EvmClassId,
		&m.EvmTxHash,
		&m.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func InsertLikeNFTMigrationActionNewClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionNewClass,
) (*model.LikeNFTMigrationActionNewClass, error) {
	row := tx.QueryRow(
		`INSERT INTO likenft_migration_action_new_class (
    cosmos_class_id,
    initial_owner,
    initial_minter,
    initial_updater,
    status,
    evm_class_id,
    evm_tx_hash,
    failed_reason
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id`,
		a.CosmosClassId,
		a.InitialOwner,
		a.InitialMintersStr,
		a.InitialUpdater,
		a.Status,
		a.EvmClassId,
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

func UpdateLikeNFTMigrationActionNewClass(
	tx TxLike,
	a *model.LikeNFTMigrationActionNewClass,
) error {
	result, err := tx.Exec(
		`UPDATE likenft_migration_action_new_class SET
    cosmos_class_id = $1,
    initial_owner = $2,
    initial_minter = $3,
    initial_updater = $4,
    status = $5,
    evm_class_id = $6,
    evm_tx_hash = $7,
    failed_reason = $8
WHERE id = $9`,
		a.CosmosClassId,
		a.InitialOwner,
		a.InitialMintersStr,
		a.InitialUpdater,
		a.Status,
		a.EvmClassId,
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
