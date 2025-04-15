package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetSnapshotNFT struct {
	Id            uint64    `json:"id"`
	NFTSnapshotId uint64    `json:"nft_snapshot_id"`
	CreatedAt     time.Time `json:"created_at"`
	CosmosClassId string    `json:"cosmos_class_id"`
	CosmosNFTId   string    `json:"cosmos_nft_id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
}

func LikeNFTAssetSnapshotNFTFromModel(nft *model.LikeNFTAssetSnapshotNFT) *LikeNFTAssetSnapshotNFT {
	return &LikeNFTAssetSnapshotNFT{
		Id:            nft.Id,
		NFTSnapshotId: nft.NFTSnapshotId,
		CreatedAt:     nft.CreatedAt,
		CosmosClassId: nft.CosmosClassId,
		CosmosNFTId:   nft.CosmosNFTId,
		Name:          nft.Name,
		Image:         nft.Image,
	}
}
