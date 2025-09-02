package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftEvmAddressGet(
	ctx context.Context,
	params api.BookNftEvmAddressGetParams,
) (*api.BookNFT, error) {
	bookNFT, err := h.bookNFTRepository.QueryBookNFT(ctx, string(params.EvmAddress))
	if err != nil {
		return nil, err
	}

	apiBookNFT := model.MakeBookNFT(bookNFT)

	return &apiBookNFT, nil
}
