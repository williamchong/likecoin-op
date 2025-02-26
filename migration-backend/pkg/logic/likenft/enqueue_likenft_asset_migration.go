package likenft

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/hibiken/asynq"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

func EnqueueLikeNFTAssetMigration(
	logger *slog.Logger,
	db *sql.DB,
	asynqClient *asynq.Client,
	likenftMigrationId uint64,
) error {
	mylogger := logger.
		WithGroup("EnqueueLikeNFTAssetMigration").
		With("likenftMigrationId", likenftMigrationId)

	classes, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(db, likenftMigrationId)
	if err != nil {
		return err
	}

	nfts, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(db, likenftMigrationId)
	if err != nil {
		return err
	}

	for _, class := range classes {
		mylogger := mylogger.WithGroup("class")
		tx, err := db.Begin()
		if err != nil {
			continue
		}

		now := time.Now().UTC()
		class.EnqueueTime = &now
		err = appdb.UpdateLikeNFTAssetMigrationClass(tx, &class)
		if err != nil {
			_ = tx.Rollback()
			continue
		}

		task, err := task.NewMigrateClassTask(class.Id)
		if err != nil {
			_ = tx.Rollback()
			continue
		}
		taskInfo, err := asynqClient.Enqueue(task, asynq.MaxRetry(0))
		if err != nil {
			_ = tx.Rollback()
			continue
		}

		_ = tx.Commit()
		mylogger.Info(
			"NewMigrateClassTask enqueued",
			"CosmosClassId",
			class.CosmosClassId,
			"taskId",
			taskInfo.ID,
			"queue",
			taskInfo.Queue,
		)
	}

	for _, nft := range nfts {
		mylogger := mylogger.WithGroup("nft")
		tx, err := db.Begin()
		if err != nil {
			continue
		}

		now := time.Now().UTC()
		nft.EnqueueTime = &now
		err = appdb.UpdateLikeNFTAssetMigrationNFT(tx, &nft)
		if err != nil {
			_ = tx.Rollback()
			continue
		}

		task, err := task.NewMigrateNFTTask(nft.Id)
		if err != nil {
			_ = tx.Rollback()
			continue
		}
		taskInfo, err := asynqClient.Enqueue(task, asynq.MaxRetry(0))
		if err != nil {
			_ = tx.Rollback()
			continue
		}
		_ = tx.Commit()
		mylogger.Info(
			"NewMigrateNFTTask enqueued",
			"CosmosClassId",
			nft.CosmosClassId,
			"CosmosNFTId",
			nft.CosmosNFTId,
			"taskId",
			taskInfo.ID,
			"queue",
			taskInfo.Queue,
		)
	}

	return nil
}
