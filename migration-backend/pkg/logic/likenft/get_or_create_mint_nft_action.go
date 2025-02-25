package likenft

import (
	"database/sql"
	"errors"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func GetOrCreateMintNFTAction(
	db *sql.DB,
	evmClassId string,
	cosmosNFTId string,
	initialBatchMintOwner string,
	evmOwner string,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	m, err := appdb.QueryLikeNFTMigrationActionMintNFT(db, appdb.QueryLikeNFTMigrationActionMintNFTFilter{
		EvmClassId:  &evmClassId,
		CosmosNFTId: &cosmosNFTId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			m = &model.LikeNFTMigrationActionMintNFT{
				EvmClassId:            evmClassId,
				CosmosNFTId:           cosmosNFTId,
				EvmOwner:              evmOwner,
				Status:                model.LikeNFTMigrationActionMintNFTStatusInit,
				InitialBatchMintOwner: initialBatchMintOwner,
			}
			m, err = appdb.InsertLikeNFTMigrationActionMintNFT(db, m)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return m, nil
}
