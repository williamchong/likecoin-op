package db

import "github.com/likecoin/like-migration-backend/pkg/model"

func QueryLatestLikeNFTAssetSnapshotByCosmosAddress(
	tx TxLike,
	cosmosAddress string,
) (*model.LikeNFTAssetSnapshot, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	created_at,
	cosmos_address,
	block_height,
	block_time,
	status,
	failed_reason
FROM likenft_asset_snapshot WHERE cosmos_address = $1 ORDER BY created_at DESC`,
		cosmosAddress,
	)

	snapshot := &model.LikeNFTAssetSnapshot{}

	err := row.Scan(
		&snapshot.Id,
		&snapshot.CreatedAt,
		&snapshot.CosmosAddress,
		&snapshot.BlockHeight,
		&snapshot.BlockTime,
		&snapshot.Status,
		&snapshot.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return snapshot, nil
}

func InsertLikeNFTAssetSnapshot(
	tx TxLike,
	snapshot *model.LikeNFTAssetSnapshot,
) (*model.LikeNFTAssetSnapshot, error) {
	row := tx.QueryRow(
		`INSERT INTO likenft_asset_snapshot (
	cosmos_address,
	block_height,
	block_time,
	status,
	failed_reason
) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		snapshot.CosmosAddress,
		snapshot.BlockHeight,
		snapshot.BlockTime,
		snapshot.Status,
		snapshot.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	n := *snapshot
	n.Id = uint64(lastInsertId)

	return &n, nil
}

func UpdateLikeNFTAssetSnapshot(
	tx TxLike,
	snapshot *model.LikeNFTAssetSnapshot,
) error {
	_, err := tx.Exec(
		`UPDATE likenft_asset_snapshot SET
	cosmos_address = $1,
	block_height = $2,
	block_time = $3,
	status = $4,
	failed_reason = $5
WHERE likenft_asset_snapshot.id = $6;`,
		snapshot.CosmosAddress,
		snapshot.BlockHeight,
		snapshot.BlockTime,
		snapshot.Status,
		snapshot.FailedReason,
		snapshot.Id,
	)

	return err
}
