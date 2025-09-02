package ent

type Client struct {
	Account      *AccountClient
	Staking      *StakingClient
	BookNFT      *BookNFTClient
	StakingEvent *StakingEventClient
}

func NewClient() *Client {
	return &Client{
		Account:      &AccountClient{},
		Staking:      &StakingClient{},
		BookNFT:      &BookNFTClient{},
		StakingEvent: &StakingEventClient{},
	}
}
