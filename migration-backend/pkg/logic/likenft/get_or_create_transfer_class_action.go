package likenft

import (
	"database/sql"
	"errors"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func GetOrCreateTransferClassAction(
	db *sql.DB,
	evmClassId string,
	cosmosOwner string,
	evmOwner string,
) (*model.LikeNFTMigrationActionTransferClass, error) {
	m, err := appdb.QueryLikeNFTMigrationActionTransferClass(db, appdb.QueryLikeNFTMigrationActionTransferClassFilter{
		EvmClassId:  &evmClassId,
		CosmosOwner: &cosmosOwner,
		EvmOwner:    &evmOwner,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			m = &model.LikeNFTMigrationActionTransferClass{
				EvmClassId:  evmClassId,
				CosmosOwner: cosmosOwner,
				EvmOwner:    evmOwner,
				Status:      model.LikeNFTMigrationActionTransferClassStatusInit,
			}
			m, err = appdb.InsertLikeNFTMigrationActionTransferClass(db, m)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return m, nil
}
