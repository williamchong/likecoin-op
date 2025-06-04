package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) GetBookNFTLatestEventBlockNumber(ctx context.Context, params api.GetBookNFTLatestEventBlockNumberParams) (*api.LatestEventBlockNumber, error) {
	nftclass, err := h.nftClassRepository.QueryNFTClassByAddress(
		ctx, params.ID,
	)

	if err != nil {
		return nil, err
	}

	return &api.LatestEventBlockNumber{
		LatestEventBlockNumber: model.MakeUint64(uint64(nftclass.LatestEventBlockNumber)),
	}, nil
}
