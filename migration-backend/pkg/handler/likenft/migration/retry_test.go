package migration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hibiken/asynq"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft/migration"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRetryMigrationRequestBody(t *testing.T) {
	Convey("Given a RetryMigrationRequestBody", t, func() {
		Convey("When the request body is empty", func() {
			req := &migration.RetryMigrationRequestBody{
				BookNFTCollection: []string{},
				BookNFT:           []migration.BookNFT{},
			}
			So(req.CheckAssets([]model.LikeNFTAssetMigrationClass{}, []model.LikeNFTAssetMigrationNFT{}), ShouldBeTrue)
			So(req.CheckAssets([]model.LikeNFTAssetMigrationClass{
				{
					LikeNFTAssetMigrationId: 1,
					CosmosClassId:           "likenft12345",
					Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
				},
			}, []model.LikeNFTAssetMigrationNFT{
				{
					LikeNFTAssetMigrationId: 1,
					CosmosClassId:           "likenft12347",
					CosmosNFTId:             "1",
					Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
				},
			}), ShouldBeFalse)
		})

		Convey("When the request body is not empty", func() {
			req := &migration.RetryMigrationRequestBody{
				BookNFTCollection: []string{"likenft12345", "likenft12346"},
				BookNFT: []migration.BookNFT{
					{
						ClassId: "likenft12347",
						NFTId:   "1",
					},
					{
						ClassId: "likenft12347",
						NFTId:   "2",
					},
				},
			}

			Convey("Should fail when no assets are provided", func() {
				So(req.CheckAssets([]model.LikeNFTAssetMigrationClass{}, []model.LikeNFTAssetMigrationNFT{}), ShouldBeFalse)
			})

			Convey("Should success when assets matched", func() {
				So(req.CheckAssets([]model.LikeNFTAssetMigrationClass{
					{
						LikeNFTAssetMigrationId: 1,
						CosmosClassId:           "likenft12345",
						Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
					},
					{
						LikeNFTAssetMigrationId: 2,
						CosmosClassId:           "likenft12346",
						Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
					},
				}, []model.LikeNFTAssetMigrationNFT{
					{
						LikeNFTAssetMigrationId: 1,
						CosmosClassId:           "likenft12347",
						CosmosNFTId:             "1",
						Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
					},
					{
						LikeNFTAssetMigrationId: 2,
						CosmosClassId:           "likenft12347",
						CosmosNFTId:             "2",
						Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
					},
				}), ShouldBeTrue)
			})

			Convey("Should fail when assets not matched", func() {
				So(req.CheckAssets([]model.LikeNFTAssetMigrationClass{
					{
						LikeNFTAssetMigrationId: 1,
						CosmosClassId:           "likenft12345",
						Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
					},
					{
						LikeNFTAssetMigrationId: 2,
						CosmosClassId:           "likenft12346",
						Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
					},
				}, []model.LikeNFTAssetMigrationNFT{
					{
						LikeNFTAssetMigrationId: 1,
						CosmosClassId:           "likenft12347",
						CosmosNFTId:             "1",
						Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
					},
					{
						LikeNFTAssetMigrationId: 2,
						CosmosClassId:           "likenft12347",
						CosmosNFTId:             "2",
						Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
					},
					{
						LikeNFTAssetMigrationId: 2,
						CosmosClassId:           "likenft12347",
						CosmosNFTId:             "3",
						Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
					},
				}), ShouldBeFalse)
			})
		})
	})
}

func TestRetryMigration(t *testing.T) {
	Convey("RetryMigration", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		s, err := appdb.InsertLikeNFTAssetSnapshot(db, &model.LikeNFTAssetSnapshot{
			CosmosAddress: "cosmosaddr",
			Status:        model.NFTSnapshotStatusInit,
		})
		if err != nil {
			t.Fatalf("appdb.InsertLikeNFTAssetSnapshot: %v", err)
		}

		m, err := appdb.InsertLikeNFTAssetMigration(db, &model.LikeNFTAssetMigration{
			LikeNFTAssetSnapshotId: s.Id,
			CosmosAddress:          "cosmosaddr",
			EthAddress:             "ethaddress",
			Status:                 model.NFTMigrationStatusFailed,
		})
		if err != nil {
			t.Fatalf("appdb.InsertLikeNFTAssetMigration: %v", err)
		}

		err = appdb.InsertLikeNFTAssetMigrationClasses(db, []model.LikeNFTAssetMigrationClass{
			{
				LikeNFTAssetMigrationId: m.Id,
				CosmosClassId:           "cosmosclass-completed",
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationClassStatusCompleted,
			},
			{
				LikeNFTAssetMigrationId: m.Id,
				CosmosClassId:           "cosmosclass-failed",
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationClassStatusFailed,
			},
		})
		if err != nil {
			t.Fatalf("appdb.InsertLikeNFTAssetMigrationClasses: %v", err)
		}

		err = appdb.InsertLikeNFTAssetMigrationNFTs(db, []model.LikeNFTAssetMigrationNFT{
			{
				LikeNFTAssetMigrationId: m.Id,
				CosmosClassId:           "cosmosclass",
				CosmosNFTId:             "nft-1",
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusCompleted,
			},
			{
				LikeNFTAssetMigrationId: m.Id,
				CosmosClassId:           "cosmosclass",
				CosmosNFTId:             "nft-2",
				Name:                    "name",
				Image:                   "image",
				Status:                  model.LikeNFTAssetMigrationNFTStatusFailed,
			},
		})
		if err != nil {
			t.Fatalf("appdb.InsertLikeNFTAssetMigrationNFTs: %v", err)
		}

		redis, _ := testutil.GetRedis(t)
		asynqClient := asynq.NewClientFromRedisClient(redis)

		h := &migration.RetryMigrationHandler{
			Db:          db,
			AsynqClient: asynqClient,
		}

		reqBody := &migration.RetryMigrationRequestBody{
			BookNFTCollection: []string{
				"cosmosclass-failed",
			},
			BookNFT: []migration.BookNFT{
				{
					ClassId: "cosmosclass",
					NFTId:   "nft-2",
				},
			},
		}
		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatalf("json.Marshal: %v", err)
		}

		req, err := http.NewRequest("PUT", "/cosmosaddr", bytes.NewReader(bodyBytes))
		if err != nil {
			t.Fatalf("http.NewRequest: %v", err)
		}

		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		m, err = appdb.QueryLikeNFTAssetMigrationById(db, m.Id)
		if err != nil {
			t.Fatalf("appdb.QueryLikeNFTAssetMigrationById: %v", err)
		}

		So(m.Status, ShouldEqual, model.NFTMigrationStatusInProgress)
	})
}
