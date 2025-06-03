package model

import "time"

type LikeNFTAssetSnapshotClass struct {
	Id            uint64
	NFTSnapshotId uint64
	CreatedAt     *time.Time
	CosmosClassId string
	Name          string
	Image         string

	EstimatedMigrationDurationNeeded time.Duration
}
