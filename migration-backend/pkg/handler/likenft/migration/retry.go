package migration

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/hibiken/asynq"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

var (
	ErrMigrationNotFailed = errors.New("migration is not failed")
	ErrAssetsNotMatched   = errors.New("assets not matched")
)

type classPrefixedNFTId string

func newClassPrefixedNFTId(classId string, nftId string) classPrefixedNFTId {
	return classPrefixedNFTId(fmt.Sprintf("%s/%s", classId, nftId))
}

type BookNFT struct {
	ClassId string `json:"class_id"`
	NFTId   string `json:"nft_id"`
}

type RetryMigrationRequestBody struct {
	BookNFTCollection []string  `json:"book_nft_collection"`
	BookNFT           []BookNFT `json:"book_nft"`
}

func (r *RetryMigrationRequestBody) CheckAssets(cs []model.LikeNFTAssetMigrationClass, ns []model.LikeNFTAssetMigrationNFT) bool {
	dbFailedClasses := make(map[string]bool)
	dbFailedNFTs := make(map[classPrefixedNFTId]bool)

	for _, c := range cs {
		if c.Status == model.LikeNFTAssetMigrationClassStatusFailed {
			dbFailedClasses[c.CosmosClassId] = true
		}
	}

	for _, n := range ns {
		if n.Status == model.LikeNFTAssetMigrationNFTStatusFailed {
			dbFailedNFTs[newClassPrefixedNFTId(n.CosmosClassId, n.CosmosNFTId)] = true
		}
	}

	reqFailedClasses := make(map[string]bool)
	reqFailedNFTs := make(map[classPrefixedNFTId]bool)

	for _, c := range r.BookNFTCollection {
		reqFailedClasses[c] = true
	}

	for _, n := range r.BookNFT {
		reqFailedNFTs[newClassPrefixedNFTId(n.ClassId, n.NFTId)] = true
	}

	return reflect.DeepEqual(reqFailedClasses, dbFailedClasses) && reflect.DeepEqual(reqFailedNFTs, dbFailedNFTs)
}

type RetryMigrationResponseBody struct {
	Migration *api_model.LikeNFTAssetMigration `json:"migration,omitempty"`
}

type RetryMigrationHandler struct {
	Db          *sql.DB
	AsynqClient *asynq.Client
}

func (h *RetryMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hub := sentry.GetHubFromContext(ctx)

	cosmosAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	decoder := json.NewDecoder(r.Body)
	var req RetryMigrationRequestBody
	err := decoder.Decode(&req)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}

	migration, err := h.handle(ctx, cosmosAddress, &req)

	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong))
		return
	}

	handler.SendJSON(w, http.StatusOK, &RetryMigrationResponseBody{
		Migration: migration,
	})

	go h.enqueueFailedMigrationTasks(migration.Id)
}

func (h *RetryMigrationHandler) handle(ctx context.Context, cosmosAddress string, req *RetryMigrationRequestBody) (*api_model.LikeNFTAssetMigration, error) {
	migration, err := db.QueryLikeNFTAssetMigrationByCosmosAddress(h.Db, cosmosAddress)

	if err != nil {
		return nil, err
	}

	if migration.Status != model.NFTMigrationStatusFailed {
		return nil, ErrMigrationNotFailed
	}

	// Check assets
	failedClasses, err := db.QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
		h.Db,
		migration.Id,
		model.LikeNFTAssetMigrationClassStatusFailed,
	)
	if err != nil {
		return nil, err
	}

	failedNFTs, err := db.QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(
		h.Db,
		migration.Id,
		model.LikeNFTAssetMigrationNFTStatusFailed,
	)
	if err != nil {
		return nil, err
	}

	if !req.CheckAssets(failedClasses, failedNFTs) {
		return nil, ErrAssetsNotMatched
	}

	pendingEstimatedDurationFromMigrationClasses, err := db.QueryTotalPendingEstimatedDurationFromMigrationClasses(ctx, h.Db)
	if err != nil {
		return nil, err
	}
	pendingEstimatedDurationFromMigrationNFTs, err := db.QueryTotalPendingEstimatedDurationFromMigrationNFTs(ctx, h.Db)
	if err != nil {
		return nil, err
	}

	pendingEstimatedDurationFromSnapshotClasses := time.Duration(0)
	for _, failedClass := range failedClasses {
		pendingEstimatedDurationFromSnapshotClasses += failedClass.EstimatedDurationNeeded
	}
	pendingEstimatedDurationFromSnapshotNFTs := time.Duration(0)
	for _, failedNFT := range failedNFTs {
		pendingEstimatedDurationFromSnapshotNFTs += failedNFT.EstimatedDurationNeeded
	}

	// Return
	classes, err := db.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(h.Db, migration.Id)
	if err != nil {
		return nil, err
	}

	nfts, err := db.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(h.Db, migration.Id)
	if err != nil {
		return nil, err
	}

	migration.Status = model.NFTMigrationStatusInProgress
	migration.EstimatedFinishedTime = time.Now().UTC().Add(
		pendingEstimatedDurationFromMigrationClasses +
			pendingEstimatedDurationFromMigrationNFTs +
			pendingEstimatedDurationFromSnapshotClasses +
			pendingEstimatedDurationFromSnapshotNFTs)

	err = db.UpdateLikeNFTAssetMigration(h.Db, migration)
	if err != nil {
		return nil, err
	}

	return api_model.LikeNFTAssetMigrationFromModel(migration, classes, nfts), nil
}

func (h *RetryMigrationHandler) enqueueFailedMigrationTasks(migrationId uint64) error {
	t, err := task.NewEnqueueFailedLikeNFTAssetMigrationTask(migrationId)
	if err != nil {
		return err
	}
	_, err = h.AsynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		return err
	}
	return nil
}
