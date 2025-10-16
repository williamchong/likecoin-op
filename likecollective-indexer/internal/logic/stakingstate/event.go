package stakingstate

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/ent/stakingevent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/evm/like_collective"
	"likecollective-indexer/internal/evm/like_stake_position"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

var ErrGetStakingEventsFromEvent = errors.New("failed to get staking events from event")

func GetStakingEventsFromLikeCollectiveEvent(
	ctx context.Context,
	evmClient evm.EVMClient,
	event *ent.EVMEvent,
) ([]*ent.StakingEvent, error) {
	log := likeCollectiveLogConverter.ConvertEvmEventToLog(event)

	if event.Name == "Staked" {
		stakedEvent := new(like_collective.LikeCollectiveStaked)
		if err := likeCollectiveLogConverter.UnpackLog(log, stakedEvent); err != nil {
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
		if err := likeCollectiveLogConverter.UnpackLog(log, unstakedEvent); err != nil {
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
		if err := likeCollectiveLogConverter.UnpackLog(log, rewardClaimedEvent); err != nil {
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
		if err := likeCollectiveLogConverter.UnpackLog(log, rewardDepositedEvent); err != nil {
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
		if err := likeCollectiveLogConverter.UnpackLog(log, allRewardsClaimedEvent); err != nil {
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

func GetStakingEventsFromLikeStakePositionEvent(
	ctx context.Context,
	evmClient evm.EVMClient,
	event *ent.EVMEvent,
) ([]*ent.StakingEvent, error) {
	log := likeStakePositionLogConverter.ConvertEvmEventToLog(event)

	if event.Name == "Transfer" {
		transferEvent := new(like_stake_position.LikeStakePositionTransfer)
		if err := likeStakePositionLogConverter.UnpackLog(log, transferEvent); err != nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				fmt.Errorf("failed to unpack log: %w", err),
			)
		}

		if transferEvent.From == common.HexToAddress("0x0") {
			// The event is mint.
			// Given that the mint event can only be triggered by `newStakePosition`
			// and no manual minting is allowed,
			// so should be handled in `staked` event and
			// no need to spawn staking event for this case
			return []*ent.StakingEvent{}, nil
		}

		if transferEvent.To == common.HexToAddress("0x0") {
			// The event is burn.
			// Given that the burn event can only be triggered by `removeStakePosition`
			// and no manual burning is allowed,
			// so should be handled in `unstaked` event and
			// no need to spawn staking event for this case
			return []*ent.StakingEvent{}, nil
		}

		position, err := evmClient.GetStakePosition(
			ctx,
			big.NewInt(0).SetUint64(uint64(event.BlockNumber)),
			transferEvent.TokenId,
		)
		if err != nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				fmt.Errorf("failed to get stake position: %w", err),
			)
		}
		nftClassAddress := position.BookNFT.Hex()
		stakedAmount, overflow := uint256.FromBig(position.StakedAmount)
		if stakedAmount == nil || overflow {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("transferEvent: failed to convert staked amount to uint256"),
			)
		}

		pendingRewardAmountBigInt, err := evmClient.GetRewardsOfPosition(
			ctx,
			big.NewInt(0).SetUint64(uint64(event.BlockNumber)),
			transferEvent.TokenId,
		)
		if err != nil {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				fmt.Errorf("failed to get rewards of position: %w", err),
			)
		}
		pendingRewardAmount, overflow := uint256.FromBig(pendingRewardAmountBigInt)
		if pendingRewardAmount == nil || overflow {
			return nil, errors.Join(
				ErrGetStakingEventsFromEvent,
				errors.New("transferEvent: failed to convert pending reward amount to uint256"),
			)
		}

		return []*ent.StakingEvent{
			{
				TransactionHash:            event.TransactionHash,
				TransactionIndex:           event.TransactionIndex,
				BlockNumber:                event.BlockNumber,
				LogIndex:                   event.LogIndex,
				EventType:                  stakingevent.EventTypeStakePositionTransferred,
				AccountEvmAddress:          transferEvent.From.Hex(),
				NftClassAddress:            nftClassAddress,
				StakedAmountRemoved:        typeutil.Uint256(stakedAmount),
				PendingRewardAmountRemoved: typeutil.Uint256(pendingRewardAmount),
				Datetime:                   event.Timestamp,
			},
			{
				TransactionHash:          event.TransactionHash,
				TransactionIndex:         event.TransactionIndex,
				BlockNumber:              event.BlockNumber,
				LogIndex:                 event.LogIndex,
				EventType:                stakingevent.EventTypeStakePositionReceived,
				AccountEvmAddress:        transferEvent.To.Hex(),
				NftClassAddress:          nftClassAddress,
				StakedAmountAdded:        typeutil.Uint256(stakedAmount),
				PendingRewardAmountAdded: typeutil.Uint256(pendingRewardAmount),
				Datetime:                 event.Timestamp,
			},
		}, nil
	}

	return nil, errors.Join(
		ErrGetStakingEventsFromEvent,
		errors.New("unknown event type"),
	)
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

func MakeRewardDepositDistributedEvent(
	stakingEvent *ent.StakingEvent,
	accountEvmAddress string,
	rewardedAmount *uint256.Int,
) *ent.StakingEvent {
	return &ent.StakingEvent{
		TransactionHash:            stakingEvent.TransactionHash,
		TransactionIndex:           stakingEvent.TransactionIndex,
		BlockNumber:                stakingEvent.BlockNumber,
		LogIndex:                   stakingEvent.LogIndex,
		EventType:                  stakingevent.EventTypeRewardDepositDistributed,
		NftClassAddress:            stakingEvent.NftClassAddress,
		AccountEvmAddress:          accountEvmAddress,
		StakedAmountAdded:          typeutil.Uint256(uint256.NewInt(0)),
		StakedAmountRemoved:        typeutil.Uint256(uint256.NewInt(0)),
		PendingRewardAmountAdded:   typeutil.Uint256(rewardedAmount),
		PendingRewardAmountRemoved: typeutil.Uint256(uint256.NewInt(0)),
		ClaimedRewardAmountAdded:   typeutil.Uint256(uint256.NewInt(0)),
		ClaimedRewardAmountRemoved: typeutil.Uint256(uint256.NewInt(0)),
	}
}
