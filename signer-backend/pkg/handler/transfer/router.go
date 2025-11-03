package transfer

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-signer-backend/pkg/evm"
)

func MakeRouter(
	db *sql.DB,
	evmClient *evm.Client,
) http.Handler {
	router := http.NewServeMux()
	router.Handle(MakeCreatePattern(), NewCreateHandler(db, evmClient))
	return router
}
