package openapi

import (
	"context"

	"likenft-indexer/ent/nft"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) Token(ctx context.Context, params api.TokenParams) (*api.NFT, error) {
	tokenId, err := model.MakeGoUint64(params.TokenID)

	if err != nil {
		return nil, err
	}

	nft, err := h.db.NFT.Query().Where(nft.ContractAddressEqualFold(params.BooknftID), nft.TokenIDEQ(typeutil.Uint64(tokenId))).Only(ctx)

	if err != nil {
		return nil, err
	}

	apiNft := model.MakeNFT(nft)
	return &apiNft, nil
}
