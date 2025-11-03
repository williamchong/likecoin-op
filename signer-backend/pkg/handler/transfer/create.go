package transfer

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/getsentry/sentry-go"
	appcontext "github.com/likecoin/like-signer-backend/pkg/context"
	"github.com/likecoin/like-signer-backend/pkg/evm"
	"github.com/likecoin/like-signer-backend/pkg/handler"
	"github.com/likecoin/like-signer-backend/pkg/logic"
)

type CreateRequestBody struct {
	ToAddress string `json:"to_address"`
	Amount    string `json:"amount"`
}

type CreateResponseBody struct {
	TransactionId *uint64 `json:"transaction_id,omitempty"`
}

type CreateHandler struct {
	db        *sql.DB
	evmClient *evm.Client
}

func NewCreateHandler(
	db *sql.DB,
	evmClient *evm.Client,
) *CreateHandler {
	return &CreateHandler{
		db:        db,
		evmClient: evmClient,
	}
}

func MakeCreatePattern() string {
	return "POST /transfer"
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := appcontext.LoggerFromContext(r.Context())
	hub := sentry.GetHubFromContext(r.Context())

	var body CreateRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("failed to decode request body", "error", err)
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}
	toAddress := common.HexToAddress(body.ToAddress)

	amount, ok := big.NewInt(0).SetString(body.Amount, 10)
	if !ok {
		logger.Error("failed to parse amount", "amount", body.Amount)
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}
	fromAddress, err := h.evmClient.SignerAddress()
	if err != nil {
		logger.Error("failed to get signer address", "error", err)
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)),
		)
		return
	}

	transactionId, err := logic.GetOrCreateTransferTransactionRequest(
		r.Context(),
		h.db,
		fromAddress,
		toAddress,
		amount,
	)

	if err != nil {
		logger.Error("failed to get or create transfer transaction request", "error", err)
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, CreateResponseBody{
		TransactionId: &transactionId,
	})

	go awaitTx(
		r.Context().(appcontext.GracefulHandleContext),
		logger,
		h.db,
		h.evmClient,
		transactionId,
	)
}

func awaitTx(
	ctx appcontext.GracefulHandleContext,
	logger *slog.Logger,
	db *sql.DB,
	evmClient *evm.Client,
	transactionId uint64,
) {
	go logic.AwaitTransaction(
		ctx,
		logger,
		db,
		evmClient,
		transactionId,
	)
}
