package likenft

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
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

func getSignMessage(
	cosmosAddress string,
	likerID string,
	ethAddress string,
	nonce string,
	issueTime time.Time,
) string {
	return fmt.Sprintf(`Liker ID: %s
Cosmos address: %s
Eth address: %s
Nonce: %s
UTC Timestamp: %d`, likerID, cosmosAddress, ethAddress, nonce, issueTime.UnixMicro())
}

func (p *CreateSigningMessageHandler) handle(
	db *sql.DB,
	cosmosAddress string,
	likerID string,
	ethAddress string,
) (string, error) {
	nonce := uuid.New().String()
	issueTime := time.Now()
	m := getSignMessage(cosmosAddress, likerID, ethAddress, nonce, issueTime)
	n := model.NFTSigningMessage{
		CosmosAddress: cosmosAddress,
		LikerID:       likerID,
		EthAddress:    ethAddress,
		Nonce:         nonce,
		Message:       m,
		IssueTime:     issueTime,
	}
	err := appdb.InsertNFTSigningMessage(db, &n)

	if err != nil {
		return "", err
	}

	return m, nil
}
