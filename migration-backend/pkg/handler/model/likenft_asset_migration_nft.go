package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetMigrationNFT struct {
	Id                  uint64                               `json:"id"`
	NFTAssetMigrationId uint64                               `json:"nft_asset_migration_id"`
	CreatedAt           time.Time                            `json:"created_at"`
	CosmosClassId       string                               `json:"cosmos_class_id"`
	CosmosNFTId         string                               `json:"cosmos_nft_id"`
	Name                string                               `json:"name"`
	Image               string                               `json:"image"`
	Status              model.LikeNFTAssetMigrationNFTStatus `json:"status"`
	EnqueueTime         *time.Time                           `json:"enqueue_time"`
	FinishTime          *time.Time                           `json:"finish_time"`
	EvmTxHash           *string                              `json:"evm_tx_hash"`
	FailedReason        *string                              `json:"failed_reason"`
}

func LikeNFTAssetMigrationNFTFromModel(nft *model.LikeNFTAssetMigrationNFT) *LikeNFTAssetMigrationNFT {
	return &LikeNFTAssetMigrationNFT{
		Id:                  nft.Id,
		NFTAssetMigrationId: nft.LikeNFTAssetMigrationId,
		CreatedAt:           nft.CreatedAt,
		CosmosClassId:       nft.CosmosClassId,
		CosmosNFTId:         nft.CosmosNFTId,
		Name:                nft.Name,
		Image:               nft.Image,
		Status:              nft.Status,
		EnqueueTime:         nft.EnqueueTime,
		FinishTime:          nft.FinishTime,
		EvmTxHash:           nft.EvmTxHash,
		FailedReason:        nft.FailedReason,
	}
}
