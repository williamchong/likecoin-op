package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	api_model "github.com/likecoin/like-migration-backend/pkg/likecoin/api/model"
)

var ErrUserProfileNotFound = errors.New("user profile not found")

func (a *LikecoinAPI) GetUserProfileViaWallet(cosmosAddress string) (*api_model.UserProfile, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/users/addr/%s/min", a.LikecoinAPIUrlBase, cosmosAddress),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrUserProfileNotFound
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var userProfile api_model.UserProfile
	err = decoder.Decode(&userProfile)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
