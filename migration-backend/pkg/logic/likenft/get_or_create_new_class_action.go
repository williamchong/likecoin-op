package likenft

import (
	"database/sql"
	"errors"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func GetOrCreateNewClassAction(
	db *sql.DB,
	cosmosClassId string,
	initialOwner string,
	initialMinter string,
	initialUpdater string,
) (*model.LikeNFTMigrationActionNewClass, error) {
	m, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		CosmosClassId: &cosmosClassId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			m = &model.LikeNFTMigrationActionNewClass{
				CosmosClassId:  cosmosClassId,
				InitialOwner:   initialOwner,
				InitialMinter:  initialMinter,
				InitialUpdater: initialUpdater,
				Status:         model.LikeNFTMigrationActionNewClassStatusInit,
			}
			m, err = appdb.InsertLikeNFTMigrationActionNewClass(db, m)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return m, nil
}
