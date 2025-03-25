package likenft

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func MigrateNFTFromAssetMigration(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
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
		ctx,
		mylogger,
		db,
		c,
		likecoinAPI,
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
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
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
		ctx,
		mylogger,
		db,
		c,
		likecoinAPI,
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
		ctx,
		mylogger,
		db,
		p,
		n,
		c,
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
