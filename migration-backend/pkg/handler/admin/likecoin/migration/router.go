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

	router.Handle("GET /migration", &ListLikeCoinMigrationHandler{
		Db: h.Db,
	})

	router.Handle("DELETE /migration/{migrationId}", &RemoveLatestLikeCoinMigrationHandler{
		Db: h.Db,
	})

	return router
}
