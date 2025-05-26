package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/model"

	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type ListLikeCoinMigrationRequestParam struct {
	Limit   int                            `json:"limit"`
	Offset  int                            `json:"offset"`
	Status  *model.LikeCoinMigrationStatus `json:"status"`
	Keyword string                         `json:"q"`
}

type ListLikeCoinMigrationResponse struct {
	Migrations []*api_model.LikeCoinMigration `json:"migrations"`
}

type ListLikeCoinMigrationHandler struct {
	Db *sql.DB
}

func (p *ListLikeCoinMigrationHandler) parseParams(r *http.Request) (*ListLikeCoinMigrationRequestParam, error) {
	var params ListLikeCoinMigrationRequestParam = ListLikeCoinMigrationRequestParam{
		Limit:   10,
		Offset:  0,
		Status:  nil,
		Keyword: "",
	}

	query := r.URL.Query()

	limitStr := query.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, errors.New("Invalid request params")
	}
	params.Limit = limit

	if offsetStr := query.Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, err
		}
		params.Offset = offset
	}

	if statusStr := query.Get("status"); statusStr != "" {
		status := model.LikeCoinMigrationStatus(statusStr)
		if !status.IsValid() {
			return nil, errors.New("Invalid request params")
		}
		params.Status = &status
	}

	if keyword := query.Get("q"); keyword != "" {
		params.Keyword = keyword
	}

	return &params, nil
}

func (p *ListLikeCoinMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	params, err := p.parseParams(r)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, handler.MakeErrorResponseBody(err))
		return
	}

	migrations, err := p.handle(params)
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, &ListLikeCoinMigrationResponse{
		Migrations: migrations,
	})
}

func (p *ListLikeCoinMigrationHandler) handle(params *ListLikeCoinMigrationRequestParam) ([]*api_model.LikeCoinMigration, error) {

	m, err := db.QueryPaginatedLikeCoinMigration(p.Db, params.Limit, params.Offset, params.Status, params.Keyword)

	if err != nil {
		return nil, err
	}

	return api_model.LikeCoinMigrationsFromModel(m), nil
}
