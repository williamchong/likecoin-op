package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/go-querystring/query"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrQueryNFTsByOwner = errors.New("err querying nfts by owner")

type QueryNFTsByOwnerRequest struct {
	QueryNFTsByOwnerPageRequest
	Owner string `url:"owner"`
}

type QueryNFTsByOwnerPageRequest struct {
	Key        string `url:"pagination.key,omitempty"`
	Offset     uint64 `url:"pagination.offset,omitempty"`
	Limit      uint64 `url:"pagination.limit,omitempty"`
	CountTotal bool   `url:"pagination.countTotal,omitempty"`
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
		return nil, errors.Join(ErrQueryNFTsByOwner, fmt.Errorf("c.HTTPClient.Get error"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryNFTsByOwner, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTsByOwnerResponse
	err = decoder.Decode(&res)
	if err != nil {
		return nil, errors.Join(ErrQueryNFTsByOwner, fmt.Errorf("decoder.Decode"), err)
	}
	return &res, nil
}

var (
	ErrQueryAllNFTsByOwner = errors.New("err query all nfts by owner")
)

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
		return nil, errors.Join(ErrQueryAllNFTsByOwner, fmt.Errorf("c.QueryNFTsByOwner limit=%d", 1), err)
	}

	limit, err := strconv.ParseInt(c1.Pagination.Total, 10, 64)

	if err != nil {
		return nil, errors.Join(ErrQueryAllNFTsByOwner, fmt.Errorf("strconv.ParseInt %s", c1.Pagination.Total), err)
	}

	c2, err := c.QueryNFTsByOwner(QueryNFTsByOwnerRequest{
		Owner: request.Owner,
		QueryNFTsByOwnerPageRequest: QueryNFTsByOwnerPageRequest{
			Limit: uint64(limit),
		},
	})

	if err != nil {
		return nil, errors.Join(ErrQueryAllNFTsByOwner, fmt.Errorf("c.QueryNFTsByOwner limit=%d", limit), err)
	}

	return &QueryAllNFTsByOwnerResponse{
		NFTs: c2.NFTs,
	}, nil
}
