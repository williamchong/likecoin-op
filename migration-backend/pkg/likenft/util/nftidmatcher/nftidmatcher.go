package nftidmatcher

import (
	"fmt"
	"regexp"
	"slices"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type CosmosNFTIDMatcher interface {
	FindCosmosNFTBySerial(nfts []cosmosmodel.NFT, serialID uint64) (*cosmosmodel.NFT, bool)
}

type nftidMatcher struct{}

func MakeNFTIDMatcher() CosmosNFTIDMatcher {
	return &nftidMatcher{}
}

func (m *nftidMatcher) FindCosmosNFTBySerial(
	cosmosNFTs []cosmosmodel.NFT,
	serialID uint64,
) (*cosmosmodel.NFT, bool) {
	regex := regexp.MustCompile(fmt.Sprintf("^(?P<prefix>[a-zA-Z0-9]+)-(?P<maybe_num>0*%d)$", serialID))
	foundIndex := slices.IndexFunc(cosmosNFTs, func(n cosmosmodel.NFT) bool {
		return regex.MatchString(n.Id)
	})
	if foundIndex == -1 {
		return nil, false
	}
	return &cosmosNFTs[foundIndex], true
}
