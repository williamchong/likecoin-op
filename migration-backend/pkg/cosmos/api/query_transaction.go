package api

import (
	"encoding/json"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/model/cosmos"
)

func (a *CosmosAPI) QueryTransaction(txHash string) (*cosmos.TxResponse, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", a.NodeURL, txHash),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var txResponse cosmos.TxResponse
	err = decoder.Decode(&txResponse)
	if err != nil {
		return nil, err
	}
	return &txResponse, nil
}
