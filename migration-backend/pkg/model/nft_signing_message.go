package model

import "time"

type NFTSigningMessage struct {
	CosmosAddress string
	LikerID       string
	EthAddress    string
	Nonce         string
	Message       string
	IssueTime     time.Time
}
