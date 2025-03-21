package evmtransactionrequest

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	appdb "github.com/likecoin/like-signer-backend/pkg/db"
	"github.com/likecoin/like-signer-backend/pkg/handler"
	"github.com/likecoin/like-signer-backend/pkg/model"
)

type GetTxHashResponseBody struct {
	TxHash           *string                            `json:"tx_hash"`
	Status           *model.EvmTransactionRequestStatus `json:"status"`
	FailedReason     *string                            `json:"failed_reason"`
	ErrorDescription string                             `json:"error_description,omitempty"`
}

type GetTxHashHandler struct {
	db *sql.DB
}

func NewGetTxHashHandler(db *sql.DB) *GetTxHashHandler {
	return &GetTxHashHandler{db: db}
}

func MakeGetTxHashPattern() string {
	return "GET /evm-transaction-request/{id}/tx-hash"
}

func GetTxHashFromPath(path string) (uint64, error) {
	paths := strings.Split(path, "/")
	transactionId := paths[len(paths)-2]
	transactionIdUint, err := strconv.ParseUint(transactionId, 10, 64)
	if err != nil {
		return 0, err
	}
	return transactionIdUint, nil
}

func (h *GetTxHashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	transactionId, err := GetTxHashFromPath(r.URL.Path)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, GetTxHashResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	transaction, err := appdb.QueryEvmTransactionRequest(h.db, &appdb.QueryEvmTransactionRequestFilter{
		Id: &transactionId,
	})
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, GetTxHashResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	reason := transaction.FailedReason

	if transaction.Status == model.EvmTransactionRequestStatusFailed {
		handler.SendJSON(w, http.StatusOK, GetTxHashResponseBody{
			TxHash:       transaction.SignedTxHash,
			Status:       &transaction.Status,
			FailedReason: reason,
		})
		return
	}

	// Only minted transaction will be returned
	if transaction.Status != model.EvmTransactionRequestStatusMined {
		handler.SendJSON(w, http.StatusOK, GetTxHashResponseBody{
			TxHash: transaction.SignedTxHash,
			Status: &transaction.Status,
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, GetTxHashResponseBody{
		TxHash: transaction.SignedTxHash,
		Status: &transaction.Status,
	})
}
