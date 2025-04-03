package evmeventprocessor

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/like_protocol"
	"likenft-indexer/internal/evm/util/logconverter"
	"likenft-indexer/internal/util/jsondatauri"

	"github.com/ethereum/go-ethereum/common"
)

type newBookNFTProcessor struct {
	httpClient         *http.Client
	evmClient          *evm.EvmClient
	nftClassRepository database.NFTClassRepository
	accountRepository  database.AccountRepository
}

func MakeNewBookNFTProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
	accountRepository database.AccountRepository,
) *newBookNFTProcessor {
	return &newBookNFTProcessor{
		httpClient:         httpClient,
		evmClient:          evmClient,
		nftClassRepository: nftClassRepository,
		accountRepository:  accountRepository,
	}
}

func (e *newBookNFTProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	mylogger := logger.WithGroup("processNewBookNFT").
		With("evmEventId", evmEvent.ID)

	logConverter := logconverter.NewLogConverter(e.evmClient.LikeProtocolABI)

	newBookNFTLog := logConverter.ConvertEvmEventToLog(evmEvent)

	newBookNFTEvent := new(like_protocol.LikeProtocolNewBookNFT)

	err := logConverter.UnpackLog(newBookNFTLog, newBookNFTEvent)

	if err != nil {
		mylogger.Error("logConverter.UnpackLog", "err", err)
		return err
	}

	contractLevelMetadata := make(map[string]any)
	err = jsondatauri.JSONDataUri(newBookNFTEvent.Config.Metadata).Resolve(e.httpClient, &contractLevelMetadata)

	if err != nil {
		mylogger.Error("JSONDataUri.Resolve", "err", err)
		return err
	}

	ownerAddress, err := e.evmClient.GetBookNFTOwner(ctx, newBookNFTEvent.BookNFT)

	if err != nil {
		mylogger.Error("e.evmClient.GetBookNFTOwner", "err", err)
		return err
	}

	account, err := e.accountRepository.GetOrCreateAccount(ctx, &ent.Account{
		CosmosAddress: nil,
		EvmAddress:    ownerAddress.Hex(),
		Likeid:        nil,
	})

	if err != nil {
		mylogger.Error("e.accountRepository.GetOrCreateAccount", "err", err)
		return err
	}

	totalSupply, err := e.evmClient.GetTotalSupply(ctx, newBookNFTEvent.BookNFT)

	if err != nil {
		mylogger.Error("e.accountRepository.GetOrCreateAccount", "err", err)
		return err
	}

	nftClass := &ent.NFTClass{
		Address:             newBookNFTEvent.BookNFT.Hex(),
		Name:                newBookNFTEvent.Config.Name,
		Symbol:              newBookNFTEvent.Config.Symbol,
		OwnerAddress:        nil,        // TODO
		MinterAddresses:     []string{}, // TODO
		TotalSupply:         totalSupply,
		MaxSupply:           typeutil.Uint64(newBookNFTEvent.Config.MaxSupply),
		Metadata:            contractLevelMetadata,
		BannerImage:         "",                                    // NO DATA
		FeaturedImage:       "",                                    // NO DATA
		DeployerAddress:     common.BytesToAddress([]byte{}).Hex(), // TODO
		DeployedBlockNumber: strconv.FormatUint(evmEvent.BlockNumber, 10),
		MintedAt:            time.Now(), // TODO
		UpdatedAt:           time.Now(), // TODO
		Edges: ent.NFTClassEdges{
			Owner: account,
		},
	}

	return e.nftClassRepository.InsertNFTClass(ctx, nftClass)
}

func init() {
	registerEventProcessor(
		evmeventprocessedblockheight.EventNewBookNFT.String(),
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeNewBookNFTProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
				inj.accountRepository,
			)
		},
	)
}
