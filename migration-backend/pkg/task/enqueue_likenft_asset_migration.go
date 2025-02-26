package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeEnqueueLikeNFTAssetMigration = "enqueue_likenft_asset_migration"

type EnqueueLikeNFTAssetMigrationPayload struct {
	LikenftAssetMigrationClassId uint64
}

func NewEnqueueLikeNFTAssetMigrationTask(likenftAssetMigrationClassId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(EnqueueLikeNFTAssetMigrationPayload{
		LikenftAssetMigrationClassId: likenftAssetMigrationClassId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEnqueueLikeNFTAssetMigration, payload), nil
}
