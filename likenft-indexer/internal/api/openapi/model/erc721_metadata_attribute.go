package model

import (
	"fmt"

	"likenft-indexer/internal/evm/model"
	"likenft-indexer/openapi/api"
)

func MakeOptErc721MetadataAttributeDisplayType(d *model.ERC721MetadataAttributeDisplayType) api.OptErc721MetadataAttributeDisplayType {
	if d == nil {
		return api.OptErc721MetadataAttributeDisplayType{
			Value: api.Erc721MetadataAttributeDisplayTypeNumber,
			Set:   false,
		}
	}
	return api.NewOptErc721MetadataAttributeDisplayType(api.Erc721MetadataAttributeDisplayType(*d))
}

func MakeErc721MetadataAttribute(a model.ERC721MetadataAttribute) api.Erc721MetadataAttribute {
	return api.Erc721MetadataAttribute{
		DisplayType: MakeOptErc721MetadataAttributeDisplayType(a.DisplayType),
		TraitType:   a.TraitType,
		Value:       fmt.Sprintf("%v", a.Value),
	}
}
