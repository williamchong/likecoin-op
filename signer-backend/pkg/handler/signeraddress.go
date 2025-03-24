package handler

import (
	"net/http"

	"github.com/likecoin/like-signer-backend/pkg/evm"
)

type SignerAddressResponseBody struct {
	SignerAddress string `json:"signer_address,omitempty"`
	FailedReason  string `json:"failed_reason,omitempty"`
}

type SignerAddressHandler struct {
	evmClient *evm.Client
}

func NewSignerAddressHandler(evmClient *evm.Client) *SignerAddressHandler {
	return &SignerAddressHandler{evmClient: evmClient}
}

func (h *SignerAddressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signerAddress, err := h.evmClient.SignerAddress()
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, SignerAddressResponseBody{
			FailedReason: err.Error(),
		})
		return
	}
	SendJSON(w, http.StatusOK, SignerAddressResponseBody{
		SignerAddress: signerAddress.Hex(),
	})
}
