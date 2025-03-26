package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeEnqueueFailedLikeNFTAssetMigration = "enqueue_failed_likenft_asset_migration"

type EnqueueFailedLikeNFTAssetMigrationPayload struct {
	LikenftAssetMigrationId uint64
}

func NewEnqueueFailedLikeNFTAssetMigrationTask(likenftAssetMigrationId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(EnqueueFailedLikeNFTAssetMigrationPayload{
		LikenftAssetMigrationId: likenftAssetMigrationId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEnqueueFailedLikeNFTAssetMigration, payload), nil
}
