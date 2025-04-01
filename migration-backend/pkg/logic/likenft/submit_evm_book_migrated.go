package likenft

import (
	"context"
	"database/sql"
	"log/slog"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func SubmitEvmBookMigrated(
	ctx context.Context,
	logger *slog.Logger,
	db *sql.DB,

	likecoinAPI *likecoin_api.LikecoinAPI,

	likeClassId string,
	evmClassId string,
) (*likecoin_api.EvmMigrateBookResponse, error) {
	mylogger := logger.WithGroup("SubmitEvmBookMigrated").
		With("likeClassId", likeClassId).
		With("evmClassId", evmClassId)
	mylogger.Info("submitting evm book migrated")

	request := &likecoin_api.EvmMigrateBookRequest{
		LikeClassID: likeClassId,
		EvmClassID:  evmClassId,
	}
	response, err := likecoinAPI.SubmitEvmBookMigrated(request)
	if err != nil {
		mylogger.Error("failed to submit evm book migrated", "error", err)
		err = submitEvmBookMigratedFailed(db, likeClassId, evmClassId, err)
		return nil, err
	}

	err = submitEvmBookMigratedSuccess(db, likeClassId, evmClassId)
	if err != nil {
		mylogger.Error("failed to submit evm book migrated", "error", err)
		return nil, err
	}
	mylogger.Info("submitted evm book migrated")

	return response, nil
}

func submitEvmBookMigratedSuccess(
	db *sql.DB,
	likeClassId string,
	evmClassId string,
) error {
	submission := &model.LikecoinEvmMigrateBookSubmission{
		LikeClassID: likeClassId,
		EvmClassID:  evmClassId,
		Status:      model.LikecoinEvmMigrateBookSubmissionStatusSuccess,
	}
	_, err := appdb.InsertLikecoinEvmMigrateBookSubmission(db, submission)
	if err != nil {
		return err
	}
	return nil
}

func submitEvmBookMigratedFailed(
	db *sql.DB,
	likeClassId string,
	evmClassId string,
	err error,
) error {
	failedReason := err.Error()
	submission := &model.LikecoinEvmMigrateBookSubmission{
		LikeClassID:  likeClassId,
		EvmClassID:   evmClassId,
		Status:       model.LikecoinEvmMigrateBookSubmissionStatusFailed,
		FailedReason: &failedReason,
	}
	_, err = appdb.InsertLikecoinEvmMigrateBookSubmission(db, submission)
	if err != nil {
		return err
	}
	return nil
}
