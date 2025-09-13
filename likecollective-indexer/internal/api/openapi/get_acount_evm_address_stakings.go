package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressStakingsGet(
	ctx context.Context,
	params api.AccountEvmAddressStakingsGetParams,
) (*api.AccountEvmAddressStakingsGetOK, error) {
	filterParams := model.StakingFilterParams{
		FilterNFTClassIn: params.FilterBookNftIn,
		FilterAccountIn:  []api.EvmAddress{params.EvmAddress},
	}

	pagination := model.StakingPagination{
		PaginationKey:   params.PaginationKey,
		PaginationLimit: params.PaginationLimit,
		Reverse:         params.Reverse,
	}

	stakings, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		filterParams.ToEntFilter(),
		pagination.ToEntPagination(),
	)
	if err != nil {
		return nil, err
	}

	apiStakings := make([]api.Staking, 0, len(stakings))
	for _, stakingEvent := range stakings {
		apiStakings = append(apiStakings, model.MakeStaking(stakingEvent))
	}

	return &api.AccountEvmAddressStakingsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakings,
	}, nil
}
