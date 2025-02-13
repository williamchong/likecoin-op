package model

type UserProfile struct {
	User         string `json:"user"`
	DisplayName  string `json:"displayName"`
	Avatar       string `json:"avatar"`
	CosmosWallet string `json:"cosmosWallet"`
	LikeWallet   string `json:"likeWallet"`
}
