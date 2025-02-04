package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/likecoin-migration-backend/pkg/logic"
)

type apiMigrationRecord struct {
	CosmosTxHash  string `json:"cosmos_tx_hash"`
	EthTxHash     string `json:"eth_tx_hash"`
	CosmosAddress string `json:"cosmos_address"`
	EthAddress    string `json:"eth_address"`
}

type responseBody struct {
	MigrationRecord  *apiMigrationRecord         `json:"migration_record,omitempty"`
	Status           logic.MigrationRecordStatus `json:"status,omitempty"`
	ErrorDescription string                      `json:"error_description,omitempty"`
}

type GetLikeCoinMigrationRecordHandler struct {
	Db        *sql.DB
	EthClient *ethclient.Client
}

func (p *GetLikeCoinMigrationRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosTxHash := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	migrationRecord, status, err := logic.GetLikeCoinMigrationRecord(p.Db, p.EthClient, cosmosTxHash)

	if err != nil {
		p.write(w, http.StatusInternalServerError, &responseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	_apiMigrationRecord := apiMigrationRecord(*migrationRecord)

	p.write(w, http.StatusOK, &responseBody{
		MigrationRecord: &_apiMigrationRecord,
		Status:          status,
	})
}

func (h *GetLikeCoinMigrationRecordHandler) write(w http.ResponseWriter, statusCode int, body *responseBody) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := encoder.Encode(body)
	if err != nil {
		panic(err)
	}
}
