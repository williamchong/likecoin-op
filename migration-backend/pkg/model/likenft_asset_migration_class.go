package model

import "time"

type LikeLikeNFTAssetAssetMigrationClassStatus string

const (
	LikeNFTAssetMigrationClassStatusInit       LikeLikeNFTAssetAssetMigrationClassStatus = "init"
	LikeNFTAssetMigrationClassStatusInProgress LikeLikeNFTAssetAssetMigrationClassStatus = "in_progress"
	LikeNFTAssetMigrationClassStatusCompleted  LikeLikeNFTAssetAssetMigrationClassStatus = "completed"
	LikeNFTAssetMigrationClassStatusFailed     LikeLikeNFTAssetAssetMigrationClassStatus = "failed"
)

type LikeNFTAssetMigrationClass struct {
	Id                      uint64
	LikeNFTAssetMigrationId uint64
	CreatedAt               time.Time
	CosmosClassId           string
	Name                    string
	Image                   string
	Status                  LikeLikeNFTAssetAssetMigrationClassStatus
	EstimatedDurationNeeded time.Duration
	EnqueueTime             *time.Time
	FinishTime              *time.Time
	EvmTxHash               *string
	FailedReason            *string
}
