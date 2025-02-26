package model

import "time"

type LikeNFTMigrationActionTransferClassStatus string

const (
	LikeNFTMigrationActionTransferClassStatusInit       LikeNFTMigrationActionTransferClassStatus = "init"
	LikeNFTMigrationActionTransferClassStatusInProgress LikeNFTMigrationActionTransferClassStatus = "in_progress"
	LikeNFTMigrationActionTransferClassStatusCompleted  LikeNFTMigrationActionTransferClassStatus = "completed"
	LikeNFTMigrationActionTransferClassStatusFailed     LikeNFTMigrationActionTransferClassStatus = "failed"
)

type LikeNFTMigrationActionTransferClass struct {
	Id           uint64
	CreatedAt    time.Time
	EvmClassId   string
	CosmosOwner  string
	EvmOwner     string
	Status       LikeNFTMigrationActionTransferClassStatus
	EvmTxHash    *string
	FailedReason *string
}
