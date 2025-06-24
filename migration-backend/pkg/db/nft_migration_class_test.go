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

func TestInsertLikeNFTAssetMigrationClasses(t *testing.T) {
	Convey("InsertLikeNFTAssetMigrationClasses", t, func() {
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

		classes := make([]model.LikeNFTAssetMigrationClass, 10000)
		for i := uint64(0); i < 10000; i++ {
			classes[i] = model.LikeNFTAssetMigrationClass{
				LikeNFTAssetMigrationId: m.Id,
				CreatedAt:               time.Now(),
				CosmosClassId:           fmt.Sprintf("cosmos-class-id-%d", i),
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationClassStatusInit,
			}
		}

		err = appdb.InsertLikeNFTAssetMigrationClasses(
			db, classes,
		)
		So(err, ShouldBeNil)

		classes, err = appdb.QueryLikeNFTAssetMigrationClassesByNFTMigrationId(db, m.Id)
		So(err, ShouldBeNil)
		So(len(classes), ShouldEqual, 10000)
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
