package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftsGet(
	ctx context.Context,
	params api.BookNftsGetParams,
) (*api.BookNftsGetOK, error) {
	filterParams := model.NFTClassFilterParams{
		TimeFrame:          params.TimeFrame,
		TimeFrameSortBy:    params.TimeFrameSortBy,
		TimeFrameSortOrder: params.TimeFrameSortOrder,
		FilterBookNftIn:    params.FilterBookNftIn,
		FilterAccountIn:    params.FilterAccountIn,
	}

	pagination := model.NFTClassPagination{
		PaginationKey:   params.PaginationKey,
		PaginationLimit: params.PaginationLimit,
		Reverse:         params.Reverse,
	}

	nftClasses, count, nextKey, err := h.nftClassRepository.QueryNFTClasses(
		ctx, filterParams.ToEntFilter(), pagination.ToEntPagination(),
	)
	if err != nil {
		return nil, err
	}

	apiBookNFTs := make([]api.BookNFT, 0, len(nftClasses))
	for _, nftClass := range nftClasses {
		apiBookNFTs = append(apiBookNFTs, model.MakeBookNFT(nftClass))
	}

	return &api.BookNftsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiBookNFTs,
	}, nil
}
