package evm

import (
	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	evmmodel "github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"
)

func ContractLevelMetadataFromCosmosClass(c *cosmosmodel.Class) *evmmodel.ContractLevelMetadata {
	return &evmmodel.ContractLevelMetadata{
		Name:         c.Name,
		Symbol:       c.Symbol,
		Description:  c.Description,
		Image:        c.Data.Metadata.Image,
		ExternalLink: c.Data.Metadata.ExternalURL,
	}
}

func ContractLevelMetadataFromCosmosClassListItem(c *cosmosmodel.ClassListItem) *evmmodel.ContractLevelMetadata {
	return &evmmodel.ContractLevelMetadata{
		Name:         c.Name,
		Symbol:       c.Symbol,
		Description:  c.Description,
		Image:        c.Metadata.Image,
		ExternalLink: c.Metadata.ExternalURL,
	}
}
