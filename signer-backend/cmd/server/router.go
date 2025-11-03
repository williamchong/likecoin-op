package main

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-signer-backend/pkg/evm"
	"github.com/likecoin/like-signer-backend/pkg/handler"
	"github.com/likecoin/like-signer-backend/pkg/handler/evmtransactionrequest"
	"github.com/likecoin/like-signer-backend/pkg/handler/transfer"
)

func MakeRouter(
	db *sql.DB,
	evmClient *evm.Client,
) *http.ServeMux {
	mainRouter := http.NewServeMux()

	mainRouter.Handle("GET /healthz", handler.MakeHealthzHandler())
	mainRouter.Handle("GET /signer-address", handler.NewSignerAddressHandler(evmClient))

	mainRouter.Handle("/evm-transaction-request", evmtransactionrequest.MakeRouter(db, evmClient))
	mainRouter.Handle("/evm-transaction-request/", evmtransactionrequest.MakeRouter(db, evmClient))

	mainRouter.Handle("/transfer", transfer.MakeRouter(db, evmClient))
	mainRouter.Handle("/transfer/", transfer.MakeRouter(db, evmClient))

	return mainRouter
}
