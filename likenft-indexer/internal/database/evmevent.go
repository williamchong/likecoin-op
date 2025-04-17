package database

import (
	"context"
	"errors"
	"math"
	"slices"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
)

type EVMEventRepository interface {
	GetAllEvmEventsAndProcess(ctx context.Context, processor func(e *ent.EVMEvent) error) error

	GetEvmEventById(ctx context.Context, id int) (*ent.EVMEvent, error)

	GetEVMEventsByStatus(ctx context.Context, status evmevent.Status) ([]*ent.EVMEvent, error)

	InsertEvmEventsIfNeeded(
		ctx context.Context,

		evmEvents []*ent.EVMEvent,
	) ([]*ent.EVMEvent, error)

	UpdateEvmEvent(
		ctx context.Context,
		evmEvent *ent.EVMEvent,
	) error

	UpdateEvmEventStatus(
		ctx context.Context,

		evmEvent *ent.EVMEvent,
		newStatus evmevent.Status,
		failedReason *string,
	) (*ent.EVMEvent, error)

	BatchUpdateEvmEventStatusByIds(
		ctx context.Context,

		evmEventIds []int,
		newStatus evmevent.Status,
	) error
}

type evmEventRepository struct {
	dbService Service
}

var _ EVMEventRepository = &evmEventRepository{}

func MakeEVMEventRepository(
	dbService Service,
) EVMEventRepository {
	return &evmEventRepository{
		dbService: dbService,
	}
}

func (s *evmEventRepository) GetAllEvmEventsAndProcess(
	ctx context.Context,
	processor func(e *ent.EVMEvent) error,
) error {
	count, err := s.dbService.Client().EVMEvent.Query().Count(ctx)
	if err != nil {
		return err
	}

	itemsPerPage := 100
	expectedNumPages := int(math.Ceil(float64(count) / float64(100)))

	for n := range expectedNumPages {
		items, err := s.dbService.Client().EVMEvent.Query().Limit(itemsPerPage).Offset(n * itemsPerPage).All(ctx)
		if err != nil {
			return err
		}
		for _, i := range items {
			err = processor(i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *evmEventRepository) GetEvmEventById(ctx context.Context, id int) (*ent.EVMEvent, error) {
	return s.dbService.Client().EVMEvent.Get(ctx, id)
}

func (s *evmEventRepository) GetEVMEventsByStatus(ctx context.Context, status evmevent.Status) ([]*ent.EVMEvent, error) {
	return s.dbService.Client().EVMEvent.Query().Where(evmevent.StatusEQ(status)).All(ctx)
}

func (s *evmEventRepository) InsertEvmEventsIfNeeded(
	ctx context.Context,

	evmEvents []*ent.EVMEvent,
) ([]*ent.EVMEvent, error) {
	resChan := make(chan []*ent.EVMEvent, 1)

	err := WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		var txHashes []string
		for _, e := range evmEvents {
			txHashes = append(txHashes, e.TransactionHash)
		}
		dbEvmEvents, err := tx.EVMEvent.Query().Where(evmevent.TransactionHashIn(txHashes...)).All(ctx)

		if err != nil {
			return err
		}

		var dbTxHashes []string
		for _, e := range dbEvmEvents {
			dbTxHashes = append(dbTxHashes, e.TransactionHash)
		}

		var eventsToBeInserted []*ent.EVMEventCreate

		for _, e := range evmEvents {
			if slices.Index(dbTxHashes, e.TransactionHash) == -1 {
				createBuilder := tx.EVMEvent.Create().
					SetAddress(e.Address).
					SetBlockHash(e.BlockHash).
					SetBlockNumber(e.BlockNumber).
					SetChainID(e.ChainID).
					SetIndexedParams(e.IndexedParams).
					SetLogIndex(e.LogIndex).
					SetName(e.Name).
					SetNonIndexedParams(e.NonIndexedParams).
					SetRemoved(e.Removed).
					SetSignature(e.Signature).
					SetStatus(e.Status).
					SetTimestamp(e.Timestamp).
					SetTopic0(e.Topic0).
					SetTopic0Hex(e.Topic0Hex).
					SetTransactionHash(e.TransactionHash).
					SetTransactionIndex(e.TransactionIndex)
				if e.Data != nil {
					createBuilder = createBuilder.SetData(*e.Data)
				}
				if e.DataHex != nil {
					createBuilder = createBuilder.SetDataHex(*e.DataHex)
				}
				if e.Topic1 != nil {
					createBuilder = createBuilder.SetTopic1(*e.Topic1)
				}
				if e.Topic1Hex != nil {
					createBuilder = createBuilder.SetTopic1Hex(*e.Topic1Hex)
				}
				if e.Topic2 != nil {
					createBuilder = createBuilder.SetTopic2(*e.Topic2)
				}
				if e.Topic2Hex != nil {
					createBuilder = createBuilder.SetTopic2Hex(*e.Topic2Hex)
				}
				if e.Topic3 != nil {
					createBuilder = createBuilder.SetTopic3(*e.Topic3)
				}
				if e.Topic3Hex != nil {
					createBuilder = createBuilder.SetTopic3Hex(*e.Topic3Hex)
				}
				eventsToBeInserted = append(eventsToBeInserted, createBuilder)
			}
		}

		err = tx.EVMEvent.CreateBulk(eventsToBeInserted...).Exec(ctx)
		if err != nil {
			return err
		}

		dbEvmEvents, err = tx.EVMEvent.Query().Where(evmevent.TransactionHashIn(txHashes...)).All(ctx)
		if err != nil {
			return err
		}

		if len(dbEvmEvents) != len(evmEvents) {
			return errors.New("err len not match")
		}

		resChan <- dbEvmEvents
		return nil
	})

	if err != nil {
		return nil, err
	}
	results := <-resChan
	return results, nil
}

func (s *evmEventRepository) UpdateEvmEvent(
	ctx context.Context,
	evmEvent *ent.EVMEvent,
) error {
	err := WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		e, err := tx.EVMEvent.Get(ctx, evmEvent.ID)
		if err != nil {
			return err
		}
		updateBuilder := e.Update().
			SetAddress(evmEvent.Address).
			SetBlockHash(evmEvent.BlockHash).
			SetBlockNumber(evmEvent.BlockNumber).
			SetChainID(evmEvent.ChainID).
			SetIndexedParams(evmEvent.IndexedParams).
			SetLogIndex(evmEvent.LogIndex).
			SetName(evmEvent.Name).
			SetNonIndexedParams(evmEvent.NonIndexedParams).
			SetRemoved(evmEvent.Removed).
			SetSignature(evmEvent.Signature).
			SetStatus(evmEvent.Status).
			SetTimestamp(evmEvent.Timestamp).
			SetTopic0(evmEvent.Topic0).
			SetTopic0Hex(evmEvent.Topic0Hex).
			SetTransactionHash(evmEvent.TransactionHash).
			SetTransactionIndex(evmEvent.TransactionIndex)
		if evmEvent.Data != nil {
			updateBuilder = updateBuilder.SetData(*evmEvent.Data)
		}
		if evmEvent.DataHex != nil {
			updateBuilder = updateBuilder.SetDataHex(*evmEvent.DataHex)
		}
		if evmEvent.Topic1 != nil {
			updateBuilder = updateBuilder.SetTopic1(*evmEvent.Topic1)
		}
		if evmEvent.Topic1Hex != nil {
			updateBuilder = updateBuilder.SetTopic1Hex(*evmEvent.Topic1Hex)
		}
		if evmEvent.Topic2 != nil {
			updateBuilder = updateBuilder.SetTopic2(*evmEvent.Topic2)
		}
		if evmEvent.Topic2Hex != nil {
			updateBuilder = updateBuilder.SetTopic2Hex(*evmEvent.Topic2Hex)
		}
		if evmEvent.Topic3 != nil {
			updateBuilder = updateBuilder.SetTopic3(*evmEvent.Topic3)
		}
		if evmEvent.Topic3Hex != nil {
			updateBuilder = updateBuilder.SetTopic3Hex(*evmEvent.Topic3Hex)
		}
		return updateBuilder.Exec(ctx)
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *evmEventRepository) UpdateEvmEventStatus(
	ctx context.Context,

	evmEvent *ent.EVMEvent,
	newStatus evmevent.Status,
	failedReason *string,
) (*ent.EVMEvent, error) {
	updatedRecordChan := make(chan *ent.EVMEvent, 1)
	err := WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.EVMEvent.Get(ctx, evmEvent.ID)

		if err != nil {
			return err
		}

		builder := tx.EVMEvent.Update().
			SetStatus(newStatus)

		if failedReason == nil {
			builder = builder.ClearFailedReason()
		} else {
			builder = builder.SetFailedReason(*failedReason)
		}

		err = builder.
			Where(evmevent.IDEQ(evmEvent.ID)).
			Exec(ctx)

		if err != nil {
			return err
		}
		updatedEvmEvent, err := tx.EVMEvent.Get(ctx, evmEvent.ID)
		if err != nil {
			return err
		}
		updatedRecordChan <- updatedEvmEvent
		return nil
	})

	if err != nil {
		return nil, err
	}
	return <-updatedRecordChan, nil
}

func (s *evmEventRepository) BatchUpdateEvmEventStatusByIds(
	ctx context.Context,

	evmEventIds []int,
	newStatus evmevent.Status,
) error {
	return s.dbService.Client().EVMEvent.Update().
		SetStatus(newStatus).
		ClearFailedReason().
		Where(evmevent.IDIn(evmEventIds...)).
		Exec(ctx)
}
