package cosmos

import (
	"encoding/json"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

func (c *LikeNFTCosmosClient) QueryNFTExternalMetadata(n *model.NFT) (*model.NFTMetadata, error) {
	if n.Uri == "" {
		return nil, nil
	}

	resp, err := c.HTTPClient.Get(n.Uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var m model.NFTMetadata

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
