package likenft

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmostoevmnftmirror"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/util/actionlifecycle"
)

type BatchMintNFTsFromCosmosAction interface {
	Act(ctx context.Context) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error)
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

func (a *batchMintNFTsFromCosmosAction) act(
	ctx context.Context,
	evmClassID string,
	expectedSupply uint64,
) ([]cosmostoevmnftmirror.CosmosToEVMNFTMirrorResult, error) {
	cosmosClassID, err := a.cosmosClassIDRetriever.GetByEvmClassId(evmClassID)
	if err != nil {
		return nil, err
	}
	return a.mirror.Mirror(ctx, cosmosClassID, evmClassID, expectedSupply)
}

func (a *batchMintNFTsFromCosmosAction) Act(ctx context.Context) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	lc := actionlifecycle.MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle(
		db.MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosRepository(a.db),
		a.evmClassId,
		a.currentSupply,
		a.expectedSupply,
		a.batchMintSize,
		a.initialBatchMintOwner,
	)

	return actionlifecycle.WithActionLifecycle[
		actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp,
		model.LikeNFTMigrationActionBatchMintNFTsFromCosmos,
		actionlifecycle.ActionLifecycle[
			actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp,
			model.LikeNFTMigrationActionBatchMintNFTsFromCosmos,
		],
	](ctx, lc, func(
		ctx context.Context,
		lc actionlifecycle.ActionLifecycle[
			actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp,
			model.LikeNFTMigrationActionBatchMintNFTsFromCosmos,
		],
	) (*actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp, error) {
		results, err := a.act(ctx, a.evmClassId, a.expectedSupply)
		if err != nil {
			return nil, err
		}
		// No minting at all
		if len(results) == 0 {
			return &actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp{
				FromID:    a.expectedSupply,
				ToID:      a.expectedSupply,
				EvmTxHash: "",
			}, nil
		}
		fromID := results[0].FromID
		toID := results[len(results)-1].ToID
		evmTxHash := results[len(results)-1].EvmTxHash
		return &actionlifecycle.LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp{
			FromID:    fromID,
			ToID:      toID,
			EvmTxHash: evmTxHash,
		}, nil
	})
}

func DoBatchMintNFTsFromCosmosAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	likeProtocolEvmClient *evm.LikeProtocol,
	bookNFTEvmClient *evm.BookNFT,
	likecoinAPI *likecoin_api.LikecoinAPI,
	likeNFTCosmosClient *cosmos.LikeNFTCosmosClient,

	evmClassId string,
	expectedSupply uint64,
	batchMintSize uint64,
	batchMintOwner string,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	currentEvmSupply, err := bookNFTEvmClient.TotalSupply(common.HexToAddress(evmClassId))
	if err != nil {
		return nil, err
	}

	doAction := MakeBatchMintNFTsFromCosmosAction(
		cosmostoevmnftmirror.MakeCosmosToEVMNFTMirror(
			logger,
			nftidmatcher.MakeNFTIDMatcher(),
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
