package model

import "time"

type NFTSigningMessage struct {
	Id            uint64
	CreatedAt     time.Time
	CosmosAddress string
	LikerID       string
	EthAddress    string
	Nonce         string
	Message       string
	IssueTime     time.Time
}
