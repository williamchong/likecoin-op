package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeMigrateNFT = "migrate_nft"

type MigrateNFTPayload struct {
	LikenftAssetMigrationNFTId uint64
}

func NewMigrateNFTTask(likenftAssetMigrationNFTId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(MigrateNFTPayload{
		LikenftAssetMigrationNFTId: likenftAssetMigrationNFTId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeMigrateNFT, payload), nil
}
