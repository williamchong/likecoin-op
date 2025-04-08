package openapi

import (
	"context"

	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) AccountByBookNFT(ctx context.Context, params api.AccountByBookNFTParams) (*api.AccountByBookNFTOK, error) {
	bookNFT, err := h.db.NFTClass.Query().Where(nftclass.AddressEqualFold(params.ID)).Only(ctx)

	if err != nil {
		return nil, err
	}

	account, err := bookNFT.QueryOwner().Only(ctx)

	if err != nil {
		return nil, err
	}

	return &api.AccountByBookNFTOK{
		Account: model.MakeAccount(account),
	}, nil
}
