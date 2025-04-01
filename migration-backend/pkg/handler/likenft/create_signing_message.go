package likenft

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	likecoin_api_model "github.com/likecoin/like-migration-backend/pkg/likecoin/api/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateSigningMessageRequestBody struct {
	CosmosAddress string `json:"cosmos_address,omitempty"`
	LikerID       string `json:"liker_id,omitempty"`
	EthAddress    string `json:"eth_address,omitempty"`
}

type CreateSigningMessageResponseBody struct {
	Message          string `json:"message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type CreateSigningMessageHandler struct {
	Db *sql.DB
}

func (p *CreateSigningMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateSigningMessageRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	m, err := p.handle(p.Db, data.CosmosAddress, data.LikerID, data.EthAddress)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateSigningMessageResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &CreateSigningMessageResponseBody{
		Message: m,
	})
}

func (p *CreateSigningMessageHandler) handle(
	db *sql.DB,
	cosmosAddress string,
	likerID string,
	ethAddress string,
) (string, error) {
	issueTime := time.Now()

	memo := likecoin_api_model.MigrateUserEVMWalletMemo{
		Action:       likecoin_api_model.MigrateUserEVMWalletMemoActionMigrate,
		CosmosWallet: cosmosAddress,
		LikerWallet:  cosmosAddress,
		Ts:           uint64(issueTime.UnixMilli()),
		EvmWallet:    ethAddress,
	}

	jsonMemo, err := json.Marshal(memo)
	if err != nil {
		return "", err
	}

	n := model.NFTSigningMessage{
		CosmosAddress: cosmosAddress,
		LikerID:       likerID,
		EthAddress:    ethAddress,
		Nonce:         strconv.FormatInt(issueTime.UnixMilli(), 10),
		Message:       string(jsonMemo),
		IssueTime:     issueTime,
	}
	err = appdb.InsertNFTSigningMessage(db, &n)

	if err != nil {
		return "", err
	}

	return string(jsonMemo), nil
}
