package signer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type EvmTransactionRequestStatus string

var (
	EvmTransactionRequestStatusInit       EvmTransactionRequestStatus = "init"
	EvmTransactionRequestStatusInProgress EvmTransactionRequestStatus = "in_progress"
	EvmTransactionRequestStatusSubmitted  EvmTransactionRequestStatus = "submitted"
	EvmTransactionRequestStatusMined      EvmTransactionRequestStatus = "mined"
	EvmTransactionRequestStatusFailed     EvmTransactionRequestStatus = "failed"
)

type GetTransactionHashResponseBody struct {
	TxHash           *string                      `json:"tx_hash"`
	Status           *EvmTransactionRequestStatus `json:"status"`
	FailedReason     *string                      `json:"failed_reason"`
	ErrorDescription string                       `json:"error_description,omitempty"`
}

func (l *SignerClient) GetTransactionHash(evmTxRequestId uint64) (*GetTransactionHashResponseBody, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%d/tx-hash", l.BaseUrl, "evm-transaction-request", evmTxRequestId), nil)
	if err != nil {
		return nil, err
	}
	l.auth(req)
	resp, err := l.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respBody GetTransactionHashResponseBody
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.ErrorDescription != "" {
		return nil, errors.New(respBody.ErrorDescription)
	}
	return &respBody, nil
}
