package cosmos

type TxBody struct {
	Memo string `json:"memo"`
}

type Tx struct {
	Body TxBody `json:"body"`
}

type TxResponse struct {
	Code   int    `json:"code"`
	RawLog string `json:"raw_log"`

	Tx Tx `json:"tx"`
}
