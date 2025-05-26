package likenft

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
)

type LikerIDMigrationRequestBody struct {
	CosmosAddress        string `json:"cosmos_address,omitempty"`
	CosmosPubKey         string `json:"cosmos_pub_key,omitempty"`
	EthAddress           string `json:"eth_address,omitempty"`
	CosmosSignature      string `json:"cosmos_signature,omitempty"`
	CosmosSigningMessage string `json:"cosmos_signing_message,omitempty"`
	EthSignature         string `json:"eth_signature,omitempty"`
	EthSigningMessage    string `json:"eth_signing_message,omitempty"`
}

type LikerIDMigrationResponseBody struct {
	Response *likecoin_api.MigrateUserEVMWalletResponse `json:"response,omitempty"`
}

type LikerIDMigrationHandler struct {
	LikecoinAPI *likecoin_api.LikecoinAPI
}

func (p *LikerIDMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	decoder := json.NewDecoder(r.Body)
	var data LikerIDMigrationRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest,
			handler.MakeErrorResponseBody(err))
		return
	}

	response, err := p.handle(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, &LikerIDMigrationResponseBody{
		Response: response,
	})
}

func (p *LikerIDMigrationHandler) handle(
	data *LikerIDMigrationRequestBody,
) (*likecoin_api.MigrateUserEVMWalletResponse, error) {
	verified, err := cosmos.VerifyArbitrarySignature(data.CosmosPubKey, data.CosmosSignature, data.CosmosSigningMessage)
	if err != nil {
		return nil, err
	}

	if !verified {
		return nil, fmt.Errorf("cosmos signature not verified")
	}

	recoveredAddr, err := ethereum.RecoverAddress(data.EthSignature, []byte(data.EthSigningMessage))
	if err != nil {
		return nil, err
	}

	// The address from request data maybe in lower case
	// while the recovered address maybe in uppercase
	if !strings.EqualFold(recoveredAddr.Hex(), data.EthAddress) {
		return nil, fmt.Errorf("ethereum signature not verified")
	}

	// The cosmos signature will be verified by the likecoin api
	response, err := p.LikecoinAPI.MigrateUserEVMWallet(&likecoin_api.MigrateUserEVMWalletRequest{
		CosmosAddress:          data.CosmosAddress,
		CosmosSignature:        data.CosmosSignature,
		CosmosPublicKey:        data.CosmosPubKey,
		CosmosSignatureContent: data.CosmosSigningMessage,
		SignMethod:             likecoin_api.MigrateUserEVMWalletRequestSignMethodMemo,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
