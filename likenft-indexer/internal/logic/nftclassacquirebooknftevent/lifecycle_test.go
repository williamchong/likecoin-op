package nftclassacquirebooknftevent_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/logic/nftclassacquirebooknftevent"

	"github.com/hibiken/asynq"
	. "github.com/smartystreets/goconvey/convey"
)

// Mock repository for testing
type mockNFTClassAcquireBookNFTEventsRepository struct {
	retrieveStateFunc func(ctx context.Context, contractAddress string) (*ent.NFTClass, error)
	enqueuedFunc      func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error
	enqueueFailedFunc func(ctx context.Context, m *ent.NFTClass, err error, retryScore float64) error
	processingFunc    func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error
	completedFunc     func(ctx context.Context, m *ent.NFTClass, lastProcessedTime time.Time, nextProcessingScore float64) error
	failedFunc        func(ctx context.Context, m *ent.NFTClass, err error, retryScore float64) error
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) RetrieveState(ctx context.Context, contractAddress string) (*ent.NFTClass, error) {
	if m.retrieveStateFunc != nil {
		return m.retrieveStateFunc(ctx, contractAddress)
	}
	return nil, errors.New("retrieveStateFunc not implemented")
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) RequestForEnqueue(ctx context.Context, count int, timeoutScore float64) ([]*ent.NFTClass, error) {
	return nil, errors.New("not implemented")
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) Enqueued(ctx context.Context, m2 *ent.NFTClass, timeoutScore float64) error {
	if m.enqueuedFunc != nil {
		return m.enqueuedFunc(ctx, m2, timeoutScore)
	}
	return nil
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) EnqueueFailed(ctx context.Context, m2 *ent.NFTClass, err error, retryScore float64) error {
	if m.enqueueFailedFunc != nil {
		return m.enqueueFailedFunc(ctx, m2, err, retryScore)
	}
	return nil
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) Processing(ctx context.Context, m2 *ent.NFTClass, timeoutScore float64) error {
	if m.processingFunc != nil {
		return m.processingFunc(ctx, m2, timeoutScore)
	}
	return nil
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) Completed(ctx context.Context, m2 *ent.NFTClass, lastProcessedTime time.Time, nextProcessingScore float64) error {
	if m.completedFunc != nil {
		return m.completedFunc(ctx, m2, lastProcessedTime, nextProcessingScore)
	}
	return nil
}

func (m *mockNFTClassAcquireBookNFTEventsRepository) Failed(ctx context.Context, m2 *ent.NFTClass, err error, retryScore float64) error {
	if m.failedFunc != nil {
		return m.failedFunc(ctx, m2, err, retryScore)
	}
	return nil
}

// Helper function to create test NFTClass
func createTestNFTClass(address string, status *nftclass.AcquireBookNftEventsStatus, failedCount int) *ent.NFTClass {
	return &ent.NFTClass{
		Address:                               address,
		AcquireBookNftEventsStatus:            status,
		AcquireBookNftEventsFailedCount:       failedCount,
		AcquireBookNftEventsWeight:            1.0,
		DeployedBlockNumber:                   100,
		AcquireBookNftEventsLastProcessedTime: nil,
		AcquireBookNftEventsScore:             nil,
		AcquireBookNftEventsFailedReason:      nil,
	}
}

func TestMakeNFTClassAcquireBookNFTEventLifecycle(t *testing.T) {
	Convey("Test MakeNFTClassAcquireBookNFTEventLifecycle", t, func() {
		Convey("Should create lifecycle with correct properties", func() {
			repo := &mockNFTClassAcquireBookNFTEventsRepository{}
			nftClass := createTestNFTClass("0x123", nil, 0)

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(lifecycle, ShouldNotBeNil)
		})
	})
}

func TestMakeNFTClassAcquireBookNFTEventLifecycles(t *testing.T) {
	Convey("Test MakeNFTClassAcquireBookNFTEventLifecycles", t, func() {
		Convey("Should create lifecycles for multiple NFT classes", func() {
			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			nftClasses := []*ent.NFTClass{
				createTestNFTClass("0x123", nil, 0),
				createTestNFTClass("0x456", nil, 0),
				createTestNFTClass("0x789", nil, 0),
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycles := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycles(
				repo,
				nftClasses,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(lifecycles, ShouldNotBeNil)
			So(len(lifecycles), ShouldEqual, 3)
		})

		Convey("Should handle empty slice", func() {
			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycles := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycles(
				repo,
				[]*ent.NFTClass{},
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(lifecycles, ShouldNotBeNil)
			So(len(lifecycles), ShouldEqual, 0)
		})

		Convey("Should handle nil slice", func() {
			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycles := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycles(
				repo,
				nil,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(lifecycles, ShouldNotBeNil)
			So(len(lifecycles), ShouldEqual, 0)
		})
	})
}

func TestMakeNFTClassAcquireBookNFTEventLifecycleFromAddress(t *testing.T) {
	Convey("Test MakeNFTClassAcquireBookNFTEventLifecycleFromAddress", t, func() {
		Convey("Should create lifecycle from address successfully", func() {
			expectedNFTClass := createTestNFTClass("0x123", nil, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				retrieveStateFunc: func(ctx context.Context, contractAddress string) (*ent.NFTClass, error) {
					So(contractAddress, ShouldEqual, "0x123")
					return expectedNFTClass, nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle, err := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycleFromAddress(
				context.Background(),
				repo,
				"0x123",
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(err, ShouldBeNil)
			So(lifecycle, ShouldNotBeNil)
		})

		Convey("Should return error when repository fails", func() {
			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				retrieveStateFunc: func(ctx context.Context, contractAddress string) (*ent.NFTClass, error) {
					return nil, errors.New("database error")
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle, err := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycleFromAddress(
				context.Background(),
				repo,
				"0x123",
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "database error")
			So(lifecycle, ShouldBeNil)
		})
	})
}

func TestNFTClassAcquireBookNFTEventLifecycle_WithEnqueueing(t *testing.T) {
	Convey("Test WithEnqueueing", t, func() {
		Convey("Should succeed when status is enqueueing", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueueing
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				enqueuedFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					So(m.Address, ShouldEqual, "0x123")
					return nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			enqueueingFn := func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
				return &asynq.TaskInfo{ID: "test-task-id"}, nil
			}

			taskInfo, err := lifecycle.WithEnqueueing(context.Background(), enqueueingFn)

			So(err, ShouldBeNil)
			So(taskInfo, ShouldNotBeNil)
			So(taskInfo.ID, ShouldEqual, "test-task-id")
		})

		Convey("Should fail when status is not enqueueing", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueued
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			enqueueingFn := func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
				return &asynq.TaskInfo{ID: "test-task-id"}, nil
			}

			taskInfo, err := lifecycle.WithEnqueueing(context.Background(), enqueueingFn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "err lifecycle is not enqueueing")
			So(taskInfo, ShouldBeNil)
		})

		Convey("Should fail when status is nil", func() {
			nftClass := createTestNFTClass("0x123", nil, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			enqueueingFn := func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
				return &asynq.TaskInfo{ID: "test-task-id"}, nil
			}

			taskInfo, err := lifecycle.WithEnqueueing(context.Background(), enqueueingFn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "err lifecycle is not enqueueing")
			So(taskInfo, ShouldBeNil)
		})

		Convey("Should handle enqueued failure", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueueing
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				enqueuedFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					return errors.New("enqueued failed")
				},
				enqueueFailedFunc: func(ctx context.Context, m *ent.NFTClass, err error, retryScore float64) error {
					So(err.Error(), ShouldContainSubstring, "enqueued failed")
					return nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			enqueueingFn := func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
				return &asynq.TaskInfo{ID: "test-task-id"}, nil
			}

			taskInfo, err := lifecycle.WithEnqueueing(context.Background(), enqueueingFn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "enqueued failed")
			So(taskInfo, ShouldBeNil)
		})

		Convey("Should handle enqueueingFn failure", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueueing
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				enqueuedFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					return nil
				},
				enqueueFailedFunc: func(ctx context.Context, m *ent.NFTClass, err error, retryScore float64) error {
					So(err.Error(), ShouldContainSubstring, "enqueueing failed")
					return nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			enqueueingFn := func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error) {
				return nil, errors.New("enqueueing failed")
			}

			taskInfo, err := lifecycle.WithEnqueueing(context.Background(), enqueueingFn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "enqueueing failed")
			So(taskInfo, ShouldBeNil)
		})
	})
}

func TestNFTClassAcquireBookNFTEventLifecycle_WithEnqueued(t *testing.T) {
	Convey("Test WithEnqueued", t, func() {
		Convey("Should succeed when status is enqueued", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueued
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				processingFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					So(m.Address, ShouldEqual, "0x123")
					return nil
				},
				completedFunc: func(ctx context.Context, m *ent.NFTClass, lastProcessedTime time.Time, nextProcessingScore float64) error {
					So(m.Address, ShouldEqual, "0x123")
					return nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return nil
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldBeNil)
		})

		Convey("Should fail when status is not enqueued", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueueing
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return nil
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "err lifecycle is not enqueued")
		})

		Convey("Should fail when status is nil", func() {
			nftClass := createTestNFTClass("0x123", nil, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return nil
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "err lifecycle is not enqueued")
		})

		Convey("Should handle processing failure", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueued
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				processingFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					return errors.New("processing failed")
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return nil
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "processing failed")
		})

		Convey("Should handle function failure", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueued
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				processingFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					return nil
				},
				failedFunc: func(ctx context.Context, m *ent.NFTClass, err error, retryScore float64) error {
					So(err.Error(), ShouldContainSubstring, "function failed")
					return nil
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return errors.New("function failed")
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "function failed")
		})

		Convey("Should handle completed failure", func() {
			status := nftclass.AcquireBookNftEventsStatusEnqueued
			nftClass := createTestNFTClass("0x123", &status, 0)

			repo := &mockNFTClassAcquireBookNFTEventsRepository{
				processingFunc: func(ctx context.Context, m *ent.NFTClass, timeoutScore float64) error {
					return nil
				},
				completedFunc: func(ctx context.Context, m *ent.NFTClass, lastProcessedTime time.Time, nextProcessingScore float64) error {
					return errors.New("completed failed")
				},
			}

			calculateScoreFn := nftclassacquirebooknftevent.MakeCalculateNextProcessingScoreFn(1.0, 10.0, 1.0, 60.0)
			calculateTimeoutScoreFn := nftclassacquirebooknftevent.MakeCalculateTimeoutScoreFn(300)
			calculateRetryScoreFn := nftclassacquirebooknftevent.MakeCalculateRetryScoreFn(60, 2.0, 3600)

			lifecycle := nftclassacquirebooknftevent.MakeNFTClassAcquireBookNFTEventLifecycle(
				repo,
				nftClass,
				calculateScoreFn,
				calculateTimeoutScoreFn,
				calculateRetryScoreFn,
			)

			fn := func(nftClass *ent.NFTClass) error {
				return nil
			}

			err := lifecycle.WithEnqueued(context.Background(), fn)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "completed failed")
		})
	})
}
