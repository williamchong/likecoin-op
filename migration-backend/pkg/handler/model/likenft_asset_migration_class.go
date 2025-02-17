package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetMigrationClass struct {
	Id                  uint64                                          `json:"id"`
	NFTAssetMigrationId uint64                                          `json:"nft_asset_migration_id"`
	CreatedAt           time.Time                                       `json:"created_at"`
	CosmosClassId       string                                          `json:"cosmos_class_id"`
	Name                string                                          `json:"name"`
	Image               string                                          `json:"image"`
	Status              model.LikeLikeNFTAssetAssetMigrationClassStatus `json:"status"`
	EnqueueTime         *time.Time                                      `json:"enqueue_time"`
	FinishTime          *time.Time                                      `json:"finish_time"`
	EvmTxHash           *string                                         `json:"evm_tx_hash"`
	FailedReason        *string                                         `json:"failed_reason"`
}

func LikeNFTAssetMigrationClassFromModel(c *model.LikeNFTAssetMigrationClass) *LikeNFTAssetMigrationClass {
	return &LikeNFTAssetMigrationClass{
		Id:                  c.Id,
		NFTAssetMigrationId: c.LikeNFTAssetMigrationId,
		CreatedAt:           c.CreatedAt,
		CosmosClassId:       c.CosmosClassId,
		Name:                c.Name,
		Image:               c.Image,
		Status:              c.Status,
		EnqueueTime:         c.EnqueueTime,
		FinishTime:          c.FinishTime,
		EvmTxHash:           c.EvmTxHash,
		FailedReason:        c.FailedReason,
	}

}
