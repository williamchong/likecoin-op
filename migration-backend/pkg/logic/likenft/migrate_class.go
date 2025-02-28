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

func MigrateClassFromAssetMigration(
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.LikeNFTClass,

	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,

	assetMigrationClassId uint64,
) (*model.LikeNFTAssetMigrationClass, error) {
	mylogger := logger.
		WithGroup("MigrateClassFromAssetMigration").
		With("assetMigrationClassId", assetMigrationClassId)

	mc, err := appdb.QueryLikeNFTAssetMigrationClassById(db, assetMigrationClassId)
	if err != nil {
		return nil, err
	}

	mc.Status = model.LikeNFTAssetMigrationClassStatusInProgress
	err = appdb.UpdateLikeNFTAssetMigrationClass(db, mc)
	if err != nil {
		return nil, err
	}

	m, err := appdb.QueryLikeNFTAssetMigrationById(db, mc.LikeNFTAssetMigrationId)
	if err != nil {
		return nil, migrateClassFromAssetMigrationFailed(db, mc, err)
	}

	defer RecalculateMigrationStatus(db, m.Id)

	lastAction, err := MigrateClass(
		mylogger,
		db,
		c,
		p,
		n,
		mc.CosmosClassId,
		initialClassOwner,
		initialClassMinter,
		initialClassUpdater,
		m.CosmosAddress,
		m.EthAddress,
	)
	if err != nil {
		return nil, migrateClassFromAssetMigrationFailed(db, mc, err)
	}

	mc.EvmTxHash = lastAction.EvmTxHash
	mc.Status = model.LikeNFTAssetMigrationClassStatusCompleted
	finishTime := time.Now().UTC()
	mc.FinishTime = &finishTime

	err = appdb.UpdateLikeNFTAssetMigrationClass(db, mc)
	if err != nil {
		return nil, migrateClassFromAssetMigrationFailed(db, mc, err)
	}

	return mc, nil
}

func MigrateClass(
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	p *evm.LikeProtocol,
	n *evm.LikeNFTClass,

	cosmosClassId string,
	initialClassOwner string,
	initialClassMinter string,
	initialClassUpdater string,

	cosmosOwner string,
	evmOwner string,
) (*model.LikeNFTMigrationActionTransferClass, error) {
	mylogger := logger.
		WithGroup("MigrateClass").
		With("cosmosClassId", cosmosClassId).
		With("initialClassOwner", initialClassOwner).
		With("initialClassMinter", initialClassMinter).
		With("initialClassUpdater", initialClassUpdater).
		With("cosmosOwner", cosmosOwner).
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

	transferClassAction, err := GetOrCreateTransferClassAction(
		db, *newClassAction.EvmClassId, cosmosOwner, evmOwner,
	)
	if err != nil {
		return nil, err
	}

	transferClassAction, err = DoTransferClassAction(
		mylogger,
		db,
		c,
		n,
		transferClassAction,
	)
	if err != nil {
		return nil, err
	}
	return transferClassAction, err
}

func migrateClassFromAssetMigrationFailed(
	db *sql.DB,
	mc *model.LikeNFTAssetMigrationClass,
	err error,
) error {
	mc.Status = model.LikeNFTAssetMigrationClassStatusFailed
	failedReason := err.Error()
	mc.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTAssetMigrationClass(db, mc))
}
