package database

import (
	"context"
	"errors"
	"slices"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
)

type EVMEventRepository interface {
	GetEvmEventById(ctx context.Context, id int) (*ent.EVMEvent, error)

	GetEVMEventsByStatus(ctx context.Context, status evmevent.Status) ([]*ent.EVMEvent, error)

	InsertEvmEventsIfNeeded(
		ctx context.Context,

		evmEvents []*ent.EVMEvent,
	) ([]*ent.EVMEvent, error)

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
