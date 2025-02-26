package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryLikeNFTAssetMigrationNFTById(
	tx TxLike,
	id uint64,
) (*model.LikeNFTAssetMigrationNFT, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	likenft_asset_migration_id,
	created_at,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
FROM likenft_asset_migration_nft WHERE id = $1`,
		id,
	)

	nft := &model.LikeNFTAssetMigrationNFT{}

	err := row.Scan(
		&nft.Id,
		&nft.LikeNFTAssetMigrationId,
		&nft.CreatedAt,
		&nft.CosmosClassId,
		&nft.CosmosNFTId,
		&nft.Name,
		&nft.Image,
		&nft.Status,
		&nft.EnqueueTime,
		&nft.FinishTime,
		&nft.EvmTxHash,
		&nft.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return nft, nil
}

func QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(
	tx TxLike,
	migrationId uint64,
) ([]model.LikeNFTAssetMigrationNFT, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	likenft_asset_migration_id,
	created_at,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	image,
	status,
	enqueue_time,
	finish_time,
	evm_tx_hash,
	failed_reason
FROM likenft_asset_migration_nft WHERE likenft_asset_migration_id = $1`,
		migrationId,
	)

	if err != nil {
		return nil, err
	}

	nfts := []model.LikeNFTAssetMigrationNFT{}
	for rows.Next() {
		nft := &model.LikeNFTAssetMigrationNFT{}

		err := rows.Scan(
			&nft.Id,
			&nft.LikeNFTAssetMigrationId,
			&nft.CreatedAt,
			&nft.CosmosClassId,
			&nft.CosmosNFTId,
			&nft.Name,
			&nft.Image,
			&nft.Status,
			&nft.EnqueueTime,
			&nft.FinishTime,
			&nft.EvmTxHash,
			&nft.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		nfts = append(nfts, *nft)
	}

	return nfts, nil
}

func InsertLikeNFTAssetMigrationNFTs(
	tx TxLike,
	nfts []model.LikeNFTAssetMigrationNFT,
) error {
	if len(nfts) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(nfts))
	numCol := 10
	valueArgs := make([]interface{}, 0, len(nfts)*numCol)

	for i, nft := range nfts {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*numCol+1, i*numCol+2, i*numCol+3, i*numCol+4, i*numCol+5, i*numCol+6, i*numCol+7, i*numCol+8, i*numCol+9, i*numCol+10))
		valueArgs = append(valueArgs, nft.LikeNFTAssetMigrationId)
		valueArgs = append(valueArgs, nft.CosmosClassId)
		valueArgs = append(valueArgs, nft.CosmosNFTId)
		valueArgs = append(valueArgs, nft.Name)
		valueArgs = append(valueArgs, nft.Image)
		valueArgs = append(valueArgs, nft.Status)
		valueArgs = append(valueArgs, nft.EnqueueTime)
		valueArgs = append(valueArgs, nft.FinishTime)
		valueArgs = append(valueArgs, nft.EvmTxHash)
		valueArgs = append(valueArgs, nft.FailedReason)
	}

	stmt := fmt.Sprintf(`INSERT INTO likenft_asset_migration_nft (
	likenft_asset_migration_id,
	cosmos_class_id,
	cosmos_nft_id,
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

func UpdateLikeNFTAssetMigrationNFT(
	tx TxLike,
	class *model.LikeNFTAssetMigrationNFT,
) error {
	_, err := tx.Exec(
		`UPDATE likenft_asset_migration_nft SET
	likenft_asset_migration_id = $1,
	cosmos_class_id = $2,
	cosmos_nft_id = $3,
	name = $4,
	image = $5,
	status = $6,
	enqueue_time = $7,
	finish_time = $8,
	evm_tx_hash = $9,
	failed_reason = $10
WHERE id = $11;`,
		class.LikeNFTAssetMigrationId,
		class.CosmosClassId,
		class.CosmosNFTId,
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
