package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftsTimeFrameDeltaGet(
	ctx context.Context,
	params api.BookNftsTimeFrameDeltaGetParams,
) (*api.BookNftsTimeFrameDeltaGetOK, error) {

	pagination := model.BookNftsTimeFrameDeltaParams{
		SortBy:          params.SortBy,
		SortOrder:       params.SortOrder,
		PaginationLimit: params.PaginationLimit,
		PaginationPage:  params.PaginationPage,
	}

	bookNFTDeltaTimeBuckets, count, err := h.bookNFTDeltaTimeBucketRepository.QueryBookNFTDeltaTimeBuckets(ctx, string(params.TimeFrame), pagination.ToEntPagination())
	if err != nil {
		return nil, err
	}

	apiBookNFTDeltaTimeBuckets := make([]api.BookNFTStakeDelta, len(bookNFTDeltaTimeBuckets))
	for i, bookNFTDeltaTimeBucket := range bookNFTDeltaTimeBuckets {
		apiBookNFTDeltaTimeBuckets[i] = model.MakeBookNFTDeltaTimeBucket(bookNFTDeltaTimeBucket)
	}

	return &api.BookNftsTimeFrameDeltaGetOK{
		Pagination: api.BookNftsTimeFrameDeltaGetOKPagination{
			Count: count,
		},
		Data: apiBookNFTDeltaTimeBuckets,
	}, nil

}
