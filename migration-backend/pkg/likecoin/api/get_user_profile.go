package api

import (
	"encoding/json"
	"errors"
	"fmt"

	api_model "github.com/likecoin/like-migration-backend/pkg/likecoin/api/model"
	"github.com/likecoin/like-migration-backend/pkg/util/httputil"
)

var (
	ErrGetUserProfile = errors.New("err getting user profile")
)

func (a *LikecoinAPI) GetUserProfileViaWallet(cosmosAddress string) (*api_model.UserProfile, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/users/addr/%s/min", a.LikecoinAPIUrlBase, cosmosAddress),
	)
	if err != nil {
		return nil, errors.Join(ErrGetUserProfile, fmt.Errorf("a.HTTPClient.Get"), err)
	}
	if err = httputil.HandleResponseStatus(resp); err != nil {
		return nil, errors.Join(ErrGetUserProfile, err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var userProfile api_model.UserProfile
	err = decoder.Decode(&userProfile)
	if err != nil {
		return nil, errors.Join(ErrGetUserProfile, fmt.Errorf("decoder.Decode"), err)
	}
	return &userProfile, nil
}
