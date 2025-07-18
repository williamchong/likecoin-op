package nftclassacquirebooknftevent

import (
	"context"
	"errors"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/database"

	"github.com/hibiken/asynq"
)

var (
	ErrLifecycIsNotEnqueueing = errors.New("err lifecycle is not enqueueing")
	ErrLifecycIsNotEnqueued   = errors.New("err lifecycle is not enqueued")
)

type NFTClassAcquireBookNFTEventLifecycle interface {
	WithEnqueueing(
		ctx context.Context,
		enqueueingFn func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error),
	) (*asynq.TaskInfo, error)

	WithEnqueued(
		ctx context.Context,
		fn func(nftClass *ent.NFTClass) error,
	) error
}

type nftClassAcquireBookNFTEventLifecycle struct {
	nftClassAcquireBookNFTEventsRepository database.NFTClassAcquireBookNFTEventsRepository
	nftClass                               *ent.NFTClass
	calculateNextScoreFn                   CalculateNextProcessingScoreFn
	calculateTimeoutScoreFn                CalculateTimeoutScoreFn
	calculateRetryScoreFn                  CalculateRetryScoreFn
}

func MakeNFTClassAcquireBookNFTEventLifecycle(
	nftClassAcquireBookNFTEventsRepository database.NFTClassAcquireBookNFTEventsRepository,
	nftClass *ent.NFTClass,
	calculateScoreFn CalculateNextProcessingScoreFn,
	calculateTimeoutScoreFn CalculateTimeoutScoreFn,
	calculateRetryScoreFn CalculateRetryScoreFn,
) NFTClassAcquireBookNFTEventLifecycle {

	return &nftClassAcquireBookNFTEventLifecycle{
		nftClassAcquireBookNFTEventsRepository,
		nftClass,
		calculateScoreFn,
		calculateTimeoutScoreFn,
		calculateRetryScoreFn,
	}
}

func MakeNFTClassAcquireBookNFTEventLifecycles(
	nftClassAcquireBookNFTEventsRepository database.NFTClassAcquireBookNFTEventsRepository,
	nftClasses []*ent.NFTClass,
	calculateScoreFn CalculateNextProcessingScoreFn,
	calculateTimeoutScoreFn CalculateTimeoutScoreFn,
	calculateRetryScoreFn CalculateRetryScoreFn,
) []NFTClassAcquireBookNFTEventLifecycle {
	objs := make([]NFTClassAcquireBookNFTEventLifecycle, len(nftClasses))

	for i, nftClass := range nftClasses {
		objs[i] = MakeNFTClassAcquireBookNFTEventLifecycle(
			nftClassAcquireBookNFTEventsRepository,
			nftClass,
			calculateScoreFn,
			calculateTimeoutScoreFn,
			calculateRetryScoreFn,
		)
	}

	return objs
}

func MakeNFTClassAcquireBookNFTEventLifecycleFromAddress(
	ctx context.Context,
	nftClassAcquireBookNFTEventsRepository database.NFTClassAcquireBookNFTEventsRepository,
	contractAddress string,
	calculateScoreFn CalculateNextProcessingScoreFn,
	calculateTimeoutScoreFn CalculateTimeoutScoreFn,
	calculateRetryScoreFn CalculateRetryScoreFn,
) (NFTClassAcquireBookNFTEventLifecycle, error) {
	nftClass, err := nftClassAcquireBookNFTEventsRepository.RetrieveState(
		ctx,
		contractAddress,
	)
	if err != nil {
		return nil, err
	}
	return MakeNFTClassAcquireBookNFTEventLifecycle(
		nftClassAcquireBookNFTEventsRepository,
		nftClass,
		calculateScoreFn,
		calculateTimeoutScoreFn,
		calculateRetryScoreFn,
	), nil
}

func (m *nftClassAcquireBookNFTEventLifecycle) WithEnqueueing(
	ctx context.Context,
	enqueueingFn func(nftClass *ent.NFTClass) (*asynq.TaskInfo, error),
) (*asynq.TaskInfo, error) {
	if err := m.ensureEnqueueing(); err != nil {
		return nil, err
	}
	// It is possible the task is executed before the enqueued
	// state is commited.
	err := m.enqueued(ctx)
	if err != nil {
		return nil, errors.Join(
			err,
			m.enqueueFailed(ctx, err),
		)
	}
	taskInfo, err := enqueueingFn(m.nftClass)
	if err != nil {
		return nil, errors.Join(
			err,
			m.enqueueFailed(ctx, err),
		)
	}
	return taskInfo, err
}

func (m *nftClassAcquireBookNFTEventLifecycle) WithEnqueued(
	ctx context.Context,
	fn func(nftClass *ent.NFTClass) error,
) error {
	if err := m.ensureEnqueued(); err != nil {
		return err
	}
	if err := m.processing(ctx); err != nil {
		return err
	}
	if err := fn(m.nftClass); err != nil {
		return errors.Join(
			err,
			m.failed(ctx, err),
		)
	}
	// May be failed to reset state due to error. The record will be in staled.
	// Should retry the nftClassAcquireBookNFTEventsRepository.Completed call.
	// otherwise should notify developers.
	return m.completed(ctx, time.Now().UTC())
}

func (m *nftClassAcquireBookNFTEventLifecycle) ensureEnqueueing() error {
	if m.nftClass.AcquireBookNftEventsStatus == nil ||
		*m.nftClass.AcquireBookNftEventsStatus != nftclass.AcquireBookNftEventsStatusEnqueueing {
		return ErrLifecycIsNotEnqueueing
	}
	return nil
}

func (m *nftClassAcquireBookNFTEventLifecycle) enqueued(
	ctx context.Context,
) error {
	now := time.Now().UTC()

	return m.nftClassAcquireBookNFTEventsRepository.Enqueued(
		ctx, m.nftClass, m.calculateTimeoutScoreFn(
			now,
		),
	)
}

func (m *nftClassAcquireBookNFTEventLifecycle) enqueueFailed(
	ctx context.Context,
	err error,
) error {
	now := time.Now().UTC()
	return m.nftClassAcquireBookNFTEventsRepository.EnqueueFailed(
		ctx, m.nftClass, err, m.calculateRetryScoreFn(
			now,
			m.nftClass.AcquireBookNftEventsFailedCount,
		),
	)
}

func (m *nftClassAcquireBookNFTEventLifecycle) ensureEnqueued() error {
	if m.nftClass.AcquireBookNftEventsStatus == nil ||
		*m.nftClass.AcquireBookNftEventsStatus != nftclass.AcquireBookNftEventsStatusEnqueued {
		return ErrLifecycIsNotEnqueued
	}
	return nil
}

func (m *nftClassAcquireBookNFTEventLifecycle) processing(
	ctx context.Context,
) error {
	now := time.Now().UTC()
	return m.nftClassAcquireBookNFTEventsRepository.Processing(
		ctx, m.nftClass, m.calculateTimeoutScoreFn(
			now,
		),
	)
}

func (m *nftClassAcquireBookNFTEventLifecycle) completed(
	ctx context.Context,
	lastProcessedTime time.Time,
) error {
	return m.nftClassAcquireBookNFTEventsRepository.Completed(
		ctx, m.nftClass, lastProcessedTime, m.calculateNextScoreFn(
			lastProcessedTime,
			uint64(m.nftClass.DeployedBlockNumber),
			m.nftClass.AcquireBookNftEventsWeight,
		),
	)
}

func (m *nftClassAcquireBookNFTEventLifecycle) failed(
	ctx context.Context,
	err error,
) error {
	now := time.Now().UTC()
	return m.nftClassAcquireBookNFTEventsRepository.Failed(
		ctx, m.nftClass, err, m.calculateRetryScoreFn(
			now,
			m.nftClass.AcquireBookNftEventsFailedCount,
		),
	)
}
