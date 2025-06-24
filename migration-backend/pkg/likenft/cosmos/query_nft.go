package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrQueryNFT = errors.New("err querying nft")

type QueryNFTRequest struct {
	ClassId string
	Id      string
}

type QueryNFTResponse struct {
	NFT *model.NFT `json:"nft,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryNFT(request QueryNFTRequest) (*QueryNFTResponse, error) {
	url := fmt.Sprintf("%s/cosmos/nft/v1beta1/nfts/%v/%v", c.NodeURL, request.ClassId, request.Id)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryNFT, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}
