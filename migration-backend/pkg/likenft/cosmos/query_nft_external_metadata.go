package cosmos

import (
	"encoding/json"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

var ErrNotOK = errors.New("err not ok")

func (c *LikeNFTCosmosClient) QueryNFTExternalMetadata(n *model.NFT) (*model.NFTMetadata, error) {
	if n.Uri == "" {
		return nil, nil
	}

	for _, urlBase := range c.nftExternalMetadataURLBaseIgnoreList {
		if strings.HasPrefix(n.Uri, urlBase) {
			return nil, nil
		}
	}

	resp, err := c.HTTPClient.Get(n.Uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.Join(ErrNotOK, errors.New(resp.Status))
	}

	var m model.NFTMetadata

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
