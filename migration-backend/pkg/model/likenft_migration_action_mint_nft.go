package model

import "time"

type LikeNFTMigrationActionMintNFTStatus string

const (
	LikeNFTMigrationActionMintNFTStatusInit       LikeNFTMigrationActionMintNFTStatus = "init"
	LikeNFTMigrationActionMintNFTStatusInProgress LikeNFTMigrationActionMintNFTStatus = "in_progress"
	LikeNFTMigrationActionMintNFTStatusCompleted  LikeNFTMigrationActionMintNFTStatus = "completed"
	LikeNFTMigrationActionMintNFTStatusFailed     LikeNFTMigrationActionMintNFTStatus = "failed"
)

type LikeNFTMigrationActionMintNFT struct {
	Id                    uint64
	CreatedAt             time.Time
	EvmClassId            string
	CosmosNFTId           string
	InitialBatchMintOwner string
	EvmOwner              string
	Status                LikeNFTMigrationActionMintNFTStatus
	EvmTxHash             *string
	FailedReason          *string
}
