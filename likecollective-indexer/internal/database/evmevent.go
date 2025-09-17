package database

import (
	"context"
	"errors"
	"slices"
	"strings"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/evmevent"
	"likecollective-indexer/ent/predicate"
	"likecollective-indexer/ent/schema/typeutil"
	slices_util "likecollective-indexer/internal/util/slices"

	"entgo.io/ent/dialect/sql"
)

type EVMEventRepository interface {
	InsertEvmEventsIfNeeded(
		ctx context.Context,

		evmEvents []*ent.EVMEvent,
	) ([]*ent.EVMEvent, error)
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

func (s *evmEventRepository) BaseQuery(q *ent.EVMEventQuery) *ent.EVMEventQuery {
	return q.Order(
		evmevent.ByBlockNumber(sql.OrderAsc()),
		evmevent.ByTransactionIndex(sql.OrderAsc()),
		evmevent.ByLogIndex(sql.OrderAsc()),
	)
}

func (s *evmEventRepository) InsertEvmEventsIfNeeded(
	ctx context.Context,

	allEvmEvents []*ent.EVMEvent,
) ([]*ent.EVMEvent, error) {
	resChan := make(chan []*ent.EVMEvent, 1)

	grouppedEvmEvents := slices_util.GroupBy(allEvmEvents, func(e *ent.EVMEvent) typeutil.Uint64 {
		return e.BlockNumber
	})

	err := WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		dbEvmEvents := make([]*ent.EVMEvent, 0)

		for _, evmEventsThisGroup := range grouppedEvmEvents {
			var txPredicates = make([]predicate.EVMEvent, len(evmEventsThisGroup))
			for i, e := range evmEventsThisGroup {
				txPredicates[i] = evmevent.And(
					evmevent.TransactionHashEqualFold(e.TransactionHash),
					evmevent.TransactionIndexEQ(e.TransactionIndex),
					evmevent.LogIndexEQ(e.LogIndex),
				)
			}

			dbEvmEventsThisGroup, err := s.BaseQuery(tx.EVMEvent.Query()).
				Where(evmevent.Or(txPredicates...)).All(ctx)

			if err != nil {
				return err
			}

			var eventsToBeInserted []*ent.EVMEventCreate

			for _, e := range evmEventsThisGroup {
				if !slices.ContainsFunc(dbEvmEventsThisGroup, func(dbEvmEvent *ent.EVMEvent) bool {
					return strings.EqualFold(dbEvmEvent.TransactionHash, e.TransactionHash) &&
						dbEvmEvent.TransactionIndex == e.TransactionIndex &&
						dbEvmEvent.LogIndex == e.LogIndex
				}) {
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

			dbEvmEventsThisGroup, err = s.BaseQuery(tx.EVMEvent.Query()).
				Where(evmevent.Or(txPredicates...)).All(ctx)
			if err != nil {
				return err
			}

			if len(dbEvmEventsThisGroup) != len(evmEventsThisGroup) {
				return errors.New("err len not match")
			}

			dbEvmEvents = append(dbEvmEvents, dbEvmEventsThisGroup...)
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
