package migration

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/hibiken/asynq"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

var (
	ErrMigrationNotFailed = errors.New("migration is not failed")
)

type RetryLikeCoinMigrationResponseBody struct {
	Migration *api_model.LikeCoinMigration `json:"migration,omitempty"`
}

type RetryLikeCoinMigrationHandler struct {
	Db          *sql.DB
	AsynqClient *asynq.Client
}

func (h *RetryLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	migrationIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	migrationId, err := strconv.ParseUint(migrationIdStr, 10, 64)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}

	migration, err := h.handle(r.Context(), migrationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(
				w,
				http.StatusNotFound,
				handler.MakeErrorResponseBody(err).
					AsError(handler.ErrNotFound),
			)
			return
		}
		handler.SendJSON(
			w,
			http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, &RetryLikeCoinMigrationResponseBody{
		Migration: migration,
	})

	go h.enqueueFailedMigrationTask(migration.Id)
}

func (h *RetryLikeCoinMigrationHandler) handle(ctx context.Context, migrationId uint64) (*api_model.LikeCoinMigration, error) {
	m, err := appdb.QueryLikeCoinMigrationById(h.Db, migrationId)
	if err != nil {
		return nil, err
	}

	if m.Status != model.LikeCoinMigrationStatusFailed {
		return nil, ErrMigrationNotFailed
	}

	err = appdb.WithTx(ctx, h.Db, func(tx *sql.Tx) error {
		m.Status = model.LikeCoinMigrationStatusVerifyingCosmosTx
		return appdb.UpdateLikeCoinMigration(tx, m)
	})

	if err != nil {
		return nil, err
	}

	return api_model.LikeCoinMigrationFromModel(m), nil
}

func (h *RetryLikeCoinMigrationHandler) enqueueFailedMigrationTask(migrationId uint64) error {
	t, err := task.NewMigrateLikeCoinTaskWithMigrationId(migrationId)
	if err != nil {
		return err
	}
	_, err = h.AsynqClient.Enqueue(t, asynq.MaxRetry(0))
	if err != nil {
		return err
	}
	return nil
}
