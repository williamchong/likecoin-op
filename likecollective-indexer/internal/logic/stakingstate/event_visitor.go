package stakingstate

import (
	"errors"
	"strings"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/stakingevent"
	"likecollective-indexer/internal/evm/like_collective"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type StakingEventVisitor interface {
	VisitAccount(account Account) (Account, error)
	VisitNFTClass(nftClass NFTClass) (NFTClass, error)
	VisitStaking(staking Staking) (Staking, error)
}

func MakeStakingEventVisitorFromEvent(event *ent.EVMEvent) (StakingEventVisitor, error) {
	log := logConverter.ConvertEvmEventToLog(event)
	stakedEvent := new(like_collective.LikeCollectiveStaked)
	if err := logConverter.UnpackLog(log, stakedEvent); err == nil {
		return NewStakedEventVisitor(stakedEvent), nil
	}
	unstakedEvent := new(like_collective.LikeCollectiveUnstaked)
	if err := logConverter.UnpackLog(log, unstakedEvent); err == nil {
		return NewUnstakedEventVisitor(unstakedEvent), nil
	}
	rewardClaimedEvent := new(like_collective.LikeCollectiveRewardClaimed)
	if err := logConverter.UnpackLog(log, rewardClaimedEvent); err == nil {
		return NewRewardClaimedEventVisitor(rewardClaimedEvent), nil
	}
	rewardDepositedEvent := new(like_collective.LikeCollectiveRewardDeposited)
	if err := logConverter.UnpackLog(log, rewardDepositedEvent); err == nil {
		return NewRewardDepositedEventVisitor(rewardDepositedEvent), nil
	}
	allRewardsClaimedEvent := new(like_collective.LikeCollectiveAllRewardClaimed)
	if err := logConverter.UnpackLog(log, allRewardsClaimedEvent); err == nil {
		return NewAllRewardsClaimedEventVisitor(allRewardsClaimedEvent), nil
	}
	return nil, errors.New("unknown event type")
}

func MakeStakingEventVisitorFromStakingEvent(stakingEvent *ent.StakingEvent) (StakingEventVisitor, error) {
	switch stakingEvent.EventType {
	case stakingevent.EventTypeStaked:
		return NewStakedEventVisitor(&like_collective.LikeCollectiveStaked{
			Account:      common.HexToAddress(stakingEvent.AccountEvmAddress),
			BookNFT:      common.HexToAddress(stakingEvent.NftClassAddress),
			StakedAmount: (*uint256.Int)(stakingEvent.StakedAmountAdded).ToBig(),
		}), nil
	case stakingevent.EventTypeUnstaked:
		return NewUnstakedEventVisitor(&like_collective.LikeCollectiveUnstaked{
			Account:      common.HexToAddress(stakingEvent.AccountEvmAddress),
			BookNFT:      common.HexToAddress(stakingEvent.NftClassAddress),
			StakedAmount: (*uint256.Int)(stakingEvent.StakedAmountRemoved).ToBig(),
		}), nil
	case stakingevent.EventTypeRewardClaimed:
		return NewRewardClaimedEventVisitor(&like_collective.LikeCollectiveRewardClaimed{
			Account:        common.HexToAddress(stakingEvent.AccountEvmAddress),
			BookNFT:        common.HexToAddress(stakingEvent.NftClassAddress),
			RewardedAmount: (*uint256.Int)(stakingEvent.PendingRewardAmountRemoved).ToBig(),
		}), nil
	case stakingevent.EventTypeRewardDeposited:
		return NewRewardDepositedEventVisitor(&like_collective.LikeCollectiveRewardDeposited{
			Account:        common.HexToAddress(stakingEvent.AccountEvmAddress),
			BookNFT:        common.HexToAddress(stakingEvent.NftClassAddress),
			RewardedAmount: (*uint256.Int)(stakingEvent.PendingRewardAmountAdded).ToBig(),
		}), nil
	case stakingevent.EventTypeAllRewardsClaimed:
		return NewAllRewardsClaimedEventVisitor(&like_collective.LikeCollectiveAllRewardClaimed{
			Account: common.HexToAddress(stakingEvent.AccountEvmAddress),
			RewardedAmount: []like_collective.LikeCollectiveRewardData{
				{
					BookNFT:        common.HexToAddress(stakingEvent.NftClassAddress),
					RewardedAmount: (*uint256.Int)(stakingEvent.PendingRewardAmountRemoved).ToBig(),
				},
			},
		}), nil
	}
	return nil, errors.New("unknown event type")
}

type StakingEventVisitors []StakingEventVisitor

func MakeStakingEventVisitorsFromEvents(events ...*ent.EVMEvent) (StakingEventVisitors, error) {
	visitors := make(StakingEventVisitors, 0)
	for _, event := range events {
		visitor, err := MakeStakingEventVisitorFromEvent(event)
		if err != nil {
			return visitors, err
		}
		visitors = append(visitors, visitor)
	}
	return visitors, nil
}

func (s StakingEventVisitors) VisitAccount(account Account) (Account, error) {
	for _, visitor := range s {
		var err error
		account, err = visitor.VisitAccount(account)
		if err != nil {
			return account, err
		}
	}
	return account, nil
}

func (s StakingEventVisitors) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	for _, visitor := range s {
		var err error
		nftClass, err = visitor.VisitNFTClass(nftClass)
		if err != nil {
			return nftClass, err
		}
	}
	return nftClass, nil
}

func (s StakingEventVisitors) VisitStaking(staking Staking) (Staking, error) {
	for _, visitor := range s {
		var err error
		staking, err = visitor.VisitStaking(staking)
		if err != nil {
			return staking, err
		}
	}
	return staking, nil
}

var ErrStakedEventVisitor = errors.New("failed to visit staked event")

type stakedEventVisitor struct {
	*like_collective.LikeCollectiveStaked
}

func NewStakedEventVisitor(
	LikeCollectiveStaked *like_collective.LikeCollectiveStaked,
) StakingEventVisitor {
	return &stakedEventVisitor{
		LikeCollectiveStaked,
	}
}

func (s *stakedEventVisitor) VisitAccount(account Account) (Account, error) {
	if !strings.EqualFold(account.EVMAddress.String(), s.Account.String()) {
		return account, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return account, errors.Join(ErrStakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	account.StakedAmount = uint256.NewInt(0).Add(account.StakedAmount, stakedAmount)
	return account, nil
}

func (s *stakedEventVisitor) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	if !strings.EqualFold(nftClass.EVMAddress.String(), s.BookNFT.String()) {
		return nftClass, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return nftClass, errors.Join(ErrStakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	nftClass.StakedAmount = uint256.NewInt(0).Add(nftClass.StakedAmount, stakedAmount)
	return nftClass, nil
}

func (s *stakedEventVisitor) VisitStaking(staking Staking) (Staking, error) {
	if !strings.EqualFold(
		staking.AccountEVMAddress.String(), s.Account.String()) ||
		!strings.EqualFold(
			staking.BookNFTEvmAddress.String(),
			s.BookNFT.String(),
		) {
		return staking, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return staking, errors.Join(ErrStakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	staking.StakedAmount = uint256.NewInt(0).Add(staking.StakedAmount, stakedAmount)
	return staking, nil
}

var ErrUnstakedEventVisitor = errors.New("failed to visit unstaked event")

type unstakedEventVisitor struct {
	*like_collective.LikeCollectiveUnstaked
}

func NewUnstakedEventVisitor(
	LikeCollectiveUnstaked *like_collective.LikeCollectiveUnstaked,
) StakingEventVisitor {
	return &unstakedEventVisitor{
		LikeCollectiveUnstaked,
	}
}

func (s *unstakedEventVisitor) VisitAccount(account Account) (Account, error) {
	if !strings.EqualFold(account.EVMAddress.String(), s.Account.String()) {
		return account, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return account, errors.Join(ErrUnstakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	if account.StakedAmount.Cmp(stakedAmount) < 0 {
		return account, errors.Join(ErrUnstakedEventVisitor, errors.New("staked amount is less than staked amount"))
	}
	account.StakedAmount = uint256.NewInt(0).Sub(account.StakedAmount, stakedAmount)
	return account, nil
}

func (s *unstakedEventVisitor) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	if !strings.EqualFold(nftClass.EVMAddress.String(), s.BookNFT.String()) {
		return nftClass, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return nftClass, errors.Join(ErrUnstakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	if nftClass.StakedAmount.Cmp(stakedAmount) < 0 {
		return nftClass, errors.Join(ErrUnstakedEventVisitor, errors.New("staked amount is less than staked amount"))
	}
	nftClass.StakedAmount = uint256.NewInt(0).Sub(nftClass.StakedAmount, stakedAmount)
	return nftClass, nil
}

func (s *unstakedEventVisitor) VisitStaking(staking Staking) (Staking, error) {
	if !strings.EqualFold(
		staking.AccountEVMAddress.String(), s.Account.String()) ||
		!strings.EqualFold(
			staking.BookNFTEvmAddress.String(),
			s.BookNFT.String(),
		) {
		return staking, nil
	}
	stakedAmount, _ := uint256.FromBig(s.StakedAmount)
	if stakedAmount == nil {
		return staking, errors.Join(ErrUnstakedEventVisitor, errors.New("failed to convert staked amount to uint256"))
	}
	if staking.StakedAmount.Cmp(stakedAmount) < 0 {
		return staking, errors.Join(ErrUnstakedEventVisitor, errors.New("staked amount is less than staked amount"))
	}
	staking.StakedAmount = uint256.NewInt(0).Sub(staking.StakedAmount, stakedAmount)
	return staking, nil
}

var ErrRewardClaimedEventVisitor = errors.New("failed to visit reward claimed event")

type rewardClaimedEventVisitor struct {
	*like_collective.LikeCollectiveRewardClaimed
}

func NewRewardClaimedEventVisitor(
	LikeCollectiveRewardClaimed *like_collective.LikeCollectiveRewardClaimed,
) StakingEventVisitor {
	return &rewardClaimedEventVisitor{
		LikeCollectiveRewardClaimed,
	}
}

func (s *rewardClaimedEventVisitor) VisitAccount(account Account) (Account, error) {
	if !strings.EqualFold(account.EVMAddress.String(), s.Account.String()) {
		return account, nil
	}
	rewardAmount, _ := uint256.FromBig(s.RewardedAmount)
	if rewardAmount == nil {
		return account, errors.Join(ErrRewardClaimedEventVisitor, errors.New("failed to convert reward amount to uint256"))
	}
	if account.PendingRewardAmount.Cmp(rewardAmount) < 0 {
		return account, errors.Join(ErrRewardClaimedEventVisitor, errors.New("pending reward amount is less than reward amount"))
	}
	account.PendingRewardAmount = uint256.NewInt(0).Sub(account.PendingRewardAmount, rewardAmount)
	account.ClaimedRewardAmount = uint256.NewInt(0).Add(account.ClaimedRewardAmount, rewardAmount)
	return account, nil
}

func (s *rewardClaimedEventVisitor) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	return nftClass, nil
}

func (s *rewardClaimedEventVisitor) VisitStaking(staking Staking) (Staking, error) {
	if !strings.EqualFold(
		staking.AccountEVMAddress.String(), s.Account.String()) ||
		!strings.EqualFold(
			staking.BookNFTEvmAddress.String(),
			s.BookNFT.String(),
		) {
		return staking, nil
	}
	rewardAmount, _ := uint256.FromBig(s.RewardedAmount)
	if rewardAmount == nil {
		return staking, errors.Join(ErrRewardClaimedEventVisitor, errors.New("failed to convert reward amount to uint256"))
	}
	if staking.PendingRewardAmount.Cmp(rewardAmount) < 0 {
		return staking, errors.Join(ErrRewardClaimedEventVisitor, errors.New("pending reward amount is less than reward amount"))
	}
	staking.PendingRewardAmount = uint256.NewInt(0).Sub(staking.PendingRewardAmount, rewardAmount)
	staking.ClaimedRewardAmount = uint256.NewInt(0).Add(staking.ClaimedRewardAmount, rewardAmount)
	return staking, nil
}

var ErrRewardDepositedEventVisitor = errors.New("failed to visit reward deposited event")

type rewardDepositedEventVisitor struct {
	*like_collective.LikeCollectiveRewardDeposited
}

func NewRewardDepositedEventVisitor(
	LikeCollectiveRewardDeposited *like_collective.LikeCollectiveRewardDeposited,
) StakingEventVisitor {
	return &rewardDepositedEventVisitor{
		LikeCollectiveRewardDeposited,
	}
}

func (s *rewardDepositedEventVisitor) VisitAccount(account Account) (Account, error) {
	rewardAmount, _ := uint256.FromBig(s.RewardedAmount)
	if rewardAmount == nil {
		return account, errors.Join(ErrRewardDepositedEventVisitor, errors.New("failed to convert reward amount to uint256"))
	}
	account.PendingRewardAmount = uint256.NewInt(0).Add(account.PendingRewardAmount, rewardAmount)
	return account, nil
}

func (s *rewardDepositedEventVisitor) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	return nftClass, nil
}

func (s *rewardDepositedEventVisitor) VisitStaking(staking Staking) (Staking, error) {
	if !strings.EqualFold(
		staking.BookNFTEvmAddress.String(), s.BookNFT.String()) ||
		!strings.EqualFold(
			staking.AccountEVMAddress.String(), s.Account.String(),
		) {
		return staking, nil
	}
	rewardAmount, _ := uint256.FromBig(s.RewardedAmount)
	if rewardAmount == nil {
		return staking, errors.Join(ErrRewardDepositedEventVisitor, errors.New("failed to convert reward amount to uint256"))
	}
	staking.PendingRewardAmount = uint256.NewInt(0).Add(staking.PendingRewardAmount, rewardAmount)
	return staking, nil
}

var ErrAllRewardsClaimedEventVisitor = errors.New("failed to visit all rewards claimed event")

type allRewardsClaimedEventVisitor struct {
	*like_collective.LikeCollectiveAllRewardClaimed
}

func NewAllRewardsClaimedEventVisitor(
	LikeCollectiveAllRewardClaimed *like_collective.LikeCollectiveAllRewardClaimed,
) StakingEventVisitor {
	return &allRewardsClaimedEventVisitor{
		LikeCollectiveAllRewardClaimed,
	}
}

func (s *allRewardsClaimedEventVisitor) VisitAccount(account Account) (Account, error) {
	if !strings.EqualFold(account.EVMAddress.String(), s.Account.String()) {
		return account, nil
	}
	for _, rewardAmount := range s.RewardedAmount {
		rewardAmount, _ := uint256.FromBig(rewardAmount.RewardedAmount)
		if rewardAmount != nil {
			account.ClaimedRewardAmount = uint256.NewInt(0).Add(account.ClaimedRewardAmount, rewardAmount)
			if account.PendingRewardAmount.Cmp(rewardAmount) < 0 {
				return account, errors.Join(ErrAllRewardsClaimedEventVisitor, errors.New("pending reward amount is less than reward amount"))
			}
			account.PendingRewardAmount = uint256.NewInt(0).Sub(account.PendingRewardAmount, rewardAmount)
		}
	}
	return account, nil
}

func (s *allRewardsClaimedEventVisitor) VisitNFTClass(nftClass NFTClass) (NFTClass, error) {
	return nftClass, nil
}

func (s *allRewardsClaimedEventVisitor) VisitStaking(staking Staking) (Staking, error) {
	for _, rewardAmount := range s.RewardedAmount {
		if strings.EqualFold(
			staking.BookNFTEvmAddress.String(), rewardAmount.BookNFT.String()) &&
			strings.EqualFold(
				staking.AccountEVMAddress.String(), s.Account.String(),
			) {
			rewardAmount, _ := uint256.FromBig(rewardAmount.RewardedAmount)
			if rewardAmount == nil {
				return staking, errors.Join(ErrAllRewardsClaimedEventVisitor, errors.New("failed to convert reward amount to uint256"))
			}
			staking.ClaimedRewardAmount = uint256.NewInt(0).Add(staking.ClaimedRewardAmount, rewardAmount)
			if staking.PendingRewardAmount.Cmp(rewardAmount) < 0 {
				return staking, errors.Join(ErrAllRewardsClaimedEventVisitor, errors.New("pending reward amount is less than reward amount"))
			}
			staking.PendingRewardAmount = uint256.NewInt(0).Sub(staking.PendingRewardAmount, rewardAmount)
		}
	}
	return staking, nil
}
