package evm

import (
	"fmt"
	"sort"
	"strings"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	evmmodel "github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"
)

func ContractLevelMetadataFromCosmosClass(c *cosmosmodel.Class) *evmmodel.ContractLevelMetadata {
	return &evmmodel.ContractLevelMetadata{
		ContractLevelMetadataOpenSea: evmmodel.ContractLevelMetadataOpenSea{
			Name:         c.Name,
			Symbol:       c.Symbol,
			Description:  c.Description,
			Image:        c.Data.Metadata.Image,
			ExternalLink: c.Data.Metadata.ExternalURL,
		},
	}
}

func ContractLevelMetadataFromCosmosClassListItem(c *cosmosmodel.ClassListItem) *evmmodel.ContractLevelMetadata {
	return &evmmodel.ContractLevelMetadata{
		ContractLevelMetadataOpenSea: evmmodel.ContractLevelMetadataOpenSea{
			Name:         c.Name,
			Symbol:       c.Symbol,
			Description:  c.Description,
			Image:        c.Metadata.Image,
			ExternalLink: c.Metadata.ExternalURL,
		},
	}
}

func ERC721MetadataFromCosmosNFT(c *cosmosmodel.NFT) *evmmodel.ERC721Metadata {
	return &evmmodel.ERC721Metadata{
		ERC721MetadataOpenSea: evmmodel.ERC721MetadataOpenSea{
			Image:       c.Data.Metadata.Image,
			ExternalUrl: c.Data.Metadata.ExternalUrl,
			Description: c.Data.Metadata.Description,
			Name:        c.Data.Metadata.Name,
			Attributes:  sortERC721MetadataAttributes(makeERC721MetadataAttribute("", c.Data.Metadata.Attributes)),
		},
	}
}

func makeERC721MetadataAttribute(prefix string, m map[string]interface{}) []evmmodel.ERC721MetadataAttribute {
	attrs := make([]evmmodel.ERC721MetadataAttribute, 0)

	for k, v := range m {
		traitType := k
		if prefix != "" {
			traitType = fmt.Sprintf("%s_%s", prefix, k)
		}

		if strVal, ok := v.(string); ok {
			attrs = append(attrs, evmmodel.ERC721MetadataAttribute{
				TraitType: traitType,
				Value:     strVal,
			})
			continue
		}

		if numVal, ok := v.(float64); ok {
			attrs = append(attrs, evmmodel.ERC721MetadataAttribute{
				DisplayType: &evmmodel.ERC721MetadataAttributeDisplayTypeNumber,
				TraitType:   traitType,
				Value:       numVal,
			})
			continue
		}

		// Flattening nested dict
		if dictVal, ok := v.(map[string]interface{}); ok {
			attrs = append(attrs, makeERC721MetadataAttribute(traitType, dictVal)...)
			continue
		}
	}

	return attrs
}

func sortERC721MetadataAttributes(attributes []evmmodel.ERC721MetadataAttribute) []evmmodel.ERC721MetadataAttribute {
	sort.Slice(attributes, func(i, j int) bool {
		return strings.Compare(attributes[i].TraitType, attributes[j].TraitType) == -1
	})
	return attributes
}
