package likenft

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type SnapshotCosmosStateLogic struct {
	DB                  *sql.DB
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient

	ClassMigrationEstimatedDuration time.Duration
	NFTMigrationEstimatedDuration   time.Duration
}

func (l *SnapshotCosmosStateLogic) Execute(ctx context.Context, cosmosAddress string) error {
	latestSnapshot, err := appdb.QueryLatestLikeNFTAssetSnapshotByCosmosAddress(l.DB, cosmosAddress)
	if err != nil {
		return fmt.Errorf("error querying latest nft snapshot: %v", err)
	}
	if latestSnapshot.Status != model.NFTSnapshotStatusInit {
		return fmt.Errorf("error latest snapshot %v is not init", latestSnapshot.Id)
	}

	latestSnapshot.Status = model.NFTSnapshotStatusInProgress
	err = appdb.UpdateLikeNFTAssetSnapshot(l.DB, latestSnapshot)
	if err != nil {
		return fmt.Errorf("error update nft snapshot: %v", err)
	}

	block, err := l.CosmosAPI.QueryLatestBlock()

	if err != nil {
		return l.failed(ctx, l.DB, latestSnapshot, fmt.Errorf("failed getting latest block: %v", err))
	}

	cosmosClasses, err := l.LikeNFTCosmosClient.QueryAllNFTClassesByOwner(
		ctx,
		cosmos.QueryAllNFTClassesByOwnerRequest{
			ISCNOwner: cosmosAddress,
		})

	if err != nil {
		return l.failed(ctx, l.DB, latestSnapshot, fmt.Errorf("failed getting classes by owner: %v", err))
	}

	cosmosNFTs, err := l.LikeNFTCosmosClient.QueryAllNFTsByOwner(cosmos.QueryAllNFTsByOwnerRequest{
		Owner: cosmosAddress,
	})

	if err != nil {
		return l.failed(ctx, l.DB, latestSnapshot, fmt.Errorf("failed getting nfts by owner: %v", err))
	}

	snapshotClasses := make([]model.LikeNFTAssetSnapshotClass, 0, len(cosmosClasses.Classes))

	estimatedTotalDuration := time.Duration(0)
	for _, cosmosClass := range cosmosClasses.Classes {
		estimatedDurationNeeded := l.ClassMigrationEstimatedDuration
		snapshotClasses = append(snapshotClasses, model.LikeNFTAssetSnapshotClass{
			NFTSnapshotId: latestSnapshot.Id,
			CosmosClassId: cosmosClass.Id,
			Name:          cosmosClass.Name,
			Image:         cosmosClass.Metadata.Image,

			EstimatedMigrationDurationNeeded: estimatedDurationNeeded,
		})

		estimatedTotalDuration += estimatedDurationNeeded
	}

	snapshotNFTs := make([]model.LikeNFTAssetSnapshotNFT, 0, len(cosmosNFTs.NFTs))

	for _, cosmosNFT := range cosmosNFTs.NFTs {
		estimatedDurationNeeded := l.NFTMigrationEstimatedDuration

		snapshotNFTs = append(snapshotNFTs, model.LikeNFTAssetSnapshotNFT{
			NFTSnapshotId: latestSnapshot.Id,
			CosmosClassId: cosmosNFT.ClassId,
			CosmosNFTId:   cosmosNFT.Id,
			Name:          cosmosNFT.Data.Metadata.Name,
			Image:         cosmosNFT.Data.Metadata.Image,

			EstimatedMigrationDurationNeeded: estimatedDurationNeeded,
		})

		estimatedTotalDuration += estimatedDurationNeeded
	}

	latestSnapshot.BlockHeight = &block.Header.Height
	latestSnapshot.BlockTime = &block.Header.Time
	latestSnapshot.Status = model.NFTSnapshotStatusCompleted
	latestSnapshot.EstimatedMigrationDurationNeeded = &estimatedTotalDuration

	err = appdb.WithTx(ctx, l.DB, func(tx *sql.Tx) error {
		err := appdb.UpdateLikeNFTAssetSnapshot(tx, latestSnapshot)
		if err != nil {
			return err
		}
		err = appdb.InsertLikeNFTAssetSnapshotClasses(tx, snapshotClasses)
		if err != nil {
			return err
		}
		err = appdb.InsertLikeNFTAssetSnapshotNFTs(tx, snapshotNFTs)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return l.failed(ctx, l.DB, latestSnapshot, err)
	}
	return nil
}

func (l *SnapshotCosmosStateLogic) failed(
	ctx context.Context,
	db *sql.DB,
	snapshot *model.LikeNFTAssetSnapshot,
	err error,
) error {
	errStr := err.Error()
	return appdb.WithTx(ctx, db, func(tx *sql.Tx) error {
		s, err := appdb.QueryLikeNFTAssetSnapshotById(tx, snapshot.Id)
		if err != nil {
			return err
		}
		s.FailedReason = &errStr
		s.Status = model.NFTSnapshotStatusFailed
		return errors.Join(err, appdb.UpdateLikeNFTAssetSnapshot(db, s))
	})
}
