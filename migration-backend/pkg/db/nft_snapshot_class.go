package db

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
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
	image
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
	valueStrings := make([]string, 0, len(classes))
	numCol := 4
	valueArgs := make([]interface{}, 0, len(classes)*numCol)

	for i, class := range classes {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*numCol+1, i*numCol+2, i*numCol+3, i*numCol+4))
		valueArgs = append(valueArgs, class.NFTSnapshotId)
		valueArgs = append(valueArgs, class.CosmosClassId)
		valueArgs = append(valueArgs, class.Name)
		valueArgs = append(valueArgs, class.Image)
	}

	stmt := fmt.Sprintf(`INSERT INTO likenft_asset_snapshot_class (
	likenft_asset_snapshot_id,
	cosmos_class_id,
	name,
	image
) VALUES %s`, strings.Join(valueStrings, ","))

	_, err := tx.Exec(stmt, valueArgs...)
	return err
}
