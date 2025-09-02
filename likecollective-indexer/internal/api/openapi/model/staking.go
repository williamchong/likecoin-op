package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/openapi/api"
)

func MakeStaking(staking *ent.Staking) api.Staking {
	return api.Staking{
		BookNft:             api.EvmAddress(staking.BookNFT),
		Account:             api.EvmAddress(staking.Account),
		StakedAmount:        api.Uint256(staking.StakedAmount),
		PendingRewardAmount: api.Uint256(staking.PendingRewardAmount),
		ClaimedRewardAmount: api.Uint256(staking.ClaimedRewardAmount),
	}
}
