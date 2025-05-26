package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type GetLikeCoinMigrationResponseBody struct {
	Migration *api_model.LikeCoinMigration `json:"migration,omitempty"`
}

type GetLikeCoinMigrationHandler struct {
	Db *sql.DB
}

func (h *GetLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())
	migrationIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	migrationId, err := strconv.ParseUint(migrationIdStr, 10, 64)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}

	migration, err := h.handle(migrationId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound,
				handler.MakeErrorResponseBody(err).
					AsError(handler.ErrNotFound),
			)
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
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
