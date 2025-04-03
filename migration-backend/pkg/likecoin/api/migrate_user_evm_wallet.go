package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MigrateUserEVMWalletRequestSignMethod string

const (
	MigrateUserEVMWalletRequestSignMethodMemo   MigrateUserEVMWalletRequestSignMethod = "memo"
	MigrateUserEVMWalletRequestSignMethodADR036 MigrateUserEVMWalletRequestSignMethod = "ADR-036"
)

type MigrateUserEVMWalletRequest struct {
	CosmosAddress          string                                `json:"cosmos_address"`
	CosmosSignature        string                                `json:"cosmos_signature"`
	CosmosPublicKey        string                                `json:"cosmos_public_key"`
	CosmosSignatureContent string                                `json:"cosmos_signature_content"`
	SignMethod             MigrateUserEVMWalletRequestSignMethod `json:"signMethod"`
}

type MigratedLikerLandUser struct {
	Id                  string `json:"id"`
	LikeWallet          string `json:"likeWallet"`
	LastLoginMethod     string `json:"lastLoginMethod"`
	RegisterLoginMethod string `json:"registerLoginMethod"`
}

type MigrateUserEVMWalletResponse struct {
	IsMigratedBookUser    bool                   `json:"isMigratedBookUser"`
	IsMigratedLikerId     bool                   `json:"isMigratedLikerId"`
	IsMigratedLikerLand   bool                   `json:"isMigratedLikerLand"`
	MigratedLikerId       string                 `json:"migratedLikerId"`
	MigratedLikerLandUser *MigratedLikerLandUser `json:"migratedLikerLandUser"`
	MigrateBookUserError  string                 `json:"migrateBookUserError"`
	MigrateLikerIdError   string                 `json:"migrateLikerIdError"`
	MigrateLikerLandError string                 `json:"migrateLikerLandError"`
}

func (a *LikecoinAPI) MigrateUserEVMWallet(
	request *MigrateUserEVMWalletRequest,
) (*MigrateUserEVMWalletResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/wallet/evm/migrate/user", a.LikecoinAPIUrlBase),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response MigrateUserEVMWalletResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
