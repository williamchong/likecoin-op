package signer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetSignerAddressResponseBody struct {
	SignerAddress string `json:"signer_address,omitempty"`
	FailedReason  string `json:"failed_reason,omitempty"`
}

func (l *SignerClient) GetSignerAddress() (*string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", l.BaseUrl, "signer-address"), nil)
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
	var respBody GetSignerAddressResponseBody
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.FailedReason != "" {
		return nil, errors.New(respBody.FailedReason)
	}
	return &respBody.SignerAddress, nil
}
