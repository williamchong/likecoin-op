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

type GetLikeNFTAssetMigrationResponseBody struct {
	Migration        *api_model.LikeNFTAssetMigration `json:"migration,omitempty"`
	Snapshot         *api_model.LikeNFTAssetSnapshot  `json:"snapshot,omitempty"`
	ErrorDescription string                           `json:"error_description,omitempty"`
}

type GetLikeNFTAssetMigrationHandler struct {
	Db *sql.DB
}

func (h *GetLikeNFTAssetMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cosmosAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	migration, snapshot, err := h.handle(cosmosAddress)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound, &GetLikeNFTAssetMigrationResponseBody{
				ErrorDescription: "Not Found",
			})
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError, &GetLikeNFTAssetMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetLikeNFTAssetMigrationResponseBody{
		Migration: migration,
		Snapshot:  snapshot,
	})
}

func (h *GetLikeNFTAssetMigrationHandler) handle(
	cosmosAddress string,
) (
	*api_model.LikeNFTAssetMigration,
	*api_model.LikeNFTAssetSnapshot,
	error,
) {
	migration, err := db.QueryLikeNFTAssetMigrationByCosmosAddress(h.Db, cosmosAddress)

	if err != nil {
		return nil, nil, err
	}

	classes, err := db.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(h.Db, migration.Id)

	if err != nil {
		return nil, nil, err
	}

	nfts, err := db.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(h.Db, migration.Id)

	if err != nil {
		return nil, nil, err
	}

	snapshot, err := db.QueryLikeNFTAssetSnapshotById(h.Db, migration.LikeNFTAssetSnapshotId)

	if err != nil {
		return nil, nil, err
	}

	sClasses, err := db.QueryLikeNFTAssetSnapshotClassesByNFTSnapshotId(h.Db, snapshot.Id)

	if err != nil {
		return nil, nil, err
	}

	sNFTs, err := db.QueryLikeNFTAssetSnapshotNFTsByNFTSnapshotId(h.Db, snapshot.Id)

	if err != nil {
		return nil, nil, err
	}

	return api_model.LikeNFTAssetMigrationFromModel(migration, classes, nfts),
		api_model.LikeNFTAssetSnapshotFromModel(snapshot, sClasses, sNFTs),
		nil
}
