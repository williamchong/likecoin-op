package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/model/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var (
	ErrQueryTransaction = errors.New("err querying transaction")
)

func (a *CosmosAPI) QueryTransaction(txHash string) (*cosmos.TxResponse, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", a.NodeURL, txHash),
	)
	if err != nil {
		return nil, err
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryTransaction, err)
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
