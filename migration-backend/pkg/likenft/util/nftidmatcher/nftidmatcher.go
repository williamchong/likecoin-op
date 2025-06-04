package nftidmatcher

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type CosmosNFTIDMatcher interface {
	ExtractSerialID(cosmosNFTID string) (uint64, bool)
	FindCosmosNFTBySerial(nfts []cosmosmodel.NFT, serialID uint64) (*cosmosmodel.NFT, bool)
}

type nftidMatcher struct {
}

func MakeNFTIDMatcher() CosmosNFTIDMatcher {
	return &nftidMatcher{}
}

func (m *nftidMatcher) ExtractSerialID(cosmosNFTID string) (uint64, bool) {
	regex := regexp.MustCompile("(?P<prefix>.+)-(?P<maybe_num>[0-9]+)")

	matches := regex.FindStringSubmatch(cosmosNFTID)
	if matches == nil {
		return 0, false
	}

	numIndex := regex.SubexpIndex("maybe_num")
	nftIdStr := matches[numIndex]
	nftId, err := strconv.ParseUint(nftIdStr, 10, 64)
	if err != nil {
		return 0, false
	}
	return nftId, true
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
