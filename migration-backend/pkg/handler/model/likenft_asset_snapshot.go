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

func LikeNFTAssetSnapshotFromModel(
	m *model.LikeNFTAssetSnapshot,
	classes []model.LikeNFTAssetSnapshotClass,
	nfts []model.LikeNFTAssetSnapshotNFT,
) *LikeNFTAssetSnapshot {
	cs := make([]LikeNFTAssetSnapshotClass, 0)

	for _, class := range classes {
		cs = append(cs, *LikeNFTAssetSnapshotClassFromModel(&class))
	}

	ns := make([]LikeNFTAssetSnapshotNFT, 0)

	for _, nft := range nfts {
		ns = append(ns, *LikeNFTAssetSnapshotNFTFromModel(&nft))
	}

	return &LikeNFTAssetSnapshot{
		Id:            m.Id,
		CreatedAt:     m.CreatedAt,
		CosmosAddress: m.CosmosAddress,
		BlockHeight:   m.BlockHeight,
		BlockTime:     m.BlockTime,
		Status:        m.Status,
		FailedReason:  m.FailedReason,
		Classes:       cs,
		NFTs:          ns,
	}
}
