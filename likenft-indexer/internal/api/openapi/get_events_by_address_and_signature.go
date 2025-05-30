package openapi

import (
	"context"
	"math"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"

	"github.com/ethereum/go-ethereum/crypto"
)

func (h *OpenAPIHandler) EventsByAddressAndSignature(ctx context.Context, params api.EventsByAddressAndSignatureParams) (*api.EventsByAddressAndSignatureOK, error) {
	signatureHash := crypto.Keccak256Hash([]byte(params.Signature))

	ps := model.OpenAPIEventParameters{
		Address:                 &params.Address,
		Signature:               &params.Signature,
		Limit:                   params.Limit,
		Page:                    params.Page,
		SortBy:                  params.SortBy,
		SortOrder:               params.SortOrder,
		FilterBlockTimestamp:    params.FilterBlockTimestamp,
		FilterBlockTimestampGte: params.FilterBlockTimestampGte,
		FilterBlockTimestampGt:  params.FilterBlockTimestampGt,
		FilterBlockTimestampLte: params.FilterBlockTimestampLte,
		FilterBlockTimestampLt:  params.FilterBlockTimestampLt,
		FilterTopic1:            params.FilterTopic1,
		FilterTopic2:            params.FilterTopic2,
		FilterTopic3:            params.FilterTopic3,
		FilterTopic0:            params.FilterTopic0,
	}
	entFilter, err := ps.ToEntFilter()
	if err != nil {
		return nil, err
	}

	events, count, err := h.evmEventRepository.GetEvmEvents(ctx, entFilter)
	if err != nil {
		return nil, err
	}

	apiEvents := make([]api.Event, len(events))

	for i, n := range events {
		apiEvent, err := model.MakeEvent(n)
		if err != nil {
			return nil, err
		}
		apiEvents[i] = apiEvent
	}

	return &api.EventsByAddressAndSignatureOK{
		Meta: api.EventQueryMetadata{
			ChainIds:      []int{},
			Address:       params.Address,
			Signature:     signatureHash.Hex(),
			Page:          params.Page.Value,
			LimitPerChain: params.Limit.Value,
			TotalItems:    count,
			TotalPages:    int(math.Ceil(float64(count) / float64(params.Limit.Value))),
		},
		Data: apiEvents,
	}, nil
}
