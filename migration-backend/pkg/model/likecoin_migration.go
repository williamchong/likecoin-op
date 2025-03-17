package model

import "time"

type LikeCoinMigrationStatus string

const (
	LikeCoinMigrationStatusPendingCosmosTxHash LikeCoinMigrationStatus = "pending_cosmos_tx_hash"
	LikeCoinMigrationStatusVerifyingCosmosTx   LikeCoinMigrationStatus = "verifying_cosmos_tx"
	LikeCoinMigrationStatusEvmMinting          LikeCoinMigrationStatus = "evm_minting"
	LikeCoinMigrationStatusEvmVerifying        LikeCoinMigrationStatus = "evm_verifying"
	LikeCoinMigrationStatusCompleted           LikeCoinMigrationStatus = "completed"
	LikeCoinMigrationStatusFailed              LikeCoinMigrationStatus = "failed"
)

type LikeCoinMigration struct {
	Id                   uint64
	CreatedAt            time.Time
	UserCosmosAddress    string
	BurningCosmosAddress string
	MintingEthAddress    string
	UserEthAddress       string
	EvmSignature         string
	EvmSignatureMessage  string
	Amount               string

	Status       LikeCoinMigrationStatus
	CosmosTxHash *string
	EvmTxHash    *string
	FailedReason *string
}
