package model

import "time"

type LikeNFTAssetMigrationStatus string

const (
	NFTMigrationStatusInit       LikeNFTAssetMigrationStatus = "init"
	NFTMigrationStatusInProgress LikeNFTAssetMigrationStatus = "in_progress"
	NFTMigrationStatusCompleted  LikeNFTAssetMigrationStatus = "completed"
	NFTMigrationStatusFailed     LikeNFTAssetMigrationStatus = "failed"
)

type LikeNFTAssetMigration struct {
	Id                     uint64
	CreatedAt              time.Time
	LikeNFTAssetSnapshotId uint64
	CosmosAddress          string
	EthAddress             string
	Status                 LikeNFTAssetMigrationStatus
	FailedReason           *string
}
