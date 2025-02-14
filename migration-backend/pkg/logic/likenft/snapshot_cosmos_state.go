package likenft

import (
	"database/sql"
	"fmt"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type SnapshotCosmosStateLogic struct {
	DB                  *sql.DB
	CosmosAPI           *api.CosmosAPI
	LikeNFTCosmosClient *cosmos.LikeNFTCosmosClient
}

func (l *SnapshotCosmosStateLogic) Execute(cosmosAddress string) error {
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
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed getting latest block: %v", err))
	}

	cosmosClasses, err := l.LikeNFTCosmosClient.QueryAllNFTClassesByOwner(cosmos.QueryAllNFTClassesByOwnerRequest{
		Owner: cosmosAddress,
	})

	if err != nil {
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed getting classes by owner: %v", err))
	}

	cosmosNFTs, err := l.LikeNFTCosmosClient.QueryAllNFTsByOwner(cosmos.QueryAllNFTsByOwnerRequest{
		Owner: cosmosAddress,
	})

	if err != nil {
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed getting nfts by owner: %v", err))
	}

	snapshotClasses := make([]model.LikeNFTAssetSnapshotClass, 0, len(cosmosClasses.Classes))

	for _, cosmosClass := range cosmosClasses.Classes {
		snapshotClasses = append(snapshotClasses, model.LikeNFTAssetSnapshotClass{
			NFTSnapshotId: latestSnapshot.Id,
			CosmosClassId: cosmosClass.Id,
			Name:          cosmosClass.Name,
			Image:         cosmosClass.Metadata.Image,
		})
	}

	snapshotNFTs := make([]model.LikeNFTAssetSnapshotNFT, 0, len(cosmosNFTs.NFTs))

	for _, cosmosNFT := range cosmosNFTs.NFTs {
		snapshotNFTs = append(snapshotNFTs, model.LikeNFTAssetSnapshotNFT{
			NFTSnapshotId: latestSnapshot.Id,
			CosmosClassId: cosmosNFT.ClassId,
			CosmosNFTId:   cosmosNFT.Id,
			Name:          cosmosNFT.Data.Metadata.Name,
			Image:         cosmosNFT.Data.Metadata.Image,
		})
	}

	latestSnapshot.BlockHeight = &block.Header.Height
	latestSnapshot.BlockTime = &block.Header.Time
	latestSnapshot.Status = model.NFTSnapshotStatusCompleted

	tx, err := l.DB.Begin()
	if err != nil {
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed begin tx: %v", err))
	}
	defer tx.Commit()
	err = appdb.UpdateLikeNFTAssetSnapshot(tx, latestSnapshot)
	if err != nil {
		tx.Rollback()
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed update nft snapshot: %v", err))
	}
	err = appdb.InsertLikeNFTAssetSnapshotClasses(tx, snapshotClasses)
	if err != nil {
		tx.Rollback()
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed insert nft snapshot classes: %v", err))
	}
	err = appdb.InsertLikeNFTAssetSnapshotNFTs(tx, snapshotNFTs)
	if err != nil {
		tx.Rollback()
		return l.failed(l.DB, latestSnapshot, fmt.Errorf("failed insert nft snapshot nfts: %v", err))
	}
	return nil
}

func (l *SnapshotCosmosStateLogic) failed(db *sql.DB, snapshot *model.LikeNFTAssetSnapshot, err error) error {
	errStr := err.Error()
	snapshot.FailedReason = &errStr
	snapshot.Status = model.NFTSnapshotStatusFailed
	updateErr := appdb.UpdateLikeNFTAssetSnapshot(db, snapshot)
	if updateErr != nil {
		return updateErr
	}
	return err
}
