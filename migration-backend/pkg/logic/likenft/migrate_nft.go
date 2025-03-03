package likenft

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func MigrateNFTFromAssetMigration(
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.BookNFT,

	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,
	initialBatchMintOwner string,

	assetMigrationNFTId uint64,
) (*model.LikeNFTAssetMigrationNFT, error) {
	mylogger := logger.
		WithGroup("MigrateNFTFromAssetMigration").
		With("assetMigrationNFTId", assetMigrationNFTId)

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
		mylogger,
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
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.BookNFT,

	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,
	initialBatchMintOwner string,

	cosmosClassId string,
	cosmosNFTId string,
	evmOwner string,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	mylogger := logger.
		WithGroup("MigrateNFT").
		With("initialClassOwner", initialClassOwner).
		With("initialClassMinter", initialClassMinter).
		With("initialClassUpdater", initialClassUpdater).
		With("initialBatchMintOwner", initialBatchMintOwner).
		With("cosmosClassId", cosmosClassId).
		With("cosmosNFTId", cosmosNFTId).
		With("evmOwner", evmOwner)

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

	newClassAction, err = DoNewClassAction(
		mylogger,
		db,
		c,
		p,
		newClassAction,
	)
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
	mintNFTAction, err = DoMintNFTAction(
		mylogger,
		db,
		p,
		n,
		mintNFTAction,
	)

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
