package migration

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type MigrationRouter struct {
	Db          *sql.DB
	AsynqClient *asynq.Client
	LikecoinAPI *likecoin_api.LikecoinAPI
}

func (h *MigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /migration/{cosmosWalletAddress}", &GetLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})
	router.Handle("POST /migration", &CreateMigrationHandler{
		Db:          h.Db,
		AsynqClient: h.AsynqClient,
		LikecoinAPI: h.LikecoinAPI,
	})
	router.Handle("PUT /migration/{cosmosWalletAddress}", &RetryMigrationHandler{
		Db:          h.Db,
		AsynqClient: h.AsynqClient,
	})

	router.Handle("GET /migration", &ListLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})

	router.Handle("DELETE /migration/{migrationId}", &RemoveLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})

	return router
}
