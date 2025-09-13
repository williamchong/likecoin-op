package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) StakingsGet(
	ctx context.Context,
	params api.StakingsGetParams,
) (*api.StakingsGetOK, error) {
	filterParams := model.StakingFilterParams{
		FilterNFTClassIn: params.FilterBookNftIn,
		FilterAccountIn:  params.FilterAccountIn,
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
