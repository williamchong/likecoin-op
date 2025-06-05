package likenft

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmostoevmnftmirror"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/util/actionlifecycle"
)

type BatchMintNFTsFromCosmosAction interface {
	Act(ctx context.Context) error
}

type batchMintNFTsFromCosmosAction struct {
	mirror                 cosmostoevmnftmirror.CosmosToEVMNFTMirror
	cosmosClassIDRetriever CosmosClassIDRetriever

	db *sql.DB

	evmClassId            string
	currentSupply         uint64
	expectedSupply        uint64
	batchMintSize         uint64
	initialBatchMintOwner string
}

func MakeBatchMintNFTsFromCosmosAction(
	mirror cosmostoevmnftmirror.CosmosToEVMNFTMirror,
	cosmosClassIDRetriever CosmosClassIDRetriever,
	db *sql.DB,
	evmClassId string,
	currentSupply uint64,
	expectedSupply uint64,
	batchMintSize uint64,
	initialBatchMintOwner string,
) BatchMintNFTsFromCosmosAction {
	return &batchMintNFTsFromCosmosAction{
		mirror:                 mirror,
		cosmosClassIDRetriever: cosmosClassIDRetriever,
		db:                     db,
		evmClassId:             evmClassId,
		currentSupply:          currentSupply,
		expectedSupply:         expectedSupply,
		batchMintSize:          batchMintSize,
		initialBatchMintOwner:  initialBatchMintOwner,
	}
}

func (a *batchMintNFTsFromCosmosAction) Act(ctx context.Context) error {
	cosmosClassID, err := a.cosmosClassIDRetriever.GetByEvmClassId(a.evmClassId)
	if err != nil {
		return err
	}

	var lc actionlifecycle.ActionLifecycle[
		actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp,
		model.LikeNFTMigrationActionBatchMintNFTsFromCosmos,
	]

	_, err = a.mirror.Mirror(
		ctx,
		cosmosClassID,
		a.evmClassId,
		a.expectedSupply,
		func(ctx context.Context, fromID, toID uint64) error {
			lc = actionlifecycle.MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle(
				db.MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosRepository(a.db),
				a.evmClassId,
				fromID,
				toID,
				a.batchMintSize,
				a.initialBatchMintOwner,
			)
			_, err := lc.Begin(ctx)
			return err
		},
		func(ctx context.Context, result cosmostoevmnftmirror.CosmosToEVMNFTMirrorResult) error {
			_, err := lc.Success(ctx, &actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp{
				FromID:    result.FromID,
				ToID:      result.ToID,
				EvmTxHash: result.EvmTxHash,
			})
			return err
		},
		func(ctx context.Context, err error) error {
			_, fErr := lc.Failed(ctx, err)
			return errors.Join(err, fErr)
		},
	)

	return err
}

func DoBatchMintNFTsFromCosmosAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	likeProtocolEvmClient *evm.LikeProtocol,
	bookNFTEvmClient *evm.BookNFT,
	likecoinAPI *likecoin_api.LikecoinAPI,
	likeNFTCosmosClient *cosmos.LikeNFTCosmosClient,
	externalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	evmClassId string,
	expectedSupply uint64,
	batchMintSize uint64,
	batchMintOwner string,
) error {
	currentEvmSupply, err := bookNFTEvmClient.TotalSupply(common.HexToAddress(evmClassId))
	if err != nil {
		return err
	}

	doAction := MakeBatchMintNFTsFromCosmosAction(
		cosmostoevmnftmirror.MakeCosmosToEVMNFTMirror(
			logger,
			nftidmatcher.MakeNFTIDMatcher(),
			externalURLBuilder,
			likeNFTCosmosClient,
			bookNFTEvmClient,
			batchMintOwner,
			batchMintSize,
		),
		MakeDBCosmosClassIDRetriever(db),
		db,
		evmClassId,
		currentEvmSupply.Uint64(),
		expectedSupply,
		batchMintSize,
		batchMintOwner,
	)

	return doAction.Act(ctx)
}
