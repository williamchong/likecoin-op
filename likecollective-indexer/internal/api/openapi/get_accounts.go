package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountsGet(
	ctx context.Context,
	params api.AccountsGetParams,
) (*api.AccountsGetOK, error) {

	accounts, count, nextKey, err := h.accountRepository.QueryAccounts(ctx)
	if err != nil {
		return nil, err
	}

	apiAccounts := make([]api.Account, 0, len(accounts))
	for _, account := range accounts {
		apiAccounts = append(apiAccounts, model.MakeAccount(account))
	}

	return &api.AccountsGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiAccounts,
	}, nil
}
