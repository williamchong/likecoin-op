package migration

import (
	"database/sql"
	"net/http"
)

type MigrationRouter struct {
	Db                           *sql.DB
}

func (h *MigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /migration", &ListLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})

	router.Handle("DELETE /migration/{migrationId}", &RemoveLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})

	return router
}
