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
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
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
	cosmosNFTIDClassifier cosmosnftidclassifier.CosmosNFTIDClassifier,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	shouldPremintAllNFTs bool,
	premintAllNFTsShouldPremintArbitraryNFTIDs bool,

	initialClassOwner string,
	initialClassMinters []string,
	initialClassUpdaters []string,
	initialBatchMintOwner string,
	defaultRoyaltyFraction *big.Int,
	batchMintPerPage uint64,

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
		cosmosNFTIDClassifier,
		erc721ExternalURLBuilder,
		shouldPremintAllNFTs,
		premintAllNFTsShouldPremintArbitraryNFTIDs,
		initialClassOwner,
		initialClassMinters,
		initialClassUpdaters,
		initialBatchMintOwner,
		defaultRoyaltyFraction,
		batchMintPerPage,
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
	cosmosNFTIDClassifier cosmosnftidclassifier.CosmosNFTIDClassifier,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	shouldPremintAllNFTs bool,
	premintAllNFTsShouldPremintArbitraryNFTIDs bool,

	initialClassOwner string,
	initialClassMinters []string,
	initialClassUpdaters []string,
	initialBatchMintOwner string,
	defaultRoyaltyFraction *big.Int,
	batchMintPerPage uint64,

	cosmosClassId string,
	cosmosNFTId string,
	evmOwner string,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	nftIDMatcher := nftidmatcher.MakeNFTIDMatcher()

	mylogger := logger.
		WithGroup("MigrateNFT").
		With("initialClassOwner", initialClassOwner).
		With("initialClassMinters", initialClassMinters).
		With("initialClassUpdaters", initialClassUpdaters).
		With("initialBatchMintOwner", initialBatchMintOwner).
		With("cosmosClassId", cosmosClassId).
		With("cosmosNFTId", cosmosNFTId).
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

	nftId, ok := nftIDMatcher.ExtractSerialID(cosmosNFTId)
	if ok {
		// nftid is serial
		// tokenid is zero based
		// So nftId=1 => expected token id should be from [0, 1]
		expectedSupply := nftId + 1
		err = DoBatchMintNFTsFromCosmosAction(
			ctx,
			logger,
			db,
			p,
			n,
			likecoinAPI,
			c,
			erc721ExternalURLBuilder,
			*newClassAction.EvmClassId,
			expectedSupply,
			batchMintPerPage,
			initialBatchMintOwner,
		)
	}

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
		erc721ExternalURLBuilder,
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
