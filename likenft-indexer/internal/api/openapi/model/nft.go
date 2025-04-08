package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/openapi/api"
)

func MakeNFT(e *ent.NFT) api.NFT {
	attributes := make([]api.Erc721MetadataAttribute, len(e.Attributes))

	for i, n := range e.Attributes {
		attributes[i] = MakeErc721MetadataAttribute(n)
	}

	return api.NFT{
		ID:              e.ID,
		ContractAddress: e.ContractAddress,
		TokenID:         MakeUint64(uint64(e.TokenID)),
		TokenURI:        MakeOptString(e.TokenURI),
		Image:           MakeOptString(e.Image),
		ImageData:       MakeOptString(e.ImageData),
		ExternalURL:     MakeOptString(e.ExternalURL),
		Description:     MakeOptString(e.Description),
		Name:            MakeOptString(e.Name),
		Attributes:      attributes,
		BackgroundColor: MakeOptString(e.BackgroundColor),
		AnimationURL:    MakeOptString(e.AnimationURL),
		YoutubeURL:      MakeOptString(e.YoutubeURL),
		OwnerAddress:    e.OwnerAddress,
		MintedAt:        e.MintedAt,
		UpdatedAt:       e.MintedAt,
	}
}
