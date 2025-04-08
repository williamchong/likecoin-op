package openapi

import (
	"context"

	"likenft-indexer/ent/account"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) TokensByAccount(ctx context.Context, params api.TokensByAccountParams) (*api.TokensByAccountOK, error) {
	account, err := h.db.Account.Query().Where(account.EvmAddressEqualFold(params.EvmAddress)).Only(ctx)

	if err != nil {
		return nil, err
	}

	tokensQ := account.QueryNfts()

	count, err := tokensQ.Count(ctx)

	if err != nil {
		return nil, err
	}

	paginatedQ := h.handleNFTPagination(tokensQ, params.PaginationLimit, params.PaginationKey, params.Reverse)

	nfts, err := paginatedQ.All(ctx)

	if err != nil {
		return nil, err
	}

	nextKey := 0

	if len(nfts) > 0 {
		nextKey = nfts[len(nfts)-1].ID
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
