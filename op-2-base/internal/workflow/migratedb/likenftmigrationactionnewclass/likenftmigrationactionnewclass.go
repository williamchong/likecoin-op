package likenftmigrationactionnewclass

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/types/commaseparatedstring"
	"github.com/likecoin/like-migration-backend/pkg/util/slice"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/sql"
	airdropnewnftclass "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/newnftclass"
)

type Input struct {
	airdropnewnftclass.Output
}

type Output struct {
	CosmosClassId          string
	InitialOwner           string
	InitialMinters         []string
	InitialUpdaters        []string
	InitialBatchMintOwner  string
	DefaultRoyaltyFraction string
	ShouldPremintAllNFTs   bool
	Status                 model.LikeNFTMigrationActionNewClassStatus
	EvmClassId             *string
	EvmTxHash              *string
}

func MakeOutput(input *Input) *Output {
	if input.CosmosClassId == nil {
		return nil
	}
	return &Output{
		CosmosClassId:          *input.CosmosClassId,
		InitialOwner:           input.InitialOwner,
		InitialMinters:         input.InitialMinters,
		InitialUpdaters:        input.InitialUpdaters,
		InitialBatchMintOwner:  input.InitialBatchMintOwner,
		DefaultRoyaltyFraction: input.DefaultRoyaltyFraction,
		ShouldPremintAllNFTs:   false,
		Status:                 model.LikeNFTMigrationActionNewClassStatusCompleted,
		EvmClassId:             &input.BaseEvmClassId,
		EvmTxHash:              &input.TxHash,
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
		return sql.Stmt(sql.Echo("No SQL emitted for likenft_migration_action_new_class"))
	}

	numCol := 10
	chunkSize := int(math.Floor(float64(appdb.PGSQL_DB_PARAMS_LIMIT) / float64(numCol)))

	stmts := make([]string, 0)

	for _, chunk := range slice.ChunkBy(o, chunkSize) {
		valueTuples := make([]string, 0, len(chunk))

		for _, output := range chunk {
			values := "(" + strings.Join([]string{
				o.MustMarshal(output.CosmosClassId),
				o.MustMarshal(output.InitialOwner),
				o.MustMarshal(string(commaseparatedstring.FromSlice(output.InitialMinters))),
				o.MustMarshal(string(commaseparatedstring.FromSlice(output.InitialUpdaters))),
				o.MustMarshal(output.InitialBatchMintOwner),
				o.MustMarshal(output.ShouldPremintAllNFTs),
				o.MustMarshal(output.DefaultRoyaltyFraction),
				o.MustMarshal(string(output.Status)),
				o.MustMarshal(*output.EvmClassId),
				o.MustMarshal(*output.EvmTxHash),
			}, ",") + ")"
			valueTuples = append(valueTuples, values)
		}

		stmt := fmt.Sprintf(`INSERT INTO likenft_migration_action_new_class (
	cosmos_class_id,
	initial_owner,
	initial_minter,
	initial_updater,
	initial_batch_mint_owner,
	should_premint_all_nfts,
	default_royalty_fraction,
	status,
	evm_class_id,
	evm_tx_hash
) VALUES
	%s`, strings.Join(valueTuples, ",\n\t"))

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
