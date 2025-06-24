package db

import (
	"fmt"
	"math"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/util/slice"
)

func QueryLikeNFTAssetSnapshotClassesByNFTSnapshotId(
	tx TxLike,
	snapshotId uint64,
) ([]model.LikeNFTAssetSnapshotClass, error) {
	rows, err := tx.Query(
		`SELECT
	id,
	likenft_asset_snapshot_id,
	created_at,
	cosmos_class_id,
	name,
	image,
	estimated_migration_duration_needed
FROM likenft_asset_snapshot_class WHERE likenft_asset_snapshot_id = $1`,
		snapshotId,
	)

	if err != nil {
		return nil, err
	}

	classes := []model.LikeNFTAssetSnapshotClass{}
	for rows.Next() {
		class := &model.LikeNFTAssetSnapshotClass{}

		err := rows.Scan(
			&class.Id,
			&class.NFTSnapshotId,
			&class.CreatedAt,
			&class.CosmosClassId,
			&class.Name,
			&class.Image,
			&class.EstimatedMigrationDurationNeeded,
		)

		if err != nil {
			return nil, err
		}

		classes = append(classes, *class)
	}

	return classes, nil
}

func InsertLikeNFTAssetSnapshotClasses(
	tx TxLike,
	classes []model.LikeNFTAssetSnapshotClass,
) error {
	if len(classes) == 0 {
		return nil
	}

	numCol := 5
	chunkSize := int(math.Floor(float64(PGSQL_DB_PARAMS_LIMIT) / float64(numCol)))

	for _, chunk := range slice.ChunkBy(classes, chunkSize) {
		valueStrings := make([]string, 0, len(classes))
		valueArgs := make([]interface{}, 0, len(classes)*numCol)

		for i, class := range chunk {
			valueStrings = append(valueStrings, fmt.Sprintf(
				"($%d, $%d, $%d, $%d, $%d)",
				i*numCol+1, i*numCol+2, i*numCol+3, i*numCol+4, i*numCol+5,
			))
			valueArgs = append(valueArgs, class.NFTSnapshotId)
			valueArgs = append(valueArgs, class.CosmosClassId)
			valueArgs = append(valueArgs, class.Name)
			valueArgs = append(valueArgs, class.Image)
			valueArgs = append(valueArgs, class.EstimatedMigrationDurationNeeded)
		}

		stmt := fmt.Sprintf(`INSERT INTO likenft_asset_snapshot_class (
	likenft_asset_snapshot_id,
	cosmos_class_id,
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
