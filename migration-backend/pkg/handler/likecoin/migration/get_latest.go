package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type GetLatestLikeCoinMigrationResponseBody struct {
	Migration        *api_model.LikeCoinMigration `json:"migration,omitempty"`
	ErrorDescription string                       `json:"error_description,omitempty"`
}

type GetLatestLikeCoinMigrationHandler struct {
	Db *sql.DB
}

func (h *GetLatestLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	migration, err := h.handle(cosmosAddress)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound, &GetLatestLikeCoinMigrationResponseBody{
				ErrorDescription: "Not Found",
			})
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError, &GetLatestLikeCoinMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetLatestLikeCoinMigrationResponseBody{
		Migration: migration,
	})
}

func (h *GetLatestLikeCoinMigrationHandler) handle(cosmosAddress string) (*api_model.LikeCoinMigration, error) {
	m, err := db.QueryLatestLikeCoinMigration(h.Db, cosmosAddress)

	if err != nil {
		return nil, err
	}

	return api_model.LikeCoinMigrationFromModel(m), nil
}
