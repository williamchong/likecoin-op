package preparememos

type Output struct {
	BookNFTId   string `json:"book_nft_id"`
	TokenId     uint64 `json:"token_id"`
	Memo        string `json:"memo"`
	BlockNumber uint64 `json:"block_number"`
}
