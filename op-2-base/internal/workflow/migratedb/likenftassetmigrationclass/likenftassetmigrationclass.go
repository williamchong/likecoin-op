package likenftassetmigrationclass

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/sql"
	airdropnewnftclass "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/newnftclass"
)

type Input struct {
	airdropnewnftclass.Output
}

type Output struct {
	CosmosClassId string
	EvmTxHash     string
}

func MakeOutput(input *Input) *Output {
	if input.CosmosClassId == nil {
		return nil
	}
	return &Output{
		CosmosClassId: *input.CosmosClassId,
		EvmTxHash:     input.TxHash,
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
		return sql.Stmt(sql.Echo("No SQL emitted for likenft_asset_migration_class"))
	}

	stmts := make([]string, 0)

	for _, output := range o {
		stmt := fmt.Sprintf(`UPDATE likenft_asset_migration_class
SET evm_tx_hash = %s
WHERE cosmos_class_id = %s
AND status = %s`,
			o.MustMarshal(output.EvmTxHash),
			o.MustMarshal(output.CosmosClassId),
			o.MustMarshal(string(model.LikeNFTAssetMigrationClassStatusCompleted)),
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
