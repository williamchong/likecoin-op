package model

import "time"

type LikeNFTAssetSnapshotNFT struct {
	Id            uint64    `json:"id"`
	NFTSnapshotId uint64    `json:"nft_snapshot_id"`
	CreatedAt     time.Time `json:"created_at"`
	CosmosClassId string    `json:"cosmos_class_id"`
	CosmosNFTId   string    `json:"cosmos_nft_id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
}
