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

func TestInsertLikeNFTAssetSnapshotNFTs(t *testing.T) {
	Convey("InsertLikeNFTAssetSnapshotNFTs", t, func() {
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

		nfts := make([]model.LikeNFTAssetSnapshotNFT, 15000)
		for i := uint64(0); i < 15000; i++ {
			nfts[i] = model.LikeNFTAssetSnapshotNFT{
				NFTSnapshotId: s.Id,
				CreatedAt:     time.Now(),
				CosmosClassId: fmt.Sprintf("cosmos-class-id-%d", i),
				CosmosNFTId:   fmt.Sprintf("%d", i),
				Name:          "name",
				Image:         "image",
			}
		}

		err = appdb.InsertLikeNFTAssetSnapshotNFTs(
			db, nfts,
		)
		So(err.Error(), ShouldContainSubstring, "but PostgreSQL only supports 65535 parameters")
	})
}
