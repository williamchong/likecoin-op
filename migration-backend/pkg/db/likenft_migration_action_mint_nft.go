package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type QueryLikeNFTMigrationActionMintNFTFilter struct {
	EvmClassId  *string
	CosmosNFTId *string
}

func makeQueryLikeNFTMigrationActionMintNFTWhereClauseFromFilter(f QueryLikeNFTMigrationActionMintNFTFilter) (string, []any) {
	wheres := make([]string, 0)
	attributes := make([]any, 0)

	if f.EvmClassId != nil {
		wheres = append(wheres, fmt.Sprintf("evm_class_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.EvmClassId)
	}

	if f.CosmosNFTId != nil {
		wheres = append(wheres, fmt.Sprintf("cosmos_nft_id = $%d", len(wheres)+1))
		attributes = append(attributes, f.CosmosNFTId)
	}

	whereClause := ""
	if len(wheres) > 0 {
		whereClause = fmt.Sprintf("%s %s", "WHERE", strings.Join(wheres, " AND "))
	}

	return whereClause, attributes
}

func QueryLikeNFTMigrationActionMintNFT(
	tx TxLike,
	filter QueryLikeNFTMigrationActionMintNFTFilter,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	whereClause, attributes := makeQueryLikeNFTMigrationActionMintNFTWhereClauseFromFilter(filter)
	row := tx.QueryRow(
		fmt.Sprintf(`SELECT
    id,
    created_at,
    evm_class_id,
    cosmos_nft_id,
    initial_batch_mint_owner,
    evm_owner,
    status,
    evm_tx_hash,
    failed_reason
FROM likenft_migration_action_mint_nft
%s
`, whereClause),
		attributes...,
	)

	m := &model.LikeNFTMigrationActionMintNFT{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.EvmClassId,
		&m.CosmosNFTId,
		&m.InitialBatchMintOwner,
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

func InsertLikeNFTMigrationActionMintNFT(
	tx TxLike,
	a *model.LikeNFTMigrationActionMintNFT,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	row := tx.QueryRow(
		`INSERT INTO likenft_migration_action_mint_nft (
    evm_class_id,
    cosmos_nft_id,
    initial_batch_mint_owner,
    evm_owner,
    status,
    evm_tx_hash,
    failed_reason
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id`,
		a.EvmClassId,
		a.CosmosNFTId,
		a.InitialBatchMintOwner,
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

func UpdateLikeNFTMigrationActionMintNFT(
	tx TxLike,
	a *model.LikeNFTMigrationActionMintNFT,
) error {
	result, err := tx.Exec(
		`UPDATE likenft_migration_action_mint_nft SET
    evm_class_id = $1,
    cosmos_nft_id = $2,
    initial_batch_mint_owner = $3,
    evm_owner = $4,
    status = $5,
    evm_tx_hash = $6,
    failed_reason = $7
WHERE id = $8`,
		a.EvmClassId,
		a.CosmosNFTId,
		a.InitialBatchMintOwner,
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
