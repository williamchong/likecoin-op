package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
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
		MinterAddresses:     e.MinterAddresses,
		UpdaterAddresses:    e.UpdaterAddresses,
		TotalSupply:         MakeBigInt(e.TotalSupply),
		MaxSupply:           MakeUint64(uint64(e.MaxSupply)),
		Metadata:            MakeOptContractLevelMetadata(opensea, metadataAdditionalProps),
		BannerImage:         e.BannerImage,
		FeaturedImage:       e.FeaturedImage,
		DeployedBlockNumber: MakeUint64(uint64(e.DeployedBlockNumber)),
		MintedAt:            e.MintedAt,
		UpdatedAt:           e.UpdatedAt,
		Owner:               MakeOptAccount(e.Edges.Owner),
	}, nil
}

type NFTClassPagination struct {
	// Pagination.limit.
	PaginationLimit api.OptInt
	// Pagination.key.
	PaginationKey api.OptInt
	// Reverse.
	Reverse api.OptBool
}

func (p *NFTClassPagination) ToEntPagination() database.NFTClassPagination {
	limit := FromOpt(p.PaginationLimit)
	if limit != nil && *limit == 0 {
		limit = nil
	}

	key := FromOpt(p.PaginationKey)
	if key != nil && *key == 0 {
		key = nil
	}

	reverse := FromOpt(p.Reverse)

	return database.NFTClassPagination{
		Limit:   limit,
		Key:     key,
		Reverse: reverse,
	}
}
