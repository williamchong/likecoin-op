package model

type UserProfile struct {
	CosmosWalletAddress string  `json:"cosmos_wallet_address"`
	LikerID             *string `json:"liker_id"`
	EthWalletAddress    *string `json:"eth_wallet_address"`
}
