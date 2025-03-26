package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

func HandleEnqueueFailedLikeNFTAssetMigration(ctx context.Context, t *asynq.Task) error {
	db := appcontext.DBFromContext(ctx)
	asynqClient := appcontext.AsynqClientFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	var p task.EnqueueFailedLikeNFTAssetMigrationPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return likenft.EnqueueFailedLikeNFTAssetMigration(logger, db, asynqClient, p.LikenftAssetMigrationId)
}
