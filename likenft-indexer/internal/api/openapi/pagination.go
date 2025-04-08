package openapi

import (
	"likenft-indexer/ent"
	"likenft-indexer/ent/nft"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/openapi/api"

	"entgo.io/ent/dialect/sql"
)

func (h *OpenAPIHandler) handleNFTClassPagination(
	q *ent.NFTClassQuery,
	limit api.OptInt,
	key api.OptInt,
	reverse api.OptBool,
) *ent.NFTClassQuery {
	if reverse.IsSet() && reverse.Value {
		q = q.Order(nftclass.ByID(sql.OrderDesc()))
	} else {
		q = q.Order(nftclass.ByID(sql.OrderAsc()))
	}

	if limit.IsSet() {
		q = q.Limit(limit.Value)
	} else {
		q = q.Limit(100)
	}

	if key.IsSet() {
		if reverse.IsSet() && reverse.Value {
			q = q.Where(nftclass.IDLT(key.Value))
		} else {
			q = q.Where(nftclass.IDGT(key.Value))
		}
	}

	return q
}

func (h *OpenAPIHandler) handleNFTPagination(
	q *ent.NFTQuery,
	limit api.OptInt,
	key api.OptInt,
	reverse api.OptBool,
) *ent.NFTQuery {
	if reverse.IsSet() && reverse.Value {
		q = q.Order(nft.ByID(sql.OrderDesc()))
	} else {
		q = q.Order(nft.ByID(sql.OrderAsc()))
	}

	if limit.IsSet() {
		q = q.Limit(limit.Value)
	} else {
		q = q.Limit(100)
	}

	if key.IsSet() {
		if reverse.IsSet() && reverse.Value {
			q = q.Where(nft.IDLT(key.Value))
		} else {
			q = q.Where(nft.IDGT(key.Value))
		}
	}

	return q
}

func (h *OpenAPIHandler) handleEventPagination(
	q *ent.EVMEventQuery,
	limit api.OptInt,
	page api.OptInt,
) *ent.EVMEventQuery {
	// page is 0 based
	return q.Limit(limit.Value).Offset(page.Value * limit.Value)
}
