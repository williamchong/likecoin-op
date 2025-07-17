package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"
)

func MakeNFT(e *ent.NFT) api.NFT {
	attributes := make([]api.Erc721MetadataAttribute, len(e.Attributes))

	for i, n := range e.Attributes {
		attributes[i] = MakeErc721MetadataAttribute(n)
	}

	return api.NFT{
		ID:              e.ID,
		ContractAddress: e.ContractAddress,
		TokenID:         MakeUint64(uint64(e.TokenID)),
		TokenURI:        MakeOptString(e.TokenURI),
		Image:           MakeOptString(e.Image),
		ImageData:       MakeOptString(e.ImageData),
		ExternalURL:     MakeOptString(e.ExternalURL),
		Description:     MakeOptString(e.Description),
		Name:            MakeOptString(e.Name),
		Attributes:      attributes,
		BackgroundColor: MakeOptString(e.BackgroundColor),
		AnimationURL:    MakeOptString(e.AnimationURL),
		YoutubeURL:      MakeOptString(e.YoutubeURL),
		OwnerAddress:    e.OwnerAddress,
		MintedAt:        e.MintedAt,
		UpdatedAt:       e.MintedAt,
	}
}

type NFTPagination struct {
	// Pagination.limit.
	PaginationLimit api.OptInt
	// Pagination.key.
	PaginationKey api.OptInt
	// Reverse.
	Reverse api.OptBool
}

func (p *NFTPagination) ToEntPagination() database.NFTPagination {
	limit := FromOpt(p.PaginationLimit)
	if limit != nil && *limit == 0 {
		limit = nil
	}

	key := FromOpt(p.PaginationKey)
	if key != nil && *key == 0 {
		key = nil
	}

	reverse := FromOpt(p.Reverse)

	return database.NFTPagination{
		Limit:   limit,
		Key:     key,
		Reverse: reverse,
	}
}
