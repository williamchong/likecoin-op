package db

import "github.com/likecoin/like-migration-backend/pkg/model"

func QueryLikeNFTAssetMigrationById(
	tx TxLike,
	id uint64,
) (*model.LikeNFTAssetMigration, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	created_at,
	likenft_asset_snapshot_id,
	cosmos_address,
	eth_address,
	status,
	failed_reason
FROM likenft_asset_migration WHERE id = $1`,
		id,
	)

	migration := &model.LikeNFTAssetMigration{}

	err := row.Scan(
		&migration.Id,
		&migration.CreatedAt,
		&migration.LikeNFTAssetSnapshotId,
		&migration.CosmosAddress,
		&migration.EthAddress,
		&migration.Status,
		&migration.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return migration, nil
}

func QueryLikeNFTAssetMigrationByCosmosAddress(
	tx TxLike,
	cosmosAddress string,
) (*model.LikeNFTAssetMigration, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	created_at,
	likenft_asset_snapshot_id,
	cosmos_address,
	eth_address,
	status,
	failed_reason
FROM likenft_asset_migration WHERE cosmos_address = $1`,
		cosmosAddress,
	)

	migration := &model.LikeNFTAssetMigration{}

	err := row.Scan(
		&migration.Id,
		&migration.CreatedAt,
		&migration.LikeNFTAssetSnapshotId,
		&migration.CosmosAddress,
		&migration.EthAddress,
		&migration.Status,
		&migration.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return migration, nil
}

func InsertLikeNFTAssetMigration(
	tx TxLike,
	migration *model.LikeNFTAssetMigration,
) (*model.LikeNFTAssetMigration, error) {
	row := tx.QueryRow(
		`INSERT INTO likenft_asset_migration (
	likenft_asset_snapshot_id,
	cosmos_address,
	eth_address,
	status,
	failed_reason
) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		migration.LikeNFTAssetSnapshotId,
		migration.CosmosAddress,
		migration.EthAddress,
		migration.Status,
		migration.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	n := *migration
	n.Id = uint64(lastInsertId)

	return &n, nil
}

func UpdateLikeNFTAssetMigration(
	tx TxLike,
	migration *model.LikeNFTAssetMigration,
) error {
	_, err := tx.Exec(
		`UPDATE likenft_asset_migration SET
	likenft_asset_snapshot_id = $1,
	cosmos_address = $2,
	eth_address = $3,
	status = $4,
	failed_reason = $5
WHERE id = $6;`,
		migration.LikeNFTAssetSnapshotId,
		migration.CosmosAddress,
		migration.EthAddress,
		migration.Status,
		migration.FailedReason,
		migration.Id,
	)

	return err
}
