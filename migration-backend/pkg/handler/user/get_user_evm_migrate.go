package user

import (
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/user/model"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type GetUserEVMMigrateResponseBody struct {
	UserProfile *api_model.UserProfile `json:"user_profile,omitempty"`
}

type GetUserEVMMigrateHandler struct {
	LikecoinAPI *likecoin_api.LikecoinAPI
}

func (h *GetUserEVMMigrateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	cosmosWalletAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	userProfile, err := h.handle(cosmosWalletAddress)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetUserEVMMigrateResponseBody{
		UserProfile: userProfile,
	})
}

func (h *GetUserEVMMigrateHandler) handle(cosmosWalletAddress string) (*api_model.UserProfile, error) {
	response, err := h.LikecoinAPI.GetUserEVMMigrate(cosmosWalletAddress)
	if err != nil {
		return nil, err
	}

	var likerId *string
	var avatar *string

	if response.LikerIdInfo != nil {
		likerId = &response.LikerIdInfo.User
		avatar = &response.LikerIdInfo.Avatar
	}

	return &api_model.UserProfile{
		CosmosWalletAddress: cosmosWalletAddress,
		LikerID:             likerId,
		Avatar:              avatar,
		EthWalletAddress:    response.EVMWallet,
	}, nil
}
