package nftclassacquirebooknftevent_test

import (
	"testing"
	"time"

	"likenft-indexer/internal/logic/nftclassacquirebooknftevent"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeCalculateNextProcessingScoreFn(t *testing.T) {
	Convey("Test MakeCalculateNextProcessingScoreFn", t, func() {
		Convey("Should create function with correct parameters", func() {
			blockHeightContribution := 1.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			So(fn, ShouldNotBeNil)
		})

		Convey("Should calculate score correctly with weight 0", func() {
			blockHeightContribution := 1.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			lastProcessedTime := time.Unix(1000, 0)
			blockHeight := uint64(100)
			weight := 0.0

			score := fn(lastProcessedTime, blockHeight, weight)

			// Expected: blockHeight * blockHeightContribution + lastProcessedTime.Unix() + lineOfWeight(weight) * weightContribution
			// lineOfWeight(0) = (1-0)*10 + 0*1 = 10
			// = 100 * 1.0 + 1000 + 10 * 60.0
			// = 100 + 1000 + 600
			// = 1700.0
			expectedScore := 1700.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should calculate score correctly with weight 1", func() {
			blockHeightContribution := 1.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			lastProcessedTime := time.Unix(1000, 0)
			blockHeight := uint64(100)
			weight := 1.0

			score := fn(lastProcessedTime, blockHeight, weight)

			// Expected: blockHeight * blockHeightContribution + lastProcessedTime.Unix() + lineOfWeight(weight) * weightContribution
			// lineOfWeight(1) = (1-1)*10 + 1*1 = 0 + 1 = 1
			// = 100 * 1.0 + 1000 + 1 * 60.0
			// = 100 + 1000 + 60
			// = 1160.0
			expectedScore := 1160.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should calculate score correctly with weight 0.5", func() {
			blockHeightContribution := 1.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			lastProcessedTime := time.Unix(1000, 0)
			blockHeight := uint64(100)
			weight := 0.5

			score := fn(lastProcessedTime, blockHeight, weight)

			// Expected: blockHeight * blockHeightContribution + lastProcessedTime.Unix() + lineOfWeight(weight) * weightContribution
			// lineOfWeight(0.5) = (1-0.5)*10 + 0.5*1 = 5 + 0.5 = 5.5
			// = 100 * 1.0 + 1000 + 5.5 * 60.0
			// = 100 + 1000 + 330
			// = 1430.0
			expectedScore := 1430.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle different block heights correctly", func() {
			blockHeightContribution := 2.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			lastProcessedTime := time.Unix(1000, 0)
			weight := 0.0

			score1 := fn(lastProcessedTime, 100, weight)
			score2 := fn(lastProcessedTime, 200, weight)

			// Difference should be (200-100) * blockHeightContribution = 100 * 2.0 = 200
			So(score2-score1, ShouldEqual, 200.0)
		})

		Convey("Should handle different timestamps correctly", func() {
			blockHeightContribution := 1.0
			weight0Constant := 10.0
			weight1Constant := 1.0
			weightContribution := 60.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			blockHeight := uint64(100)
			weight := 0.0

			time1 := time.Unix(1000, 0)
			time2 := time.Unix(2000, 0)

			score1 := fn(time1, blockHeight, weight)
			score2 := fn(time2, blockHeight, weight)

			// Difference should be 2000 - 1000 = 1000
			So(score2-score1, ShouldEqual, 1000.0)
		})

		Convey("Should handle edge case with zero values", func() {
			blockHeightContribution := 0.0
			weight0Constant := 0.0
			weight1Constant := 0.0
			weightContribution := 0.0

			fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(
				blockHeightContribution,
				weight0Constant,
				weight1Constant,
				weightContribution,
			)

			lastProcessedTime := time.Unix(1000, 0)
			blockHeight := uint64(100)
			weight := 0.5

			score := fn(lastProcessedTime, blockHeight, weight)

			// Should only be the timestamp contribution
			expectedScore := float64(lastProcessedTime.Unix())
			So(score, ShouldEqual, expectedScore)
		})
	})
}

func TestMakeCalculateTimeoutScoreFn(t *testing.T) {
	Convey("Test MakeCalculateTimeoutScoreFn", t, func() {
		Convey("Should create function with correct parameters", func() {
			timeoutSeconds := 300

			fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(timeoutSeconds)

			So(fn, ShouldNotBeNil)
		})

		Convey("Should calculate timeout score correctly", func() {
			timeoutSeconds := 300

			fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(timeoutSeconds)

			asOf := time.Unix(1000, 0)
			score := fn(asOf)

			// Expected: asOf.Unix() + timeoutSeconds = 1000 + 300 = 1300
			expectedScore := 1300.0
			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle different timestamps correctly", func() {
			timeoutSeconds := 300

			fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(timeoutSeconds)

			time1 := time.Unix(1000, 0)
			time2 := time.Unix(2000, 0)

			score1 := fn(time1)
			score2 := fn(time2)

			// Difference should be 2000 - 1000 = 1000
			So(score2-score1, ShouldEqual, 1000.0)
		})

		Convey("Should handle zero timeout", func() {
			timeoutSeconds := 0

			fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(timeoutSeconds)

			asOf := time.Unix(1000, 0)
			score := fn(asOf)

			// Should just be the timestamp
			expectedScore := 1000.0
			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle negative timeout", func() {
			timeoutSeconds := -100

			fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(timeoutSeconds)

			asOf := time.Unix(1000, 0)
			score := fn(asOf)

			// Should be timestamp minus 100
			expectedScore := 900.0
			So(score, ShouldEqual, expectedScore)
		})
	})
}

func TestMakeCalculateRetryScoreFn(t *testing.T) {
	Convey("Test MakeCalculateRetryScoreFn", t, func() {
		Convey("Should create function with correct parameters", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			So(fn, ShouldNotBeNil)
		})

		Convey("Should calculate retry score correctly for first failure", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 1

			score := fn(asOf, failedCount)

			// Expected: asOf.Unix() + initialTimeoutSeconds * exponentialBackoffCoeff^failedCount
			// = 1000 + 60 * 2.0^1 = 1000 + 60 * 2 = 1000 + 120 = 1120
			expectedScore := 1120.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should calculate retry score correctly for multiple failures", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 3

			score := fn(asOf, failedCount)

			// Expected: asOf.Unix() + initialTimeoutSeconds * exponentialBackoffCoeff^failedCount
			// = 1000 + 60 * 2.0^3 = 1000 + 60 * 8 = 1000 + 480 = 1480
			expectedScore := 1480.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should respect max timeout limit", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 300 // Lower than what would be calculated

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 5

			score := fn(asOf, failedCount)

			// Expected backoff time would be 60 * 2^5 = 1920, but max is 300
			// So score should be 1000 + 300 = 1300
			expectedScore := 1300.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle zero failed count", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 0

			score := fn(asOf, failedCount)

			// Expected: asOf.Unix() + initialTimeoutSeconds * exponentialBackoffCoeff^0
			// = 1000 + 60 * 1 = 1000 + 60 = 1060
			expectedScore := 1060.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle different exponential backoff coefficients", func() {
			initialTimeoutSeconds := 60
			exponentialBackoffCoeff := 1.5
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 2

			score := fn(asOf, failedCount)

			// Expected: asOf.Unix() + initialTimeoutSeconds * exponentialBackoffCoeff^failedCount
			// = 1000 + 60 * 1.5^2 = 1000 + 60 * 2.25 = 1000 + 135 = 1135
			expectedScore := 1135.0

			So(score, ShouldEqual, expectedScore)
		})

		Convey("Should handle edge case with zero initial timeout", func() {
			initialTimeoutSeconds := 0
			exponentialBackoffCoeff := 2.0
			maxTimeoutSeconds := 3600

			fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(
				initialTimeoutSeconds,
				exponentialBackoffCoeff,
				maxTimeoutSeconds,
			)

			asOf := time.Unix(1000, 0)
			failedCount := 5

			score := fn(asOf, failedCount)

			// Should just be the timestamp since backoff time is 0
			expectedScore := 1000.0
			So(score, ShouldEqual, expectedScore)
		})
	})
}

// Benchmark tests for performance
func BenchmarkMakeCalculateNextProcessingScoreFn(b *testing.B) {
	fn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
	lastProcessedTime := time.Unix(1000, 0)
	blockHeight := uint64(100)
	weight := 0.5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(lastProcessedTime, blockHeight, weight)
	}
}

func BenchmarkMakeCalculateTimeoutScoreFn(b *testing.B) {
	fn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
	asOf := time.Unix(1000, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(asOf)
	}
}

func BenchmarkMakeCalculateRetryScoreFn(b *testing.B) {
	fn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)
	asOf := time.Unix(1000, 0)
	failedCount := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(asOf, failedCount)
	}
}
