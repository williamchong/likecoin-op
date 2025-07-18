package database_test

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"

	. "github.com/smartystreets/goconvey/convey"
)

type nftClassAcquireBookNFTEventsTestUtil struct{}

func (util *nftClassAcquireBookNFTEventsTestUtil) makeNFTClass(
	address string,
	status *nftclass.AcquireBookNftEventsStatus,
	score *float64,
	failedCount int,
	failedReason *string,
) *ent.NFTClass {
	return &ent.NFTClass{
		Address:                          address,
		Name:                             "Test NFT Class",
		Symbol:                           "TEST",
		DeployerAddress:                  "0xDeployer",
		DeployedBlockNumber:              typeutil.Uint64(1000),
		LatestEventBlockNumber:           typeutil.Uint64(2000),
		TotalSupply:                      big.NewInt(1000),
		MaxSupply:                        typeutil.Uint64(10000),
		BannerImage:                      "banner.jpg",
		FeaturedImage:                    "featured.jpg",
		MintedAt:                         time.Now().UTC(),
		UpdatedAt:                        time.Now().UTC(),
		DisabledForIndexing:              false,
		AcquireBookNftEventsWeight:       1.0,
		AcquireBookNftEventsStatus:       status,
		AcquireBookNftEventsScore:        score,
		AcquireBookNftEventsFailedCount:  failedCount,
		AcquireBookNftEventsFailedReason: failedReason,
	}
}

func TestNFTClassAcquireBookNFTEventsRepository(t *testing.T) {
	util := &nftClassAcquireBookNFTEventsTestUtil{}

	Convey("Test RetrieveState", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class
		nftClass := util.makeNFTClass("0xTestAddress", nil, nil, 0, nil)
		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test successful retrieval
		retrieved, err := repo.RetrieveState(context.Background(), "0xTestAddress")
		So(err, ShouldBeNil)
		So(retrieved, ShouldNotBeNil)
		So(retrieved.Address, ShouldEqual, "0xTestAddress")

		// Test case-insensitive address matching
		retrieved, err = repo.RetrieveState(context.Background(), "0xtestaddress")
		So(err, ShouldBeNil)
		So(retrieved, ShouldNotBeNil)
		So(retrieved.Address, ShouldEqual, "0xTestAddress")

		// Test non-existent address
		retrieved, err = repo.RetrieveState(context.Background(), "0xNonExistent")
		So(err, ShouldNotBeNil)
		So(retrieved, ShouldBeNil)
	})

	Convey("Test RequestForEnqueue", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT classes with different scores
		statusCompleted := nftclass.AcquireBookNftEventsStatusCompleted
		statusFailed := nftclass.AcquireBookNftEventsStatusFailed
		score1 := 1.0
		score2 := 2.0
		score3 := 3.0

		nftClass1 := util.makeNFTClass("0xAddress1", &statusCompleted, &score1, 0, nil)
		nftClass2 := util.makeNFTClass("0xAddress2", &statusFailed, &score2, 1, nil)
		nftClass3 := util.makeNFTClass("0xAddress3", nil, &score3, 0, nil)

		// Insert test data
		for _, nc := range []*ent.NFTClass{nftClass1, nftClass2, nftClass3} {
			_, err := dbService.Client().NFTClass.Create().
				SetAddress(nc.Address).
				SetName(nc.Name).
				SetSymbol(nc.Symbol).
				SetDeployerAddress(nc.DeployerAddress).
				SetDeployedBlockNumber(nc.DeployedBlockNumber).
				SetLatestEventBlockNumber(nc.LatestEventBlockNumber).
				SetTotalSupply(nc.TotalSupply).
				SetMaxSupply(nc.MaxSupply).
				SetBannerImage(nc.BannerImage).
				SetFeaturedImage(nc.FeaturedImage).
				SetMintedAt(nc.MintedAt).
				SetUpdatedAt(nc.UpdatedAt).
				SetDisabledForIndexing(nc.DisabledForIndexing).
				SetAcquireBookNftEventsWeight(nc.AcquireBookNftEventsWeight).
				SetNillableAcquireBookNftEventsStatus(nc.AcquireBookNftEventsStatus).
				SetNillableAcquireBookNftEventsScore(nc.AcquireBookNftEventsScore).
				SetAcquireBookNftEventsFailedCount(nc.AcquireBookNftEventsFailedCount).
				Save(context.Background())
			So(err, ShouldBeNil)
		}

		// Test requesting 2 items (should return lowest scores first)
		results, err := repo.RequestForEnqueue(context.Background(), 2, 10)
		So(err, ShouldBeNil)
		So(results, ShouldHaveLength, 2)

		// Should be ordered by score (nulls first, then ascending)
		// nftClass3 has score 3.0, nftClass1 has score 1.0, nftClass2 has score 2.0
		// So order should be: nftClass1 (1.0), nftClass2 (2.0)
		So(results[0].Address, ShouldEqual, "0xAddress1")
		So(*results[0].AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueing)
		So(*results[0].AcquireBookNftEventsScore, ShouldEqual, 10.0)
		So(results[1].Address, ShouldEqual, "0xAddress2")
		So(*results[1].AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueing)
		So(*results[1].AcquireBookNftEventsScore, ShouldEqual, 10.0)

		// Test requesting more than available
		// nftClass3 has score 3.0, nftClass1 has score 10.0, nftClass2 has score 10.0
		// So order should be: nftClass3 (3.0), nftClass1 (10.0), nftClass2 (10.0)
		results, err = repo.RequestForEnqueue(context.Background(), 10, 10)
		So(err, ShouldBeNil)
		So(results, ShouldHaveLength, 3)
		So(results[0].Address, ShouldEqual, "0xAddress3")
		So(*results[0].AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueing)
		So(*results[0].AcquireBookNftEventsScore, ShouldEqual, 10.0)
		So(results[1].Address, ShouldEqual, "0xAddress1")
		So(*results[1].AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueing)
		So(*results[1].AcquireBookNftEventsScore, ShouldEqual, 10.0)
		So(results[2].Address, ShouldEqual, "0xAddress2")
		So(*results[2].AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueing)
		So(*results[2].AcquireBookNftEventsScore, ShouldEqual, 10.0)
	})

	Convey("Test RequestForEnqueue with disabled items", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT classes - one disabled, one enabled
		statusCompleted := nftclass.AcquireBookNftEventsStatusCompleted
		score1 := 1.0
		score2 := 2.0

		nftClass1 := util.makeNFTClass("0xAddress1", &statusCompleted, &score1, 0, nil)
		nftClass1.DisabledForIndexing = true

		nftClass2 := util.makeNFTClass("0xAddress2", &statusCompleted, &score2, 0, nil)
		nftClass2.DisabledForIndexing = false

		// Insert test data
		for _, nc := range []*ent.NFTClass{nftClass1, nftClass2} {
			_, err := dbService.Client().NFTClass.Create().
				SetAddress(nc.Address).
				SetName(nc.Name).
				SetSymbol(nc.Symbol).
				SetDeployerAddress(nc.DeployerAddress).
				SetDeployedBlockNumber(nc.DeployedBlockNumber).
				SetLatestEventBlockNumber(nc.LatestEventBlockNumber).
				SetTotalSupply(nc.TotalSupply).
				SetMaxSupply(nc.MaxSupply).
				SetBannerImage(nc.BannerImage).
				SetFeaturedImage(nc.FeaturedImage).
				SetMintedAt(nc.MintedAt).
				SetUpdatedAt(nc.UpdatedAt).
				SetDisabledForIndexing(nc.DisabledForIndexing).
				SetAcquireBookNftEventsWeight(nc.AcquireBookNftEventsWeight).
				SetNillableAcquireBookNftEventsStatus(nc.AcquireBookNftEventsStatus).
				SetNillableAcquireBookNftEventsScore(nc.AcquireBookNftEventsScore).
				SetAcquireBookNftEventsFailedCount(nc.AcquireBookNftEventsFailedCount).
				Save(context.Background())
			So(err, ShouldBeNil)
		}

		// Test requesting items (should only return enabled items)
		results, err := repo.RequestForEnqueue(context.Background(), 5, 10)
		So(err, ShouldBeNil)
		So(results, ShouldHaveLength, 1)
		So(results[0].Address, ShouldEqual, "0xAddress2")
		So(results[0].DisabledForIndexing, ShouldBeFalse)
	})

	Convey("Test Enqueued", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class
		statusEnqueueing := nftclass.AcquireBookNftEventsStatusEnqueueing
		nftClass := util.makeNFTClass("0xTestAddress", &statusEnqueueing, nil, 0, nil)

		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			SetNillableAcquireBookNftEventsStatus(nftClass.AcquireBookNftEventsStatus).
			SetNillableAcquireBookNftEventsScore(nftClass.AcquireBookNftEventsScore).
			SetAcquireBookNftEventsFailedCount(nftClass.AcquireBookNftEventsFailedCount).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test successful enqueue
		timeoutScore := 100.0
		err = repo.Enqueued(context.Background(), nftClass, timeoutScore)
		So(err, ShouldBeNil)

		// Verify the update
		updated, err := dbService.Client().NFTClass.Query().
			Where(nftclass.AddressEqualFold("0xTestAddress")).
			Only(context.Background())
		So(err, ShouldBeNil)
		So(*updated.AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueued)
		So(*updated.AcquireBookNftEventsScore, ShouldEqual, timeoutScore)
	})

	Convey("Test EnqueueFailed", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class
		statusEnqueueing := nftclass.AcquireBookNftEventsStatusEnqueueing
		nftClass := util.makeNFTClass("0xTestAddress", &statusEnqueueing, nil, 0, nil)

		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			SetNillableAcquireBookNftEventsStatus(nftClass.AcquireBookNftEventsStatus).
			SetNillableAcquireBookNftEventsScore(nftClass.AcquireBookNftEventsScore).
			SetAcquireBookNftEventsFailedCount(nftClass.AcquireBookNftEventsFailedCount).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test enqueue failure
		retryScore := 200.0
		testError := errors.New("test enqueue error")
		err = repo.EnqueueFailed(context.Background(), nftClass, testError, retryScore)
		So(err, ShouldBeNil)

		// Verify the update
		updated, err := dbService.Client().NFTClass.Query().
			Where(nftclass.AddressEqualFold("0xTestAddress")).
			Only(context.Background())
		So(err, ShouldBeNil)
		So(*updated.AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusEnqueueFailed)
		So(*updated.AcquireBookNftEventsScore, ShouldEqual, retryScore)
		So(*updated.AcquireBookNftEventsFailedReason, ShouldEqual, testError.Error())
		So(updated.AcquireBookNftEventsFailedCount, ShouldEqual, 1)
	})

	Convey("Test Processing", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class
		statusEnqueued := nftclass.AcquireBookNftEventsStatusEnqueued
		nftClass := util.makeNFTClass("0xTestAddress", &statusEnqueued, nil, 0, nil)

		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			SetNillableAcquireBookNftEventsStatus(nftClass.AcquireBookNftEventsStatus).
			SetNillableAcquireBookNftEventsScore(nftClass.AcquireBookNftEventsScore).
			SetAcquireBookNftEventsFailedCount(nftClass.AcquireBookNftEventsFailedCount).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test processing
		timeoutScore := 300.0
		err = repo.Processing(context.Background(), nftClass, timeoutScore)
		So(err, ShouldBeNil)

		// Verify the update
		updated, err := dbService.Client().NFTClass.Query().
			Where(nftclass.AddressEqualFold("0xTestAddress")).
			Only(context.Background())
		So(err, ShouldBeNil)
		So(*updated.AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusProcessing)
		So(*updated.AcquireBookNftEventsScore, ShouldEqual, timeoutScore)
	})

	Convey("Test Completed", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class with failed state
		statusProcessing := nftclass.AcquireBookNftEventsStatusProcessing
		failedReason := "previous error"
		nftClass := util.makeNFTClass("0xTestAddress", &statusProcessing, nil, 2, &failedReason)

		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			SetNillableAcquireBookNftEventsStatus(nftClass.AcquireBookNftEventsStatus).
			SetNillableAcquireBookNftEventsScore(nftClass.AcquireBookNftEventsScore).
			SetAcquireBookNftEventsFailedCount(nftClass.AcquireBookNftEventsFailedCount).
			SetNillableAcquireBookNftEventsFailedReason(nftClass.AcquireBookNftEventsFailedReason).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test completion
		lastProcessedTime := time.Now().UTC()
		nextProcessingScore := 400.0
		err = repo.Completed(context.Background(), nftClass, lastProcessedTime, nextProcessingScore)
		So(err, ShouldBeNil)

		// Verify the update
		updated, err := dbService.Client().NFTClass.Query().
			Where(nftclass.AddressEqualFold("0xTestAddress")).
			Only(context.Background())
		So(err, ShouldBeNil)
		So(*updated.AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusCompleted)
		So(*updated.AcquireBookNftEventsScore, ShouldEqual, nextProcessingScore)
		So(updated.AcquireBookNftEventsLastProcessedTime.Unix(), ShouldEqual, lastProcessedTime.Unix())
		So(updated.AcquireBookNftEventsFailedReason, ShouldBeNil)   // Should be cleared
		So(updated.AcquireBookNftEventsFailedCount, ShouldEqual, 0) // Should be reset
	})

	Convey("Test Failed", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT class
		statusProcessing := nftclass.AcquireBookNftEventsStatusProcessing
		nftClass := util.makeNFTClass("0xTestAddress", &statusProcessing, nil, 1, nil)

		nftClass, err := dbService.Client().NFTClass.Create().
			SetAddress(nftClass.Address).
			SetName(nftClass.Name).
			SetSymbol(nftClass.Symbol).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetLatestEventBlockNumber(nftClass.LatestEventBlockNumber).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetBannerImage(nftClass.BannerImage).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMintedAt(nftClass.MintedAt).
			SetUpdatedAt(nftClass.UpdatedAt).
			SetDisabledForIndexing(nftClass.DisabledForIndexing).
			SetAcquireBookNftEventsWeight(nftClass.AcquireBookNftEventsWeight).
			SetNillableAcquireBookNftEventsStatus(nftClass.AcquireBookNftEventsStatus).
			SetNillableAcquireBookNftEventsScore(nftClass.AcquireBookNftEventsScore).
			SetAcquireBookNftEventsFailedCount(nftClass.AcquireBookNftEventsFailedCount).
			Save(context.Background())
		So(err, ShouldBeNil)

		// Test failure
		retryScore := 500.0
		testError := errors.New("test processing error")
		err = repo.Failed(context.Background(), nftClass, testError, retryScore)
		So(err, ShouldBeNil)

		// Verify the update
		updated, err := dbService.Client().NFTClass.Query().
			Where(nftclass.AddressEqualFold("0xTestAddress")).
			Only(context.Background())
		So(err, ShouldBeNil)
		So(*updated.AcquireBookNftEventsStatus, ShouldEqual, nftclass.AcquireBookNftEventsStatusFailed)
		So(*updated.AcquireBookNftEventsScore, ShouldEqual, retryScore)
		So(*updated.AcquireBookNftEventsFailedReason, ShouldEqual, testError.Error())
		So(updated.AcquireBookNftEventsFailedCount, ShouldEqual, 2) // Incremented from 1
	})

	Convey("Test RequestForEnqueue ordering with null scores", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		repo := database.MakeNFTClassAcquireBookNFTEventsRepository(dbService)

		// Create test NFT classes with null scores and different statuses
		statusCompleted := nftclass.AcquireBookNftEventsStatusCompleted
		statusFailed := nftclass.AcquireBookNftEventsStatusFailed
		score1 := 1.0

		// nftClass1: has score, completed status
		nftClass1 := util.makeNFTClass("0xAddress1", &statusCompleted, &score1, 0, nil)

		// nftClass2: null score, failed status
		nftClass2 := util.makeNFTClass("0xAddress2", &statusFailed, nil, 0, nil)

		// nftClass3: null score, completed status
		nftClass3 := util.makeNFTClass("0xAddress3", &statusCompleted, nil, 0, nil)

		// Insert test data
		for _, nc := range []*ent.NFTClass{nftClass1, nftClass2, nftClass3} {
			_, err := dbService.Client().NFTClass.Create().
				SetAddress(nc.Address).
				SetName(nc.Name).
				SetSymbol(nc.Symbol).
				SetDeployerAddress(nc.DeployerAddress).
				SetDeployedBlockNumber(nc.DeployedBlockNumber).
				SetLatestEventBlockNumber(nc.LatestEventBlockNumber).
				SetTotalSupply(nc.TotalSupply).
				SetMaxSupply(nc.MaxSupply).
				SetBannerImage(nc.BannerImage).
				SetFeaturedImage(nc.FeaturedImage).
				SetMintedAt(nc.MintedAt).
				SetUpdatedAt(nc.UpdatedAt).
				SetDisabledForIndexing(nc.DisabledForIndexing).
				SetAcquireBookNftEventsWeight(nc.AcquireBookNftEventsWeight).
				SetNillableAcquireBookNftEventsStatus(nc.AcquireBookNftEventsStatus).
				SetNillableAcquireBookNftEventsScore(nc.AcquireBookNftEventsScore).
				SetAcquireBookNftEventsFailedCount(nc.AcquireBookNftEventsFailedCount).
				Save(context.Background())
			So(err, ShouldBeNil)
		}

		// Test ordering: nulls first, then by score ascending
		results, err := repo.RequestForEnqueue(context.Background(), 3, 10)
		So(err, ShouldBeNil)
		So(results, ShouldHaveLength, 3)

		// Order should be: nftClass2 (null score), nftClass3 (null score), nftClass1 (score 1.0)
		So(results[0].Address, ShouldEqual, "0xAddress2")
		So(results[1].Address, ShouldEqual, "0xAddress3")
		So(results[2].Address, ShouldEqual, "0xAddress1")
	})
}
