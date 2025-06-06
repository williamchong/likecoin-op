package likenft

import (
	"database/sql"
	"errors"
	"math/big"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/types/commaseparatedstring"
)

func GetOrCreateNewClassAction(
	db *sql.DB,
	cosmosClassId string,
	initialOwner string,
	initialClassMinters []string,
	initialUpdater string,
	initialBatchMintOwner string,
	shouldPremintAllNFTs bool,
	defaultRoyaltyFraction *big.Int,
) (*model.LikeNFTMigrationActionNewClass, error) {
	m, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		CosmosClassId: &cosmosClassId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			m = &model.LikeNFTMigrationActionNewClass{
				CosmosClassId:          cosmosClassId,
				InitialOwner:           initialOwner,
				InitialMintersStr:      commaseparatedstring.FromSlice(initialClassMinters),
				InitialUpdater:         initialUpdater,
				InitialBatchMintOwner:  initialBatchMintOwner,
				ShouldPremintAllNFTs:   shouldPremintAllNFTs,
				DefaultRoyaltyFraction: defaultRoyaltyFraction,
				Status:                 model.LikeNFTMigrationActionNewClassStatusInit,
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
