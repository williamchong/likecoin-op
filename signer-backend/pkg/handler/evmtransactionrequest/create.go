package evmtransactionrequest

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/getsentry/sentry-go"
	appcontext "github.com/likecoin/like-signer-backend/pkg/context"
	"github.com/likecoin/like-signer-backend/pkg/evm"
	"github.com/likecoin/like-signer-backend/pkg/handler"
	"github.com/likecoin/like-signer-backend/pkg/logic"
)

type CreateRequestBody struct {
	ContractAddress string `json:"contract_address"`
	Method          string `json:"method"`
	ParamsHex       string `json:"params_hex"`
	DataHex         string `json:"data_hex"`
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
	return "POST /evm-transaction-request"
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
	contractAddress := common.HexToAddress(body.ContractAddress)

	data, err := hex.DecodeString(body.DataHex)
	if err != nil {
		logger.Error("failed to decode data hex", "error", err)
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

	transactionId, err := logic.GetOrCreateTransactionRequest(
		r.Context(),
		h.db,
		fromAddress,
		contractAddress,
		body.Method,
		body.ParamsHex,
		data,
	)

	if err != nil {
		logger.Error("failed to get or create transaction request", "error", err)
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
