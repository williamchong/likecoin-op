package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/openapi/api"
)

func MakeOptAccount(e *ent.Account) api.OptAccount {
	if e == nil {
		return api.OptAccount{
			Value: api.Account{},
			Set:   false,
		}
	}
	return api.NewOptAccount(MakeAccount(e))
}

func MakeAccount(e *ent.Account) api.Account {
	return api.Account{
		ID:            e.ID,
		CosmosAddress: MakeOptString(e.CosmosAddress),
		EvmAddress:    e.EvmAddress,
		Likeid:        MakeOptString(e.Likeid),
	}
}
