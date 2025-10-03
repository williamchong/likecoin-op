package likenftmigrationactionmintnft

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/util/slice"
	airdropmintnfts "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/mintnfts"
)

type Input struct {
	airdropmintnfts.Output
}

type Output struct {
	EvmClassId            string
	CosmosNFTId           string
	InitialBatchMintOwner string
	EvmOwner              string
	Status                model.LikeNFTMigrationActionMintNFTStatus
	EvmTxHash             *string
}

func MakeOutput(input *Input) *Output {
	if input.CosmosNFTId == nil {
		return nil
	}
	return &Output{
		EvmClassId:            input.EvmClassId,
		CosmosNFTId:           *input.CosmosNFTId,
		InitialBatchMintOwner: input.InitialBatchMintOwner,
		EvmOwner:              input.EvmOwner,
		Status:                model.LikeNFTMigrationActionMintNFTStatusCompleted,
		EvmTxHash:             &input.TxHash,
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
		return ""
	}

	numCol := 6
	chunkSize := int(math.Floor(float64(appdb.PGSQL_DB_PARAMS_LIMIT) / float64(numCol)))

	stmts := make([]string, 0)

	for _, chunk := range slice.ChunkBy(o, chunkSize) {
		valueTuples := make([]string, 0, len(chunk))

		for _, output := range chunk {
			values := "(" + strings.Join([]string{
				o.MustMarshal(output.EvmClassId),
				o.MustMarshal(output.CosmosNFTId),
				o.MustMarshal(output.InitialBatchMintOwner),
				o.MustMarshal(output.EvmOwner),
				o.MustMarshal(string(output.Status)),
				o.MustMarshal(*output.EvmTxHash),
			}, ",") + ")"
			valueTuples = append(valueTuples, values)
		}

		stmt := fmt.Sprintf(`INSERT INTO likenft_migration_action_mint_nft (
	evm_class_id,
	cosmos_nft_id,
	initial_batch_mint_owner,
	evm_owner,
	status,
	evm_tx_hash
) VALUES
	%s`, strings.Join(valueTuples, ",\n\t"))

		stmts = append(stmts, stmt)
	}

	return strings.Join(stmts, ";\n") + ";"
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
