package db

import (
	"fmt"
	"math"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/util/slice"
)

func QueryLikeNFTAssetSnapshotNFTsByNFTSnapshotId(
	tx TxLike,
	snapshotId uint64,
) ([]model.LikeNFTAssetSnapshotNFT, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	likenft_asset_snapshot_id,
	created_at,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	image,
	estimated_migration_duration_needed
FROM likenft_asset_snapshot_nft WHERE likenft_asset_snapshot_id = $1`,
		snapshotId,
	)

	if err != nil {
		return nil, err
	}

	nfts := []model.LikeNFTAssetSnapshotNFT{}
	for rows.Next() {
		nft := &model.LikeNFTAssetSnapshotNFT{}

		err := rows.Scan(
			&nft.Id,
			&nft.NFTSnapshotId,
			&nft.CreatedAt,
			&nft.CosmosClassId,
			&nft.CosmosNFTId,
			&nft.Name,
			&nft.Image,
			&nft.EstimatedMigrationDurationNeeded,
		)

		if err != nil {
			return nil, err
		}

		nfts = append(nfts, *nft)
	}

	return nfts, nil
}

func InsertLikeNFTAssetSnapshotNFTs(
	tx TxLike,
	nfts []model.LikeNFTAssetSnapshotNFT,
) error {
	if len(nfts) == 0 {
		return nil
	}

	numCol := 6
	chunkSize := int(math.Floor(float64(PGSQL_DB_PARAMS_LIMIT) / float64(numCol)))

	for _, chunk := range slice.ChunkBy(nfts, chunkSize) {
		valueStrings := make([]string, 0, len(nfts))
		valueArgs := make([]interface{}, 0, len(nfts)*numCol)

		for i, nft := range chunk {
			valueStrings = append(valueStrings, fmt.Sprintf(
				"($%d, $%d, $%d, $%d, $%d, $%d)",
				i*numCol+1, i*numCol+2, i*numCol+3, i*numCol+4, i*numCol+5, i*numCol+6,
			))
			valueArgs = append(valueArgs, nft.NFTSnapshotId)
			valueArgs = append(valueArgs, nft.CosmosClassId)
			valueArgs = append(valueArgs, nft.CosmosNFTId)
			valueArgs = append(valueArgs, nft.Name)
			valueArgs = append(valueArgs, nft.Image)
			valueArgs = append(valueArgs, nft.EstimatedMigrationDurationNeeded)
		}

		stmt := fmt.Sprintf(`INSERT INTO likenft_asset_snapshot_nft (
	likenft_asset_snapshot_id,
	cosmos_class_id,
	cosmos_nft_id,
	name,
	image,
	estimated_migration_duration_needed
) VALUES %s`, strings.Join(valueStrings, ","))

		_, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			return err

		}
	}
	return nil

}
