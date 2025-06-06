package cosmosnftidclassifier

import "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"

type CosmosNFTIDClassifier interface {
	Classify(nft model.NFT, nfts ...model.NFT) *result
}

type cosmosNFTIDClassifier struct {
	cosmosNFTIDMatcher cosmosNFTIDMatcher
}

func MakeCosmosNFTIDClassifier(
	cosmosNFTIDMatcher cosmosNFTIDMatcher,
) CosmosNFTIDClassifier {
	return &cosmosNFTIDClassifier{
		cosmosNFTIDMatcher,
	}
}

func (c *cosmosNFTIDClassifier) Classify(first model.NFT, others ...model.NFT) *result {
	nfts := append([]model.NFT{first}, others...)
	serialNFTIDs := make(SerialNFTIDs, 0)
	arbitraryNFTIDs := make([]string, 0)

	for _, nft := range nfts {
		serialID, ok := c.cosmosNFTIDMatcher.ExtractSerialID(nft.Id)
		if ok {
			serialNFTIDs = append(serialNFTIDs, serialID)
		} else {
			arbitraryNFTIDs = append(arbitraryNFTIDs, nft.Id)
		}

	}
	return &result{
		SerialNFTIDs:    serialNFTIDs,
		ArbitraryNFTIDs: arbitraryNFTIDs,
	}
}
