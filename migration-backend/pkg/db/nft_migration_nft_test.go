package db_test

import (
	"testing"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(t *testing.T) {
	Convey("QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		s1, err := appdb.InsertLikeNFTAssetSnapshot(db, &model.LikeNFTAssetSnapshot{
			CosmosAddress: "cosmos123",
			Status:        model.NFTSnapshotStatusInit,
		})
		if err != nil {
			t.Fatalf("failed insert snapshot: %v", err)
		}
		m1, err := appdb.InsertLikeNFTAssetMigration(
			db, &model.LikeNFTAssetMigration{
				LikeNFTAssetSnapshotId: s1.Id,
				CosmosAddress:          "cosmos123",
				EthAddress:             "eth123",
				Status:                 model.NFTMigrationStatusInit,
			},
		)
		if err != nil {
			t.Fatalf("failed insert asset migration: %v", err)
		}
		err = appdb.InsertLikeNFTAssetMigrationNFTs(db, []model.LikeNFTAssetMigrationNFT{
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass123",
				CosmosNFTId:             "nft-0001",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass123",
				CosmosNFTId:             "nft-0002",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusInit,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass123",
				CosmosNFTId:             "nft-0003",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusInit,
			},
		})
		if err != nil {
			t.Fatalf("failed insert asset migration nfts: %v", err)
		}

		{
			nfts, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(
				db,
				m1.Id,
				model.LikeNFTAssetMigrationNFTStatusInit,
			)

			So(err, ShouldBeNil)
			So(len(nfts), ShouldEqual, 2)
		}

		{
			nfts, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(
				db,
				m1.Id,
				model.LikeNFTAssetMigrationNFTStatusFailed,
			)

			So(err, ShouldBeNil)
			So(len(nfts), ShouldEqual, 1)
		}

		{
			classes, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationIdAndStatus(
				db,
				uint64(10000),
				model.LikeNFTAssetMigrationNFTStatusFailed,
			)

			So(err, ShouldBeNil)
			So(len(classes), ShouldEqual, 0)
		}
	})
}
