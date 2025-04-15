package model

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTAssetSnapshotClass struct {
	Id            uint64     `json:"id"`
	NFTSnapshotId uint64     `json:"nft_snapshot_id"`
	CreatedAt     *time.Time `json:"created_at"`
	CosmosClassId string     `json:"cosmos_class_id"`
	Name          string     `json:"name"`
	Image         string     `json:"image"`
}

func LikeNFTAssetSnapshotClassFromModel(c *model.LikeNFTAssetSnapshotClass) *LikeNFTAssetSnapshotClass {
	return &LikeNFTAssetSnapshotClass{
		Id:            c.Id,
		NFTSnapshotId: c.NFTSnapshotId,
		CreatedAt:     c.CreatedAt,
		CosmosClassId: c.CosmosClassId,
		Name:          c.Name,
		Image:         c.Image,
	}
}
