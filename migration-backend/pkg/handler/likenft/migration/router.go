package migration

import (
	"database/sql"
	"net/http"
)

type MigrationRouter struct {
	Db *sql.DB
}

func (h *MigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /migration/{cosmosWalletAddress}", &GetLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})
	router.Handle("POST /migration", &CreateMigrationHandler{
		Db: h.Db,
	})

	return router
}
