package workflow

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/signer"
	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/cli/context"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/mintnfts"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/airdrop/newnftclass"
)

var AirdropCmd = &cobra.Command{
	Use:   "airdrop <params-file> ",
	Short: "Airdrop",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envCfg := context.ConfigFromContext(cmd.Context())
		logger := slog.New(slog.Default().Handler()).WithGroup("AirdropCmd")

		paramsFile := args[0]
		paramsFileBytes, err := os.ReadFile(paramsFile)
		if err != nil {
			panic(err)
		}

		var params struct {
			NewNFTClassParams []*newnftclass.Input `json:"new_nft_class_params"`
			MintNFTsParams    []*mintnfts.Input    `json:"mint_nfts_params"`
		}
		if err := json.Unmarshal(paramsFileBytes, &params); err != nil {
			panic(err)
		}

		baseEthClient, err := ethclient.Dial(envCfg.BaseEthNetworkPublicRPCURL)
		if err != nil {
			panic(err)
		}

		baseSigner := signer.NewSignerClient(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			envCfg.BaseEthSignerBaseUrl,
			envCfg.BaseEthSignerAPIKey,
		)

		baseLikeProtocol := evm.NewLikeProtocol(
			logger,
			baseEthClient,
			baseSigner,
			common.HexToAddress(envCfg.BaseEthLikeNFTContractAddress),
		)

		airdropNFTClass := newnftclass.NewNewNFTClassAirdrop(&baseLikeProtocol)
		airdropMintNFTs := mintnfts.NewMintNFTsAirdrop(baseEthClient, baseSigner)

		logger = logger.With(
			"newNFTClassParamsCount", len(params.NewNFTClassParams),
			"mintNFTsParamsCount", len(params.MintNFTsParams),
		)

		airdropNFTClassOutputs := make([]*newnftclass.Output, 0)
		for i, param := range params.NewNFTClassParams {
			logger := logger.With("newNFTClassParamIndex", i)
			output, err := airdropNFTClass.Airdrop(
				cmd.Context(),
				logger,
				param,
			)
			if err != nil {
				panic(err)
			}
			airdropNFTClassOutputs = append(airdropNFTClassOutputs, output)
		}

		nftsMap := make(map[string][]*mintnfts.Input)
		for _, param := range params.MintNFTsParams {
			if _, ok := nftsMap[param.EvmClassId]; !ok {
				nftsMap[param.EvmClassId] = make([]*mintnfts.Input, 0)
			}
			nftsMap[param.EvmClassId] = append(nftsMap[param.EvmClassId], param)
		}

		airdropMintNFTsOutputs := make([]*mintnfts.Output, 0)
		for classId, paramsOfBookNFT := range nftsMap {
			chunks := slices.Collect(slices.Chunk(paramsOfBookNFT, 10))
			logger := logger.With("classId", classId, "numOfChunks", len(chunks))
			for i, chunk := range chunks {
				logger := logger.With(
					"chunkIndex", i,
				)
				logger.Info(
					"Minting NFTs...",
					"fromTokenId", chunk[0].TokenId,
					"toTokenId", chunk[len(chunk)-1].TokenId,
				)
				output, err := airdropMintNFTs.Airdrop(cmd.Context(), logger, chunk)
				if err != nil {
					panic(err)
				}
				airdropMintNFTsOutputs = append(airdropMintNFTsOutputs, output...)
			}
		}

		json.NewEncoder(os.Stdout).Encode(struct {
			NewNFTClassOutputs []*newnftclass.Output `json:"new_nft_class_outputs"`
			MintNFTsOutputs    []*mintnfts.Output    `json:"mint_nfts_outputs"`
		}{
			NewNFTClassOutputs: airdropNFTClassOutputs,
			MintNFTsOutputs:    airdropMintNFTsOutputs,
		})
	},
}

func init() {
	WorkflowCmd.AddCommand(AirdropCmd)
}
