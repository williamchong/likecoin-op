package migration_preview

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

type MigrationPreviewRouter struct {
	Db                  *sql.DB
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient

	ClassMigrationEstimatedDuration time.Duration
	NFTMigrationEstimatedDuration   time.Duration
}

func (h *MigrationPreviewRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /migration-preview/{cosmosWalletAddress}", &GetMigrationPreviewHandler{
		Db: h.Db,
	})
	router.Handle("POST /migration-preview", &CreateMigrationPreviewHandler{
		Db:                  h.Db,
		CosmosAPI:           h.CosmosAPI,
		LikeNFTCosmosClient: h.LikeNFTCosmosClient,

		ClassMigrationEstimatedDuration: h.ClassMigrationEstimatedDuration,
		NFTMigrationEstimatedDuration:   h.NFTMigrationEstimatedDuration,
	})

	return router
}
