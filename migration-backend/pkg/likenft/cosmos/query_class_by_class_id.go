package cosmos

import (
	"encoding/json"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type QueryClassByClassIdRequest struct {
	ClassId string
}

type QueryClassByClassIdResponse struct {
	Class *model.Class `json:"class"`
}

func (c *LikeNFTCosmosClient) QueryClassByClassId(request QueryClassByClassIdRequest) (*QueryClassByClassIdResponse, error) {
	url := fmt.Sprintf("%s/cosmos/nft/v1beta1/classes/%v", c.NodeURL, request.ClassId)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryClassByClassIdResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}
