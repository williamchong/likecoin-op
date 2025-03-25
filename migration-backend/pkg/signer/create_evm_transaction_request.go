package signer

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type CreateEvmTransactionRequestRequestBody struct {
	ContractAddress string `json:"contract_address"`
	Method          string `json:"method"`
	ParamsHex       string `json:"params_hex"`
	DataHex         string `json:"data_hex"`
}

func MakeCreateEvmTransactionRequestRequestBody(
	metadata *bind.MetaData,
	method string,
	args ...interface{},
) func(contractAddress string) (*CreateEvmTransactionRequestRequestBody, error) {
	return func(contractAddress string) (*CreateEvmTransactionRequestRequestBody, error) {
		abi, err := metadata.GetAbi()
		if err != nil {
			return nil, err
		}
		paramsBytes, err := abi.Methods[method].Inputs.Pack(args...)
		if err != nil {
			return nil, err
		}
		paramsHex := hex.EncodeToString(paramsBytes)
		data, err := abi.Pack(method, args...)
		if err != nil {
			return nil, err
		}
		dataHex := hex.EncodeToString(data)
		return &CreateEvmTransactionRequestRequestBody{
			ContractAddress: contractAddress,
			Method:          method,
			ParamsHex:       paramsHex,
			DataHex:         dataHex,
		}, nil
	}
}

type CreateEvmTransactionRequestResponseBody struct {
	TransactionId    *uint64 `json:"transaction_id,omitempty"`
	ErrorDescription string  `json:"error_description,omitempty"`
}

func (l *SignerClient) CreateEvmTransactionRequest(reqBody *CreateEvmTransactionRequestRequestBody) (*CreateEvmTransactionRequestResponseBody, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", l.BaseUrl, "evm-transaction-request"), bytes.NewReader(body))
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
	return &respBody, nil
}
