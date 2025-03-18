package likecoin

import (
	"database/sql"
	"fmt"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

var ErrMigrationNotPendingCosmosTxHash = fmt.Errorf("err likecoin migration not pending cosmos tx hash")

func UpdateCosmosTxHash(
	db *sql.DB,
	cosmosAddress string,
	cosmosTxHash string,
) (*model.LikeCoinMigration, error) {
	latestMigration, err := appdb.QueryLatestLikeCoinMigration(db, cosmosAddress)

	if err != nil {
		return nil, err
	}

	if latestMigration.Status != model.LikeCoinMigrationStatusPendingCosmosTxHash {
		return nil, ErrMigrationNotPendingCosmosTxHash
	}

	latestMigration.CosmosTxHash = &cosmosTxHash
	latestMigration.Status = model.LikeCoinMigrationStatusVerifyingCosmosTx

	err = appdb.UpdateLikeCoinMigration(db, latestMigration)

	if err != nil {
		return nil, err
	}

	return latestMigration, nil
}
