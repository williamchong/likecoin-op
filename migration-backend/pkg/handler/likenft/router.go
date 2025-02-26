package likenft

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft/migration"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft/migration_preview"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

type LikeNFTRouter struct {
	Db                  *sql.DB
	AsynqClient         *asynq.Client
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
	LikerlandUrlBase    string

	InitialNewClassOwner     string
	InitialBatchMintNFTOwner string
}

func (h *LikeNFTRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("POST /signing_message", &CreateSigningMessageHandler{
		Db: h.Db,
	})
	router.Handle("POST /likerid/migration", &LikerIDMigrationHandler{})

	migrationPreviewRouter := &migration_preview.MigrationPreviewRouter{
		Db:                  h.Db,
		CosmosAPI:           h.CosmosAPI,
		LikeNFTCosmosClient: h.LikeNFTCosmosClient,
		LikerlandUrlBase:    h.LikerlandUrlBase,
	}

	// FIXME: Find a way to handle CRUD paths
	// This is for paths without trailing /. e.g. GET / POST
	router.Handle("/migration-preview", migrationPreviewRouter.Router())
	// This is for paths with trailing/intermediate /, e.g. GET / PUT
	router.Handle("/migration-preview/", migrationPreviewRouter.Router())

	migrationRouter := &migration.MigrationRouter{
		Db:                       h.Db,
		AsynqClient:              h.AsynqClient,
		InitialNewClassOwner:     h.InitialNewClassOwner,
		InitialBatchMintNFTOwner: h.InitialBatchMintNFTOwner,
	}

	// This is for paths without trailing /. e.g. GET / POST
	router.Handle("/migration", migrationRouter.Router())
	// This is for paths with trailing/intermediate /, e.g. GET / PUT
	router.Handle("/migration/", migrationRouter.Router())

	return router
}
