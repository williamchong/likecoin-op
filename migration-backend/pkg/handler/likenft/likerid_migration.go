package likenft

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/handler"
)

type MigrateLikerIDWithEthAddressRequestBody struct {
	CosmosPubKey    string `json:"cosmos_pub_key,omitempty"`
	LikerID         string `json:"liker_id,omitempty"`
	EthAddress      string `json:"eth_address,omitempty"`
	CosmosSignature string `json:"cosmos_signature,omitempty"`
	EthSignature    string `json:"eth_signature,omitempty"`
	SigningMessage  string `json:"signing_message,omitempty"`
}

type MigrateLikerIDWithEthAddressResponseBody struct {
	Message          string `json:"message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type LikerIDMigrationHandler struct {
}

func (p *LikerIDMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data MigrateLikerIDWithEthAddressRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	err = p.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &CreateSigningMessageResponseBody{
		Message: "OK",
	})
}

func (p *LikerIDMigrationHandler) handle(data *MigrateLikerIDWithEthAddressRequestBody) error {
	verified, err := cosmos.VerifyArbitrarySignature(data.CosmosPubKey, data.CosmosSignature, data.SigningMessage)
	if err != nil {
		return err
	}

	// FIXME: check cosmos signature verified
	// ```
	// if !verified {
	//   return fmt.Errorf("cosmos signature not verified")
	// }
	// ```
	// For details, see VerifyArbitrarySignature
	fmt.Printf("cosmos signature verified: %v\n", verified)

	recoveredAddr, err := ethereum.RecoverAddress(data.EthSignature, []byte(data.SigningMessage))
	if err != nil {
		return err
	}

	if recoveredAddr.Hex() != data.EthAddress {
		return fmt.Errorf("ethereum signature not verified")
	}

	// TODO send to likerland server

	return nil
}
