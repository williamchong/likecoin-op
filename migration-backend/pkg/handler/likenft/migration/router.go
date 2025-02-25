package migration

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"
)

type MigrationRouter struct {
	Db                       *sql.DB
	AsynqClient              *asynq.Client
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
		InitialNewClassOwner:     h.InitialNewClassOwner,
		InitialBatchMintNFTOwner: h.InitialBatchMintNFTOwner,
	})

	return router
}
