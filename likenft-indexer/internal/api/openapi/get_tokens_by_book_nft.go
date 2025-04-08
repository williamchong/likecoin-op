package openapi

import (
	"context"

	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) TokensByBookNFT(
	ctx context.Context,
	params api.TokensByBookNFTParams,
) (*api.TokensByBookNFTOK, error) {
	bookNFT, err := h.db.NFTClass.Query().Where(nftclass.AddressEqualFold(params.ID)).Only(ctx)

	if err != nil {
		return nil, err
	}

	nftsQ := bookNFT.QueryNfts()

	count, err := nftsQ.Count(ctx)

	if err != nil {
		return nil, err
	}

	paginatedQ := h.handleNFTPagination(
		nftsQ,
		params.PaginationLimit,
		params.PaginationKey,
		params.Reverse,
	)

	nftClasses, err := paginatedQ.All(ctx)

	if err != nil {
		return nil, err
	}

	nextKey := 0

	if len(nftClasses) > 0 {
		nextKey = nftClasses[len(nftClasses)-1].ID
	}

	apiNFTClasses := make([]api.NFT, len(nftClasses))

	for i, n := range nftClasses {
		apiNFTClasses[i] = model.MakeNFT(n)
	}

	return &api.TokensByBookNFTOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiNFTClasses,
	}, nil
}
