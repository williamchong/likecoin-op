package migration_preview

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type CreateMigrationPreviewRequestBody struct {
	CosmosAddress string `json:"cosmos_address"`
}

type CreateMigrationPreviewResponseBody struct {
	Preview *api_model.LikeNFTAssetSnapshot `json:"preview,omitempty"`
}

type CreateMigrationPreviewHandler struct {
	Db                  *sql.DB
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
	LikerlandUrlBase    string
}

var ErrLatestSnapshotInProgress = fmt.Errorf("error latest snapshot in progress")

func (h *CreateMigrationPreviewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	decoder := json.NewDecoder(r.Body)
	var data CreateMigrationPreviewRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest,
			handler.MakeErrorResponseBody(err))
		return
	}

	snapshot, err := h.handle(data)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusCreated, &CreateMigrationPreviewResponseBody{
		Preview: snapshot,
	})

	// Trigger job
	// TODO: use job queue
	go func() {
		logic := likenft.SnapshotCosmosStateLogic{
			DB:                  h.Db,
			CosmosAPI:           h.CosmosAPI,
			LikeNFTCosmosClient: h.LikeNFTCosmosClient,
		}
		err := logic.Execute(context.Background(), snapshot.CosmosAddress)
		if err != nil {
			fmt.Printf("error executing logic SnapshotCosmosStateLogic: %v", err)
		}
	}()
}

func (h *CreateMigrationPreviewHandler) handle(req CreateMigrationPreviewRequestBody) (*api_model.LikeNFTAssetSnapshot, error) {
	latestSnapshot, err := db.QueryLatestLikeNFTAssetSnapshotByCosmosAddress(h.Db, req.CosmosAddress)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			latestSnapshot = &model.LikeNFTAssetSnapshot{
				CosmosAddress: req.CosmosAddress,
				Status:        model.NFTSnapshotStatusInit,
			}
			snapshot, err := db.InsertLikeNFTAssetSnapshot(h.Db, latestSnapshot)
			if err != nil {
				return nil, err
			}
			return h.mapNFTSnapshot(snapshot), err
		} else {
			return nil, err
		}
	}

	if !(latestSnapshot.Status == model.NFTSnapshotStatusCompleted || latestSnapshot.Status == model.NFTSnapshotStatusFailed) {
		return nil, ErrLatestSnapshotInProgress
	}

	snapshot, err := db.InsertLikeNFTAssetSnapshot(h.Db, &model.LikeNFTAssetSnapshot{
		CosmosAddress: req.CosmosAddress,
		Status:        model.NFTSnapshotStatusInit,
	})
	if err != nil {
		return nil, err
	}
	return h.mapNFTSnapshot(snapshot), nil
}

func (h *CreateMigrationPreviewHandler) mapNFTSnapshot(nftSnapshot *model.LikeNFTAssetSnapshot) *api_model.LikeNFTAssetSnapshot {
	return &api_model.LikeNFTAssetSnapshot{
		Id:            nftSnapshot.Id,
		CreatedAt:     nftSnapshot.CreatedAt,
		CosmosAddress: nftSnapshot.CosmosAddress,
		BlockHeight:   nftSnapshot.BlockHeight,
		BlockTime:     nftSnapshot.BlockTime,
		Status:        nftSnapshot.Status,
		FailedReason:  nftSnapshot.FailedReason,
		Classes:       make([]api_model.LikeNFTAssetSnapshotClass, 0),
		NFTs:          make([]api_model.LikeNFTAssetSnapshotNFT, 0),
	}
}
