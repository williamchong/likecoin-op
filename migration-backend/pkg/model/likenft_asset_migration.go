package model

import "time"

type LikeNFTAssetMigrationStatus string

const (
	NFTMigrationStatusInit       LikeNFTAssetMigrationStatus = "init"
	NFTMigrationStatusInProgress LikeNFTAssetMigrationStatus = "in_progress"
	NFTMigrationStatusCompleted  LikeNFTAssetMigrationStatus = "completed"
	NFTMigrationStatusFailed     LikeNFTAssetMigrationStatus = "failed"
)

func (s *LikeNFTAssetMigrationStatus) IsValid() bool {
	switch *s {
	case NFTMigrationStatusInit,
		NFTMigrationStatusInProgress,
		NFTMigrationStatusCompleted,
		NFTMigrationStatusFailed:
		return true
	}
	return false
}

type LikeNFTAssetMigration struct {
	Id                     uint64
	CreatedAt              time.Time
	LikeNFTAssetSnapshotId uint64
	CosmosAddress          string
	EthAddress             string
	Status                 LikeNFTAssetMigrationStatus
	EstimatedFinishedTime  time.Time
	FailedReason           *string
}
