package user

import (
	"net/http"

	likerland_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type UserRouter struct {
	LikecoinAPI *likerland_api.LikecoinAPI
}

func (h *UserRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /profile/", &GetUserEVMMigrateHandler{
		LikecoinAPI: h.LikecoinAPI,
	})

	return router
}
