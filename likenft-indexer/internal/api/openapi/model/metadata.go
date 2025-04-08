package model

import (
	"likenft-indexer/internal/evm/model"
	"likenft-indexer/openapi/api"
)

func MakeOptContractLevelMetadata(e *model.ContractLevelMetadata) api.OptContractLevelMetadata {
	if e == nil {
		return api.OptContractLevelMetadata{
			Value: api.ContractLevelMetadata{},
			Set:   false,
		}
	}
	return api.NewOptContractLevelMetadata(api.ContractLevelMetadata{
		Name:          MakeOptString(&e.Name),
		Symbol:        MakeOptString(&e.Symbol),
		Description:   MakeOptString(&e.Description),
		Image:         MakeOptString(&e.Image),
		BannerImage:   MakeOptString(&e.BannerImage),
		FeaturedImage: MakeOptString(&e.FeaturedImage),
		ExternalLink:  MakeOptString(&e.ExternalLink),
		Collaborators: e.Collaborators,
	})
}
