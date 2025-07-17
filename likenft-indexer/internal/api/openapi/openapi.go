package openapi

import (
	"log"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/internal/api/openapi/middleware"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"

	"github.com/hibiken/asynq"
)

type OpenAPIHandler struct {
	likeProtocolAddress string

	db                     *ent.Client
	nftClassRepository     database.NFTClassRepository
	nftRepository          database.NFTRepository
	evmEventRepository     database.EVMEventRepository
	accountRepository      database.AccountRepository
	likeProtocolRepository database.LikeProtocolRepository

	asynqClient *asynq.Client
}

var _ api.Handler = &OpenAPIHandler{}

func NewOpenAPIHandler(
	indexActionApiKey string,

	likeProtocolAddress string,

	db database.Service,

	asyncClient *asynq.Client,
) http.Handler {
	handler := &OpenAPIHandler{
		likeProtocolAddress: likeProtocolAddress,

		db:                     db.Client(),
		nftClassRepository:     database.MakeNFTClassRepository(db),
		nftRepository:          database.MakeNFTRepository(db),
		evmEventRepository:     database.MakeEVMEventRepository(db),
		accountRepository:      database.MakeAccountRepository(db),
		likeProtocolRepository: database.MakeLikeProtocolRepository(db),

		asynqClient: asyncClient,
	}

	srv, err := api.NewServer(
		handler,
		api.WithMiddleware(
			middleware.MakeHeaderApiKeyAuthMiddleware(
				middleware.IndexActionApiKeyAuthHeaderName,
				indexActionApiKey),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return srv
}
