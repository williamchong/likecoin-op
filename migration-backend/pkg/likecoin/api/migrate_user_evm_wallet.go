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

type MigrateUserEVMWalletResponse struct {
	IsMigratedLikerId     bool   `json:"isMigratedLikerId"`
	IsMigratedLikerLand   bool   `json:"isMigratedLikerLand"`
	MigratedLikerId       string `json:"migratedLikerId"`
	MigratedLikerLandUser string `json:"migratedLikerLandUser"`
	MigrateLikerIdError   string `json:"migrateLikerIdError"`
	MigrateLikerLandError string `json:"migrateLikerLandError"`
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
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
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
