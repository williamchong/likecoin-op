package cosmos

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type QueryNFTsByOwnerRequest struct {
	QueryNFTsByOwnerPageRequest
	Owner string `url:"owner"`
}

type QueryNFTsByOwnerPageRequest struct {
	Key        string `url:"pagination.key"`
	Offset     uint64 `url:"pagination.offset"`
	Limit      uint64 `url:"pagination.limit"`
	CountTotal bool   `url:"pagination.countTotal"`
}

type QueryNFTsByOwnerResponse struct {
	NFTs       []model.NFT                  `json:"nfts,omitempty"`
	Pagination QueryNFTsByOwnerPageResponse `json:"pagination,omitempty"`
}

type QueryNFTsByOwnerPageResponse struct {
	NextKey string `json:"next_key,omitempty"`
	Total   string `json:"total,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryNFTsByOwner(request QueryNFTsByOwnerRequest) (*QueryNFTsByOwnerResponse, error) {
	v, _ := query.Values(request)
	url := fmt.Sprintf("%s/cosmos/nft/v1beta1/nfts?%v", c.NodeURL, v.Encode())
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTsByOwnerResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}

type QueryAllNFTsByOwnerRequest struct {
	Owner string `url:"owner"`
}

type QueryAllNFTsByOwnerResponse struct {
	NFTs []model.NFT `json:"nfts,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryAllNFTsByOwner(request QueryAllNFTsByOwnerRequest) (*QueryAllNFTsByOwnerResponse, error) {
	c1, err := c.QueryNFTsByOwner(QueryNFTsByOwnerRequest{
		Owner: request.Owner,
		QueryNFTsByOwnerPageRequest: QueryNFTsByOwnerPageRequest{
			Limit:      1,
			CountTotal: true,
		},
	})

	if err != nil {
		return nil, err
	}

	limit, err := strconv.ParseInt(c1.Pagination.Total, 10, 64)

	if err != nil {
		return nil, err
	}

	c2, err := c.QueryNFTsByOwner(QueryNFTsByOwnerRequest{
		Owner: request.Owner,
		QueryNFTsByOwnerPageRequest: QueryNFTsByOwnerPageRequest{
			Limit: uint64(limit),
		},
	})

	if err != nil {
		return nil, err
	}

	return &QueryAllNFTsByOwnerResponse{
		NFTs: c2.NFTs,
	}, nil
}
