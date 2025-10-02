package workflow

import (
	"encoding/json"
	"log"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/cli/context"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareactions"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparememos"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparenfts"
)

var PrepareActionsCmd = &cobra.Command{
	Use:   "prepare-actions <indexer-dump-path> <nfts-dump-path> <memos-dump-path>",
	Short: "Prepare actions",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		indexerDumpPath := args[0]
		nftsDumpPath := args[1]
		memosDumpPath := args[2]

		envCfg := context.ConfigFromContext(cmd.Context())

		bytecodeFile, err := cmd.Flags().GetString("bytecode-file")
		if err != nil {
			cmd.PrintErrf("failed to get flag bytecodeFile: %v\n", err)
			return
		}
		byteCodeData, err := os.ReadFile(bytecodeFile)
		if err != nil {
			cmd.PrintErrf("failed to read bytecode file '%s': %v\n", bytecodeFile, err)
			return
		}

		likeProtocolAddressStr := envCfg.BaseEthLikeNFTContractAddress
		likeProtocolAddress := common.HexToAddress(likeProtocolAddressStr)

		signerAddressStr := envCfg.BaseEthSignerAddress
		signerAddress := common.HexToAddress(signerAddressStr)

		logger := slog.New(slog.Default().Handler()).
			WithGroup("PrepareActionsCmd").
			With("indexerDumpPath", indexerDumpPath).
			With("nftsDumpPath", nftsDumpPath)

		defaultRoyaltyFraction, ok := big.NewInt(0).SetString("500", 10)
		if !ok {
			panic("failed to set default royalty fraction")
		}

		indexerDump, err := os.ReadFile(indexerDumpPath)
		if err != nil {
			log.Fatalf("failed to read indexer dump: %v", err)
			panic(err)
		}

		nftsDump, err := os.ReadFile(nftsDumpPath)
		if err != nil {
			log.Fatalf("failed to read nfts dump: %v", err)
			panic(err)
		}

		var nfts []preparenfts.Output
		err = json.Unmarshal(nftsDump, &nfts)
		if err != nil {
			log.Fatalf("failed to unmarshal nfts dump: %v", err)
		}
		nftsMap := make(map[string][]preparenfts.Output)
		for _, nft := range nfts {
			if _, ok := nftsMap[nft.ContractAddress]; !ok {
				nftsMap[nft.ContractAddress] = make([]preparenfts.Output, 0)
			}
			nftsMap[nft.ContractAddress] = append(nftsMap[nft.ContractAddress], nft)
		}

		memosDump, err := os.ReadFile(memosDumpPath)
		if err != nil {
			log.Fatalf("failed to read memos dump: %v", err)
			panic(err)
		}

		var memos []preparememos.Output
		err = json.Unmarshal(memosDump, &memos)
		if err != nil {
			log.Fatalf("failed to unmarshal memos dump: %v", err)
			panic(err)
		}

		memosByBookNFTId := make(map[string][]preparememos.Output)
		for _, memo := range memos {
			if _, ok := memosByBookNFTId[memo.BookNFTId]; !ok {
				memosByBookNFTId[memo.BookNFTId] = make([]preparememos.Output, 0)
			}
			memosByBookNFTId[memo.BookNFTId] = append(memosByBookNFTId[memo.BookNFTId], memo)
		}

		prepareActions := prepareactions.NewPrepareNewNFTClassAction(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			defaultRoyaltyFraction,
			creationcode.NewCreationCode(byteCodeData),
			likeProtocolAddress,
			signerAddress,
		)

		bookNFTInputs := make([]*prepareactions.BookNFTInput, 0)

		err = json.Unmarshal(indexerDump, &bookNFTInputs)
		if err != nil {
			log.Fatalf("failed to unmarshal indexer dump: %v", err)
		}

		logger = logger.With("count", len(bookNFTInputs))

		outputs := make([]*prepareactions.Output, 0)
		for i, bookNFTInput := range bookNFTInputs {
			logger := logger.With("index", i)
			nfts, ok := nftsMap[bookNFTInput.OpAddress]
			if !ok {
				log.Fatalf("nfts not found for contract address: %s", bookNFTInput.OpAddress)
			}
			memosOfBookNFTId, ok := memosByBookNFTId[bookNFTInput.OpAddress]
			if !ok {
				log.Fatalf("memos not found for contract address: %s", bookNFTInput.OpAddress)
				memosOfBookNFTId = make([]preparememos.Output, 0)
			}
			input := &prepareactions.Input{
				BookNFTInput: *bookNFTInput,
				NFTsInput: prepareactions.NFTsInput{
					NFTs: nfts,
				},
				MemosInput: prepareactions.MemosInput{
					Memos: memosOfBookNFTId,
				},
			}
			output, err := prepareActions.Prepare(cmd.Context(), logger, input)
			if err != nil {
				log.Fatalf("failed to prepare actions: %v", err)
			}
			outputs = append(outputs, output)
		}
		prepareActionsOutput := new(prepareactions.Output).Merge(outputs...)

		json.NewEncoder(os.Stdout).Encode(prepareActionsOutput)
	},
}

func init() {
	PrepareActionsCmd.Flags().String(
		"bytecode-file",
		"BeaconProxy.creationCode",
		"Path to bytecode file (default: BeaconProxy.creationCode)",
	)
	WorkflowCmd.AddCommand(PrepareActionsCmd)
}
