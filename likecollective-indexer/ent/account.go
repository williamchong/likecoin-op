package ent

import (
	"errors"
	"io"
	"os"

	goyaml "gopkg.in/yaml.v2"
)

type Account struct {
	EvmAddress          string `json:"evmaddress"`
	StakedAmount        string `json:"stakedamount"`
	PendingRewardAmount string `json:"pendingrewardamount"`
	ClaimedRewardAmount string `json:"claimedrewardamount"`
}

type AccountClient struct {
}

func (c *AccountClient) Query() ([]*Account, error) {
	f, err := os.Open("dbmockdata/account.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := goyaml.NewDecoder(f)

	accounts := make([]*Account, 0)

	for {
		var account Account
		if err := decoder.Decode(&account); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	return accounts, nil
}
