package likenft

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/hibiken/asynq"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

func EnqueueFailedLikeNFTAssetMigration(
	logger *slog.Logger,
	db *sql.DB,
	asynqClient *asynq.Client,
	likenftMigrationId uint64,
) error {
	mylogger := logger.
		WithGroup("EnqueueFailedLikeNFTAssetMigration").
		With("likenftMigrationId", likenftMigrationId)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	likenftMigration, err := appdb.QueryLikeNFTAssetMigrationById(tx, likenftMigrationId)
	if err != nil {
		mylogger.Error("failed to query likenft migration", "error", err)
		return errors.Join(err, tx.Rollback())
	}

	failedClasses, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
		tx,
		likenftMigrationId,
		model.LikeNFTAssetMigrationClassStatusFailed,
	)
	if err != nil {
		mylogger.Error("failed to query likenft migration classes", "error", err)
		return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
	}

	failedNFTs, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(
		tx,
		likenftMigrationId,
		model.LikeNFTAssetMigrationNFTStatusFailed,
	)
	if err != nil {
		mylogger.Error("failed to query likenft migration nfts", "error", err)
		return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
	}

	tasks := make([]*asynq.Task, 0, len(failedClasses)+len(failedNFTs))

	for _, failedClass := range failedClasses {
		now := time.Now().UTC()
		failedClass.EnqueueTime = &now
		err = appdb.UpdateLikeNFTAssetMigrationClass(tx, &failedClass)
		if err != nil {
			return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
		}

		task, err := task.NewMigrateClassTask(failedClass.Id)
		if err != nil {
			return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
		}
		tasks = append(tasks, task)
	}

	for _, failedNFT := range failedNFTs {
		now := time.Now().UTC()
		failedNFT.EnqueueTime = &now
		err = appdb.UpdateLikeNFTAssetMigrationNFT(tx, &failedNFT)
		if err != nil {
			return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
		}

		task, err := task.NewMigrateNFTTask(failedNFT.Id)
		if err != nil {
			return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
		}
		tasks = append(tasks, task)
	}

	err = tx.Commit()
	if err != nil {
		return enqueueFailedLikeNFTAssetMigrationFailed(db, likenftMigration, errors.Join(err, tx.Rollback()))
	}

	for _, task := range tasks {
		taskInfo, err := asynqClient.Enqueue(task, asynq.MaxRetry(0))
		if err != nil {
			mylogger.Error("failed to enqueue task", "error", err)
		}
		mylogger.Info(
			"enqueued",
			"taskId",
			taskInfo.ID,
			"queue",
			taskInfo.Queue,
		)
	}

	return nil
}

func enqueueFailedLikeNFTAssetMigrationFailed(
	db *sql.DB,
	a *model.LikeNFTAssetMigration,
	err error,
) error {
	a.Status = model.NFTMigrationStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTAssetMigration(db, a))
}
