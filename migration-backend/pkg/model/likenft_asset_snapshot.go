package model

import "time"

type LikeNFTAssetSnapshotStatus string

const (
	NFTSnapshotStatusInit       LikeNFTAssetSnapshotStatus = "init"
	NFTSnapshotStatusInProgress LikeNFTAssetSnapshotStatus = "in_progress"
	NFTSnapshotStatusCompleted  LikeNFTAssetSnapshotStatus = "completed"
	NFTSnapshotStatusFailed     LikeNFTAssetSnapshotStatus = "failed"
)

type LikeNFTAssetSnapshot struct {
	Id            uint64
	CreatedAt     time.Time
	CosmosAddress string
	BlockHeight   *string
	BlockTime     *time.Time
	Status        LikeNFTAssetSnapshotStatus
	FailedReason  *string

	EstimatedMigrationDurationNeeded *time.Duration
}
