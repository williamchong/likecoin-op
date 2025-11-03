package signer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type CreateEvmTransferTransactionRequestRequestBody struct {
	ToAddress common.Address `json:"to_address"`
	Amount    string         `json:"amount"`
}

func MakeCreateEvmTransferTransactionRequestRequestBody(
	toAddress common.Address,
	amount *big.Int,
) (*CreateEvmTransferTransactionRequestRequestBody, error) {
	return &CreateEvmTransferTransactionRequestRequestBody{
		ToAddress: toAddress,
		Amount:    amount.String(),
	}, nil
}

func (l *SignerClient) CreateEvmTransferTransactionRequest(
	reqBody *CreateEvmTransferTransactionRequestRequestBody,
) (*CreateEvmTransactionRequestResponseBody, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", l.BaseUrl, "transfer"), bytes.NewReader(body))
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
	var respBody CreateEvmTransactionRequestResponseBody
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.ErrorResponseBody != nil {
		return nil, respBody.ErrorResponseBody
	}
	return &respBody, nil
}
