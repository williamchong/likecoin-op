package preparenfts

type Output struct {
	ContractAddress string `json:"contract_address"`
	TokenId         uint64 `json:"token_id"`
	TokenURI        string `json:"token_uri"`
	OwnerAddress    string `json:"owner_address"`
}
