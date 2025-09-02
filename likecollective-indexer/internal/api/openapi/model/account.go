package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/openapi/api"
)

func MakeAccount(account *ent.Account) api.Account {
	return api.Account{
		EvmAddress:          api.EvmAddress(account.EvmAddress),
		StakedAmount:        api.Uint256(account.StakedAmount),
		PendingRewardAmount: api.Uint256(account.PendingRewardAmount),
		ClaimedRewardAmount: api.Uint256(account.ClaimedRewardAmount),
	}
}
