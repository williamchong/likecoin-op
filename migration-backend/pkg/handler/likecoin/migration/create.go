package migration

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

type CreateLikeCoinMigrationRequestBody struct {
	EthAddress          string `json:"eth_address"`
	CosmosAddress       string `json:"cosmos_address"`
	EvmSignature        string `json:"evm_signature"`
	EvmSignatureMessage string `json:"evm_signature_message"`
	Amount              string `json:"amount"`
}

type CreateLikeCoinMigrationResponseBody struct {
	Migration        *model.LikeCoinMigration `json:"migration,omitempty"`
	ErrorDescription string                   `json:"error_description,omitempty"`
}

type CreateLikeCoinMigrationHandler struct {
	Db                           *sql.DB
	Signer                       *signer.SignerClient
	LikecoinBurningCosmosAddress string
}

func (p *CreateLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateLikeCoinMigrationRequestBody
	err := decoder.Decode(&data)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateLikeCoinMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	migration, err := p.handle(&data)
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateLikeCoinMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &CreateLikeCoinMigrationResponseBody{
		Migration: migration,
	})
}

func (p *CreateLikeCoinMigrationHandler) handle(req *CreateLikeCoinMigrationRequestBody) (*model.LikeCoinMigration, error) {
	mintingEthAddress, err := p.Signer.GetSignerAddress()

	if err != nil {
		return nil, err
	}

	m, err := likecoin.CreateIfAllEnded(
		p.Db,
		*mintingEthAddress,
		req.EthAddress,
		req.EvmSignature,
		req.EvmSignatureMessage,
		req.CosmosAddress,
		p.LikecoinBurningCosmosAddress,
		req.Amount,
	)

	if err != nil {
		return nil, err
	}

	return model.LikeCoinMigrationFromModel(m), nil
}
