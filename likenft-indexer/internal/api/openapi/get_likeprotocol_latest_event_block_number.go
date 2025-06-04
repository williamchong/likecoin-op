package openapi

import (
	"context"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) GetLikeProtocolLatestEventBlockNumber(ctx context.Context) (*api.LatestEventBlockNumber, error) {
	likeprotocol, err := h.likeProtocolRepository.GetLikeProtocol(ctx, h.likeProtocolAddress)

	if err != nil {
		return nil, err
	}

	return &api.LatestEventBlockNumber{
		LatestEventBlockNumber: model.MakeUint64(uint64(likeprotocol.LatestEventBlockNumber)),
	}, nil
}
