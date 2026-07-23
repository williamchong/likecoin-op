package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) TokenBookNFTsByAccount(ctx context.Context, params api.TokenBookNFTsByAccountParams) (*api.TokenBookNFTsByAccountOK, error) {
	pagination := model.NFTClassPagination{
		PaginationLimit: params.PaginationLimit,
		PaginationKey:   params.PaginationKey,
		Reverse:         params.Reverse,
	}

	metadataEQ := database.ContractLevelMetadataFilterEquatable(
		params.ContractLevelMetadataEq.Or(api.ContractLevelMetadataEQ{}),
	)

	metadataNEQ := database.ContractLevelMetadataFilterEquatable(
		params.ContractLevelMetadataNeq.Or(api.ContractLevelMetadataNEQ{}),
	)

	nftClassWithTokenIDs, count, nextKey, err := h.nftClassRepository.QueryNFTClassesByAccountTokensWithNFTID(
		ctx,
		params.EvmAddress,
		metadataEQ,
		metadataNEQ,
		pagination.ToEntPagination(),
	)

	if err != nil {
		return nil, err
	}

	apiBookNFTs := make([]api.BookNFT, len(nftClassWithTokenIDs))

	for i, nftClassWithTokenID := range nftClassWithTokenIDs {
		apiNFTClass, err := model.MakeNFTClassWithToken(nftClassWithTokenID)
		if err != nil {
			return nil, err
		}
		apiBookNFTs[i] = *apiNFTClass
	}

	return &api.TokenBookNFTsByAccountOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiBookNFTs,
	}, nil
}
