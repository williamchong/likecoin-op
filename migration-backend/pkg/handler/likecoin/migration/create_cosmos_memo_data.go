package migration

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
)

type CreateCosmosMemoDataRequestBody struct {
	Signature  string `json:"signature"`
	EthAddress string `json:"ethAddress"`
	Amount     string `json:"amount"`
}

type CreateCosmosMemoDataResponseBody struct {
	MemoData         string `json:"memo_data,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type CreateCosmosMemoDataHandler struct {
}

func (h *CreateCosmosMemoDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateCosmosMemoDataRequestBody
	err := decoder.Decode(&data)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	m, err := h.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateCosmosMemoDataResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusAccepted, CreateCosmosMemoDataResponseBody{
		MemoData: m,
	})
}

func (h *CreateCosmosMemoDataHandler) handle(req *CreateCosmosMemoDataRequestBody) (string, error) {
	amount, err := types.ParseCoinNormalized(req.Amount)
	if err != nil {
		return "", err
	}

	return likecoin.EncodeCosmosMemoData(&likecoin.MemoData{
		Signature:  req.Signature,
		EthAddress: req.EthAddress,
		Amount:     amount,
	})
}
