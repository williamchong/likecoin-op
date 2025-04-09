package migration

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/model"

	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
)

type ListLikeNFTAssetMigrationRequestParam struct {
	Limit   int                                `json:"limit"`
	Offset  int                                `json:"offset"`
	Status  *model.LikeNFTAssetMigrationStatus `json:"status"`
	Keyword string                             `json:"q"`
}

type ListLikeNFTAssetMigrationResponse struct {
	Migrations       []*api_model.LikeNFTAssetMigrationBase `json:"migrations"`
	ErrorDescription string                                 `json:"error_description"`
}

type ListLikeNFTAssetMigrationHandler struct {
	Db *sql.DB
}

func (p *ListLikeNFTAssetMigrationHandler) parseParams(r *http.Request) (*ListLikeNFTAssetMigrationRequestParam, error) {
	var params ListLikeNFTAssetMigrationRequestParam = ListLikeNFTAssetMigrationRequestParam{
		Limit:  10,
		Offset: 0,
		Status: nil,
	}

	query := r.URL.Query()

	limitStr := query.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, errors.New("invalid request params")
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
		status := model.LikeNFTAssetMigrationStatus(statusStr)
		if !status.IsValid() {
			return nil, errors.New("invalid request params")
		}
		params.Status = &status
	}

	if keyword := query.Get("q"); keyword != "" {
		params.Keyword = keyword
	}

	return &params, nil
}

func (p *ListLikeNFTAssetMigrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params, err := p.parseParams(r)
	if err != nil {
		handler.SendJSON(w, http.StatusBadRequest, &ListLikeNFTAssetMigrationResponse{
			ErrorDescription: err.Error(),
		})
		return
	}

	migrations, err := p.handle(params)
	if err != nil {
		handler.SendJSON(w, http.StatusInternalServerError, &ListLikeNFTAssetMigrationResponse{
			ErrorDescription: err.Error(),
		})
		return
	}

	handler.SendJSON(w, http.StatusOK, &ListLikeNFTAssetMigrationResponse{
		Migrations: migrations,
	})
}

func (p *ListLikeNFTAssetMigrationHandler) handle(params *ListLikeNFTAssetMigrationRequestParam) ([]*api_model.LikeNFTAssetMigrationBase, error) {
	m, err := db.QueryPaginatedLikeNFTAssetMigration(p.Db, params.Limit, params.Offset, params.Status, params.Keyword)

	if err != nil {
		return nil, err
	}

	return api_model.LikeNFTAssetMigrationBasesFromModel(m), nil
}
