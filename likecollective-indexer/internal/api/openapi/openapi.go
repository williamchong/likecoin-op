package openapi

import (
	"log"
	"net/http"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

type openAPIHandler struct {
	db *ent.Client

	stakingRepository      database.StakingRepository
	accountRepository      database.AccountRepository
	bookNFTRepository      database.BookNFTRepository
	stakingEventRepository database.StakingEventRepository
}

var _ api.Handler = &openAPIHandler{}

func NewOpenAPIHandler(db database.Service) http.Handler {
	handler := &openAPIHandler{
		db: db.Client(),

		stakingRepository:      database.MakeStakingRepository(db),
		accountRepository:      database.MakeAccountRepository(db),
		bookNFTRepository:      database.MakeBookNFTRepository(db),
		stakingEventRepository: database.MakeStakingEventRepository(db),
	}

	srv, err := api.NewServer(
		handler,
	)
	if err != nil {
		log.Fatal(err)
	}
	return srv
}
