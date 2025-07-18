package database

import (
	"context"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"

	"entgo.io/ent/dialect/sql"
)

type NFTClassAcquireBookNFTEventsRepository interface {
	RetrieveState(
		ctx context.Context,
		contractAddress string,
	) (*ent.NFTClass, error)

	RequestForEnqueue(
		ctx context.Context,
		count int,
		timeoutScore float64,
	) ([]*ent.NFTClass, error)

	Enqueued(
		ctx context.Context,
		m *ent.NFTClass,
		timeoutScore float64,
	) error

	// Should reset the state
	// and to be enqueued in next batch
	EnqueueFailed(
		ctx context.Context,
		m *ent.NFTClass,
		err error,
		retryScore float64,
	) error

	Processing(
		ctx context.Context,
		m *ent.NFTClass,
		timeoutScore float64,
	) error

	// Should reset the state
	// and update attributes
	Completed(
		ctx context.Context,
		m *ent.NFTClass,
		lastProcessedTime time.Time,
		nextProcessingScore float64,
	) error

	// Should reset the state
	// and to be enqueued in next batch
	Failed(
		ctx context.Context,
		m *ent.NFTClass,
		err error,
		retryTimeoutScore float64,
	) error
}

type nftClassAcquireBookNFTEventsRepository struct {
	dbService Service
}

func MakeNFTClassAcquireBookNFTEventsRepository(
	dbService Service,
) NFTClassAcquireBookNFTEventsRepository {
	return &nftClassAcquireBookNFTEventsRepository{
		dbService: dbService,
	}
}

func (r *nftClassAcquireBookNFTEventsRepository) RetrieveState(
	ctx context.Context,
	contractAddress string,
) (*ent.NFTClass, error) {
	return r.dbService.Client().NFTClass.Query().Where(
		nftclass.AddressEqualFold(contractAddress),
	).Only(ctx)
}

func (r *nftClassAcquireBookNFTEventsRepository) RequestForEnqueue(
	ctx context.Context,
	count int,
	timeoutScore float64,
) ([]*ent.NFTClass, error) {
	resChan := make(chan []*ent.NFTClass, 1)

	err := WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		var addressRows []struct {
			Address string `json:"address"`
		}
		err := tx.NFTClass.Query().
			Select(nftclass.FieldAddress).
			Where(
				nftclass.DisabledForIndexingEQ(false),
			).Order(
			nftclass.ByAcquireBookNftEventsScore(sql.OrderNullsFirst(), sql.OrderAsc()),
		).Limit(count).
			Select(nftclass.FieldAddress).
			Scan(ctx, &addressRows)

		if err != nil {
			return err
		}

		if len(addressRows) == 0 {
			resChan <- []*ent.NFTClass{}
			return nil
		}

		addresses := make([]string, len(addressRows))
		for i, addressRow := range addressRows {
			addresses[i] = addressRow.Address
		}

		if err := tx.NFTClass.Update().
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusEnqueueing).
			SetAcquireBookNftEventsScore(timeoutScore).
			Where(
				nftclass.AddressIn(addresses...),
			).
			Exec(ctx); err != nil {
			return err
		}

		nftClasses, err := tx.NFTClass.Query().
			Where(
				nftclass.AddressIn(addresses...),
			).Order(
			nftclass.ByAcquireBookNftEventsScore(sql.OrderNullsFirst(), sql.OrderAsc()),
		).All(ctx)

		nftClassByAddress := make(map[string]*ent.NFTClass)
		for _, nftClass := range nftClasses {
			nftClassByAddress[nftClass.Address] = nftClass
		}

		orderedNftClasses := make([]*ent.NFTClass, len(addresses))
		for i, address := range addresses {
			orderedNftClasses[i] = nftClassByAddress[address]
		}

		resChan <- orderedNftClasses
		return nil
	})

	if err != nil {
		return nil, err
	}
	return <-resChan, nil
}

func (r *nftClassAcquireBookNFTEventsRepository) Enqueued(
	ctx context.Context,
	m *ent.NFTClass,
	timeoutScore float64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusEnqueued).
			SetAcquireBookNftEventsScore(timeoutScore).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) EnqueueFailed(
	ctx context.Context,
	m *ent.NFTClass,
	err error,
	retryScore float64,
) error {

	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(
				nftclass.AcquireBookNftEventsStatusEnqueueFailed,
			).
			SetAcquireBookNftEventsFailedReason(
				err.Error(),
			).
			SetAcquireBookNftEventsScore(retryScore).
			SetAcquireBookNftEventsFailedCount(m.AcquireBookNftEventsFailedCount + 1).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Processing(
	ctx context.Context,
	m *ent.NFTClass,
	timeoutScore float64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusProcessing).
			SetAcquireBookNftEventsScore(timeoutScore).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Completed(
	ctx context.Context,
	m *ent.NFTClass,
	lastProcessedTime time.Time,
	nextProcessingScore float64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusCompleted).
			SetAcquireBookNftEventsLastProcessedTime(lastProcessedTime).
			SetAcquireBookNftEventsScore(nextProcessingScore).
			ClearAcquireBookNftEventsFailedReason().
			SetAcquireBookNftEventsFailedCount(0).
			Exec(ctx)
	})
}

func (r *nftClassAcquireBookNFTEventsRepository) Failed(
	ctx context.Context,
	m *ent.NFTClass,
	err error,
	retryScore float64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		return tx.NFTClass.UpdateOne(m).
			Where(
				nftclass.AddressEqualFold(m.Address),
			).
			SetAcquireBookNftEventsStatus(nftclass.AcquireBookNftEventsStatusFailed).
			SetAcquireBookNftEventsFailedReason(err.Error()).
			SetAcquireBookNftEventsScore(retryScore).
			SetAcquireBookNftEventsFailedCount(m.AcquireBookNftEventsFailedCount + 1).
			Exec(ctx)
	})
}
