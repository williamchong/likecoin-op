package migration

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateMigrationRequestBody struct {
	AssetSnapshotId uint64 `json:"asset_snapshot_id"`
	CosmosAddress   string `json:"cosmos_address"`
	EthAddress      string `json:"eth_address"`
}

type CreateMigrationResponseBody struct {
	Migration        *api_model.LikeNFTAssetMigration `json:"migration,omitempty"`
	ErrorDescription string                           `json:"error_description,omitempty"`
}

type CreateMigrationHandler struct {
	Db *sql.DB
}

var ErrMigrationExists = fmt.Errorf("error migration exists")
var ErrSignedEthAddressNotMatch = fmt.Errorf("error signed eth address not match")

func (h *CreateMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data CreateMigrationRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &CreateMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	migration, err := h.handle(data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &CreateMigrationResponseBody{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusCreated, &CreateMigrationResponseBody{
		Migration: migration,
	})

	// Trigger job
	// TODO: use job queue
}

func (h *CreateMigrationHandler) handle(req CreateMigrationRequestBody) (*api_model.LikeNFTAssetMigration, error) {
	signingMessage, err := db.QueryNFTSigningMessageByCosmosAddress(h.Db, req.CosmosAddress)

	if err != nil {
		return nil, err
	}

	if signingMessage.EthAddress != req.EthAddress {
		return nil, ErrSignedEthAddressNotMatch
	}

	_, err = db.QueryLikeNFTAssetMigrationByCosmosAddress(h.Db, req.CosmosAddress)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			classes, err := db.QueryLikeNFTAssetSnapshotClassesByNFTSnapshotId(h.Db, req.AssetSnapshotId)

			if err != nil {
				return nil, err
			}

			nfts, err := db.QueryLikeNFTAssetSnapshotNFTsByNFTSnapshotId(h.Db, req.AssetSnapshotId)

			if err != nil {
				return nil, err
			}

			tx, err := h.Db.Begin()
			if err != nil {
				return nil, err
			}
			defer tx.Commit()

			migration, err := db.InsertLikeNFTAssetMigration(tx, &model.LikeNFTAssetMigration{
				LikeNFTAssetSnapshotId: req.AssetSnapshotId,
				CosmosAddress:          req.CosmosAddress,
				EthAddress:             req.EthAddress,
				Status:                 model.NFTMigrationStatusInit,
			})
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			migrationClasses := make([]model.LikeNFTAssetMigrationClass, 0)

			for _, class := range classes {
				migrationClasses = append(migrationClasses, model.LikeNFTAssetMigrationClass{
					LikeNFTAssetMigrationId: migration.Id,
					CosmosClassId:           class.CosmosClassId,
					Name:                    class.Name,
					Image:                   class.Image,
					Status:                  model.LikeNFTAssetMigrationClassStatusInit,
				})
			}

			migrationNFTs := make([]model.LikeNFTAssetMigrationNFT, 0)

			for _, nft := range nfts {
				migrationNFTs = append(migrationNFTs, model.LikeNFTAssetMigrationNFT{
					LikeNFTAssetMigrationId: migration.Id,
					CosmosClassId:           nft.CosmosClassId,
					CosmosNFTId:             nft.CosmosNFTId,
					Name:                    nft.Name,
					Image:                   nft.Image,
					Status:                  model.LikeNFTAssetMigrationNFTStatusInit,
				})
			}

			err = db.InsertLikeNFTAssetMigrationClasses(tx, migrationClasses)

			if err != nil {
				tx.Rollback()
				return nil, err
			}

			err = db.InsertLikeNFTAssetMigrationNFTs(tx, migrationNFTs)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			return api_model.LikeNFTAssetMigrationFromModel(migration, migrationClasses, migrationNFTs), nil
		} else {
			return nil, err
		}
	}

	return nil, ErrMigrationExists
}
