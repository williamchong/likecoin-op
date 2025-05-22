package migration

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
)

type CreateSigningMessageRequestBody struct {
	Amount string `json:"amount"`
}

type CreateSigningMessageResponseBody struct {
	SigningMessage string `json:"signing_message,omitempty"`
}

type CreateEthSigningMessageHandler struct {
}

func (h *CreateEthSigningMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	decoder := json.NewDecoder(r.Body)
	var data CreateSigningMessageRequestBody
	err := decoder.Decode(&data)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}

	m, err := h.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
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
