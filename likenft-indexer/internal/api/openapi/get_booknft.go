package openapi

import (
	"context"

	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) BookNFT(ctx context.Context, params api.BookNFTParams) (*api.BookNFT, error) {
	nftclass, err := h.db.NFTClass.Query().Where(nftclass.AddressEqualFold(params.ID)).Only(ctx)

	if err != nil {
		return nil, err
	}

	metadataAdditionalProps, err := model.MakeAPIAdditionalProps(nftclass.Metadata.AdditionalProps)

	if err != nil {
		return nil, err
	}

	apiNftClass := model.MakeNFTClass(nftclass, metadataAdditionalProps)
	return &apiNftClass, nil
}
