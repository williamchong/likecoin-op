package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) TokensByAccount(ctx context.Context, params api.TokensByAccountParams) (*api.TokensByAccountOK, error) {
	ps := model.NFTPagination{
		PaginationLimit: params.PaginationLimit,
		PaginationKey:   params.PaginationKey,
		Reverse:         params.Reverse,
	}

	contractMetadataEQ := database.ContractLevelMetadataFilterEquatable(
		params.ContractLevelMetadataEq.Or(api.ContractLevelMetadataEQ{}),
	)

	nfts, count, nextKey, err := h.nftRepository.QueryNFTsByEvmAddress(
		ctx,
		params.EvmAddress,
		contractMetadataEQ,
		ps.ToEntPagination(),
	)

	if err != nil {
		return nil, err
	}

	apiNFTClasses := make([]api.NFT, len(nfts))

	for i, n := range nfts {
		apiNFTClasses[i] = model.MakeNFT(n)
	}

	return &api.TokensByAccountOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiNFTClasses,
	}, nil
}
