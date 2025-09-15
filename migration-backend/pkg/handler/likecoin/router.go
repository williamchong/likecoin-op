package likecoin

import (
	"database/sql"
	"net/http"

	"github.com/hibiken/asynq"

	"github.com/likecoin/like-migration-backend/pkg/handler/likecoin/migration"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type LikeCoinRouter struct {
	Db                           *sql.DB
	AsynqClient                  *asynq.Client
	Signer                       *signer.SignerClient
	EvmLikeCoinClient            *evm.LikeCoin
	CosmosLikcCoinClient         *cosmos.LikeCoin
	LikecoinBurningCosmosAddress string
}

func (h *LikeCoinRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	migrationRouter := migration.MigrationRouter{
		Db:                           h.Db,
		AsynqClient:                  h.AsynqClient,
		Signer:                       h.Signer,
		EvmLikeCoinClient:            h.EvmLikeCoinClient,
		CosmosLikcCoinClient:         h.CosmosLikcCoinClient,
		LikecoinBurningCosmosAddress: h.LikecoinBurningCosmosAddress,
	}
	// FIXME: Find a way to handle CRUD paths
	// This is for paths without trailing /. e.g. GET / POST
	router.Handle("/migration", migrationRouter.Router())
	// This is for paths with trailing/intermediate /, e.g. GET / PUT
	router.Handle("/migration/", migrationRouter.Router())

	return router
}
