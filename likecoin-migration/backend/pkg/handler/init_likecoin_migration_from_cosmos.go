package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/likecoin-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/likecoin-migration-backend/pkg/logic"
)

type RequestBody struct {
	CosmosTxHash string `json:"cosmos_tx_hash"`
}

type ResponseBody struct {
	EthTxHash        string `json:"eth_tx_hash,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type InitLikeCoinMigrationFromCosmosHandler struct {
	Db                     *sql.DB
	EthClient              *ethclient.Client
	CosmosAPI              *api.CosmosAPI
	EthWalletPrivateKey    string
	EthNetworkPublicRPCURL string
	EthTokenAddress        string
}

func (p *InitLikeCoinMigrationFromCosmosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data RequestBody
	err := decoder.Decode(&data)
	if err != nil {
		p.write(w, http.StatusInternalServerError, &ResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	// TODO: lock CosmosTxHash for this operation to prevent multiple trigger

	tx, err := logic.MigrateLikeCoinFromCosmos(
		p.Db,
		p.EthClient,
		p.CosmosAPI,
		p.EthNetworkPublicRPCURL,
		p.EthWalletPrivateKey,
		p.EthTokenAddress,
		data.CosmosTxHash,
	)
	if err != nil {
		p.write(w, http.StatusInternalServerError, &ResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	p.write(w, http.StatusAccepted, &ResponseBody{
		EthTxHash: tx.Hash().Hex(),
	})
}

func (h *InitLikeCoinMigrationFromCosmosHandler) write(w http.ResponseWriter, statusCode int, body *ResponseBody) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := encoder.Encode(body)
	if err != nil {
		panic(err)
	}
}
