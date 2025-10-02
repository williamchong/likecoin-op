package workflow

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/mintnfts"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/newnftclass"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/migratedb/likenftassetmigrationclass"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/migratedb/likenftassetmigrationnft"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/migratedb/likenftmigrationactionmintnft"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/migratedb/likenftmigrationactionnewclass"
)

var MigratedbCmd = &cobra.Command{
	Use:   "migratedb <airdrop-output-file>",
	Short: "Migrate database from airdrop output file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		airdropOutputFile := args[0]
		airdropOutputBytes, err := os.ReadFile(airdropOutputFile)
		if err != nil {
			panic(err)
		}

		var airdropOutput struct {
			MintNFTsOutputs    []*mintnfts.Output    `json:"mint_nfts_outputs"`
			NewNFTClassOutputs []*newnftclass.Output `json:"new_nft_class_outputs"`
		}
		err = json.Unmarshal(airdropOutputBytes, &airdropOutput)

		var newClassInputs []*likenftmigrationactionnewclass.Input
		for _, newNFTClassOutput := range airdropOutput.NewNFTClassOutputs {
			newClassInputs = append(newClassInputs, &likenftmigrationactionnewclass.Input{
				Output: *newNFTClassOutput,
			})
		}
		insertNewClassSqls := likenftmigrationactionnewclass.MakeOutputs(newClassInputs).ToSQL()

		var mintNFTsInputs []*likenftmigrationactionmintnft.Input
		for _, mintNFTOutput := range airdropOutput.MintNFTsOutputs {
			mintNFTsInputs = append(mintNFTsInputs, &likenftmigrationactionmintnft.Input{
				Output: *mintNFTOutput,
			})
		}
		insertMintNFTsSqls := likenftmigrationactionmintnft.MakeOutputs(mintNFTsInputs).ToSQL()

		var assertMigrationClassInputs []*likenftassetmigrationclass.Input
		for _, assertMigrationClassOutput := range airdropOutput.NewNFTClassOutputs {
			assertMigrationClassInputs = append(assertMigrationClassInputs, &likenftassetmigrationclass.Input{
				Output: *assertMigrationClassOutput,
			})
		}
		insertAssertMigrationClassSqls := likenftassetmigrationclass.MakeOutputs(assertMigrationClassInputs).ToSQL()

		migrationNFTInputs := make([]*likenftassetmigrationnft.Input, 0)
		evmClassIdMap := make(map[string]*newnftclass.Output)
		for _, newNFTClassOutput := range airdropOutput.NewNFTClassOutputs {
			if newNFTClassOutput.CosmosClassId == nil {
				continue
			}
			evmClassIdMap[newNFTClassOutput.BaseEvmClassId] = newNFTClassOutput
		}

		for _, nft := range airdropOutput.MintNFTsOutputs {
			newNFTClassOutput, ok := evmClassIdMap[nft.EvmClassId]
			if !ok {
				continue
			}

			migrationNFTInputs = append(migrationNFTInputs, &likenftassetmigrationnft.Input{
				NewNFTClassOutput: *newNFTClassOutput,
				MintNFTsOutput:    *nft,
			})
		}

		insertNftMigrationSqls := likenftassetmigrationnft.MakeOutputs(migrationNFTInputs).ToSQL()

		fmt.Println(insertNewClassSqls)
		fmt.Println(insertMintNFTsSqls)
		fmt.Println(insertAssertMigrationClassSqls)
		fmt.Println(insertNftMigrationSqls)
	},
}

func init() {
	WorkflowCmd.AddCommand(MigratedbCmd)
}
