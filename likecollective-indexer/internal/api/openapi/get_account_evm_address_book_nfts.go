package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressBookNftsGet(
	ctx context.Context,
	params api.AccountEvmAddressBookNftsGetParams,
) (*api.AccountEvmAddressBookNftsGetOK, error) {
	var filterBookNFTIn *[]string
	filterAccountIn := []string{string(params.EvmAddress)}
	if len(params.FilterBookNftIn) > 0 {
		_filterBookNFTIn := make([]string, len(params.FilterBookNftIn))
		for _, bookNFT := range params.FilterBookNftIn {
			_filterBookNFTIn = append(_filterBookNFTIn, string(bookNFT))
		}
		filterBookNFTIn = &_filterBookNFTIn
	}

	stakings, count, nextKey, err := h.stakingRepository.QueryStakings(
		ctx,
		database.NewStakingFilter(filterBookNFTIn, &filterAccountIn),
	)
	if err != nil {
		return nil, err
	}

	apiStakings := make([]api.Staking, 0, len(stakings))

	for _, staking := range stakings {
		apiStakings = append(apiStakings, model.MakeStaking(staking))
	}

	return &api.AccountEvmAddressBookNftsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakings,
	}, nil
}
