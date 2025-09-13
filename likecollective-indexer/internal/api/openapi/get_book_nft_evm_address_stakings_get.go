package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftEvmAddressStakingsGet(
	ctx context.Context,
	params api.BookNftEvmAddressStakingsGetParams,
) (*api.BookNftEvmAddressStakingsGetOK, error) {
	filterParams := model.StakingFilterParams{
		FilterNFTClassIn: []api.EvmAddress{params.EvmAddress},
		FilterAccountIn:  params.FilterAccountIn,
	}

	pagination := model.StakingPagination{
		PaginationKey:   params.PaginationKey,
		PaginationLimit: params.PaginationLimit,
		Reverse:         params.Reverse,
	}

	stakingEvents, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		filterParams.ToEntFilter(),
		pagination.ToEntPagination(),
	)
	if err != nil {
		return nil, err
	}

	apiStakingEvents := make([]api.Staking, 0, len(stakingEvents))
	for _, stakingEvent := range stakingEvents {
		apiStakingEvents = append(apiStakingEvents, model.MakeStaking(stakingEvent))
	}

	return &api.BookNftEvmAddressStakingsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakingEvents,
	}, nil
}
