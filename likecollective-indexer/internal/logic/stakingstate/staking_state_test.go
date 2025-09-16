package stakingstate_test

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/ent/stakingevent"
	"likecollective-indexer/internal/logic/stakingstate"

	"github.com/holiman/uint256"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	goyaml "gopkg.in/yaml.v2"
)

type StakingStateTestCaseEventType string

const (
	StakingStateTestCaseEventTypeStaked            StakingStateTestCaseEventType = "staked"
	StakingStateTestCaseEventTypeUnstaked          StakingStateTestCaseEventType = "unstaked"
	StakingStateTestCaseEventTypeRewardDeposited   StakingStateTestCaseEventType = "rewardDeposited"
	StakingStateTestCaseEventTypeRewardClaimed     StakingStateTestCaseEventType = "rewardClaimed"
	StakingStateTestCaseEventTypeAllRewardsClaimed StakingStateTestCaseEventType = "allRewardsClaimed"
)

type StakingStateTestCaseEventStaked struct {
	BookNFT      string `yaml:"bookNFT"`
	Account      string `yaml:"account"`
	StakedAmount string `yaml:"stakedAmount"`
}

func (e *StakingStateTestCaseEventStaked) ToEnt() *ent.StakingEvent {
	stakedAmount, err := uint256.FromDecimal(e.StakedAmount)
	if err != nil {
		panic("failed to convert staked amount to big.Int")
	}
	return &ent.StakingEvent{
		EventType:         stakingevent.EventTypeStaked,
		AccountEvmAddress: e.Account,
		NftClassAddress:   e.BookNFT,
		StakedAmountAdded: typeutil.Uint256(stakedAmount),
	}
}

type StakingStateTestCaseEventUnstaked struct {
	BookNFT      string `yaml:"bookNFT"`
	Account      string `yaml:"account"`
	StakedAmount string `yaml:"stakedAmount"`
}

func (e *StakingStateTestCaseEventUnstaked) ToEnt() *ent.StakingEvent {
	stakedAmount, err := uint256.FromDecimal(e.StakedAmount)
	if err != nil {
		panic("failed to convert staked amount to big.Int")
	}
	return &ent.StakingEvent{
		EventType:           stakingevent.EventTypeUnstaked,
		AccountEvmAddress:   e.Account,
		NftClassAddress:     e.BookNFT,
		StakedAmountRemoved: typeutil.Uint256(stakedAmount),
	}
}

type StakingStateTestCaseEventRewardDeposited struct {
	BookNFT      string `yaml:"bookNFT"`
	Account      string `yaml:"account"`
	RewardAmount string `yaml:"rewardAmount"`
}

func (e *StakingStateTestCaseEventRewardDeposited) ToEnt() *ent.StakingEvent {
	rewardAmount, err := uint256.FromDecimal(e.RewardAmount)
	if err != nil {
		panic("failed to convert reward amount to big.Int")
	}
	return &ent.StakingEvent{
		EventType:                stakingevent.EventTypeRewardDeposited,
		AccountEvmAddress:        e.Account,
		NftClassAddress:          e.BookNFT,
		PendingRewardAmountAdded: typeutil.Uint256(rewardAmount),
	}
}

type StakingStateTestCaseEventRewardClaimed struct {
	BookNFT      string `yaml:"bookNFT"`
	Account      string `yaml:"account"`
	RewardAmount string `yaml:"rewardAmount"`
}

func (e *StakingStateTestCaseEventRewardClaimed) ToEnt() *ent.StakingEvent {
	rewardAmount, err := uint256.FromDecimal(e.RewardAmount)
	if err != nil {
		panic("failed to convert reward amount to big.Int")
	}
	return &ent.StakingEvent{
		EventType:                  stakingevent.EventTypeRewardClaimed,
		AccountEvmAddress:          e.Account,
		NftClassAddress:            e.BookNFT,
		PendingRewardAmountRemoved: typeutil.Uint256(rewardAmount),
	}
}

type StakingStateTestCaseEventAllRewardsClaimed struct {
	BookNFT      string `yaml:"bookNFT"`
	Account      string `yaml:"account"`
	RewardAmount string `yaml:"rewardAmount"`
}

func (e *StakingStateTestCaseEventAllRewardsClaimed) ToEnt() *ent.StakingEvent {
	rewardAmount, err := uint256.FromDecimal(e.RewardAmount)
	if err != nil {
		panic("failed to convert reward amount to big.Int")
	}
	return &ent.StakingEvent{
		EventType:                  stakingevent.EventTypeAllRewardsClaimed,
		AccountEvmAddress:          e.Account,
		NftClassAddress:            e.BookNFT,
		PendingRewardAmountRemoved: typeutil.Uint256(rewardAmount),
		ClaimedRewardAmountAdded:   typeutil.Uint256(rewardAmount),
	}
}

type StakingStateTestCaseEvent struct {
	Type                       StakingStateTestCaseEventType               `yaml:"type"`
	StakedEventData            *StakingStateTestCaseEventStaked            `yaml:"stakedEventData"`
	UnstakedEventData          *StakingStateTestCaseEventUnstaked          `yaml:"unstakedEventData"`
	RewardDepositedEventData   *StakingStateTestCaseEventRewardDeposited   `yaml:"rewardDepositedEventData"`
	RewardClaimedEventData     *StakingStateTestCaseEventRewardClaimed     `yaml:"rewardClaimedEventData"`
	AllRewardsClaimedEventData *StakingStateTestCaseEventAllRewardsClaimed `yaml:"allRewardsClaimedEventData"`
}

func (e *StakingStateTestCaseEvent) ToStakingEvent() *ent.StakingEvent {
	switch e.Type {
	case StakingStateTestCaseEventTypeStaked:
		return e.StakedEventData.ToEnt()
	case StakingStateTestCaseEventTypeUnstaked:
		return e.UnstakedEventData.ToEnt()
	case StakingStateTestCaseEventTypeRewardDeposited:
		return e.RewardDepositedEventData.ToEnt()
	case StakingStateTestCaseEventTypeRewardClaimed:
		return e.RewardClaimedEventData.ToEnt()
	case StakingStateTestCaseEventTypeAllRewardsClaimed:
		return e.AllRewardsClaimedEventData.ToEnt()
	}
	panic("unknown event type")
}

type StakingStateAccount struct {
	EVMAddress          string `yaml:"address"`
	StakedAmount        string `yaml:"stakedAmount"`
	PendingRewardAmount string `yaml:"pendingRewardAmount"`
	ClaimedRewardAmount string `yaml:"claimedRewardAmount"`
}

type StakingStateBookNFT struct {
	EVMAddress   string `yaml:"address"`
	StakedAmount string `yaml:"stakedAmount"`
}

type StakingStateStaking struct {
	AccountEVMAddress   string `yaml:"accountAddress"`
	BookNFTEvmAddress   string `yaml:"bookNFTAddress"`
	StakedAmount        string `yaml:"stakedAmount"`
	PendingRewardAmount string `yaml:"pendingRewardAmount"`
	ClaimedRewardAmount string `yaml:"claimedRewardAmount"`
}

type StakingState struct {
	Accounts   map[string]StakingStateAccount            `yaml:"accounts"`
	NftClasses map[string]StakingStateBookNFT            `yaml:"nftClasses"`
	Stakings   map[string]map[string]StakingStateStaking `yaml:"stakings"`
}

type StakingStateTestCaseStep struct {
	Events       []StakingStateTestCaseEvent `yaml:"events"`
	StakingState *StakingState               `yaml:"stakingstate"`
	Error        string                      `yaml:"error"`
}

type StakingStateTestCaseSnapshottedAccounts struct {
	Address             string `yaml:"address"`
	StakedAmount        string `yaml:"stakedAmount"`
	PendingRewardAmount string `yaml:"pendingRewardAmount"`
	ClaimedRewardAmount string `yaml:"claimedRewardAmount"`
}

func (s *StakingStateTestCaseSnapshottedAccounts) ToEnt() *ent.Account {
	stakedAmount, err := uint256.FromDecimal(s.StakedAmount)
	if err != nil {
		panic("failed to convert staked amount to big.Int")
	}
	pendingRewardAmount, err := uint256.FromDecimal(s.PendingRewardAmount)
	if err != nil {
		panic("failed to convert pending reward amount to big.Int")
	}
	claimedRewardAmount, err := uint256.FromDecimal(s.ClaimedRewardAmount)
	if err != nil {
		panic("failed to convert claimed reward amount to big.Int")
	}
	return &ent.Account{
		EvmAddress:          s.Address,
		StakedAmount:        typeutil.Uint256(stakedAmount),
		PendingRewardAmount: typeutil.Uint256(pendingRewardAmount),
		ClaimedRewardAmount: typeutil.Uint256(claimedRewardAmount),
	}
}

type StakingStateTestCaseSnapshottedNftClasses struct {
	Address      string `yaml:"address"`
	StakedAmount string `yaml:"stakedAmount"`
}

func (s *StakingStateTestCaseSnapshottedNftClasses) ToEnt() *ent.NFTClass {
	stakedAmount, err := uint256.FromDecimal(s.StakedAmount)
	if err != nil {
		panic("failed to convert staked amount to big.Int")
	}
	return &ent.NFTClass{
		Address:      s.Address,
		StakedAmount: typeutil.Uint256(stakedAmount),
	}
}

type StakingStateTestCaseSnapshottedStakings struct {
	AccountAddress      string `yaml:"accountAddress"`
	BookNFTAddress      string `yaml:"bookNFTAddress"`
	StakedAmount        string `yaml:"stakedAmount"`
	PendingRewardAmount string `yaml:"pendingRewardAmount"`
	ClaimedRewardAmount string `yaml:"claimedRewardAmount"`
}

func (s *StakingStateTestCaseSnapshottedStakings) ToEnt() *ent.Staking {
	stakedAmount, err := uint256.FromDecimal(s.StakedAmount)
	if err != nil {
		panic("failed to convert staked amount to big.Int")
	}
	pendingRewardAmount, err := uint256.FromDecimal(s.PendingRewardAmount)
	if err != nil {
		panic("failed to convert pending reward amount to big.Int")
	}
	claimedRewardAmount, err := uint256.FromDecimal(s.ClaimedRewardAmount)
	if err != nil {
		panic("failed to convert claimed reward amount to big.Int")
	}
	return &ent.Staking{
		Edges: ent.StakingEdges{
			Account: &ent.Account{
				EvmAddress: s.AccountAddress,
			},
			NftClass: &ent.NFTClass{
				Address: s.BookNFTAddress,
			},
		},
		StakedAmount:        typeutil.Uint256(stakedAmount),
		PendingRewardAmount: typeutil.Uint256(pendingRewardAmount),
		ClaimedRewardAmount: typeutil.Uint256(claimedRewardAmount),
	}
}

type StakingStateTestCaseParameters struct {
	SnapshottedAccounts   []StakingStateTestCaseSnapshottedAccounts   `yaml:"snapshottedAccounts"`
	SnapshottedNftClasses []StakingStateTestCaseSnapshottedNftClasses `yaml:"snapshottedNftClasses"`
	SnapshottedStakings   []StakingStateTestCaseSnapshottedStakings   `yaml:"snapshottedStakings"`
}

type StakingStateTestCase struct {
	Name       string                          `yaml:"name"`
	Parameters *StakingStateTestCaseParameters `yaml:"parameters"`
	Steps      []StakingStateTestCaseStep      `yaml:"steps"`
}

func TestStakingStateFromTestData(t *testing.T) {
	Convey("Test Staking State From Test Data", t, func() {
		rootDir := "staking_state_testdata/"
		entries, err := os.ReadDir(rootDir)
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			f, err := os.Open(rootDir + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			decoder := goyaml.NewDecoder(f)

			Convey(entry.Name(), func() {
				for {
					var testCase StakingStateTestCase
					err := decoder.Decode(&testCase)
					if errors.Is(err, io.EOF) {
						break
					}
					if err != nil {
						t.Fatal(err)
					}

					Convey(testCase.Name, func() {
						snapshottedAccounts := make([]*ent.Account, 0)
						for _, account := range testCase.Parameters.SnapshottedAccounts {
							snapshottedAccounts = append(snapshottedAccounts, account.ToEnt())
						}
						snapshottedNftClasses := make([]*ent.NFTClass, 0)
						for _, nftClass := range testCase.Parameters.SnapshottedNftClasses {
							snapshottedNftClasses = append(snapshottedNftClasses, nftClass.ToEnt())
						}
						snapshottedStakings := make([]*ent.Staking, 0)
						for _, staking := range testCase.Parameters.SnapshottedStakings {
							snapshottedStakings = append(snapshottedStakings, staking.ToEnt())
						}

						stakingState := stakingstate.NewStakingStateFromEnt(
							snapshottedAccounts,
							snapshottedNftClasses,
							snapshottedStakings,
						)

						for _, step := range testCase.Steps {
							stakingEvents := make([]*ent.StakingEvent, 0)
							for _, event := range step.Events {
								stakingEvents = append(stakingEvents, event.ToStakingEvent())
							}
							var err error
							stakingState, err = stakingState.HandleStakingEvents(stakingEvents)
							if step.Error != "" {
								So(err, ShouldNotBeNil)
								So(err.Error(), ShouldContainSubstring, step.Error)
								break
							}

							actualStakingStateBytes, err := json.Marshal(stakingState)
							if err != nil {
								t.Fatal(err)
							}

							expectedStakingStateBytes, err := json.Marshal(step.StakingState)
							if err != nil {
								t.Fatal(err)
							}

							require.JSONEq(t, string(expectedStakingStateBytes), string(actualStakingStateBytes))
						}
					})
				}
			})
		}

	})
}
