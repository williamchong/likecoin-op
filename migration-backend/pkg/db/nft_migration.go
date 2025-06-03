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
	estimated_finished_time,
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
		&migration.EstimatedFinishedTime,
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
	estimated_finished_time,
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
		&migration.EstimatedFinishedTime,
		&migration.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return migration, nil
}

func QueryPaginatedLikeNFTAssetMigration(
	tx TxLike,
	limit int,
	offset int,
	status *model.LikeNFTAssetMigrationStatus,
	keyword string,
) ([]*model.LikeNFTAssetMigration, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	created_at,
	likenft_asset_snapshot_id,
	cosmos_address,
	eth_address,
	status,
	estimated_finished_time,
	failed_reason
FROM likenft_asset_migration
WHERE ($3::text IS NULL OR status = $3) AND
(
	$4::text = '' OR 
	failed_reason ILIKE '%' || $4 || '%' OR 
	cosmos_address ILIKE '%' || $4 || '%' OR 
	eth_address ILIKE '%' || $4 || '%'
)
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`,
		limit,
		offset,
		status,
		keyword,
	)

	if err != nil {
		return nil, err
	}

	migrations := []*model.LikeNFTAssetMigration{}

	for rows.Next() {
		m := &model.LikeNFTAssetMigration{}
		err := rows.Scan(
			&m.Id,
			&m.CreatedAt,
			&m.LikeNFTAssetSnapshotId,
			&m.CosmosAddress,
			&m.EthAddress,
			&m.Status,
			&m.EstimatedFinishedTime,
			&m.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		migrations = append(migrations, m)
	}

	return migrations, nil
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
	estimated_finished_time,
	failed_reason
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		migration.LikeNFTAssetSnapshotId,
		migration.CosmosAddress,
		migration.EthAddress,
		migration.Status,
		migration.EstimatedFinishedTime,
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
	estimated_finished_time = $5,
	failed_reason = $6
WHERE id = $7;`,
		migration.LikeNFTAssetSnapshotId,
		migration.CosmosAddress,
		migration.EthAddress,
		migration.Status,
		migration.EstimatedFinishedTime,
		migration.FailedReason,
		migration.Id,
	)

	return err
}

func RemoveLikeNFTAssetMigration(
	tx TxLike,
	id uint64,
) error {
	_, err := tx.Exec(
		`DELETE FROM likenft_asset_migration WHERE id = $1;`,
		id,
	)

	return err
}
