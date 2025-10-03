package prepareactions

import (
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparebooknfts"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparememos"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparenfts"
)

type BookNFTInput struct {
	preparebooknfts.Output
}

type NFTsInput struct {
	NFTs []preparenfts.Output
}

type MemosInput struct {
	Memos []preparememos.Output
}

type Input struct {
	BookNFTInput
	NFTsInput
	MemosInput
}

type PreparedNewClassAction struct {
	CosmosClassId          *string  `json:"cosmos_class_id"`
	InitialOwner           string   `json:"initial_owner"`
	InitialMinters         []string `json:"initial_minters"`
	InitialUpdaters        []string `json:"initial_updaters"`
	InitialBatchMintOwner  string   `json:"initial_batch_mint_owner"`
	DefaultRoyaltyFraction string   `json:"default_royalty_fraction"`
	ShouldPremintAllNFTs   bool     `json:"should_premint_all_nfts"`
}

type PrepareNewClassActionOutput struct {
	BookNFTInput
	PreparedNewClassAction
}

type PrepareMintNFTActionInput struct {
	MetadataStr string   `json:"metadata"`
	TokenId     uint64   `json:"token_id"`
	Memos       []string `json:"memos"`
}

type PrepareMintNFTActionOutput struct {
	PrepareMintNFTActionInput
	EvmClassId            string  `json:"evm_class_id"`
	CosmosNFTId           *string `json:"cosmos_nft_id"`
	InitialBatchMintOwner string  `json:"initial_batch_mint_owner"`
	EvmOwner              string  `json:"evm_owner"`
}

type Output struct {
	NewClassAction *PrepareNewClassActionOutput  `json:"new_class_action"`
	MintNFTActions []*PrepareMintNFTActionOutput `json:"mint_nft_actions"`
}

func (o *Output) Merge(others ...*Output) *Output {
	mintNFTActions := o.MintNFTActions
	for _, other := range others {
		mintNFTActions = append(mintNFTActions, other.MintNFTActions...)
	}
	return &Output{
		NewClassAction: o.NewClassAction,
		MintNFTActions: mintNFTActions,
	}
}
