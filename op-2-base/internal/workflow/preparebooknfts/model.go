package preparebooknfts

import "github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"

type Output struct {
	OpAddress    string                       `json:"address"`
	Name         string                       `json:"name"`
	OwnerAddress string                       `json:"owner_address"`
	Metadata     *model.ContractLevelMetadata `json:"metadata"`
	Salt         *string                      `json:"salt"`
	Salt2        *string                      `json:"salt2"`
	Count        uint64                       `json:"count"`
	MaxSupply    uint64                       `json:"max_supply"`
}
