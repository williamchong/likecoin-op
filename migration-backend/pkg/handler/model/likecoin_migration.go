package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeCoinMigration struct {
	Id                uint64    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UserCosmosAddress string    `json:"user_cosmos_address"`
	UserEthAddress    string    `json:"user_eth_address"`
	EvmSignature      string    `json:"evm_signature"`
	Amount            string    `json:"amount"`

	Status       model.LikeCoinMigrationStatus `json:"status"`
	CosmosTxHash *string                       `json:"cosmos_tx_hash"`
	EvmTxHash    *string                       `json:"evm_tx_hash"`
	FailedReason *string                       `json:"failed_reason"`
}

func LikeCoinMigrationFromModel(c *model.LikeCoinMigration) *LikeCoinMigration {
	return &LikeCoinMigration{
		Id:                c.Id,
		CreatedAt:         c.CreatedAt,
		UserEthAddress:    c.UserEthAddress,
		EvmSignature:      c.EvmSignature,
		Amount:            c.Amount,
		UserCosmosAddress: c.UserCosmosAddress,
		Status:            c.Status,
		CosmosTxHash:      c.CosmosTxHash,
		EvmTxHash:         c.EvmTxHash,
		FailedReason:      c.FailedReason,
	}

}

func LikeCoinMigrationsFromModel(cs []*model.LikeCoinMigration) []*LikeCoinMigration {
	migrations := make([]*LikeCoinMigration, len(cs))
	for i, c := range cs {
		migrations[i] = LikeCoinMigrationFromModel(c)
	}
	return migrations
}
