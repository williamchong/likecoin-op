package user

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/user/model"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type GetUserProfileResponseBody struct {
	UserProfile      *api_model.UserProfile `json:"user_profile,omitempty"`
	ErrorDescription string                 `json:"error_description,omitempty"`
}

type GetUserProfileHandler struct {
	LikecoinAPI *likecoin_api.LikecoinAPI
	Db          *sql.DB
}

func (h *GetUserProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosWalletAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	userProfile, err := h.handle(cosmosWalletAddress)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &GetUserProfileResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetUserProfileResponseBody{
		UserProfile: userProfile,
	})
}

func (h *GetUserProfileHandler) handle(cosmosWalletAddress string) (*api_model.UserProfile, error) {
	var userProfile = &api_model.UserProfile{
		CosmosWalletAddress: cosmosWalletAddress,
	}

	likecoinUserProfile, err := h.LikecoinAPI.GetUserProfileViaWallet(cosmosWalletAddress)

	if err != nil {
		if errors.Is(err, likecoin_api.ErrUserProfileNotFound) {
			// continue
		} else {
			return nil, err
		}
	} else {
		userProfile.LikerID = &likecoinUserProfile.User
	}

	// TODO: the eth address should be retrieved from likerland
	sign, err := appdb.QueryNFTSigningMessageByCosmosAddress(h.Db, cosmosWalletAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// continue
		} else {
			return nil, err
		}
	} else {
		userProfile.EthWalletAddress = &sign.EthAddress
	}

	return userProfile, nil
}
