package cosmos

type TxBody struct {
	Memo string `json:"memo"`
}

type Tx struct {
	Body TxBody `json:"body"`
}

type TxResponse struct {
	Tx Tx `json:"tx"`
}
