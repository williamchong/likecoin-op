package migration

import (
	"database/sql"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type GetEvmPoolBalanceResponseBody struct {
	Denom    string `json:"denom,omitempty"`
	Amount   string `json:"amount,omitempty"`
	Decimals uint8  `json:"decimals,omitempty"`
}

type GetEvmPoolBalanceHandler struct {
	Db                *sql.DB
	EvmLikeCoinClient *evm.LikeCoin
	Signer            *signer.SignerClient
}

func (h *GetEvmPoolBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signerAddress, err := h.Signer.GetSignerAddress()
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, handler.MakeErrorResponseBody(err))
		return
	}

	symbol, err := h.EvmLikeCoinClient.Symbol(r.Context())
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, handler.MakeErrorResponseBody(err))
		return
	}

	balance, err := h.EvmLikeCoinClient.BalanceOf(r.Context(), common.HexToAddress(*signerAddress))
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, handler.MakeErrorResponseBody(err))
		return
	}
	balanceString := balance.String()

	decimals, err := h.EvmLikeCoinClient.Decimals()
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, handler.MakeErrorResponseBody(err))
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetEvmPoolBalanceResponseBody{
		Denom:    symbol,
		Amount:   balanceString,
		Decimals: decimals,
	})
}
