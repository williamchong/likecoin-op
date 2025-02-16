package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetSnapshot struct {
	Id            uint64                           `json:"id"`
	CreatedAt     time.Time                        `json:"created_at"`
	CosmosAddress string                           `json:"cosmos_address"`
	BlockHeight   *string                          `json:"block_height"`
	BlockTime     *time.Time                       `json:"block_time"`
	Status        model.LikeNFTAssetSnapshotStatus `json:"status"`
	FailedReason  *string                          `json:"failed_reason"`
	Classes       []LikeNFTAssetSnapshotClass      `json:"classes"`
	NFTs          []LikeNFTAssetSnapshotNFT        `json:"nfts"`
}
