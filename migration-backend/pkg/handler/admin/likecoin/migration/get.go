package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"strconv"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type GetLikeCoinMigrationResponseBody struct {
	Migration        *api_model.LikeCoinMigration `json:"migration,omitempty"`
	ErrorDescription string                       `json:"error_description,omitempty"`
}

type GetLikeCoinMigrationHandler struct {
	Db *sql.DB
}

func (h *GetLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	migrationIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	
	migrationId, err := strconv.ParseUint(migrationIdStr, 10, 64)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &GetLikeCoinMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	migration, err := h.handle(migrationId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound, &GetLikeCoinMigrationResponseBody{
				ErrorDescription: "Not Found",
			})
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError, &GetLikeCoinMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetLikeCoinMigrationResponseBody{
		Migration: migration,
	})
}

func (h *GetLikeCoinMigrationHandler) handle(migrationId uint64) (*api_model.LikeCoinMigration, error) {
	m, err := db.QueryLikeCoinMigrationById(h.Db, migrationId)

	if err != nil {
		return nil, err
	}

	return api_model.LikeCoinMigrationFromModel(m), nil
}
