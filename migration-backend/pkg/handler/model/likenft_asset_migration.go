package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetMigrationBase struct {
	Id                     uint64                            `json:"id"`
	CreatedAt              time.Time                         `json:"created_at"`
	LikeNFTAssetSnapshotId uint64                            `json:"likenft_asset_snapshot_id"`
	CosmosAddress          string                            `json:"cosmos_address"`
	EthAddress             string                            `json:"eth_address"`
	Status                 model.LikeNFTAssetMigrationStatus `json:"status"`
	EstimatedFinishedTime  time.Time                         `json:"estimated_finished_time"`
	FailedReason           *string                           `json:"failed_reason"`
}

type LikeNFTAssetMigration struct {
	LikeNFTAssetMigrationBase
	Classes []LikeNFTAssetMigrationClass `json:"classes"`
	NFTs    []LikeNFTAssetMigrationNFT   `json:"nfts"`
}

func LikeNFTAssetMigrationFromModel(m *model.LikeNFTAssetMigration, classes []model.LikeNFTAssetMigrationClass, nfts []model.LikeNFTAssetMigrationNFT) *LikeNFTAssetMigration {
	cs := make([]LikeNFTAssetMigrationClass, 0)

	for _, class := range classes {
		cs = append(cs, *LikeNFTAssetMigrationClassFromModel(&class))
	}

	ns := make([]LikeNFTAssetMigrationNFT, 0)

	for _, nft := range nfts {
		ns = append(ns, *LikeNFTAssetMigrationNFTFromModel(&nft))
	}

	return &LikeNFTAssetMigration{
		LikeNFTAssetMigrationBase: LikeNFTAssetMigrationBase{
			Id:                     m.Id,
			CreatedAt:              m.CreatedAt,
			LikeNFTAssetSnapshotId: m.LikeNFTAssetSnapshotId,
			CosmosAddress:          m.CosmosAddress,
			EthAddress:             m.EthAddress,
			Status:                 m.Status,
			EstimatedFinishedTime:  m.EstimatedFinishedTime,
			FailedReason:           m.FailedReason,
		},
		Classes: cs,
		NFTs:    ns,
	}
}

func LikeNFTAssetMigrationBasesFromModel(m []*model.LikeNFTAssetMigration) []*LikeNFTAssetMigrationBase {
	migrations := make([]*LikeNFTAssetMigrationBase, 0)

	for _, migration := range m {
		migrations = append(migrations, &LikeNFTAssetMigrationBase{
			Id:                     migration.Id,
			CreatedAt:              migration.CreatedAt,
			LikeNFTAssetSnapshotId: migration.LikeNFTAssetSnapshotId,
			CosmosAddress:          migration.CosmosAddress,
			EthAddress:             migration.EthAddress,
			Status:                 migration.Status,
			EstimatedFinishedTime:  migration.EstimatedFinishedTime,
			FailedReason:           migration.FailedReason,
		})
	}

	return migrations
}
