package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/openapi/api"
)

func MakeStakingEvent(stakingEvent *ent.StakingEvent) api.StakingEvent {
	switch stakingEvent.EventType {
	case ent.StakingEventTypeStaked:
		return api.StakingEvent{
			Type: api.StakingEventStakedStakingEvent,
			StakingEventStaked: api.StakingEventStaked{
				EventType: api.StakingEventStakedEventTypeStaked,
				BookNft:   api.EvmAddress(stakingEvent.BookNFT),
				Amount:    api.Uint256(stakingEvent.StakedAmountAdded),
				Datetime:  stakingEvent.DateTime,
			},
		}
	case ent.StakingEventTypeUnstaked:
		return api.StakingEvent{
			Type: api.StakingEventUnstakedStakingEvent,
			StakingEventUnstaked: api.StakingEventUnstaked{
				EventType: api.StakingEventUnstakedEventTypeUnstaked,
				BookNft:   api.EvmAddress(stakingEvent.BookNFT),
				Amount:    api.Uint256(stakingEvent.StakedAmountRemoved),
				Datetime:  stakingEvent.DateTime,
			},
		}
	case ent.StakingEventTypeRewardAdded:
		return api.StakingEvent{
			Type: api.StakingEventRewardAddedStakingEvent,
			StakingEventRewardAdded: api.StakingEventRewardAdded{
				EventType: api.StakingEventRewardAddedEventTypeRewardAdded,
				BookNft:   api.EvmAddress(stakingEvent.BookNFT),
				Amount:    api.Uint256(stakingEvent.RewardAmountAdded),
				Datetime:  stakingEvent.DateTime,
			},
		}
	case ent.StakingEventTypeRewardClaimed:
		return api.StakingEvent{
			Type: api.StakingEventRewardClaimedStakingEvent,
			StakingEventRewardClaimed: api.StakingEventRewardClaimed{
				EventType: api.StakingEventRewardClaimedEventTypeRewardClaimed,
				BookNft:   api.EvmAddress(stakingEvent.BookNFT),
				Amount:    api.Uint256(stakingEvent.RewardAmountRemoved),
				Datetime:  stakingEvent.DateTime,
			},
		}
	case ent.StakingEventTypeRewardDeposited:
		return api.StakingEvent{
			Type: api.StakingEventRewardDepositedStakingEvent,
			StakingEventRewardDeposited: api.StakingEventRewardDeposited{
				EventType: api.StakingEventRewardDepositedEventTypeRewardDeposited,
				BookNft:   api.EvmAddress(stakingEvent.BookNFT),
				Amount:    api.Uint256(stakingEvent.StakedAmountAdded),
				Datetime:  stakingEvent.DateTime,
			},
		}
	case ent.StakingEventTypeAllRewardsClaimed:
		return api.StakingEvent{
			Type: api.StakingEventAllRewardsClaimedStakingEvent,
			StakingEventAllRewardsClaimed: api.StakingEventAllRewardsClaimed{
				EventType: api.StakingEventAllRewardsClaimedEventTypeAllRewardsClaimed,
				ClaimedAmountList: []api.StakingEventAllRewardsClaimedClaimedAmountListItem{
					{
						BookNft: api.EvmAddress(stakingEvent.BookNFT),
						Amount:  api.Uint256(stakingEvent.RewardAmountRemoved),
					},
				},
				Datetime: stakingEvent.DateTime,
			},
		}
	}

	panic("unknown staking event type")
}
