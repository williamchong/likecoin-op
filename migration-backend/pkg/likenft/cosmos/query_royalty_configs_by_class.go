package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrQueryRoyaltyConfigsByClassId = errors.New("err querying royalty configs by class id")

type QueryRoyaltyConfigsByClassIdRequest struct {
	ClassId string
}

type QueryRoyaltyConfigsByClassIdResponse struct {
	RoyaltyConfig *model.RoyaltyConfig `json:"royalty_config,omitempty"`
}

// e.g. 'https://node.testnet.like.co/likechain/likenft/v1/royalty_configs/likenft1trdzjmnnyr73jqh4mmkgp0kj8zxh6m598duwtvum820zzday4muquqme0r
func (c *LikeNFTCosmosClient) QueryRoyaltyConfigsByClassId(request QueryRoyaltyConfigsByClassIdRequest) (*QueryRoyaltyConfigsByClassIdResponse, error) {
	url, err := url.Parse("/likechain/likenft/v1/royalty_configs")

	if err != nil {
		return nil, err
	}

	url = url.JoinPath(request.ClassId)

	base, err := url.Parse(c.NodeURL)

	if err != nil {
		return nil, err
	}

	fmt.Println(base.ResolveReference(url).String())

	resp, err := c.HTTPClient.Get(base.ResolveReference(url).String())

	if err != nil {
		fmt.Printf("c.HTTPClient.Get error %v\n", err)
		return nil, err
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryRoyaltyConfigsByClassId, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryRoyaltyConfigsByClassIdResponse
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Printf("decoder.Decode error %v\n", err)
		return nil, err
	}
	return &res, nil
}
