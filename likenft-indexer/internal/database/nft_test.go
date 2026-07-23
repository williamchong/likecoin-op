package database_test

import (
	"context"
	"io"
	"log/slog"
	"math/big"
	"testing"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
)

type nftTestUtil struct {
}

func slogNop() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func (nftTestUtil) makeNFT(
	nftRepository database.NFTRepository,
	nftClass *ent.NFTClass,
	tokenID int64,
	ownerAddress string,
	owner *ent.Account,
) (*ent.NFT, error) {
	return nftRepository.GetOrCreate(
		context.Background(),
		nftClass.Address,
		big.NewInt(tokenID),
		"https://example.com/token",
		"https://example.com/image.png",
		nil,
		nil,
		"Test NFT",
		"Test NFT",
		nil,
		nil,
		nil,
		nil,
		ownerAddress,
		owner,
		nftClass,
	)
}

func TestNFTUpdateOwnerSetsUpdatedAt(t *testing.T) {
	classUtil := &nftClassUtil{}
	util := &nftTestUtil{}

	Convey("Test NFT UpdateOwner sets updated_at", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()

		accountRepository := database.MakeAccountRepository(dbService)
		nftClassRepository := database.MakeNFTClassRepository(dbService)
		nftRepository := database.MakeNFTRepository(dbService)

		accountOrigin, err := accountRepository.GetOrCreateAccount(context.Background(), &ent.Account{
			EvmAddress: common.HexToAddress("0x1000000000000000000000000000000000000001").String(),
		})
		So(err, ShouldBeNil)

		accountNew, err := accountRepository.GetOrCreateAccount(context.Background(), &ent.Account{
			EvmAddress: common.HexToAddress("0x1000000000000000000000000000000000000002").String(),
		})
		So(err, ShouldBeNil)

		nftClassAddress := common.HexToAddress("0x8Ed946DF65a6dD97fcD071c995710b9bf14361c3").String()
		err = classUtil.makeNFTClass(nftClassRepository, accountOrigin, nftClassAddress, 100)
		So(err, ShouldBeNil)
		nftClass, err := nftClassRepository.QueryNFTClassByAddress(context.Background(), nftClassAddress)
		So(err, ShouldBeNil)

		_, err = util.makeNFT(nftRepository, nftClass, 0, accountOrigin.EvmAddress, accountOrigin)
		So(err, ShouldBeNil)

		transferredAt := time.Date(2026, 7, 1, 12, 0, 0, 0, time.UTC)
		err = nftRepository.UpdateOwner(
			context.Background(),
			nftClassAddress,
			big.NewInt(0),
			accountNew.EvmAddress,
			accountNew,
			transferredAt,
		)
		So(err, ShouldBeNil)

		n := dbService.Client().NFT.Query().OnlyX(context.Background())
		So(n.OwnerAddress, ShouldEqual, accountNew.EvmAddress)
		So(n.UpdatedAt.Unix(), ShouldEqual, transferredAt.Unix())
	})
}

func TestQueryNFTClassesByAccountTokensWithNFTIDTokenUpdatedAt(t *testing.T) {
	classUtil := &nftClassUtil{}
	util := &nftTestUtil{}

	Convey("Test token_updated_at is max across owned copies", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()

		accountRepository := database.MakeAccountRepository(dbService)
		nftClassRepository := database.MakeNFTClassRepository(dbService)
		nftRepository := database.MakeNFTRepository(dbService)

		account, err := accountRepository.GetOrCreateAccount(context.Background(), &ent.Account{
			EvmAddress: common.HexToAddress("0x1000000000000000000000000000000000000001").String(),
		})
		So(err, ShouldBeNil)

		nftClassAddress := common.HexToAddress("0x8Ed946DF65a6dD97fcD071c995710b9bf14361c3").String()
		err = classUtil.makeNFTClass(nftClassRepository, account, nftClassAddress, 100)
		So(err, ShouldBeNil)
		nftClass, err := nftClassRepository.QueryNFTClassByAddress(context.Background(), nftClassAddress)
		So(err, ShouldBeNil)

		earlier := time.Date(2026, 7, 1, 12, 0, 0, 0, time.UTC)
		later := time.Date(2026, 7, 20, 12, 0, 0, 0, time.UTC)

		for _, token := range []struct {
			tokenID    int64
			acquiredAt time.Time
		}{
			{tokenID: 0, acquiredAt: later},
			{tokenID: 1, acquiredAt: earlier},
		} {
			_, err = util.makeNFT(nftRepository, nftClass, token.tokenID, account.EvmAddress, account)
			So(err, ShouldBeNil)
			err = nftRepository.UpdateOwner(
				context.Background(),
				nftClassAddress,
				big.NewInt(token.tokenID),
				account.EvmAddress,
				account,
				token.acquiredAt,
			)
			So(err, ShouldBeNil)
		}

		rows, count, _, err := nftClassRepository.QueryNFTClassesByAccountTokensWithNFTID(
			context.Background(),
			account.EvmAddress,
			nil,
			nil,
			database.NFTClassPagination{},
		)
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 1)
		So(rows[0].TokenUpdatedAt, ShouldNotBeNil)
		So(rows[0].TokenUpdatedAt.Unix(), ShouldEqual, later.Unix())
	})
}

func TestBackfillUpdatedAtFromTransferEvents(t *testing.T) {
	classUtil := &nftClassUtil{}
	util := &nftTestUtil{}

	Convey("Test backfill from transfer events", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()

		accountRepository := database.MakeAccountRepository(dbService)
		nftClassRepository := database.MakeNFTClassRepository(dbService)
		nftRepository := database.MakeNFTRepository(dbService)
		evmEventRepository := database.MakeEVMEventRepository(dbService)

		account, err := accountRepository.GetOrCreateAccount(context.Background(), &ent.Account{
			EvmAddress: common.HexToAddress("0x1000000000000000000000000000000000000001").String(),
		})
		So(err, ShouldBeNil)

		nftClassAddress := common.HexToAddress("0x8Ed946DF65a6dD97fcD071c995710b9bf14361c3").String()
		err = classUtil.makeNFTClass(nftClassRepository, account, nftClassAddress, 100)
		So(err, ShouldBeNil)
		nftClass, err := nftClassRepository.QueryNFTClassByAddress(context.Background(), nftClassAddress)
		So(err, ShouldBeNil)

		_, err = util.makeNFT(nftRepository, nftClass, 0, account.EvmAddress, account)
		So(err, ShouldBeNil)

		mintedAt := time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC)
		transferredAt := time.Date(2026, 7, 20, 12, 0, 0, 0, time.UTC)
		tokenIDStr := "0"
		events := []*ent.EVMEvent{}
		for i, ts := range []time.Time{mintedAt, transferredAt} {
			events = append(events, &ent.EVMEvent{
				BlockNumber:      typeutil.Uint64(uint64(100 + i)),
				TransactionIndex: 0,
				LogIndex:         0,
				TransactionHash:  common.HexToHash(big.NewInt(int64(i)).String()).String(),
				BlockHash:        common.HexToHash(big.NewInt(int64(i)).String()).String(),
				Address:          nftClassAddress,
				ChainID:          typeutil.Uint64(99999),
				Topic0:           "Transfer",
				Topic0Hex:        "topic0hex",
				Topic3:           &tokenIDStr,
				Removed:          false,
				Status:           evmevent.StatusProcessed,
				Name:             "Transfer",
				Signature:        "Transfer(address,address,uint256)",
				Timestamp:        ts,
			})
		}
		_, err = evmEventRepository.InsertEvmEventsIfNeeded(context.Background(), events)
		So(err, ShouldBeNil)

		updatedCount, err := nftRepository.BackfillUpdatedAtFromTransferEvents(
			context.Background(),
			slogNop(),
		)
		So(err, ShouldBeNil)
		So(updatedCount, ShouldEqual, 1)

		n := dbService.Client().NFT.Query().OnlyX(context.Background())
		So(n.UpdatedAt.Unix(), ShouldEqual, transferredAt.Unix())
	})
}
