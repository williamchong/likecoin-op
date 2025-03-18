package migration

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
)

type CreateSigningMessageRequestBody struct {
	Amount string `json:"amount"`
}

type CreateSigningMessageResponseBody struct {
	SigningMessage   string `json:"signing_message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type CreateEthSigningMessageHandler struct {
}

func (h *CreateEthSigningMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateSigningMessageRequestBody
	err := decoder.Decode(&data)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	m, err := h.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusAccepted, CreateSigningMessageResponseBody{
		SigningMessage: m,
	})
}

func (h *CreateEthSigningMessageHandler) handle(req *CreateSigningMessageRequestBody) (string, error) {
	amount, err := types.ParseCoinNormalized(req.Amount)
	if err != nil {
		return "", err
	}
	return likecoin.GetEthSigningMessage(amount), nil
}
