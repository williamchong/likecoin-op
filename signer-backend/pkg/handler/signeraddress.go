package handler

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-signer-backend/pkg/evm"
)

type SignerAddressResponseBody struct {
	SignerAddress string `json:"signer_address,omitempty"`
}

type SignerAddressHandler struct {
	evmClient *evm.Client
}

func NewSignerAddressHandler(evmClient *evm.Client) *SignerAddressHandler {
	return &SignerAddressHandler{evmClient: evmClient}
}

func (h *SignerAddressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	signerAddress, err := h.evmClient.SignerAddress()
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, MakeErrorResponseBody(err).
			WithSentryReported(hub.CaptureException(err)),
		)
		return
	}
	SendJSON(w, http.StatusOK, SignerAddressResponseBody{
		SignerAddress: signerAddress.Hex(),
	})
}
