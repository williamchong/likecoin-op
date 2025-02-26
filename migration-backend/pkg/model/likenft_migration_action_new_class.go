package model

import "time"

type LikeNFTMigrationActionNewClassStatus string

const (
	LikeNFTMigrationActionNewClassStatusInit       LikeNFTMigrationActionNewClassStatus = "init"
	LikeNFTMigrationActionNewClassStatusInProgress LikeNFTMigrationActionNewClassStatus = "in_progress"
	LikeNFTMigrationActionNewClassStatusCompleted  LikeNFTMigrationActionNewClassStatus = "completed"
	LikeNFTMigrationActionNewClassStatusFailed     LikeNFTMigrationActionNewClassStatus = "failed"
)

type LikeNFTMigrationActionNewClass struct {
	Id             uint64
	CreatedAt      time.Time
	CosmosClassId  string
	InitialOwner   string
	InitialMinter  string
	InitialUpdater string
	Status         LikeNFTMigrationActionNewClassStatus
	EvmClassId     *string
	EvmTxHash      *string
	FailedReason   *string
}
