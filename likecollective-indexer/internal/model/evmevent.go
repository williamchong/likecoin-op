package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/util/ordered"
)

var cmpBlockNumber ordered.Comparator[*ent.EVMEvent] = func(a, b *ent.EVMEvent) int {
	return ordered.Normalize(a.BlockNumber, b.BlockNumber)
}
var cmpTransactionIndex ordered.Comparator[*ent.EVMEvent] = func(a, b *ent.EVMEvent) int {
	return ordered.Normalize(a.TransactionIndex, b.TransactionIndex)
}

var topic0OrderMap map[string]int = map[string]int{}

var cmpTopic0 ordered.Comparator[*ent.EVMEvent] = func(a, b *ent.EVMEvent) int {
	orderA, ok := topic0OrderMap[a.Topic0]
	if !ok {
		return 1
	}
	orderB, ok := topic0OrderMap[b.Topic0]
	if !ok {
		return -1
	}
	return ordered.Normalize(orderA, orderB)
}

var cmpLogIndex ordered.Comparator[*ent.EVMEvent] = func(a, b *ent.EVMEvent) int {
	return ordered.Normalize(a.LogIndex, b.LogIndex)
}

var EvmEventsProcessingComparator = ordered.CombineComparators(
	cmpBlockNumber,
	cmpTransactionIndex,
	cmpTopic0,
	cmpLogIndex,
)
