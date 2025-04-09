package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type RemoveLikeNFTAssetMigrationRequestBody struct {
	MigrationId uint64 `json:"migration_id"`
}

type RemoveLikeNFTAssetMigrationResponseBody struct {
	Migration        *api_model.LikeNFTAssetMigration `json:"migration,omitempty"`
	ErrorDescription string                           `json:"error_description,omitempty"`
}

type RemoveLikeNFTAssetMigrationHandler struct {
	Db *sql.DB
}

func (h *RemoveLikeNFTAssetMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	migrationIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	migrationId, err := strconv.ParseUint(migrationIdStr, 10, 64)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &RemoveLikeNFTAssetMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	migration, err := h.handle(migrationId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound, &RemoveLikeNFTAssetMigrationResponseBody{
				ErrorDescription: "Not Found",
			})
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError, &RemoveLikeNFTAssetMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &RemoveLikeNFTAssetMigrationResponseBody{
		Migration: migration,
	})
}

func (h *RemoveLikeNFTAssetMigrationHandler) handle(migrationId uint64) (*api_model.LikeNFTAssetMigration, error) {
	migration, err := db.QueryLikeNFTAssetMigrationById(h.Db, migrationId)

	if err != nil {
		return nil, err
	}

	classes, err := db.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(h.Db, migration.Id)
	if err != nil {
		return nil, err
	}

	nfts, err := db.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(h.Db, migration.Id)
	if err != nil {
		return nil, err
	}

	tx, err := h.Db.Begin()
	if err != nil {
		return nil, err
	}

	err = db.RemoveLikeNFTAssetMigrationClassByMigrationId(tx, migration.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = db.RemoveLikeNFTAssetMigrationNFTByMigrationId(tx, migration.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = db.RemoveLikeNFTAssetMigration(tx, migration.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return api_model.LikeNFTAssetMigrationFromModel(migration, classes, nfts), nil
}
