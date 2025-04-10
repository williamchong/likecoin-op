package cosmos

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type QueryNFTClassesByOwnerRequest struct {
	QueryNFTClassesByOwnerPageRequest
	ISCNOwner string `url:"iscn_owner"`
}

type QueryNFTClassesByOwnerPageRequest struct {
	Key        int    `url:"pagination.key"`
	Offset     uint64 `url:"pagination.offset"`
	Limit      uint64 `url:"pagination.limit"`
	CountTotal bool   `url:"pagination.countTotal"`
	Reverse    bool   `url:"reverse"`
}

type QueryNFTClassesByOwnerResponse struct {
	Classes    []model.ClassListItem              `json:"classes,omitempty"`
	Pagination QueryNFTClassesByOwnerPageResponse `json:"pagination,omitempty"`
}

type QueryNFTClassesByOwnerPageResponse struct {
	NextKey int    `json:"next_key,omitempty"`
	Count   uint64 `json:"count,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryNFTClassesByOwner(request QueryNFTClassesByOwnerRequest) (*QueryNFTClassesByOwnerResponse, error) {
	v, _ := query.Values(request)
	url := fmt.Sprintf("%s/likechain/likenft/v1/class?%v", c.NodeURL, v.Encode())
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTClassesByOwnerResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}

type QueryAllNFTClassesByOwnerRequest struct {
	ISCNOwner string `url:"iscn_owner"`
}

type QueryAllNFTClasssesByOwnerResponse struct {
	Classes []model.ClassListItem `json:"classes,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryAllNFTClassesByOwner(request QueryAllNFTClassesByOwnerRequest) (*QueryAllNFTClasssesByOwnerResponse, error) {
	c1, err := c.QueryNFTClassesByOwner(QueryNFTClassesByOwnerRequest{
		ISCNOwner: request.ISCNOwner,
		QueryNFTClassesByOwnerPageRequest: QueryNFTClassesByOwnerPageRequest{
			Limit:      1,
			CountTotal: true,
		},
	})

	if err != nil {
		return nil, err
	}

	c2, err := c.QueryNFTClassesByOwner(QueryNFTClassesByOwnerRequest{
		ISCNOwner: request.ISCNOwner,
		QueryNFTClassesByOwnerPageRequest: QueryNFTClassesByOwnerPageRequest{
			Limit: c1.Pagination.Count,
		},
	})

	if err != nil {
		return nil, err
	}

	return &QueryAllNFTClasssesByOwnerResponse{
		Classes: c2.Classes,
	}, nil
}
