package likenft

import (
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func RecalculateMigrationStatus(
	tx appdb.TxLike,
	migrationId uint64,
) error {
	migration, err := appdb.QueryLikeNFTAssetMigrationById(tx, migrationId)
	if err != nil {
		return err
	}

	classes, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(tx, migrationId)
	if err != nil {
		return err
	}

	nfts, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(tx, migrationId)
	if err != nil {
		return err
	}

	migration.Status = DetermineMigrationStatus(classes, nfts)
	return appdb.UpdateLikeNFTAssetMigration(tx, migration)
}

func DetermineMigrationStatus(
	classes []model.LikeNFTAssetMigrationClass,
	nfts []model.LikeNFTAssetMigrationNFT,
) model.LikeNFTAssetMigrationStatus {
	var allCompleted = true
	var someFailed = false
	var someInProgress = false

	for _, class := range classes {
		allCompleted = allCompleted && class.Status == model.LikeNFTAssetMigrationClassStatusCompleted
		someFailed = someFailed || class.Status == model.LikeNFTAssetMigrationClassStatusFailed
		someInProgress = someInProgress ||
			(class.Status == model.LikeNFTAssetMigrationClassStatusInProgress ||
				class.Status == model.LikeNFTAssetMigrationClassStatusInit)
	}

	for _, nft := range nfts {
		allCompleted = allCompleted && nft.Status == model.LikeNFTAssetMigrationNFTStatusCompleted
		someFailed = someFailed || nft.Status == model.LikeNFTAssetMigrationNFTStatusFailed
		someInProgress = someInProgress ||
			(nft.Status == model.LikeNFTAssetMigrationNFTStatusInProgress ||
				nft.Status == model.LikeNFTAssetMigrationNFTStatusInit)
	}

	if someInProgress {
		return model.NFTMigrationStatusInProgress
	}
	if someFailed {
		return model.NFTMigrationStatusFailed
	}
	if allCompleted {
		return model.NFTMigrationStatusCompleted
	}
	panic("unknwon migration status")
}
