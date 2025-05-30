package model

import "time"

type LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus string

const (
	LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusInit       LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus = "init"
	LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusInProgress LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus = "in_progress"
	LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusCompleted  LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus = "completed"
	LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusFailed     LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus = "failed"
)

type LikeNFTMigrationActionBatchMintNFTsFromCosmos struct {
	Id                    uint64
	CreatedAt             time.Time
	EvmClassId            string
	CurrentSupply         uint64
	ExpectedSupply        uint64
	BatchMintAmount       uint64
	InitialBatchMintOwner string
	Status                LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus
	FromID                *uint64
	ToID                  *uint64
	EvmTxHash             *string
	FailedReason          *string
}
