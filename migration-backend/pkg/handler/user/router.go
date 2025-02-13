package user

import (
	"database/sql"
	"net/http"

	likerland_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type UserRouter struct {
	Db          *sql.DB
	LikecoinAPI *likerland_api.LikecoinAPI
}

func (h *UserRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /profile/", &GetUserProfileHandler{
		Db:          h.Db,
		LikecoinAPI: h.LikecoinAPI,
	})

	return router
}
