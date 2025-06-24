package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var ErrQueryClassByClassId = errors.New("err querying class by class id")

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
		return nil, errors.Join(ErrQueryClassByClassId, fmt.Errorf("c.HTTPClient.Get"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryClassByClassId, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var res QueryClassByClassIdResponse
	err = decoder.Decode(&res)
	if err != nil {
		return nil, errors.Join(ErrQueryClassByClassId, fmt.Errorf("decoder.Decode"), err)
	}
	return &res, nil
}
