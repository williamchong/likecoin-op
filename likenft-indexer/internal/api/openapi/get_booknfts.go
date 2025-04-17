package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) BookNFTs(ctx context.Context, params api.BookNFTsParams) (*api.BookNFTsOK, error) {
	bookNFTsQ := h.db.NFTClass.Query()

	count, err := bookNFTsQ.Count(ctx)

	if err != nil {
		return nil, err
	}

	paginatedBookNFTsQ := h.handleNFTClassPagination(
		bookNFTsQ,
		params.PaginationLimit,
		params.PaginationKey,
		params.Reverse,
	)

	dbBookNFTs, err := paginatedBookNFTsQ.All(ctx)

	if err != nil {
		return nil, err
	}

	nextKey := 0

	if len(dbBookNFTs) > 0 {
		nextKey = dbBookNFTs[len(dbBookNFTs)-1].ID
	}

	apiBookNFTs := make([]api.BookNFT, len(dbBookNFTs))

	for i, n := range dbBookNFTs {
		metadataAdditionalProps, err := model.MakeAPIAdditionalProps(n.Metadata.AdditionalProps)

		if err != nil {
			return nil, err
		}

		apiBookNFTs[i] = model.MakeNFTClass(n, metadataAdditionalProps)
	}

	return &api.BookNFTsOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiBookNFTs,
	}, nil
}
