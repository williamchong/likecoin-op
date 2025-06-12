package model

import (
	"math/big"
	"time"

	"github.com/likecoin/like-migration-backend/pkg/types/commaseparatedstring"
)

type LikeNFTMigrationActionNewClassStatus string

const (
	LikeNFTMigrationActionNewClassStatusInit       LikeNFTMigrationActionNewClassStatus = "init"
	LikeNFTMigrationActionNewClassStatusInProgress LikeNFTMigrationActionNewClassStatus = "in_progress"
	LikeNFTMigrationActionNewClassStatusCompleted  LikeNFTMigrationActionNewClassStatus = "completed"
	LikeNFTMigrationActionNewClassStatusFailed     LikeNFTMigrationActionNewClassStatus = "failed"
)

type LikeNFTMigrationActionNewClass struct {
	Id                      uint64
	CreatedAt               time.Time
	CosmosClassId           string
	InitialOwner            string
	InitialMintersStr       commaseparatedstring.CommaSeparatedString
	InitialUpdatersStr      commaseparatedstring.CommaSeparatedString
	InitialBatchMintOwner   string
	DefaultRoyaltyFraction  *big.Int
	ShouldPremintAllNFTs    bool
	Status                  LikeNFTMigrationActionNewClassStatus
	NumberOfCosmosNFTsFound *uint64
	EvmClassId              *string
	EvmTxHash               *string
	FailedReason            *string
}
