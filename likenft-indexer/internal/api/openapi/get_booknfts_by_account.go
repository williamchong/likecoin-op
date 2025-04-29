package openapi

import (
	"context"

	"likenft-indexer/ent/account"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) BookNFTsByAccount(
	ctx context.Context,
	params api.BookNFTsByAccountParams,
) (*api.BookNFTsByAccountOK, error) {
	account, err := h.db.Account.Query().Where(account.EvmAddressEqualFold(params.EvmAddress)).Only(ctx)

	if err != nil {
		return nil, err
	}

	bookNFTsQ := account.QueryNftClasses()

	count, err := bookNFTsQ.Count(ctx)

	if err != nil {
		return nil, err
	}

	paginatedQ := h.handleNFTClassPagination(
		bookNFTsQ,
		params.PaginationLimit,
		params.PaginationKey,
		params.Reverse,
	)

	bookNFTs, err := paginatedQ.All(ctx)

	if err != nil {
		return nil, err
	}

	nextKey := 0

	if len(bookNFTs) > 0 {
		nextKey = bookNFTs[len(bookNFTs)-1].ID
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
