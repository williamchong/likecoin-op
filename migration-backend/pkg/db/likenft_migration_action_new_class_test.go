package db_test

import (
	"testing"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	"github.com/likecoin/like-migration-backend/pkg/types/commaseparatedstring"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLikeNFTMigrationActionNewClass(t *testing.T) {
	Convey("QueryLikeCoinMigrationById", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		insertResult, err := appdb.InsertLikeNFTMigrationActionNewClass(db, &model.LikeNFTMigrationActionNewClass{
			CosmosClassId:     "cosmosclassid",
			InitialOwner:      "initialowner",
			InitialMintersStr: commaseparatedstring.CommaSeparatedString("initialminter1,initialminter2"),
			InitialUpdater:    "initialupdater",
			Status:            model.LikeNFTMigrationActionNewClassStatusInit,
		})
		if err != nil {
			t.Fatalf("failed to insert MigrationActionNewClass: %v", err)
		}

		cosmosClassId := insertResult.CosmosClassId

		queryResult, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
			CosmosClassId: &cosmosClassId,
		})

		if err != nil {
			t.Fatalf("failed to query MigrationActionNewClass: %v", err)
		}

		So(queryResult.InitialMintersStr, ShouldEqual, commaseparatedstring.CommaSeparatedString("initialminter1,initialminter2"))
	})
}
