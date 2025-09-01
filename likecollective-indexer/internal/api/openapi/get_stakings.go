package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) StakingsGet(
	ctx context.Context,
	params api.StakingsGetParams,
) (*api.StakingsGetOK, error) {
	var filterBookNFTIn *[]string
	var filterAccountIn *[]string
	if len(params.FilterBookNftIn) > 0 {
		_filterBookNFTIn := make([]string, len(params.FilterBookNftIn))
		for _, bookNFT := range params.FilterBookNftIn {
			_filterBookNFTIn = append(_filterBookNFTIn, string(bookNFT))
		}
		filterBookNFTIn = &_filterBookNFTIn
	}
	if len(params.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(params.FilterAccountIn))
		for _, account := range params.FilterAccountIn {
			_filterAccountIn = append(_filterAccountIn, string(account))
		}
		filterAccountIn = &_filterAccountIn
	}

	stakings, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		database.NewStakingFilter(filterBookNFTIn, filterAccountIn),
	)

	if err != nil {
		return nil, err
	}

	apiStakings := make([]api.Staking, 0, len(stakings))
	for _, staking := range stakings {
		apiStakings = append(apiStakings, model.MakeStaking(staking))
	}

	return &api.StakingsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakings,
	}, nil
}
