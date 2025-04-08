package db_test

import (
	"database/sql"
	"errors"
	"testing"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryPaginatedLikeCoinMigration(t *testing.T) {
	Convey("QueryPaginatedLikeCoinMigration", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		m1, err := appdb.InsertLikeCoinMigration(db, &model.LikeCoinMigration{
			UserCosmosAddress:    "cosmos123",
			BurningCosmosAddress: "cosmos123",
			MintingEthAddress:    "eth123",
			Status:               model.LikeCoinMigrationStatusPendingCosmosTxHash,
		})
		if err != nil {
			t.Fatalf("failed to insert likecoin migration: %v", err)
		}
		m2, err := appdb.InsertLikeCoinMigration(db, &model.LikeCoinMigration{
			UserCosmosAddress:    "cosmos456",
			BurningCosmosAddress: "cosmos456",
			MintingEthAddress:    "eth456",
			Status:               model.LikeCoinMigrationStatusCompleted,
		})
		if err != nil {
			t.Fatalf("failed to insert likecoin migration: %v", err)
		}

		{
			// limit 10
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, nil, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 2)
		}

		{
			// limit 0
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 0, 0, nil, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}

		{
			// limit 1, offset 0
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 1, 0, nil, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// limit 1, offset 1
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 1, 1, nil, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m1.Id)
		}

		{
			// status = completed
			status := model.LikeCoinMigrationStatusCompleted
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, &status, "")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// status = failed
			status := model.LikeCoinMigrationStatusFailed
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, &status, "")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}

		{
			// keyword = cosmos123
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, nil, "cosmos123")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m1.Id)
		}

		{
			// keyword = eth456
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, nil, "eth456")
			So(err, ShouldBeNil)
			So(migrations[0].Id, ShouldEqual, m2.Id)
		}

		{
			// keyword = helloworld
			migrations, err := appdb.QueryPaginatedLikeCoinMigration(db, 10, 0, nil, "helloworld")
			So(err, ShouldBeNil)
			So(len(migrations), ShouldEqual, 0)
		}
	})
}

func TestRemoveLikeCoinMigration(t *testing.T) {
	Convey("RemoveLikeCoinMigration", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		address := "cosmos123"

		m1, err := appdb.InsertLikeCoinMigration(db, &model.LikeCoinMigration{
			UserCosmosAddress:    address,
			BurningCosmosAddress: address,
			MintingEthAddress:    "eth123",
			Status:               model.LikeCoinMigrationStatusPendingCosmosTxHash,
		})
		if err != nil {
			t.Fatalf("failed to insert likecoin migration: %v", err)
		}

		{

			m, err := appdb.QueryLatestLikeCoinMigration(db, address)
			So(err, ShouldBeNil)
			So(m.Id, ShouldEqual, m1.Id)

			err = appdb.RemoveLikeCoinMigration(db, m1.Id)
			So(err, ShouldBeNil)
			_, err = appdb.QueryLatestLikeCoinMigration(db, address)
			So(errors.Is(err, sql.ErrNoRows), ShouldBeTrue)
		}
	})
}
