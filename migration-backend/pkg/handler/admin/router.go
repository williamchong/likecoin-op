package admin

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"

	"github.com/likecoin/like-migration-backend/pkg/handler/admin/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/handler/admin/likenft"
)

type AdminRouter struct {
	Db          *sql.DB
	AsynqClient *asynq.Client
}

func (h *AdminRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	likecoinRouter := likecoin.LikeCoinRouter{
		Db:          h.Db,
		AsynqClient: h.AsynqClient,
	}

	likenftRouter := likenft.LikeNFTRouter{
		Db: h.Db,
	}

	router.Handle("/likecoin/", http.StripPrefix("/likecoin", likecoinRouter.Router()))
	router.Handle("/likenft/", http.StripPrefix("/likenft", likenftRouter.Router()))

	return router
}
