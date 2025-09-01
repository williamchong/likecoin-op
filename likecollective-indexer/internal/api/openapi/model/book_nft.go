package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/openapi/api"
)

func MakeBookNFT(bookNFT *ent.BookNFT) api.BookNFT {
	return api.BookNFT{
		EvmAddress:      api.EvmAddress(bookNFT.EvmAddress),
		StakedAmount:    api.Uint256(bookNFT.StakedAmount),
		LastStakedAt:    api.NewNilDateTime(bookNFT.LastStakedAt),
		NumberOfStakers: bookNFT.NumberOfStakers,
	}
}
