package migration

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hibiken/asynq"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

type UpdateLikeCoinMigrationCosmosHandlerRequestBody struct {
	CosmosTxHash string `json:"cosmos_tx_hash"`
}

type UpdateLikeCoinMigrationCosmosHandlerResponseBody struct {
	Migration        *model.LikeCoinMigration `json:"migration,omitempty"`
	ErrorDescription string                   `json:"error_description,omitempty"`
}

type UpdateLikeCoinMigrationCosmosHandler struct {
	Db          *sql.DB
	AsynqClient *asynq.Client
}

func (p *UpdateLikeCoinMigrationCosmosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	cosmosAddress := paths[len(paths)-2]

	decoder := json.NewDecoder(r.Body)
	var req UpdateLikeCoinMigrationCosmosHandlerRequestBody
	err := decoder.Decode(&req)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &UpdateLikeCoinMigrationCosmosHandlerResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	m, err := p.handle(cosmosAddress, &req)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &UpdateLikeCoinMigrationCosmosHandlerResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &UpdateLikeCoinMigrationCosmosHandlerResponseBody{
		Migration: m,
	})

	// TODO enqueue a job to continue migration
	go p.enqueue(m.UserCosmosAddress)
}

func (p *UpdateLikeCoinMigrationCosmosHandler) handle(cosmosAddress string, req *UpdateLikeCoinMigrationCosmosHandlerRequestBody) (*model.LikeCoinMigration, error) {
	m, err := likecoin.UpdateCosmosTxHash(p.Db, cosmosAddress, req.CosmosTxHash)

	if err != nil {
		return nil, err
	}

	return model.LikeCoinMigrationFromModel(m), nil
}

func (p *UpdateLikeCoinMigrationCosmosHandler) enqueue(cosmosAddress string) error {
	t, err := task.NewMigrateLikeCoinTask(cosmosAddress)
	if err != nil {
		return err
	}
	_, err = p.AsynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		return err
	}
	return nil
}
