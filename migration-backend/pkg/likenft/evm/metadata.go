package evm

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	evmmodel "github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
)

func ContractLevelMetadataFromCosmosClassAndISCN(
	c *cosmosmodel.Class,
	iscn *cosmosmodel.ISCN,
	royaltyConfig *cosmosmodel.RoyaltyConfig,
) *evmmodel.ContractLevelMetadata {
	iscnRecord := iscn.Records[0].Data

	attributes := make([]evmmodel.ContractLevelMetadataAttributes, 0)

	if iscnRecord.ContentMetadata.Author != nil {
		author := iscnRecord.ContentMetadata.Author.Name()
		attributes = append(attributes, evmmodel.ContractLevelMetadataAttributes{
			TraitType: "Author",
			Value:     author,
		})
	}

	if iscnRecord.ContentMetadata.Publisher != "" {
		attributes = append(attributes, evmmodel.ContractLevelMetadataAttributes{
			TraitType: "Publisher",
			Value:     iscnRecord.ContentMetadata.Publisher,
		})
	}

	if iscnRecord.ContentMetadata.DatePublished != nil && !iscnRecord.ContentMetadata.DatePublished.IsEmpty() {
		attributes = append(attributes, evmmodel.ContractLevelMetadataAttributes{
			TraitType:   "Publish Date",
			DisplayType: "date",
			Value:       iscnRecord.ContentMetadata.DatePublished.GetEpochSeconds(),
		})
	}

	return &evmmodel.ContractLevelMetadata{
		ContractLevelMetadataOpenSea: evmmodel.ContractLevelMetadataOpenSea{
			Name:         iscnRecord.ContentMetadata.Name,
			Symbol:       c.Symbol,
			Description:  iscnRecord.ContentMetadata.Description,
			Image:        c.Data.Metadata.Image,
			ExternalLink: iscnRecord.ContentMetadata.Url,
		},
		MetadataAdditional: evmmodel.MetadataAdditional{
			Message:                      c.Data.Metadata.Message,
			NftMetaCollectionId:          c.Data.Metadata.NFTMetaCollectionID,
			NftMetaCollectionName:        c.Data.Metadata.NFTMetaCollectionName,
			NftMetaCollectionDescription: c.Data.Metadata.GetNFTMetaCollectionDescription(),
		},
		MetadataISCN: evmmodel.MakeMetadataISCNFromCosmosISCN(iscn),
		Attributes:   attributes,
		LikeCoin: &evmmodel.ContractLevelMetadataLikeCoin{
			ISCNIdPrefix: c.Data.Parent.IscnIdPrefix,
			ClassId:      c.Id,
		},
		RoyaltyConfig: royaltyConfig,
	}
}

func ERC721OpenSeaMetadataFromCosmosNFTMetadata(m *cosmosmodel.NFTMetadata) *evmmodel.ERC721MetadataOpenSea {
	return &evmmodel.ERC721MetadataOpenSea{
		Image:       m.Image,
		ExternalUrl: m.ExternalUrl,
		Description: m.Description,
		Name:        m.Name,
		Attributes:  sortERC721MetadataAttributes(makeERC721MetadataAttribute("", m.Attributes)),
	}
}

func ERC721MetadataFromCosmosNFTAndClassAndISCNDataArbitrary(
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,
	n *cosmosmodel.NFT,
	c *cosmosmodel.Class,
	iscn *cosmosmodel.ISCN,
	cosmosMetadataOverride *cosmosmodel.NFTMetadata,
	evmClassId string,
) *evmmodel.ERC721Metadata {
	iscnAttributes := makeERC721MetadataAttributeFromISCN(iscn)

	var metadataOverride *evmmodel.ERC721MetadataOpenSea
	if cosmosMetadataOverride != nil {
		metadataOverride = ERC721OpenSeaMetadataFromCosmosNFTMetadata(cosmosMetadataOverride)
	}

	metadata := &evmmodel.ERC721Metadata{
		ERC721MetadataOpenSea: evmmodel.OverrideERC721MetadataOpenSea(
			evmmodel.ERC721MetadataOpenSea{
				Image:       c.Data.Metadata.Image,
				ExternalUrl: erc721ExternalURLBuilder.BuildArbitrary(evmClassId),
				Description: c.Description,
				Name:        c.Name,
				Attributes: sortERC721MetadataAttributes(
					slices.Concat(
						makeERC721MetadataAttribute("", n.Data.Metadata.Attributes),
						iscnAttributes,
					),
				),
				AnimationUrl: n.Data.Metadata.AnimationUrl,
			}, metadataOverride),
		LikeCoin: &evmmodel.ERC721MetadataLikeCoin{
			ISCNIdPrefix: c.Data.Parent.IscnIdPrefix,
			ClassId:      n.ClassId,
			NFTId:        n.Id,
		},
	}

	return metadata
}

func ERC721MetadataFromCosmosNFTAndClassAndISCNData(
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,
	n *cosmosmodel.NFT,
	c *cosmosmodel.Class,
	iscn *cosmosmodel.ISCN,
	cosmosMetadataOverride *cosmosmodel.NFTMetadata,
	evmClassId string,
	evmTokenId uint64,
) *evmmodel.ERC721Metadata {
	iscnRecord := iscn.Records[0].Data

	iscnAttributes := makeERC721MetadataAttributeFromISCN(iscn)

	var metadataOverride *evmmodel.ERC721MetadataOpenSea
	if cosmosMetadataOverride != nil {
		metadataOverride = ERC721OpenSeaMetadataFromCosmosNFTMetadata(cosmosMetadataOverride)
	}

	metadata := &evmmodel.ERC721Metadata{
		ERC721MetadataOpenSea: evmmodel.OverrideERC721MetadataOpenSea(
			evmmodel.ERC721MetadataOpenSea{
				Image:       c.Data.Metadata.Image,
				ExternalUrl: erc721ExternalURLBuilder.BuildSerial(evmClassId, evmTokenId),
				Description: fmt.Sprintf("Copy #%s of %s", strconv.FormatUint(evmTokenId, 10), iscnRecord.ContentMetadata.Name),
				Name:        fmt.Sprintf("%s #%s", iscnRecord.ContentMetadata.Name, strconv.FormatUint(evmTokenId, 10)),
				Attributes: sortERC721MetadataAttributes(
					slices.Concat(
						makeERC721MetadataAttribute("", n.Data.Metadata.Attributes),
						iscnAttributes,
					),
				),
				AnimationUrl: n.Data.Metadata.AnimationUrl,
			}, metadataOverride),
		LikeCoin: &evmmodel.ERC721MetadataLikeCoin{
			ISCNIdPrefix: c.Data.Parent.IscnIdPrefix,
			ClassId:      n.ClassId,
			NFTId:        n.Id,
		},
	}

	return metadata
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

func makeERC721MetadataAttributeFromISCN(iscn *cosmosmodel.ISCN) []evmmodel.ERC721MetadataAttribute {
	iscnRecord := iscn.Records[0].Data

	attributes := make([]evmmodel.ERC721MetadataAttribute, 0)

	if iscnRecord.ContentMetadata.Author != nil {
		author := iscnRecord.ContentMetadata.Author.Name()
		if author != "" {
			attributes = append(attributes, evmmodel.ERC721MetadataAttribute{
				TraitType: "Author",
				Value:     author,
			})
		}
	}

	if iscnRecord.ContentMetadata.Publisher != "" {
		attributes = append(attributes, evmmodel.ERC721MetadataAttribute{
			TraitType: "Publisher",
			Value:     iscnRecord.ContentMetadata.Publisher,
		})
	}

	if iscnRecord.ContentMetadata.DatePublished != nil && !iscnRecord.ContentMetadata.DatePublished.IsEmpty() {
		attributes = append(attributes, evmmodel.ERC721MetadataAttribute{
			TraitType:   "Publish Date",
			DisplayType: &evmmodel.ERC721MetadataAttributeDisplayTypeDate,
			Value:       iscnRecord.ContentMetadata.DatePublished.GetEpochSeconds(),
		})
	}
	return attributes
}

func sortERC721MetadataAttributes(attributes []evmmodel.ERC721MetadataAttribute) []evmmodel.ERC721MetadataAttribute {
	sort.Slice(attributes, func(i, j int) bool {
		return strings.Compare(attributes[i].TraitType, attributes[j].TraitType) == -1
	})
	return attributes
}
