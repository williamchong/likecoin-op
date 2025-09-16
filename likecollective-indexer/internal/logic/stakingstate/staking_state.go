package stakingstate

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
)

type StakingState struct {
	Accounts   map[string]Account
	NftClasses map[string]NFTClass
	Stakings   map[string]map[string]Staking
}

func NewStakingStateFromEnt(
	snapshottedAccounts []*ent.Account,
	snapshottedNftClasses []*ent.NFTClass,
	snapshottedStakings []*ent.Staking,
) *StakingState {
	entAccountMap := make(map[string]*ent.Account)
	for _, account := range snapshottedAccounts {
		entAccountMap[account.EvmAddress] = account
	}
	entNftClassMap := make(map[string]*ent.NFTClass)
	for _, nftClass := range snapshottedNftClasses {
		entNftClassMap[nftClass.Address] = nftClass
	}
	entStakingMap := make(map[database.StakingKey]*ent.Staking)
	for _, staking := range snapshottedStakings {
		entStakingMap[database.NewStakingKey(staking.Edges.Account.EvmAddress, staking.Edges.NftClass.Address)] = staking
	}

	accounts := make(map[string]Account)
	for _, account := range entAccountMap {
		accounts[account.EvmAddress] = NewAccountFromEnt(account)
	}

	nftClasses := make(map[string]NFTClass)
	for _, nftClass := range entNftClassMap {
		nftClasses[nftClass.Address] = NewNFTClassFromEnt(nftClass)
	}

	stakings := make(map[string]map[string]Staking)
	for _, staking := range entStakingMap {
		if _, ok := stakings[staking.Edges.Account.EvmAddress]; !ok {
			stakings[staking.Edges.Account.EvmAddress] = make(map[string]Staking)
		}
		stakings[staking.Edges.Account.EvmAddress][staking.Edges.NftClass.Address] = NewStakingFromEnt(staking)
	}

	return &StakingState{
		Accounts:   accounts,
		NftClasses: nftClasses,
		Stakings:   stakings,
	}
}

func (s *StakingState) HandleStakingEvents(stakingEvents []*ent.StakingEvent) (*StakingState, error) {
	accountKeys := GetAccountKeysFromEvents(stakingEvents)
	for _, accountKey := range accountKeys {
		if _, ok := s.Accounts[accountKey]; !ok {
			s.Accounts[accountKey] = NewAccount(accountKey)
		}
	}

	bookNFTKeys := GetBookNFTKeysFromEvents(stakingEvents)
	for _, bookNFTKey := range bookNFTKeys {
		if _, ok := s.NftClasses[bookNFTKey]; !ok {
			s.NftClasses[bookNFTKey] = NewNFTClass(bookNFTKey)
		}
	}

	stakingKeys := GetStakingKeysFromEvents(stakingEvents)
	for _, stakingKey := range stakingKeys {
		if _, ok := s.Stakings[stakingKey.AccountEVMAddress]; !ok {
			s.Stakings[stakingKey.AccountEVMAddress] = make(map[string]Staking)
		}
		if _, ok := s.Stakings[stakingKey.AccountEVMAddress][stakingKey.BookNFTEVMAddress]; !ok {
			s.Stakings[stakingKey.AccountEVMAddress][stakingKey.BookNFTEVMAddress] = NewStakingFromStakingKey(stakingKey)
		}
	}

	visitors := make(StakingEventVisitors, 0)
	for _, stakingEvent := range stakingEvents {
		visitor, err := MakeStakingEventVisitorFromStakingEvent(stakingEvent)
		if err != nil {
			return nil, err
		}
		visitors = append(visitors, visitor)
	}

	return s.Run(visitors)
}

func (s *StakingState) Run(visitors StakingEventVisitors) (*StakingState, error) {
	accounts := make(map[string]Account)
	var err error
	for key, account := range s.Accounts {
		accounts[key], err = visitors.VisitAccount(account)
		if err != nil {
			return nil, err
		}
	}

	nftClasses := make(map[string]NFTClass)
	for key, nftClass := range s.NftClasses {
		nftClasses[key], err = visitors.VisitNFTClass(nftClass)
		if err != nil {
			return nil, err
		}
	}

	stakings := make(map[string]map[string]Staking)
	for accountKey, staking := range s.Stakings {
		for bookNFTKey, staking := range staking {
			if _, ok := stakings[accountKey]; !ok {
				stakings[accountKey] = make(map[string]Staking)
			}
			stakings[accountKey][bookNFTKey], err = visitors.VisitStaking(staking)
			if err != nil {
				return nil, err
			}
		}
	}

	return &StakingState{
		Accounts:   accounts,
		NftClasses: nftClasses,
		Stakings:   stakings,
	}, nil
}

func (s *StakingState) Persist(
	ctx context.Context,
	stakingEvents []*ent.StakingEvent,
	persistor StakingStatePersistor,
) error {
	accounts := make([]Account, 0)
	for _, account := range s.Accounts {
		accounts = append(accounts, account)
	}
	nftClasses := make([]NFTClass, 0)
	for _, nftClass := range s.NftClasses {
		nftClasses = append(nftClasses, nftClass)
	}
	stakings := make([]Staking, 0)
	for _, staking := range s.Stakings {
		for _, staking := range staking {
			stakings = append(stakings, staking)
		}
	}
	return persistor.Persist(
		ctx,
		stakingEvents,
		accounts,
		nftClasses,
		stakings,
	)
}
