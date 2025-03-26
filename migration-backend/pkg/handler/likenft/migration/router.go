package migration

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type MigrationRouter struct {
	Db                       *sql.DB
	AsynqClient              *asynq.Client
	LikecoinAPI              *likecoin_api.LikecoinAPI
	InitialNewClassOwner     string
	InitialBatchMintNFTOwner string
}

func (h *MigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /migration/{cosmosWalletAddress}", &GetLikeNFTAssetMigrationHandler{
		Db: h.Db,
	})
	router.Handle("POST /migration", &CreateMigrationHandler{
		Db:                       h.Db,
		AsynqClient:              h.AsynqClient,
		LikecoinAPI:              h.LikecoinAPI,
		InitialNewClassOwner:     h.InitialNewClassOwner,
		InitialBatchMintNFTOwner: h.InitialBatchMintNFTOwner,
	})
	router.Handle("PUT /migration/{cosmosWalletAddress}", &RetryMigrationHandler{
		Db:          h.Db,
		AsynqClient: h.AsynqClient,
	})

	return router
}
