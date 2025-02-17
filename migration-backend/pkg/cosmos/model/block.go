package model

import "time"

type BlockHeader struct {
	Height string    `json:"height"`
	Time   time.Time `json:"time"`
}

type Block struct {
	Header BlockHeader `json:"header"`
}
