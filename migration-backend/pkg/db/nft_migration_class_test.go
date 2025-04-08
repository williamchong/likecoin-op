package db_test

import (
	"testing"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(t *testing.T) {
	Convey("QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus", t, func() {
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
		err = appdb.InsertLikeNFTAssetMigrationClasses(db, []model.LikeNFTAssetMigrationClass{
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass123",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass124",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusInit,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass125",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusInit,
			},
		})
		if err != nil {
			t.Fatalf("failed insert asset migration classes: %v", err)
		}

		{
			classes, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
				db,
				m1.Id,
				model.LikeNFTAssetMigrationClassStatusInit,
			)

			So(err, ShouldBeNil)
			So(len(classes), ShouldEqual, 2)
		}

		{
			classes, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
				db,
				m1.Id,
				model.LikeNFTAssetMigrationClassStatusFailed,
			)

			So(err, ShouldBeNil)
			So(len(classes), ShouldEqual, 1)
		}

		{
			classes, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationIdAndStatus(
				db,
				uint64(10000),
				model.LikeNFTAssetMigrationClassStatusFailed,
			)

			So(err, ShouldBeNil)
			So(len(classes), ShouldEqual, 0)
		}
	})
}

func TestRemoveLikeNFTAssetMigrationClassByMigrationId(t *testing.T) {
	Convey("RemoveLikeNFTAssetMigrationClassByMigrationId", t, func() {
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
		err = appdb.InsertLikeNFTAssetMigrationClasses(db, []model.LikeNFTAssetMigrationClass{
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass123",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass124",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusInit,
			},
			{
				LikeNFTAssetMigrationId: m1.Id,
				CosmosClassId:           "cosmosclass125",
				Name:                    "Name",
				Image:                   "Image",
				Status:                  model.LikeNFTAssetMigrationClassStatusInit,
			},
		})
		if err != nil {
			t.Fatalf("failed insert asset migration classes: %v", err)
		}

		{
			mc, err := appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			So(len(mc), ShouldEqual, 3)
			err = appdb.RemoveLikeNFTAssetMigrationClassByMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			mc, err = appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(db, m1.Id)
			So(err, ShouldBeNil)
			So(len(mc), ShouldEqual, 0)
		}
	})
}
