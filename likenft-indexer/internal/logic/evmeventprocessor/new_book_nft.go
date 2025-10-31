package evmeventprocessor

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/like_protocol"
	"likenft-indexer/internal/evm/model"
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

	blockHeader, err := e.evmClient.GetHeaderByBlockNumber(
		ctx,
		uint64(evmEvent.BlockNumber))
	if err != nil {
		mylogger.Error("e.evmClient.GetHeaderByBlockNumber", "err", err)
		return err
	}

	contractLevelMetadata := new(model.ContractLevelMetadata)
	err = jsondatauri.JSONDataUri(newBookNFTEvent.Config.Metadata).Resolve(e.httpClient, &contractLevelMetadata)

	if err != nil {
		mylogger.Error("JSONDataUri.Resolve", "err", err)
		return err
	}

	ownerAddress, err := e.evmClient.GetBookNFTOwner(ctx, newBookNFTEvent.BookNFT)
	ownerAddressStr := ownerAddress.Hex()

	if err != nil {
		mylogger.Error("e.evmClient.GetBookNFTOwner", "err", err)
		return err
	}

	account, err := e.accountRepository.GetOrCreateAccount(ctx, &ent.Account{
		CosmosAddress: nil,
		EvmAddress:    ownerAddressStr,
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

	return e.nftClassRepository.InsertNFTClass(
		ctx,
		newBookNFTEvent.BookNFT.Hex(),
		newBookNFTEvent.Config.Name,
		newBookNFTEvent.Config.Symbol,
		&ownerAddressStr, // TODO
		[]string{},       // TODO
		totalSupply,
		typeutil.Uint64(newBookNFTEvent.Config.MaxSupply),
		contractLevelMetadata,
		"",
		"",
		common.BytesToAddress([]byte{}).Hex(),
		evmEvent.BlockNumber,
		evmEvent.BlockNumber,
		time.Unix(int64(blockHeader.Time), 0),
		account,
	)
}

func init() {
	registerEventProcessor(
		"NewBookNFT",
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
