package migration_preview

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	api_model "github.com/likecoin/like-migration-backend/pkg/handler/model"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type GetMigrationPreviewResponseBody struct {
	Preview *api_model.LikeNFTAssetSnapshot `json:"preview,omitempty"`
}

type GetMigrationPreviewHandler struct {
	Db *sql.DB
}

func (h *GetMigrationPreviewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())

	cosmosAddress := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	snapshot, err := h.handle(cosmosAddress)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handler.SendJSON(w, http.StatusNotFound,
				handler.MakeErrorResponseBody(err).AsError(handler.ErrNotFound))
			return
		}
		handler.SendJSON(w, http.StatusInternalServerError,
			handler.MakeErrorResponseBody(err).
				WithSentryReported(hub.CaptureException(err)).
				AsError(handler.ErrSomethingWentWrong),
		)
		return
	}

	handler.SendJSON(w, http.StatusOK, &GetMigrationPreviewResponseBody{
		Preview: snapshot,
	})
}

func (h *GetMigrationPreviewHandler) handle(cosmosAddress string) (*api_model.LikeNFTAssetSnapshot, error) {
	snapshot, err := db.QueryLatestLikeNFTAssetSnapshotByCosmosAddress(h.Db, cosmosAddress)

	if err != nil {
		return nil, err
	}

	classes, err := db.QueryLikeNFTAssetSnapshotClassesByNFTSnapshotId(h.Db, snapshot.Id)

	if err != nil {
		return nil, err
	}

	nfts, err := db.QueryLikeNFTAssetSnapshotNFTsByNFTSnapshotId(h.Db, snapshot.Id)

	if err != nil {
		return nil, err
	}

	cs := make([]api_model.LikeNFTAssetSnapshotClass, 0)
	ns := make([]api_model.LikeNFTAssetSnapshotNFT, 0)

	for _, class := range classes {
		cs = append(cs, *h.mapNFTSnapshotClass(&class))
	}

	for _, nft := range nfts {
		ns = append(ns, *h.mapNFTSnapshotNFT(&nft))
	}

	return &api_model.LikeNFTAssetSnapshot{
		Id:            snapshot.Id,
		CreatedAt:     snapshot.CreatedAt,
		CosmosAddress: snapshot.CosmosAddress,
		BlockHeight:   snapshot.BlockHeight,
		BlockTime:     snapshot.BlockTime,
		Status:        snapshot.Status,
		FailedReason:  snapshot.FailedReason,
		Classes:       cs,
		NFTs:          ns,
	}, nil
}

func (h *GetMigrationPreviewHandler) mapNFTSnapshotClass(c *model.LikeNFTAssetSnapshotClass) *api_model.LikeNFTAssetSnapshotClass {
	return &api_model.LikeNFTAssetSnapshotClass{
		Id:            c.Id,
		NFTSnapshotId: c.NFTSnapshotId,
		CreatedAt:     c.CreatedAt,
		CosmosClassId: c.CosmosClassId,
		Name:          c.Name,
		Image:         c.Image,
	}
}

func (h *GetMigrationPreviewHandler) mapNFTSnapshotNFT(n *model.LikeNFTAssetSnapshotNFT) *api_model.LikeNFTAssetSnapshotNFT {
	return &api_model.LikeNFTAssetSnapshotNFT{
		Id:            n.Id,
		NFTSnapshotId: n.NFTSnapshotId,
		CreatedAt:     n.CreatedAt,
		CosmosClassId: n.CosmosClassId,
		CosmosNFTId:   n.CosmosNFTId,
		Name:          n.Name,
		Image:         n.Image,
	}
}
