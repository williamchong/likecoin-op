package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var (
	ErrQueryLatestBlock = errors.New("err querying latest block")
)

type queryLatestBlockResponse struct {
	Block model.Block `json:"block"`
}

func (a *CosmosAPI) QueryLatestBlock() (*model.Block, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/cosmos/base/tendermint/v1beta1/blocks/latest", a.NodeURL),
	)
	if err != nil {
		return nil, err
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryLatestBlock, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response queryLatestBlockResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response.Block, nil
}
