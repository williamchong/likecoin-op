package prepareactions

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"
	"github.com/likecoin/like-migration-backend/pkg/util/jsondatauri"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparememos"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparenfts"
)

type PrepareNewNFTClassAction interface {
	Prepare(
		ctx context.Context,
		logger *slog.Logger,
		input *Input,
	) (*Output, error)
}

type prepareNewNFTClassAction struct {
	httpClient             *http.Client
	defaultRoyaltyFraction *big.Int
	creationCode           creationcode.CreationCode
	protocolAddress        common.Address
	signerAddress          common.Address
}

func NewPrepareNewNFTClassAction(
	httpClient *http.Client,
	defaultRoyaltyFraction *big.Int,
	creationCode creationcode.CreationCode,
	protocolAddress common.Address,
	signerAddress common.Address,
) PrepareNewNFTClassAction {
	return &prepareNewNFTClassAction{
		httpClient,
		defaultRoyaltyFraction,
		creationCode,
		protocolAddress,
		signerAddress,
	}
}

func (p *prepareNewNFTClassAction) Prepare(
	ctx context.Context,
	logger *slog.Logger,
	input *Input,
) (*Output, error) {
	mylogger := logger.
		WithGroup("PrepareNewNFTClassAction").
		With("opAddress", input.OpAddress)
	mylogger.Info("PrepareNewNFTClassAction")

	var cosmosClassId *string
	if input.Metadata.LikeCoin != nil {
		cosmosClassId = &input.Metadata.LikeCoin.ClassId
	}

	name := input.Metadata.Name
	symbol := input.Metadata.Symbol

	initCodeHash, err := p.creationCode.MakeInitCodeHash(p.protocolAddress, name, symbol)
	if err != nil {
		return nil, fmt.Errorf("creationCode.MakeInitCodeHash: %v", err)
	}

	salt, err := evm.ComputeSaltDataFromCandidates(
		p.signerAddress,
		[2]byte{0, 0},
		input.Salt,
		input.Salt2,
	)
	if err != nil {
		return nil, fmt.Errorf("evm.ComputeNewBookNFTSalt: %v", err)
	}

	bookNFTAddress := crypto.CreateAddress2(p.signerAddress, salt, initCodeHash)

	initialOwner := input.OwnerAddress

	// TODO
	// Setting signer as initial minters and updaters
	// For nft token migration
	initialMinters := []string{p.signerAddress.Hex()}
	initialUpdaters := []string{p.signerAddress.Hex()}
	initialBatchMintOwner := p.signerAddress.Hex()
	defaultRoyaltyFraction := p.defaultRoyaltyFraction
	shouldPremintAllNFTs := false

	mylogger = mylogger.With("totalSupply", input.Count)

	memosOfBookNFTIdMap := make(map[string]map[uint64][]preparememos.Output)
	for _, memo := range input.Memos {
		if _, ok := memosOfBookNFTIdMap[memo.BookNFTId]; !ok {
			memosOfBookNFTIdMap[memo.BookNFTId] = make(map[uint64][]preparememos.Output)
		}
		if _, ok := memosOfBookNFTIdMap[memo.BookNFTId][memo.TokenId]; !ok {
			memosOfBookNFTIdMap[memo.BookNFTId][memo.TokenId] = make([]preparememos.Output, 0)
		}
		memosOfBookNFTIdMap[memo.BookNFTId][memo.TokenId] = append(
			memosOfBookNFTIdMap[memo.BookNFTId][memo.TokenId],
			memo,
		)
	}

	memosOfBookNFTId, ok := memosOfBookNFTIdMap[input.OpAddress]
	if !ok {
		memosOfBookNFTId = make(map[uint64][]preparememos.Output)
	}

	mintNFTActions := make([]*PrepareMintNFTActionOutput, 0)
	for i := uint64(0); i < input.Count; i++ {
		nft := input.NFTs[i]

		memosOfTokenId, ok := memosOfBookNFTId[nft.TokenId]
		if !ok {
			memosOfTokenId = make([]preparememos.Output, 0)
		}
		slices.SortFunc(memosOfTokenId, func(a, b preparememos.Output) int {
			aBlockNumber := big.NewInt(0).SetUint64(a.BlockNumber)
			bBlockNumber := big.NewInt(0).SetUint64(b.BlockNumber)
			return aBlockNumber.Cmp(bBlockNumber)
		})

		memos := make([]string, 0)
		for _, memo := range memosOfTokenId {
			memos = append(memos, memo.Memo)
		}

		mintNFTAction, err := p.prepareMintNFTAction(
			ctx,
			mylogger,
			common.HexToAddress(input.OpAddress),
			bookNFTAddress,
			nft,
			memos,
		)
		if err != nil {
			return nil, fmt.Errorf("prepareMintNFTAction: %v", err)
		}
		mintNFTActions = append(mintNFTActions, mintNFTAction)
	}

	output := &Output{
		NewClassActions: []*PrepareNewClassActionOutput{
			{
				BookNFTInput: input.BookNFTInput,
				PreparedNewClassAction: PreparedNewClassAction{
					CosmosClassId:          cosmosClassId,
					InitialOwner:           initialOwner,
					InitialMinters:         initialMinters,
					InitialUpdaters:        initialUpdaters,
					InitialBatchMintOwner:  initialBatchMintOwner,
					DefaultRoyaltyFraction: defaultRoyaltyFraction.String(),
					ShouldPremintAllNFTs:   shouldPremintAllNFTs,
				},
			},
		},
		MintNFTActions: mintNFTActions,
	}

	mylogger.Info("PrepareNewNFTClassAction completed")
	return output, nil
}

func (p *prepareNewNFTClassAction) prepareMintNFTAction(
	ctx context.Context,
	logger *slog.Logger,
	opClassId common.Address,
	baseClassId common.Address,
	nft preparenfts.Output,
	memos []string,
) (*PrepareMintNFTActionOutput, error) {
	mylogger := logger.WithGroup("prepareMintNFTAction").
		With("tokenId", nft.TokenId)

	mylogger.Info("prepareMintNFTAction")

	tokenUriStr := nft.TokenURI

	tokenUriJsonDataUri := jsondatauri.JSONDataUri(tokenUriStr)

	metadataStr, err := tokenUriJsonDataUri.Decoded(p.httpClient)
	if err != nil {
		return nil, fmt.Errorf("jsondatauri.JSONDataUri.Decoded: %v", err)
	}

	metadata := new(model.ERC721Metadata)
	err = tokenUriJsonDataUri.Resolve(p.httpClient, metadata)
	if err != nil {
		return nil, fmt.Errorf("jsondatauri.JSONDataUri.Resolve: %v", err)
	}

	var cosmosNFTId *string
	if metadata.LikeCoin != nil {
		cosmosNFTId = &metadata.LikeCoin.NFTId
	}

	return &PrepareMintNFTActionOutput{
		PrepareMintNFTActionInput: PrepareMintNFTActionInput{
			MetadataStr: metadataStr,
			TokenId:     nft.TokenId,
			Memos:       memos,
		},
		EvmClassId:            baseClassId.Hex(),
		CosmosNFTId:           cosmosNFTId,
		InitialBatchMintOwner: p.signerAddress.Hex(),
		EvmOwner:              nft.OwnerAddress,
	}, nil
}
