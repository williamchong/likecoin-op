package workflow

import (
	"encoding/json"
	"log"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/cli/context"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareactions"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparememos"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparenfts"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparerolechangeevents"
)

var PrepareActionsCmd = &cobra.Command{
	Use:   "prepare-actions <indexer-dump-path> <nfts-dump-path> <memos-dump-path> <role-change-events-dump-path> <book-nft-id>",
	Short: "Prepare actions",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		indexerDumpPath := args[0]
		nftsDumpPath := args[1]
		memosDumpPath := args[2]
		roleChangeEventsDumpPath := args[3]
		bookNFTId := args[4]

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

		bookNFTInputs := make([]*prepareactions.BookNFTInput, 0)

		err = json.Unmarshal(indexerDump, &bookNFTInputs)
		if err != nil {
			log.Fatalf("failed to unmarshal indexer dump: %v", err)
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

		roleChangeEventsDump, err := os.ReadFile(roleChangeEventsDumpPath)
		if err != nil {
			log.Fatalf("failed to read role change events dump: %v", err)
			panic(err)
		}

		var roleChangeEvents []preparerolechangeevents.Output

		err = json.Unmarshal(roleChangeEventsDump, &roleChangeEvents)
		if err != nil {
			log.Fatalf("failed to unmarshal role change events dump: %v", err)
			panic(err)
		}

		// End of reading files

		var bookNFTInput *prepareactions.BookNFTInput = nil
		for _, i := range bookNFTInputs {
			if strings.EqualFold(i.OpAddress, bookNFTId) {
				bookNFTInput = i
				break
			}
		}

		if bookNFTInput == nil {
			log.Fatalf("book nft input not found for book nft id: %s", bookNFTId)
		}

		nftsOfBookNFTId := make([]preparenfts.Output, 0)
		for _, nft := range nfts {
			if strings.EqualFold(nft.ContractAddress, bookNFTInput.OpAddress) {
				nftsOfBookNFTId = append(nftsOfBookNFTId, nft)
			}
		}

		slices.SortFunc(nftsOfBookNFTId, func(a, b preparenfts.Output) int {
			if a.TokenId < b.TokenId {
				return -1
			}
			if a.TokenId > b.TokenId {
				return 1
			}
			return 0
		})

		memosOfBookNFTId := make([]preparememos.Output, 0)
		for _, memo := range memos {
			if strings.EqualFold(memo.BookNFTId, bookNFTInput.OpAddress) {
				memosOfBookNFTId = append(memosOfBookNFTId, memo)
			}
		}

		roleChangeEventsOfBookNFTId := make([]preparerolechangeevents.Output, 0)
		for _, roleChangeEvent := range roleChangeEvents {
			if strings.EqualFold(roleChangeEvent.BookNFTId, bookNFTInput.OpAddress) {
				roleChangeEventsOfBookNFTId = append(roleChangeEventsOfBookNFTId, roleChangeEvent)
			}
		}

		// End of processing data

		prepareActions := prepareactions.NewPrepareNewNFTClassAction(
			&http.Client{
				Timeout: 10 * time.Second,
			},
			defaultRoyaltyFraction,
			creationcode.NewCreationCode(byteCodeData),
			likeProtocolAddress,
			signerAddress,
		)

		logger = logger.With("count", len(bookNFTInputs))

		input := &prepareactions.Input{
			BookNFTInput: *bookNFTInput,
			NFTsInput: prepareactions.NFTsInput{
				NFTs: nftsOfBookNFTId,
			},
			MemosInput: prepareactions.MemosInput{
				Memos: memosOfBookNFTId,
			},
			RoleChangeEventsInput: prepareactions.RoleChangeEventsInput{
				RoleChangeEvents: roleChangeEventsOfBookNFTId,
			},
		}
		output, err := prepareActions.Prepare(cmd.Context(), logger, input)
		if err != nil {
			log.Fatalf("failed to prepare actions: %v", err)
		}

		json.NewEncoder(os.Stdout).Encode(output)
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
