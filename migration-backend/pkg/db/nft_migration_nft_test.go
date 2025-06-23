package db_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
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

func TestInsertLikeNFTAssetMigrationNFTs(t *testing.T) {
	Convey("InsertLikeNFTAssetMigrationNFTs", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		var err error

		s := &model.LikeNFTAssetSnapshot{
			CreatedAt:     time.Now(),
			CosmosAddress: "cosmos-address",
			Status:        model.NFTSnapshotStatusInProgress,
		}
		s, err = appdb.InsertLikeNFTAssetSnapshot(db, s)
		So(err, ShouldBeNil)

		m := &model.LikeNFTAssetMigration{
			CreatedAt:              time.Now(),
			LikeNFTAssetSnapshotId: s.Id,
			CosmosAddress:          "cosmos-address",
			EthAddress:             "eth-address",
			Status:                 model.NFTMigrationStatusInit,
			EstimatedFinishedTime:  time.Now(),
		}
		m, err = appdb.InsertLikeNFTAssetMigration(db, m)
		So(err, ShouldBeNil)

		nfts := make([]model.LikeNFTAssetMigrationNFT, 10000)
		for i := uint64(0); i < 10000; i++ {
			nfts[i] = model.LikeNFTAssetMigrationNFT{
				LikeNFTAssetMigrationId: m.Id,
				CreatedAt:               time.Now(),
				CosmosClassId:           fmt.Sprintf("cosmos-class-id-%d", i),
				CosmosNFTId:             fmt.Sprintf("%d", i),
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusInit,
			}
		}

		err = appdb.InsertLikeNFTAssetMigrationNFTs(
			db, nfts,
		)
		So(err.Error(), ShouldContainSubstring, "but PostgreSQL only supports 65535 parameters")
	})
}

func TestRemoveLikeNFTAssetMigrationNFTByMigrationId(t *testing.T) {
	Convey("RemoveLikeNFTAssetMigrationNFTByMigrationId", t, func() {
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
			mn, err := appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			So(len(mn), ShouldEqual, 3)
			err = appdb.RemoveLikeNFTAssetMigrationNFTByMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			mn, err = appdb.QueryLikeNFTAssetMigrationNFTsByNFTMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			So(len(mn), ShouldEqual, 0)
		}
	})
}
