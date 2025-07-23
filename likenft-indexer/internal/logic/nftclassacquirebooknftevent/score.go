package nftclassacquirebooknftevent

import (
	"math"
	"time"
)

type CalculateNextProcessingScoreFn func(
	lastProcessedTime time.Time,
	blockHeight uint64,
	weight float64,
) float64

type CalculateNextProcessingScoreFnFactory interface {
	CalculateScoreFn(
		lastProcessedTime time.Time,
		blockHeight uint64,
		weight float64,
	) float64
}

type calculateNextProcessingScoreFnFactory struct {
	blockHeightWeight float64
	timeFloor         float64
	timeCeiling       float64
	timeWeight        float64
}

// Calculate score based on the following formula
//
//	score = time + weight * weightContribution
//
// The weight coeff is interpolated by weight0Constant and weight1Constant,
// And finally multiplied by weightContribution to contribute to the final score
//
// Adjust accordingly with respect to the scheduler config such that the score
// can be consumed by the desire frequency.
//
// e.g. with scheduler cron * * * * *
//
//	weight0Constant = 10 // minutes
//	weight1Constant = 1 // minutes
//	weightContribution = 60 // 1 minute
//
// The jobs are expected to have a diff of t + 60 for weight = 1, and a diff of t + 600 for weight = 0
// when the score is re-calculated.
func MakeCalculateNextProcessingScoreFn(
	blockHeightWeight float64,
	timeFloor float64,
	timeCeiling float64,
	timeWeight float64,
) CalculateNextProcessingScoreFn {
	factory := &calculateNextProcessingScoreFnFactory{
		blockHeightWeight,
		timeFloor,
		timeCeiling,
		timeWeight,
	}
	return factory.CalculateScoreFn
}

func (f *calculateNextProcessingScoreFnFactory) CalculateScoreFn(
	lastProcessedTime time.Time,
	blockHeight uint64,
	weight float64,
) float64 {
	// ↓ last_processed_time -> ↓ score
	// ↓ block_height -> ↓ score

	// As item of higher weight will be distributed more compatly
	// The slope (rate of change with respect of time delta) should decrease while weight increase
	// i.e. w=0 => high, w=1 => low
	lineOfWeight := func(w float64) float64 {
		return (1-w)*f.timeCeiling + w*f.timeFloor
	}

	return float64(blockHeight)*f.blockHeightWeight +
		float64(lastProcessedTime.Unix()) +
		lineOfWeight(weight)*f.timeWeight
}

type CalculateTimeoutScoreFn func(
	asOf time.Time,
) float64

type calculateTimeoutScoreFnFactory struct {
	timeoutSeconds int
}

func MakeCalculateTimeoutScoreFn(
	timeoutSeconds int,
) CalculateTimeoutScoreFn {
	factory := &calculateTimeoutScoreFnFactory{
		timeoutSeconds,
	}
	return factory.CalculateScoreFn
}

func (f *calculateTimeoutScoreFnFactory) CalculateScoreFn(
	asOf time.Time,
) float64 {
	return float64(asOf.Unix()) + float64(f.timeoutSeconds)
}

type CalculateRetryScoreFn func(
	asOf time.Time,
	failedCount int,
) float64

type calculateRetryScoreFnFactory struct {
	initialTimeoutSeconds   int
	exponentialBackoffCoeff float64
	maxTimeoutSeconds       int
}

func MakeCalculateRetryScoreFn(
	initialTimeoutSeconds int,
	exponentialBackoffCoeff float64,
	maxTimeoutSeconds int,
) CalculateRetryScoreFn {
	factory := &calculateRetryScoreFnFactory{
		initialTimeoutSeconds,
		exponentialBackoffCoeff,
		maxTimeoutSeconds,
	}
	return factory.CalculateScoreFn
}

func (f *calculateRetryScoreFnFactory) CalculateScoreFn(
	asOf time.Time,
	failedCount int,
) float64 {
	backoffTime := float64(f.initialTimeoutSeconds) * math.Pow(f.exponentialBackoffCoeff, float64(failedCount))
	return float64(asOf.Unix()) + math.Min(backoffTime, float64(f.maxTimeoutSeconds))
}
