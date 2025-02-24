package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeMigrateClass = "migrate_class"

type MigrateClassPayload struct {
	LikenftAssetMigrationClassId uint64
}

func NewMigrateClassTask(likenftAssetMigrationClassId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(MigrateClassPayload{
		LikenftAssetMigrationClassId: likenftAssetMigrationClassId,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeMigrateClass, payload), nil
}
