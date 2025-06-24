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

var ErrQueryNFTsByClassId = errors.New("err querying nfts by class id")

type QueryNFTsByClassIdRequest struct {
	QueryNFTsByClassIdPageRequest
	ClassId string `url:"class_id"`
}

type QueryNFTsByClassIdPageRequest struct {
	Key        string `url:"pagination.key,omitempty"`
	Offset     uint64 `url:"pagination.offset,omitempty"`
	Limit      uint64 `url:"pagination.limit,omitempty"`
	CountTotal bool   `url:"pagination.countTotal,omitempty"`
}

type QueryNFTsByClassIdResponse struct {
	NFTs       []model.NFT                    `json:"nfts,omitempty"`
	Pagination QueryNFTsByClassIdPageResponse `json:"pagination,omitempty"`
}

type QueryNFTsByClassIdPageResponse struct {
	NextKey string `json:"next_key,omitempty"`
	Total   string `json:"total,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryNFTsByClassId(request QueryNFTsByClassIdRequest) (*QueryNFTsByClassIdResponse, error) {
	v, _ := query.Values(request)
	url := fmt.Sprintf("%s/cosmos/nft/v1beta1/nfts?%v", c.NodeURL, v.Encode())
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryNFTsByClassId, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryNFTsByClassIdResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}

type QueryAllNFTsByClassIdRequest struct {
	ClassId string `url:"class_id"`
}

type QueryAllNFTsByClassIdResponse struct {
	NFTs []model.NFT `json:"nfts,omitempty"`
}

func (c *LikeNFTCosmosClient) QueryAllNFTsByClassId(request QueryAllNFTsByClassIdRequest) (*QueryAllNFTsByClassIdResponse, error) {
	c1, err := c.QueryNFTsByClassId(QueryNFTsByClassIdRequest{
		ClassId: request.ClassId,
		QueryNFTsByClassIdPageRequest: QueryNFTsByClassIdPageRequest{
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

	c2, err := c.QueryNFTsByClassId(QueryNFTsByClassIdRequest{
		ClassId: request.ClassId,
		QueryNFTsByClassIdPageRequest: QueryNFTsByClassIdPageRequest{
			Limit: uint64(limit),
		},
	})

	if err != nil {
		return nil, err
	}

	return &QueryAllNFTsByClassIdResponse{
		NFTs: c2.NFTs,
	}, nil
}
