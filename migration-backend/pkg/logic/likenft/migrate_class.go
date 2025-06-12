package likenft

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"math/big"
	"time"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmosnftidclassifier"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func MigrateClassFromAssetMigration(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
	p *evm.LikeProtocol,
	n *evm.BookNFT,
	cosmosNFTIDClassifier cosmosnftidclassifier.CosmosNFTIDClassifier,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	shouldPremintAllNFTs bool,
	premintAllNFTsShouldPremintArbitraryNFTIDs bool,

	initialClassOwner string,
	initialClassMinters []string,
	initialClassUpdaters []string,
	initialBatchMintOwner string,
	batchMintPerPage uint64,
	defaultRoyaltyFraction *big.Int,

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

	lastActionEvmTxHash, err := MigrateClass(
		ctx,
		mylogger,
		db,
		c,
		likecoinAPI,
		p,
		n,
		cosmosNFTIDClassifier,
		erc721ExternalURLBuilder,
		shouldPremintAllNFTs,
		premintAllNFTsShouldPremintArbitraryNFTIDs,
		mc.CosmosClassId,
		initialClassOwner,
		initialClassMinters,
		initialClassUpdaters,
		initialBatchMintOwner,
		batchMintPerPage,
		defaultRoyaltyFraction,
		m.CosmosAddress,
		m.EthAddress,
	)
	if err != nil {
		return nil, migrateClassFromAssetMigrationFailed(db, mc, err)
	}

	mc.EvmTxHash = lastActionEvmTxHash
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
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
	p *evm.LikeProtocol,
	n *evm.BookNFT,
	cosmosNFTIDClassifier cosmosnftidclassifier.CosmosNFTIDClassifier,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	shouldPremintAllNFTs bool,
	premintAllNFTsShouldPremintArbitraryNFTIDs bool,

	cosmosClassId string,
	initialClassOwner string,
	initialClassMinters []string,
	initialClassUpdaters []string,
	initialBatchMintOwner string,
	batchMintPerPage uint64,
	defaultRoyaltyFraction *big.Int,

	cosmosOwner string,
	evmOwner string,
) (lastTxEvmHash *string, err error) {
	mylogger := logger.
		WithGroup("MigrateClass").
		With("cosmosClassId", cosmosClassId).
		With("initialClassOwner", initialClassOwner).
		With("initialClassMinters", initialClassMinters).
		With("initialClassUpdaters", initialClassUpdaters).
		With("cosmosOwner", cosmosOwner).
		With("evmOwner", evmOwner)

	newClassAction, err := GetOrCreateNewClassAction(
		db,
		cosmosClassId,
		initialClassOwner,
		initialClassMinters,
		initialClassUpdaters,
		initialBatchMintOwner,
		shouldPremintAllNFTs,
		defaultRoyaltyFraction,
	)
	if err != nil {
		return nil, err
	}

	// Also sync fields if the class already created
	if shouldPremintAllNFTs != newClassAction.ShouldPremintAllNFTs ||
		initialBatchMintOwner != newClassAction.InitialBatchMintOwner {
		newClassAction.ShouldPremintAllNFTs = shouldPremintAllNFTs
		newClassAction.InitialBatchMintOwner = initialBatchMintOwner
		err := appdb.UpdateLikeNFTMigrationActionNewClass(
			db, newClassAction,
		)
		if err != nil {
			return nil, err
		}
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

	lastTxEvmHash = newClassAction.EvmTxHash

	err = DoPremintAllNFTsActionIfNeeded(
		ctx,
		mylogger,
		db,
		c,
		likecoinAPI,
		p,
		n,
		cosmosNFTIDClassifier,
		erc721ExternalURLBuilder,
		premintAllNFTsShouldPremintArbitraryNFTIDs,
		batchMintPerPage,
		newClassAction,
	)

	if err != nil {
		return nil, err
	}

	if evmOwner != initialClassOwner {
		transferClassAction, err := GetOrCreateTransferClassAction(
			db, *newClassAction.EvmClassId, cosmosOwner, evmOwner,
		)
		if err != nil {
			return nil, err
		}
		transferClassAction, err = DoTransferClassAction(
			ctx,
			mylogger,
			db,
			c,
			n,
			transferClassAction,
		)
		if err != nil {
			return nil, err
		}

		lastTxEvmHash = transferClassAction.EvmTxHash
	} else {
		mylogger.Info("initial class owner and evm owner are the same. skip.")
	}

	return lastTxEvmHash, err
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
