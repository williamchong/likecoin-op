package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeMigrateLikeCoin = "migrate_likecoin"

type MigrateLikeCoinPayload struct {
	CosmosAddress string
}

func NewMigrateLikeCoinTask(cosmosAddress string) (*asynq.Task, error) {
	payload, err := json.Marshal(MigrateLikeCoinPayload{
		CosmosAddress: cosmosAddress,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeMigrateLikeCoin, payload), nil
}
