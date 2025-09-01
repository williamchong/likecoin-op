package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftsGet(
	ctx context.Context,
	params api.BookNftsGetParams,
) (*api.BookNftsGetOK, error) {
	var timeFrameSortBy *database.BookNFTsRequestTimeFrameSortBy
	var timeFrameSortOrder *database.BookNFTsRequestTimeFrameSortOrder
	if params.TimeFrameSortBy.IsSet() {
		_timeFrameSortBy := database.BookNFTsRequestTimeFrameSortBy(params.TimeFrameSortBy.Value)
		timeFrameSortBy = &_timeFrameSortBy
	}
	if params.TimeFrameSortOrder.IsSet() {
		_timeFrameSortOrder := database.BookNFTsRequestTimeFrameSortOrder(params.TimeFrameSortOrder.Value)
		timeFrameSortOrder = &_timeFrameSortOrder
	}
	bookNFTs, count, nextKey, err := h.bookNFTRepository.QueryBookNFTs(
		ctx, database.NewQueryBookNFTsFilter(
			timeFrameSortBy,
			timeFrameSortOrder,
		),
	)
	if err != nil {
		return nil, err
	}

	apiBookNFTs := make([]api.BookNFT, 0, len(bookNFTs))
	for _, bookNFT := range bookNFTs {
		apiBookNFTs = append(apiBookNFTs, model.MakeBookNFT(bookNFT))
	}

	return &api.BookNftsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiBookNFTs,
	}, nil
}
