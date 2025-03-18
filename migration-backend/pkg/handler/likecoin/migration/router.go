package migration

import (
	"database/sql"
	"net/http"
)

type MigrationRouter struct {
	Db                           *sql.DB
	EthWalletPrivateKey          string
	LikecoinBurningCosmosAddress string
}

func (h *MigrationRouter) Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("POST /migration/eth-signing-message", &CreateEthSigningMessageHandler{})
	router.Handle("POST /migration/cosmos-memo-data", &CreateCosmosMemoDataHandler{})
	router.Handle("GET /migration/{cosmosWalletAddress}", &GetLatestLikeCoinMigrationHandler{
		Db: h.Db,
	})
	router.Handle("POST /migration", &CreateLikeCoinMigrationHandler{
		Db:                           h.Db,
		EthWalletPrivateKey:          h.EthWalletPrivateKey,
		LikecoinBurningCosmosAddress: h.LikecoinBurningCosmosAddress,
	})
	router.Handle("PUT /migration/{cosmosWalletAddress}/cosmos-tx-hash", &UpdateLikeCoinMigrationCosmosHandler{
		Db: h.Db,
	})

	return router
}
