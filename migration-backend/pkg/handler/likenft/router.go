package likenft

import (
	"database/sql"
	"net/http"
)

type LikeNFTRouter struct {
	Db *sql.DB
}

func (h *LikeNFTRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("POST /signing_message", &CreateSigningMessageHandler{
		Db: h.Db,
	})
	router.Handle("POST /likerid/migration", &LikerIDMigrationHandler{})

	return router
}
