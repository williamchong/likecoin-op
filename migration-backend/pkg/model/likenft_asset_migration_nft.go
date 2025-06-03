package model

import "time"

type LikeNFTAssetMigrationNFTStatus string

const (
	LikeNFTAssetMigrationNFTStatusInit       LikeNFTAssetMigrationNFTStatus = "init"
	LikeNFTAssetMigrationNFTStatusInProgress LikeNFTAssetMigrationNFTStatus = "in_progress"
	LikeNFTAssetMigrationNFTStatusCompleted  LikeNFTAssetMigrationNFTStatus = "completed"
	LikeNFTAssetMigrationNFTStatusFailed     LikeNFTAssetMigrationNFTStatus = "failed"
)

type LikeNFTAssetMigrationNFT struct {
	Id                      uint64
	LikeNFTAssetMigrationId uint64
	CreatedAt               time.Time
	CosmosClassId           string
	CosmosNFTId             string
	Name                    string
	Image                   string
	Status                  LikeNFTAssetMigrationNFTStatus
	EstimatedDurationNeeded time.Duration
	EnqueueTime             *time.Time
	FinishTime              *time.Time
	EvmTxHash               *string
	FailedReason            *string
}
