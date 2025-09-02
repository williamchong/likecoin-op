package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressGet(
	ctx context.Context,
	params api.AccountEvmAddressGetParams,
) (*api.Account, error) {
	account, err := h.accountRepository.QueryAccount(ctx, string(params.EvmAddress))
	if err != nil {
		return nil, err
	}

	apiAccount := model.MakeAccount(account)

	return &apiAccount, nil
}
