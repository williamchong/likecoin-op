package workflow

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareactions"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareparam/mintnfts"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareparam/newnftclass"
)

var PrepareParamsCmd = &cobra.Command{
	Use:   "prepare-params <actions-path>",
	Short: "Prepare params for migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		actionsPath := args[0]

		logger := slog.New(slog.Default().Handler()).
			WithGroup("PrepareParamsCmd").
			With("actionsPath", actionsPath)

		actionsBytes, err := os.ReadFile(actionsPath)
		if err != nil {
			panic(err)
		}

		actions := new(prepareactions.Output)
		err = json.Unmarshal(actionsBytes, &actions)
		if err != nil {
			panic(err)
		}

		logger = logger.With("newClassActionsCount", len(actions.NewClassActions))
		logger = logger.With("mintNFTsCount", len(actions.MintNFTActions))

		prepareNewNFTClassParam := newnftclass.NewPrepareNewNFTClassParam(
			&http.Client{
				Timeout: 10 * time.Second,
			},
		)

		prepareMintNFTsParam := mintnfts.NewPrepareMintNFTsParam()

		newNFTClassParams := make([]*newnftclass.Output, 0)
		for i, action := range actions.NewClassActions {
			logger := logger.With("newClassActionIndex", i)
			newNFTClassParam, err := prepareNewNFTClassParam.Prepare(
				cmd.Context(),
				logger,
				&newnftclass.Input{
					PrepareNewClassActionOutput: *action,
				},
			)
			if err != nil {
				panic(err)
			}
			newNFTClassParams = append(newNFTClassParams, newNFTClassParam)
		}
		mintNFTsParams := make([]*mintnfts.Output, 0)
		for i, action := range actions.MintNFTActions {
			logger := logger.With("mintNFTsIndex", i)
			mintNFTsParam, err := prepareMintNFTsParam.Prepare(
				cmd.Context(),
				logger,
				&mintnfts.Input{
					PrepareMintNFTActionOutput: *action,
				})
			if err != nil {
				panic(err)
			}
			mintNFTsParams = append(mintNFTsParams, mintNFTsParam)
		}

		json.NewEncoder(os.Stdout).Encode(struct {
			NewNFTClassParams []*newnftclass.Output `json:"new_nft_class_params"`
			MintNFTsParams    []*mintnfts.Output    `json:"mint_nfts_params"`
		}{
			NewNFTClassParams: newNFTClassParams,
			MintNFTsParams:    mintNFTsParams,
		})
	},
}

func init() {
	WorkflowCmd.AddCommand(PrepareParamsCmd)
}
