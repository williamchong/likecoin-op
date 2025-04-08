package db_test

import (
	"testing"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryPaginatedLikeNFTAssetMigration(t *testing.T) {
	Convey("QueryPaginatedLikeNFTAssetMigration", t, func() {
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

		s2, err := appdb.InsertLikeNFTAssetSnapshot(db, &model.LikeNFTAssetSnapshot{
			CosmosAddress: "cosmos456",
			Status:        model.NFTSnapshotStatusInit,
		})
		if err != nil {
			t.Fatalf("failed insert snapshot: %v", err)
		}
		m2, err := appdb.InsertLikeNFTAssetMigration(
			db, &model.LikeNFTAssetMigration{
				LikeNFTAssetSnapshotId: s2.Id,
				CosmosAddress:          "cosmos456",
				EthAddress:             "eth456",
				Status:                 model.NFTMigrationStatusCompleted,
			},
		)
		if err != nil {
			t.Fatalf("failed insert asset migration: %v", err)
		}

		{
			// limit 10
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, nil, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 2)
		}

		{
			// limit 0
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 0, 0, nil, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}

		{
			// limit 1, offset 0
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 1, 0, nil, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// limit 1, offset 1
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 1, 1, nil, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m1.Id)
		}

		{
			// status = completed
			status := model.NFTMigrationStatusCompleted
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, &status, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// status = failed
			status := model.NFTMigrationStatusFailed
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, &status, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}

		{
			// keyword = cosmos123
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, nil, "cosmos123")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m1.Id)
		}

		{
			// keyword = eth456
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, nil, "eth456")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// keyword = helloworld
			migrations, err := appdb.QueryPaginatedLikeNFTAssetMigration(db, 10, 0, nil, "helloworld")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}
	})
}
