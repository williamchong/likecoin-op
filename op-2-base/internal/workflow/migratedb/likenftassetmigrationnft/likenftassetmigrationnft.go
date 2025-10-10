package likenftassetmigrationnft

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/sql"
	airdropmintnfts "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/mintnfts"
	airdropnewnftclass "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/newnftclass"
)

type Input struct {
	NewNFTClassOutput airdropnewnftclass.Output
	MintNFTsOutput    airdropmintnfts.Output
}

type Output struct {
	CosmosClassId string
	CosmosNFTId   string
	EvmTxHash     string
}

func MakeOutput(input *Input) *Output {
	if input.NewNFTClassOutput.CosmosClassId == nil || input.MintNFTsOutput.CosmosNFTId == nil {
		return nil
	}
	return &Output{
		CosmosClassId: *input.NewNFTClassOutput.CosmosClassId,
		CosmosNFTId:   *input.MintNFTsOutput.CosmosNFTId,
		EvmTxHash:     input.MintNFTsOutput.TxHash,
	}
}

type Outputs []Output

func MakeOutputs(inputs []*Input) Outputs {
	outputs := make(Outputs, 0)
	for _, input := range inputs {
		output := MakeOutput(input)
		if output == nil {
			continue
		}
		outputs = append(outputs, *output)
	}
	return outputs
}

func (o Outputs) ToSQL() string {
	if len(o) == 0 {
		return sql.Stmt(sql.Echo("No SQL emitted for likenft_asset_migration_nft"))
	}

	stmts := make([]string, 0)

	for _, output := range o {
		stmt := fmt.Sprintf(`UPDATE likenft_asset_migration_nft
SET evm_tx_hash = %s
WHERE cosmos_class_id = %s
  AND cosmos_nft_id = %s
  AND status = %s`,
			o.MustMarshal(output.EvmTxHash),
			o.MustMarshal(output.CosmosClassId),
			o.MustMarshal(output.CosmosNFTId),
			o.MustMarshal(string(model.LikeNFTAssetMigrationNFTStatusCompleted)),
		)
		stmts = append(stmts, sql.Stmt(sql.Echo(stmt)), sql.Stmt(stmt))
	}

	return strings.Join(stmts, "\n")
}

func (o Outputs) MustMarshal(v any) string {
	if s, ok := v.(string); ok {
		return `'` + s + `'`
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
