package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/evm/model"
	"likenft-indexer/openapi/api"
)

func MakeNFTClass(e *ent.NFTClass, metadataAdditionalProps APIAdditionalProps) api.BookNFT {
	var opensea *model.ContractLevelMetadataOpenSea
	if e.Metadata != nil {
		opensea = &e.Metadata.ContractLevelMetadataOpenSea
	}
	return api.BookNFT{
		ID:                  e.ID,
		Address:             e.Address,
		Name:                e.Name,
		Symbol:              e.Symbol,
		OwnerAddress:        MakeOptString(e.OwnerAddress),
		TotalSupply:         MakeBigInt(e.TotalSupply),
		MaxSupply:           MakeUint64(uint64(e.MaxSupply)),
		Metadata:            MakeOptContractLevelMetadata(opensea, metadataAdditionalProps),
		BannerImage:         e.BannerImage,
		FeaturedImage:       e.FeaturedImage,
		DeployerAddress:     e.DeployerAddress,
		DeployedBlockNumber: MakeUint64(uint64(e.DeployedBlockNumber)),
		MintedAt:            e.MintedAt,
		UpdatedAt:           e.UpdatedAt,
		Owner:               MakeOptAccount(e.Edges.Owner),
	}
}
