package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryLikeNFTAssetMigrationClassById(
	tx TxLike,
	id uint64,
) (*model.LikeNFTAssetMigrationClass, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	likenft_asset_migration_id,
	created_at,
	cosmos_class_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
FROM likenft_asset_migration_class WHERE id = $1`,
		id,
	)

	class := &model.LikeNFTAssetMigrationClass{}
	err := row.Scan(
		&class.Id,
		&class.LikeNFTAssetMigrationId,
		&class.CreatedAt,
		&class.CosmosClassId,
		&class.Name,
		&class.Image,
		&class.Status,
		&class.EnqueueTime,
		&class.FinishTime,
		&class.EvmTxHash,
		&class.FailedReason,
	)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func QueryLikeNFTAssetMigrationClassesByNFTMigrationId(
	tx TxLike,
	migrationId uint64,
) ([]model.LikeNFTAssetMigrationClass, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	likenft_asset_migration_id,
	created_at,
	cosmos_class_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
FROM likenft_asset_migration_class WHERE likenft_asset_migration_id = $1`,
		migrationId,
	)

	if err != nil {
		return nil, err
	}

	classes := []model.LikeNFTAssetMigrationClass{}
	for rows.Next() {
		class := &model.LikeNFTAssetMigrationClass{}

		err := rows.Scan(
			&class.Id,
			&class.LikeNFTAssetMigrationId,
			&class.CreatedAt,
			&class.CosmosClassId,
			&class.Name,
			&class.Image,
			&class.Status,
			&class.EnqueueTime,
			&class.FinishTime,
			&class.EvmTxHash,
			&class.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		classes = append(classes, *class)
	}

	return classes, nil
}

func QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
	tx TxLike,
	migrationId uint64,
	status model.LikeLikeNFTAssetAssetMigrationClassStatus,
) ([]model.LikeNFTAssetMigrationClass, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	likenft_asset_migration_id,
	created_at,
	cosmos_class_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
FROM likenft_asset_migration_class
WHERE likenft_asset_migration_id = $1
	AND status = $2`,
		migrationId,
		status,
	)

	if err != nil {
		return nil, err
	}

	classes := []model.LikeNFTAssetMigrationClass{}
	for rows.Next() {
		class := &model.LikeNFTAssetMigrationClass{}

		err := rows.Scan(
			&class.Id,
			&class.LikeNFTAssetMigrationId,
			&class.CreatedAt,
			&class.CosmosClassId,
			&class.Name,
			&class.Image,
			&class.Status,
			&class.EnqueueTime,
			&class.FinishTime,
			&class.EvmTxHash,
			&class.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		classes = append(classes, *class)
	}

	return classes, nil
}

func InsertLikeNFTAssetMigrationClasses(
	tx TxLike,
	classes []model.LikeNFTAssetMigrationClass,
) error {
	if len(classes) == 0 {
		return nil
	}
	valueStrings := make([]string, 0, len(classes))
	numCol := 9
	valueArgs := make([]interface{}, 0, len(classes)*numCol)

	for i, class := range classes {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*numCol+1, i*numCol+2, i*numCol+3, i*numCol+4, i*numCol+5, i*numCol+6, i*numCol+7, i*numCol+8, i*numCol+9))
		valueArgs = append(valueArgs, class.LikeNFTAssetMigrationId)
		valueArgs = append(valueArgs, class.CosmosClassId)
		valueArgs = append(valueArgs, class.Name)
		valueArgs = append(valueArgs, class.Image)
		valueArgs = append(valueArgs, class.Status)
		valueArgs = append(valueArgs, class.EnqueueTime)
		valueArgs = append(valueArgs, class.FinishTime)
		valueArgs = append(valueArgs, class.EvmTxHash)
		valueArgs = append(valueArgs, class.FailedReason)
	}

	stmt := fmt.Sprintf(`INSERT INTO likenft_asset_migration_class (
	likenft_asset_migration_id,
	cosmos_class_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
) VALUES %s`, strings.Join(valueStrings, ","))

	_, err := tx.Exec(stmt, valueArgs...)
	return err
}

func UpdateLikeNFTAssetMigrationClass(
	tx TxLike,
	class *model.LikeNFTAssetMigrationClass,
) error {
	_, err := tx.Exec(
		`UPDATE likenft_asset_migration_class SET
	likenft_asset_migration_id = $1,
	cosmos_class_id = $2,
	name = $3,
	image = $4,
	status = $5,
	enqueue_time = $6,
	finish_time = $7,
	evm_tx_hash = $8,
	failed_reason = $9
WHERE id = $10;`,
		class.LikeNFTAssetMigrationId,
		class.CosmosClassId,
		class.Name,
		class.Image,
		class.Status,
		class.EnqueueTime,
		class.FinishTime,
		class.EvmTxHash,
		class.FailedReason,
		class.Id,
	)

	return err
}
