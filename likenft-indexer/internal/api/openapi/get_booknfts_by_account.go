package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) BookNFTsByAccount(
	ctx context.Context,
	params api.BookNFTsByAccountParams,
) (*api.BookNFTsByAccountOK, error) {
	ps := model.NFTClassPagination{
		PaginationLimit: params.PaginationLimit,
		PaginationKey:   params.PaginationKey,
		Reverse:         params.Reverse,
	}

	metadataEQ := database.ContractLevelMetadataFilterEquatable(
		params.ContractLevelMetadataEq.Or(api.ContractLevelMetadataEQ{}),
	)

	bookNFTs, count, nextKey, err := h.nftClassRepository.QueryNFTClassesByEvmAddress(
		ctx,
		params.EvmAddress,
		metadataEQ,
		ps.ToEntPagination(),
	)

	if err != nil {
		return nil, err
	}

	apiBookNFTs := make([]api.BookNFT, len(bookNFTs))

	for i, n := range bookNFTs {
		apin, err := model.MakeNFTClass(n)

		if err != nil {
			return nil, err
		}

		apiBookNFTs[i] = *apin
	}

	return &api.BookNFTsByAccountOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiBookNFTs,
	}, nil
}
