package cosmos

import "encoding/json"

type TxBody struct {
	Messages []json.RawMessage `json:"messages"`
	Memo     string            `json:"memo"`
}

type Tx struct {
	Body TxBody `json:"body"`
}

type TxResponse struct {
	Code   int    `json:"code"`
	RawLog string `json:"raw_log"`

	Tx Tx `json:"tx"`
}

// MsgSend represents the cosmos.bank.v1beta1.MsgSend message
type MsgSend struct {
	TypeURL string `json:"@type"`
	From    string `json:"from_address"`
	To      string `json:"to_address"`
	Amount  []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
}
