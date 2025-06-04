package model

import "time"

type LikeNFTAssetSnapshotNFT struct {
	Id            uint64
	NFTSnapshotId uint64
	CreatedAt     time.Time
	CosmosClassId string
	CosmosNFTId   string
	Name          string
	Image         string

	EstimatedMigrationDurationNeeded time.Duration
}
