package stakingstate

import (
	"context"
	"fmt"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate/loader"
	"likecollective-indexer/internal/logic/stakingstate/model"
	"likecollective-indexer/internal/logic/stakingstate/persistor"

	"github.com/ethereum/go-ethereum/common"
)

type StakingState interface {
	Process(stakingEvents []*ent.StakingEvent) (*stakingState, []*ent.StakingEvent, error)
	Persist(ctx context.Context, stakingEvents []*ent.StakingEvent, persistor persistor.StakingStatePersistor) error
}

type stakingState struct {
	accounts   []*model.Account
	nftClasses []*model.NFTClass
	stakings   []*model.Staking
}

func LoadStakingState(
	ctx context.Context,
	stakingStateLoader loader.StakingStateLoader,
	stakingEvents []*ent.StakingEvent,
) (StakingState, error) {
	loadStates := make([]*loader.LoadState, 0)
	for _, stakingEvent := range stakingEvents {
		loadStateFactory, err := loader.MakeLoadStateFactory(stakingEvent)
		if err != nil {
			return nil, err
		}
		loadStates = append(loadStates, loadStateFactory.MakeLoadState())
	}

	accounts, nftClasses, stakings, err := stakingStateLoader.Load(ctx, loader.CombineLoadStates(loadStates))
	if err != nil {
		return nil, err
	}

	return &stakingState{
		accounts:   accounts,
		nftClasses: nftClasses,
		stakings:   stakings,
	}, nil
}

func (s *stakingState) Process(stakingEvents []*ent.StakingEvent) (*stakingState, []*ent.StakingEvent, error) {
	applications := make([]StakingEventApplication, 0)
	for _, stakingEvent := range stakingEvents {
		application, err := MakeStakingEventApplication(stakingEvent)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to make staking event application: %w", err)
		}
		applications = append(applications, application)
	}

	return s.run(applications)
}

func (s *stakingState) run(applications []StakingEventApplication) (*stakingState, []*ent.StakingEvent, error) {
	processedStakingEvents := make([]*ent.StakingEvent, 0)
	newState := s
	for _, application := range applications {
		var (
			furtherApplications []StakingEventApplication
			err                 error
		)
		newState, furtherApplications, err = application.Apply(newState)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to apply staking event application: %w", err)
		}
		processedStakingEvents = append(
			processedStakingEvents, application.GetStakingEvent(),
		)

		if len(furtherApplications) > 0 {
			var (
				furtherProcessedStakingEvents []*ent.StakingEvent
			)
			newState, furtherProcessedStakingEvents, err = s.run(furtherApplications)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to run further applications: %w", err)
			}
			processedStakingEvents = append(processedStakingEvents, furtherProcessedStakingEvents...)
		}
	}
	return newState, processedStakingEvents, nil
}

func (s *stakingState) Persist(
	ctx context.Context,
	stakingEvents []*ent.StakingEvent,
	persistor persistor.StakingStatePersistor,
) error {
	return persistor.Persist(
		ctx,
		stakingEvents,
		s.accounts,
		s.nftClasses,
		s.stakings,
	)
}

func (s *stakingState) GetAccountByAddress(evmAddress common.Address) (*model.Account, bool) {
	for _, account := range s.accounts {
		if account.EVMAddress.String() == evmAddress.String() {
			return account, true
		}
	}
	return nil, false
}

func (s *stakingState) GetNFTClassByAddress(evmAddress common.Address) (*model.NFTClass, bool) {
	for _, nftClass := range s.nftClasses {
		if nftClass.EVMAddress.String() == evmAddress.String() {
			return nftClass, true
		}
	}
	return nil, false
}

func (s *stakingState) GetStakingByAddress(accountEVMAddress common.Address, bookNFTEVMAddress common.Address) (*model.Staking, bool) {
	for _, staking := range s.stakings {
		if staking.AccountEVMAddress.String() == accountEVMAddress.String() && staking.BookNFTEvmAddress.String() == bookNFTEVMAddress.String() {
			return staking, true
		}
	}
	return nil, false
}

func (s *stakingState) GetOrCreateAccount(evmAddress common.Address) *model.Account {
	for _, account := range s.accounts {
		if account.EVMAddress.String() == evmAddress.String() {
			return account
		}
	}
	account := model.NewAccount(evmAddress.String())
	s.accounts = append(s.accounts, account)
	return account
}

func (s *stakingState) GetOrCreateNFTClass(evmAddress common.Address) *model.NFTClass {
	for _, nftClass := range s.nftClasses {
		if nftClass.EVMAddress.String() == evmAddress.String() {
			return nftClass
		}
	}
	nftClass := model.NewNFTClass(evmAddress.String())
	s.nftClasses = append(s.nftClasses, nftClass)
	return nftClass
}

func (s *stakingState) GetOrCreateStaking(
	accountEVMAddress common.Address,
	bookNFTEVMAddress common.Address,
) *model.Staking {
	for _, staking := range s.stakings {
		if staking.AccountEVMAddress.String() == accountEVMAddress.String() && staking.BookNFTEvmAddress.String() == bookNFTEVMAddress.String() {
			return staking
		}
	}
	staking := model.NewStakingFromStakingKey(database.NewStakingKey(accountEVMAddress.String(), bookNFTEVMAddress.String()))
	s.stakings = append(s.stakings, staking)
	return staking
}

func (s *stakingState) GetStakingsByNFTClassAddress(nftClassAddress common.Address) []*model.Staking {
	stakings := make([]*model.Staking, 0)
	for _, staking := range s.stakings {
		if staking.BookNFTEvmAddress.String() == nftClassAddress.String() {
			stakings = append(stakings, staking)
		}
	}
	return stakings
}
