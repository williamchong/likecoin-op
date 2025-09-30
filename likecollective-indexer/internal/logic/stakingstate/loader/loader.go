package loader

import (
	"context"
	"errors"
	"fmt"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/stakingevent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate/model"

	"github.com/ethereum/go-ethereum/common"
)

type addressList map[common.Address]struct{}

func MakeAddressList(addresses ...common.Address) addressList {
	addressList := make(map[common.Address]struct{})
	for _, address := range addresses {
		addressList[address] = struct{}{}
	}
	return addressList
}

func (l addressList) Combine(other addressList) addressList {
	for address := range other {
		l[address] = struct{}{}
	}
	return l
}

type stakingKeyList map[database.StakingKey]struct{}

func MakeStakingKeyList(stakingKeys ...database.StakingKey) stakingKeyList {
	stakingKeyList := make(map[database.StakingKey]struct{})
	for _, stakingKey := range stakingKeys {
		stakingKeyList[stakingKey] = struct{}{}
	}
	return stakingKeyList
}

func (l stakingKeyList) Combine(other stakingKeyList) stakingKeyList {
	for stakingKey := range other {
		l[stakingKey] = struct{}{}
	}
	return l
}

type LoadState struct {
	AccountByAddress           addressList
	AccountByBookNFT           addressList
	BookNFTByAddress           addressList
	StakingByAccountAndBookNFT stakingKeyList
	StakingByBookNFT           addressList
}

type LoadStateFactory interface {
	MakeLoadState() *LoadState
}

func CombineLoadStates(
	loadStates []*LoadState,
) *LoadState {
	accountByAddress := MakeAddressList()
	accountByBookNFT := MakeAddressList()
	bookNFTByAddress := MakeAddressList()
	stakingByAccountAndBookNFT := MakeStakingKeyList()
	stakingByBookNFT := MakeAddressList()
	for _, loadState := range loadStates {
		accountByAddress = accountByAddress.Combine(loadState.AccountByAddress)
		accountByBookNFT = accountByBookNFT.Combine(loadState.AccountByBookNFT)
		bookNFTByAddress = bookNFTByAddress.Combine(loadState.BookNFTByAddress)
		stakingByAccountAndBookNFT = stakingByAccountAndBookNFT.Combine(loadState.StakingByAccountAndBookNFT)
		stakingByBookNFT = stakingByBookNFT.Combine(loadState.StakingByBookNFT)
	}
	return &LoadState{
		AccountByAddress:           accountByAddress,
		AccountByBookNFT:           accountByBookNFT,
		BookNFTByAddress:           bookNFTByAddress,
		StakingByAccountAndBookNFT: stakingByAccountAndBookNFT,
		StakingByBookNFT:           stakingByBookNFT,
	}
}

type StakingStateLoader interface {
	Load(ctx context.Context, loadState *LoadState) ([]*model.Account, []*model.NFTClass, []*model.Staking, error)
}

type stakingStateLoader struct {
	accountRepository  database.AccountRepository
	nftClassRepository database.NFTClassRepository
	stakingRepository  database.StakingRepository
}

func MakeStakingStateLoader(
	accountRepository database.AccountRepository,
	nftClassRepository database.NFTClassRepository,
	stakingRepository database.StakingRepository,
) StakingStateLoader {
	return &stakingStateLoader{
		accountRepository,
		nftClassRepository,
		stakingRepository,
	}
}

func (l *stakingStateLoader) Load(ctx context.Context, loadState *LoadState) ([]*model.Account, []*model.NFTClass, []*model.Staking, error) {
	accounts := make([]*model.Account, 0)
	accountAddedMap := make(map[string]struct{})

	if len(loadState.AccountByAddress) > 0 {
		accountEvmAddresses := make([]string, 0)
		for address := range loadState.AccountByAddress {
			accountEvmAddresses = append(accountEvmAddresses, address.String())
		}
		dbAccounts, err := l.accountRepository.QueryAccountsByEvmAddresses(ctx, accountEvmAddresses)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to query accounts by evm addresses: %w", err)
		}
		for _, dbAccount := range dbAccounts {
			if _, ok := accountAddedMap[dbAccount.EvmAddress]; !ok {
				accountAddedMap[dbAccount.EvmAddress] = struct{}{}
				accounts = append(accounts, model.NewAccountFromEnt(dbAccount))
			}
		}
	}

	if len(loadState.AccountByBookNFT) > 0 {
		addressesOfAccountsByBookNFTAddress := make([]string, 0)
		for address := range loadState.AccountByBookNFT {
			addressesOfAccountsByBookNFTAddress = append(
				addressesOfAccountsByBookNFTAddress,
				address.String(),
			)
		}
		dbAccountsByBookNFTAddress, err := l.accountRepository.QueryAccountsByNFTClassAddresses(
			ctx,
			addressesOfAccountsByBookNFTAddress,
		)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to query accounts by nft class addresses: %w", err)
		}
		for _, dbAccount := range dbAccountsByBookNFTAddress {
			if _, ok := accountAddedMap[dbAccount.EvmAddress]; !ok {
				accountAddedMap[dbAccount.EvmAddress] = struct{}{}
				accounts = append(accounts, model.NewAccountFromEnt(dbAccount))
			}
		}
	}

	nftClasses := make([]*model.NFTClass, 0)
	nftClassAddedMap := make(map[string]struct{})

	if len(loadState.BookNFTByAddress) > 0 {
		nftClassAddresses := make([]string, 0)
		for address := range loadState.BookNFTByAddress {
			nftClassAddresses = append(
				nftClassAddresses,
				address.String(),
			)
		}
		dbNftClasses, err := l.nftClassRepository.QueryNFTClassesByAddresses(ctx, nftClassAddresses)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to query stakings by keys: %w", err)
		}
		for _, dbNftClass := range dbNftClasses {
			if _, ok := nftClassAddedMap[dbNftClass.Address]; !ok {
				nftClassAddedMap[dbNftClass.Address] = struct{}{}
				nftClasses = append(nftClasses, model.NewNFTClassFromEnt(dbNftClass))
			}
		}
	}

	stakings := make([]*model.Staking, 0)
	stakingAddedMap := make(map[database.StakingKey]struct{})

	if len(loadState.StakingByAccountAndBookNFT) > 0 {
		stakingKeys := make([]database.StakingKey, 0)
		for stakingKey := range loadState.StakingByAccountAndBookNFT {
			stakingKeys = append(stakingKeys, stakingKey)
		}
		dbStakings, err := l.stakingRepository.QueryStakingsByKeys(ctx, stakingKeys)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to query stakings by keys: %w", err)
		}
		for _, dbStaking := range dbStakings {
			stakingKey := database.NewStakingKey(
				dbStaking.Edges.Account.EvmAddress,
				dbStaking.Edges.NftClass.Address,
			)
			if _, ok := stakingAddedMap[stakingKey]; !ok {
				stakingAddedMap[stakingKey] = struct{}{}
				stakings = append(stakings, model.NewStakingFromEnt(dbStaking))
			}
		}
	}

	if len(loadState.StakingByBookNFT) > 0 {
		addressesOfStakingsByBookNFTAddress := make([]string, 0)
		for address := range loadState.StakingByBookNFT {
			addressesOfStakingsByBookNFTAddress = append(
				addressesOfStakingsByBookNFTAddress,
				address.String(),
			)
		}
		dbStakingsByBookNFTAddress, err := l.stakingRepository.QueryStakingsByNFTClassAddresses(
			ctx,
			addressesOfStakingsByBookNFTAddress,
		)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to query stakings by nft class addresses: %w", err)
		}
		for _, dbStaking := range dbStakingsByBookNFTAddress {
			stakingKey := database.NewStakingKey(
				dbStaking.Edges.Account.EvmAddress,
				dbStaking.Edges.NftClass.Address,
			)
			if _, ok := stakingAddedMap[stakingKey]; !ok {
				stakingAddedMap[stakingKey] = struct{}{}
				stakings = append(stakings, model.NewStakingFromEnt(dbStaking))
			}
		}
	}

	return accounts, nftClasses, stakings, nil
}

func MakeLoadStateFactory(stakingEvent *ent.StakingEvent) (LoadStateFactory, error) {
	switch stakingEvent.EventType {
	case stakingevent.EventTypeStaked:
		return MakeLoadStakedStateFactory(stakingEvent), nil
	case stakingevent.EventTypeUnstaked:
		return MakeLoadUnstakedStateFactory(stakingEvent), nil
	case stakingevent.EventTypeRewardClaimed:
		return MakeLoadRewardClaimedStateFactory(stakingEvent), nil
	case stakingevent.EventTypeRewardDeposited:
		return MakeLoadRewardDepositedStateFactory(stakingEvent), nil
	case stakingevent.EventTypeRewardDepositDistributed:
		return MakeLoadRewardDepositDistributedStateFactory(stakingEvent), nil
	case stakingevent.EventTypeAllRewardsClaimed:
		return MakeLoadAllRewardsClaimedStateFactory(stakingEvent), nil
	default:
		return nil, errors.New("invalid staking event type")
	}
}
