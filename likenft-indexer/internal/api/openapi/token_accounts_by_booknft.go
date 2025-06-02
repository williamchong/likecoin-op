package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) TokenAccountsByBookNFT(ctx context.Context, params api.TokenAccountsByBookNFTParams) (*api.TokenAccountsByBookNFTOK, error) {
	pagination := model.AccountPagination{
		PaginationLimit: params.PaginationLimit,
		PaginationKey:   params.PaginationKey,
		Reverse:         params.Reverse,
	}

	accounts, count, nextKey, err := h.accountRepository.GetTokenAccountsByBookNFT(
		ctx,
		params.ID,
		pagination.ToEntPagination(),
	)

	if err != nil {
		return nil, err
	}

	apiAccounts := make([]api.Account, len(accounts))

	for i, account := range accounts {
		apiAccounts[i] = model.MakeAccount(account)
	}

	return &api.TokenAccountsByBookNFTOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiAccounts,
	}, nil
}
