package likenft

import (
	"database/sql"
	"errors"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func MigrateNFTFromAssetMigration(
	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.LikeNFTClass,

	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,
	initialBatchMintOwner string,

	assetMigrationNFTId uint64,
) (*model.LikeNFTAssetMigrationNFT, error) {
	mn, err := appdb.QueryLikeNFTAssetMigrationNFTById(db, assetMigrationNFTId)
	if err != nil {
		return nil, err
	}

	mn.Status = model.LikeNFTAssetMigrationNFTStatusInProgress
	err = appdb.UpdateLikeNFTAssetMigrationNFT(db, mn)
	if err != nil {
		return nil, err
	}

	m, err := appdb.QueryLikeNFTAssetMigrationById(db, mn.LikeNFTAssetMigrationId)
	if err != nil {
		return nil, migrateNFTFromAssetMigrationFailed(db, mn, err)
	}
	defer RecalculateMigrationStatus(db, m.Id)

	lastAction, err := MigrateNFT(
		db,
		c,
		p,
		n,
		initialClassOwner,
		initialClassMinter,
		initialClassUpdater,
		initialBatchMintOwner,
		mn.CosmosClassId,
		mn.CosmosNFTId,
		m.EthAddress,
	)

	if err != nil {
		return nil, migrateNFTFromAssetMigrationFailed(db, mn, err)
	}

	mn.EvmTxHash = lastAction.EvmTxHash
	mn.Status = model.LikeNFTAssetMigrationNFTStatusCompleted
	finishTime := time.Now().UTC()
	mn.FinishTime = &finishTime
	err = appdb.UpdateLikeNFTAssetMigrationNFT(db, mn)

	if err != nil {
		return nil, migrateNFTFromAssetMigrationFailed(db, mn, err)
	}

	return mn, nil
}

func MigrateNFT(
	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.LikeNFTClass,

	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,
	initialBatchMintOwner string,

	cosmosClassId string,
	cosmosNFTId string,
	evmOwner string,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	newClassAction, err := GetOrCreateNewClassAction(
		db,
		cosmosClassId,
		initialClassOwner,
		initialClassMinter,
		initialClassUpdater,
	)
	if err != nil {
		return nil, err
	}

	newClassAction, err = DoNewClassAction(db, c, p, newClassAction)
	if err != nil {
		return nil, err
	}

	mintNFTAction, err := GetOrCreateMintNFTAction(
		db,
		*newClassAction.EvmClassId,
		cosmosNFTId,
		initialBatchMintOwner,
		evmOwner,
	)

	if err != nil {
		return nil, err
	}
	mintNFTAction, err = DoMintNFTAction(db, p, n, mintNFTAction)

	if err != nil {
		return nil, err
	}
	return mintNFTAction, nil
}

func migrateNFTFromAssetMigrationFailed(
	db *sql.DB,
	mc *model.LikeNFTAssetMigrationNFT,
	err error,
) error {
	mc.Status = model.LikeNFTAssetMigrationNFTStatusFailed
	failedReason := err.Error()
	mc.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTAssetMigrationNFT(db, mc))
}
