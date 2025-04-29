package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/evm/model"
	"likenft-indexer/openapi/api"

	"github.com/go-faster/jx"
)

func MakeNFTClass(e *ent.NFTClass) (*api.BookNFT, error) {
	var (
		opensea                 *model.ContractLevelMetadataOpenSea
		metadataAdditionalProps = make(map[string]jx.Raw)
		err                     error
	)
	if e.Metadata != nil {
		opensea = &e.Metadata.ContractLevelMetadataOpenSea
		metadataAdditionalProps, err = MakeAPIAdditionalProps(e.Metadata.AdditionalProps)
	}
	if err != nil {
		return nil, err
	}
	return &api.BookNFT{
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
	}, nil
}
