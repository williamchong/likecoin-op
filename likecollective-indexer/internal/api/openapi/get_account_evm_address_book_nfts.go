package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressBookNftsGet(
	ctx context.Context,
	params api.AccountEvmAddressBookNftsGetParams,
) (*api.AccountEvmAddressBookNftsGetOK, error) {
	filterParams := model.StakingFilterParams{
		FilterNFTClassIn: params.FilterBookNftIn,
		FilterAccountIn:  []api.EvmAddress{params.EvmAddress},
	}

	pagination := model.StakingPagination{
		PaginationKey:   params.PaginationKey,
		PaginationLimit: params.PaginationLimit,
		Reverse:         params.Reverse,
	}

	dbFilter := filterParams.ToEntFilter()

	stakings, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		dbFilter,
		pagination.ToEntPagination(),
	)
	if err != nil {
		return nil, err
	}

	apiStakings := make([]api.AccountBookNFT, 0, len(stakings))

	for _, staking := range stakings {
		apiStaking, err := model.MakeAccountBookNFT(staking)
		if err != nil {
			return nil, err
		}
		apiStakings = append(apiStakings, apiStaking)
	}

	return &api.AccountEvmAddressBookNftsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakings,
	}, nil
}
