package api

import (
	"encoding/json"
	"fmt"

	api_model "github.com/likecoin/like-migration-backend/pkg/likecoin/api/model"
)

type GetUserEVMMigrateResponse struct {
	LikerIdInfo *api_model.LikerIdInfo `json:"likerIdInfo"`
	EVMWallet   *string                `json:"evmWallet"`
}

func (a *LikecoinAPI) GetUserEVMMigrate(cosmosAddress string) (*GetUserEVMMigrateResponse, error) {
	resp, err := a.HTTPClient.Get(
		fmt.Sprintf("%s/wallet/evm/migrate/user/addr/%s", a.LikecoinAPIUrlBase, cosmosAddress),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response GetUserEVMMigrateResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
