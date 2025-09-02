package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftEvmAddressStakingsGet(
	ctx context.Context,
	params api.BookNftEvmAddressStakingsGetParams,
) (*api.BookNftEvmAddressStakingsGetOK, error) {
	filterBookNFTIn := []string{string(params.EvmAddress)}
	var filterAccountIn *[]string
	if len(params.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(params.FilterAccountIn))
		for _, account := range params.FilterAccountIn {
			_filterAccountIn = append(_filterAccountIn, string(account))
		}
		filterAccountIn = &_filterAccountIn
	}

	stakingEvents, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		database.NewStakingFilter(&filterBookNFTIn, filterAccountIn),
	)
	if err != nil {
		return nil, err
	}

	apiStakingEvents := make([]api.Staking, 0, len(stakingEvents))
	for _, stakingEvent := range stakingEvents {
		if stakingEvent.BookNFT == string(params.EvmAddress) {
			apiStakingEvents = append(apiStakingEvents, model.MakeStaking(stakingEvent))
		}
	}

	return &api.BookNftEvmAddressStakingsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakingEvents,
	}, nil
}
