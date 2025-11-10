package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var (
	ErrQueryTransaction = errors.New("err querying transaction")
	BlockTime           = 6 * time.Second
)

func (a *CosmosAPI) QueryTransaction(txHash string) (*cosmos.TxResponse, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/cosmos/tx/v1beta1/txs/%s", a.NodeURL, txHash),
	)
	if err != nil {
		return nil, errors.Join(ErrQueryTransaction, fmt.Errorf("a.HTTPClient.Get"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrQueryTransaction, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var txResponse cosmos.TxResponse
	err = decoder.Decode(&txResponse)
	if err != nil {
		return nil, errors.Join(ErrQueryTransaction, fmt.Errorf("decoder.Decode"), err)
	}
	return &txResponse, nil
}

func (a *CosmosAPI) QueryTransactionWithRetry(txHash string, blockCount int) (*cosmos.TxResponse, error) {
	txResponse, err := a.QueryTransaction(txHash)
	if err != nil && errors.Is(err, httputil.ErrNotFound) {
		time.Sleep(time.Duration(blockCount) * BlockTime)
		return a.QueryTransaction(txHash)
	}
	return txResponse, err
}
