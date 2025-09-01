package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressStakingsGet(
	ctx context.Context,
	params api.AccountEvmAddressStakingsGetParams,
) (*api.AccountEvmAddressStakingsGetOK, error) {
	var filterBookNFTIn *[]string
	if len(params.FilterBookNftIn) > 0 {
		_filterBookNFTIn := make([]string, len(params.FilterBookNftIn))
		for _, bookNFT := range params.FilterBookNftIn {
			_filterBookNFTIn = append(_filterBookNFTIn, string(bookNFT))
		}
		filterBookNFTIn = &_filterBookNFTIn
	}

	stakings, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		database.NewStakingFilter(filterBookNFTIn, nil),
	)
	if err != nil {
		return nil, err
	}

	apiStakings := make([]api.Staking, 0, len(stakings))
	for _, stakingEvent := range stakings {
		if stakingEvent.Account == string(params.EvmAddress) {
			apiStakings = append(apiStakings, model.MakeStaking(stakingEvent))
		}
	}

	return &api.AccountEvmAddressStakingsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakings,
	}, nil
}
