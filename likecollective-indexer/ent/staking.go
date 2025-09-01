package ent

import (
	"errors"
	"fmt"
	"io"
	"os"

	goyaml "gopkg.in/yaml.v2"
)

type Staking struct {
	ID                  string `json:"id" `
	BookNFT             string `json:"book_nft"`
	Account             string `json:"account"`
	PoolShare           string `json:"poolshare"`
	StakedAmount        string `json:"stakedamount"`
	PendingRewardAmount string `json:"pendingrewardamount"`
	ClaimedRewardAmount string `json:"claimedrewardamount"`
}

type StakingClient struct {
}

func (c *StakingClient) Query() ([]*Staking, error) {
	f, err := os.Open("dbmockdata/staking.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := goyaml.NewDecoder(f)

	stakings := make([]*Staking, 0)

	for {
		var staking Staking
		if err := decoder.Decode(&staking); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to decode staking: %w", err)
		}

		stakings = append(stakings, &staking)
	}

	return stakings, nil
}
