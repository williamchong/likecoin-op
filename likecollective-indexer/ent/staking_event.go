package ent

import (
	"errors"
	"io"
	"os"
	"time"

	goyaml "gopkg.in/yaml.v2"
)

type StakingEventType string

const (
	StakingEventTypeStaked            StakingEventType = "staked"
	StakingEventTypeUnstaked          StakingEventType = "unstaked"
	StakingEventTypeRewardAdded       StakingEventType = "reward-added"
	StakingEventTypeRewardClaimed     StakingEventType = "reward-claimed"
	StakingEventTypeRewardDeposited   StakingEventType = "reward-deposited"
	StakingEventTypeAllRewardsClaimed StakingEventType = "all-rewards-claimed"
)

type StakingEvent struct {
	EventType           StakingEventType `json:"event_type"`
	BookNFT             string           `json:"book_nft"`
	Account             string           `json:"account"`
	StakedAmountAdded   string           `json:"stakedamountadded"`
	StakedAmountRemoved string           `json:"stakedamountremoved"`
	RewardAmountAdded   string           `json:"rewardamountadded"`
	RewardAmountRemoved string           `json:"rewardamountremoved"`
	DateTime            time.Time        `json:"datetime"`
}

type StakingEventClient struct {
}

func (c *StakingEventClient) Query() ([]*StakingEvent, error) {
	f, err := os.Open("dbmockdata/staking_event.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := goyaml.NewDecoder(f)

	stakingEvents := make([]*StakingEvent, 0)

	for {
		var stakingEvent StakingEvent
		if err := decoder.Decode(&stakingEvent); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		stakingEvents = append(stakingEvents, &stakingEvent)
	}

	return stakingEvents, nil
}
