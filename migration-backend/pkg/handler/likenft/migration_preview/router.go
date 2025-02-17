package migration_preview

import (
	"database/sql"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
)

type MigrationPreviewRouter struct {
	Db                  *sql.DB
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
	LikerlandUrlBase    string
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
		LikerlandUrlBase:    h.LikerlandUrlBase,
	})

	return router
}
