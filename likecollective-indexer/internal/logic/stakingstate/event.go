package stakingstate

import (
	"errors"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/ent/stakingevent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/like_collective"

	"github.com/holiman/uint256"
)

var ErrGetStakingEventsFromEvent = errors.New("failed to get staking events from event")

func GetStakingEventsFromEvent(event *ent.EVMEvent) ([]*ent.StakingEvent, error) {
	log := logConverter.ConvertEvmEventToLog(event)

	if event.Name == "Staked" {
		stakedEvent := new(like_collective.LikeCollectiveStaked)
		if err := logConverter.UnpackLog(log, stakedEvent); err != nil {
			return nil, err
		}
		stakedAmount, _ := uint256.FromBig(stakedEvent.StakedAmount)
		if stakedAmount == nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("failed to convert staked amount to uint256"),
			)
		}
		return []*ent.StakingEvent{
				{
					TransactionHash:   event.TransactionHash,
					TransactionIndex:  event.TransactionIndex,
					BlockNumber:       event.BlockNumber,
					LogIndex:          event.LogIndex,
					EventType:         stakingevent.EventTypeStaked,
					AccountEvmAddress: stakedEvent.Account.String(),
					NftClassAddress:   stakedEvent.BookNFT.String(),
					StakedAmountAdded: typeutil.Uint256(stakedAmount),
					Datetime:          event.Timestamp,
				},
			},
			nil
	}

	if event.Name == "Unstaked" {
		unstakedEvent := new(like_collective.LikeCollectiveUnstaked)
		if err := logConverter.UnpackLog(log, unstakedEvent); err != nil {
			return nil, err
		}
		unstakedAmount, _ := uint256.FromBig(unstakedEvent.StakedAmount)
		if unstakedAmount == nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("failed to convert staked amount to uint256"),
			)
		}
		return []*ent.StakingEvent{
			{
				TransactionHash:     event.TransactionHash,
				TransactionIndex:    event.TransactionIndex,
				BlockNumber:         event.BlockNumber,
				LogIndex:            event.LogIndex,
				EventType:           stakingevent.EventTypeUnstaked,
				NftClassAddress:     unstakedEvent.BookNFT.String(),
				AccountEvmAddress:   unstakedEvent.Account.String(),
				StakedAmountRemoved: typeutil.Uint256(unstakedAmount),
				Datetime:            event.Timestamp,
			},
		}, nil
	}

	if event.Name == "RewardClaimed" {
		rewardClaimedEvent := new(like_collective.LikeCollectiveRewardClaimed)
		if err := logConverter.UnpackLog(log, rewardClaimedEvent); err != nil {
			return nil, err
		}
		rewardAmount, _ := uint256.FromBig(rewardClaimedEvent.RewardedAmount)
		if rewardAmount == nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("rewardClaimedEvent: failed to convert reward amount to uint256"),
			)
		}
		return []*ent.StakingEvent{
			{
				TransactionHash:            event.TransactionHash,
				TransactionIndex:           event.TransactionIndex,
				BlockNumber:                event.BlockNumber,
				LogIndex:                   event.LogIndex,
				EventType:                  stakingevent.EventTypeRewardClaimed,
				NftClassAddress:            rewardClaimedEvent.BookNFT.String(),
				AccountEvmAddress:          rewardClaimedEvent.Account.String(),
				PendingRewardAmountRemoved: typeutil.Uint256(rewardAmount),
				ClaimedRewardAmountAdded:   typeutil.Uint256(rewardAmount),
				Datetime:                   event.Timestamp,
			},
		}, nil
	}

	if event.Name == "RewardDeposited" {
		rewardDepositedEvent := new(like_collective.LikeCollectiveRewardDeposited)
		if err := logConverter.UnpackLog(log, rewardDepositedEvent); err != nil {
			return nil, err
		}
		rewardAmount, _ := uint256.FromBig(rewardDepositedEvent.RewardedAmount)
		if rewardAmount == nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("rewardDepositedEvent: failed to convert reward amount to uint256"),
			)
		}
		return []*ent.StakingEvent{
			{
				TransactionHash:          event.TransactionHash,
				TransactionIndex:         event.TransactionIndex,
				BlockNumber:              event.BlockNumber,
				LogIndex:                 event.LogIndex,
				EventType:                stakingevent.EventTypeRewardDeposited,
				NftClassAddress:          rewardDepositedEvent.BookNFT.String(),
				AccountEvmAddress:        rewardDepositedEvent.Account.String(),
				PendingRewardAmountAdded: typeutil.Uint256(rewardAmount),
				Datetime:                 event.Timestamp,
			},
		}, nil
	}

	if event.Name == "AllRewardClaimed" {
		allRewardsClaimedEvent := new(like_collective.LikeCollectiveAllRewardClaimed)
		if err := logConverter.UnpackLog(log, allRewardsClaimedEvent); err != nil {
			return nil, err
		}
		stakingEvents := make([]*ent.StakingEvent, len(allRewardsClaimedEvent.RewardedAmount))
		for i, rewardAmountItem := range allRewardsClaimedEvent.RewardedAmount {
			rewardAmount, _ := uint256.FromBig(rewardAmountItem.RewardedAmount)
			if rewardAmount == nil {
				return nil, errors.Join(
					ErrGetStakingEventsFromEvent,
					errors.New("allRewardsClaimedEvent: failed to convert reward amount to uint256"),
				)
			}
			stakingEvents[i] = &ent.StakingEvent{
				TransactionHash:            event.TransactionHash,
				TransactionIndex:           event.TransactionIndex,
				BlockNumber:                event.BlockNumber,
				LogIndex:                   event.LogIndex,
				EventType:                  stakingevent.EventTypeAllRewardsClaimed,
				AccountEvmAddress:          allRewardsClaimedEvent.Account.String(),
				NftClassAddress:            rewardAmountItem.BookNFT.String(),
				PendingRewardAmountRemoved: typeutil.Uint256(rewardAmount),
				ClaimedRewardAmountAdded:   typeutil.Uint256(rewardAmount),
				Datetime:                   event.Timestamp,
			}
		}
		return stakingEvents, nil
	}

	return nil, errors.Join(
		ErrGetStakingEventsFromEvent,
		errors.New("unknown event type"),
	)
}

func GetAccountKeysFromEvents(events []*ent.StakingEvent) []string {
	accountKeys := make(map[string]struct{})
	for _, event := range events {
		accountKeys[event.AccountEvmAddress] = struct{}{}
	}
	keys := make([]string, 0)
	for key := range accountKeys {
		keys = append(keys, key)
	}
	return keys
}

func GetBookNFTKeysFromEvent(event *ent.EVMEvent) []string {
	log := logConverter.ConvertEvmEventToLog(event)
	stakedEvent := new(like_collective.LikeCollectiveStaked)
	if err := logConverter.UnpackLog(log, stakedEvent); err == nil {
		return []string{stakedEvent.BookNFT.String()}
	}
	unstakedEvent := new(like_collective.LikeCollectiveUnstaked)
	if err := logConverter.UnpackLog(log, unstakedEvent); err == nil {
		return []string{unstakedEvent.BookNFT.String()}
	}
	rewardClaimedEvent := new(like_collective.LikeCollectiveRewardClaimed)
	if err := logConverter.UnpackLog(log, rewardClaimedEvent); err == nil {
		return []string{rewardClaimedEvent.BookNFT.String()}
	}
	rewardDepositedEvent := new(like_collective.LikeCollectiveRewardDeposited)
	if err := logConverter.UnpackLog(log, rewardDepositedEvent); err == nil {
		return []string{rewardDepositedEvent.BookNFT.String()}
	}
	allRewardsClaimedEvent := new(like_collective.LikeCollectiveAllRewardClaimed)
	if err := logConverter.UnpackLog(log, allRewardsClaimedEvent); err == nil {
		bookNFTEvmAddresses := make([]string, 0)
		for _, claimedAmount := range allRewardsClaimedEvent.RewardedAmount {
			bookNFTEvmAddresses = append(bookNFTEvmAddresses, claimedAmount.BookNFT.String())
		}
		return bookNFTEvmAddresses
	}
	return []string{}
}

func GetBookNFTKeysFromEvents(events []*ent.StakingEvent) []string {
	bookNFTKeys := make(map[string]struct{})
	for _, event := range events {
		bookNFTKeys[event.NftClassAddress] = struct{}{}
	}

	keys := make([]string, 0)
	for key := range bookNFTKeys {
		keys = append(keys, key)
	}
	return keys
}

func GetStakingKeysFromEvent(event *ent.EVMEvent) []database.StakingKey {
	log := logConverter.ConvertEvmEventToLog(event)
	stakedEvent := new(like_collective.LikeCollectiveStaked)
	if err := logConverter.UnpackLog(log, stakedEvent); err == nil {
		return []database.StakingKey{
			database.NewStakingKey(
				stakedEvent.Account.String(),
				stakedEvent.BookNFT.String(),
			),
		}
	}
	unstakedEvent := new(like_collective.LikeCollectiveUnstaked)
	if err := logConverter.UnpackLog(log, unstakedEvent); err == nil {
		return []database.StakingKey{
			database.NewStakingKey(
				unstakedEvent.Account.String(),
				unstakedEvent.BookNFT.String(),
			),
		}
	}
	rewardClaimedEvent := new(like_collective.LikeCollectiveRewardClaimed)
	if err := logConverter.UnpackLog(log, rewardClaimedEvent); err == nil {
		return []database.StakingKey{
			database.NewStakingKey(
				rewardClaimedEvent.Account.String(),
				rewardClaimedEvent.BookNFT.String(),
			),
		}
	}
	rewardDepositedEvent := new(like_collective.LikeCollectiveRewardDeposited)
	if err := logConverter.UnpackLog(log, rewardDepositedEvent); err == nil {
		return []database.StakingKey{
			database.NewStakingKey(
				rewardDepositedEvent.Account.String(),
				rewardDepositedEvent.BookNFT.String(),
			),
		}
	}
	allRewardsClaimedEvent := new(like_collective.LikeCollectiveAllRewardClaimed)
	if err := logConverter.UnpackLog(log, allRewardsClaimedEvent); err == nil {
		stakingKeys := make([]database.StakingKey, 0)
		for _, claimedAmount := range allRewardsClaimedEvent.RewardedAmount {
			stakingKeys = append(stakingKeys, database.NewStakingKey(
				allRewardsClaimedEvent.Account.String(),
				claimedAmount.BookNFT.String(),
			))
		}
		return stakingKeys
	}
	return []database.StakingKey{}
}

func GetStakingKeysFromEvents(events []*ent.StakingEvent) []database.StakingKey {
	stakingKeys := make(map[string]map[string]struct{})
	for _, event := range events {
		if _, ok := stakingKeys[event.AccountEvmAddress]; !ok {
			stakingKeys[event.AccountEvmAddress] = make(map[string]struct{})
		}
		stakingKeys[event.AccountEvmAddress][event.NftClassAddress] = struct{}{}
	}

	keys := make([]database.StakingKey, 0)
	for accountEVMAddress, bookNFTEVMAddresses := range stakingKeys {
		for bookNFTEVMAddress := range bookNFTEVMAddresses {
			keys = append(keys, database.NewStakingKey(
				accountEVMAddress,
				bookNFTEVMAddress,
			))
		}
	}
	return keys
}
