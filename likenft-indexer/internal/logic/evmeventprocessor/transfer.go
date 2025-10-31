package evmeventprocessor

import (
	"context"
	"log/slog"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/model"
	"likenft-indexer/internal/evm/util/logconverter"
)

type transferProcessor struct {
	httpClient         *http.Client
	evmClient          *evm.EvmClient
	nftClassRepository database.NFTClassRepository
	nftRepository      database.NFTRepository
	accountRepository  database.AccountRepository
}

func MakeTransferProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
	nftRepository database.NFTRepository,
	accountRepository database.AccountRepository,
) *transferProcessor {
	return &transferProcessor{
		httpClient:         httpClient,
		evmClient:          evmClient,
		nftClassRepository: nftClassRepository,
		nftRepository:      nftRepository,
		accountRepository:  accountRepository,
	}
}

func (e *transferProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	mylogger := logger.WithGroup("processTransfer").
		With("evmEventId", evmEvent.ID)

	logConverter := logconverter.NewLogConverter(e.evmClient.BookNFTABI)

	newBookNFTLog := logConverter.ConvertEvmEventToLog(evmEvent)

	transferEvent := new(book_nft.BookNftTransfer)
	err := logConverter.UnpackLog(newBookNFTLog, transferEvent)
	if err != nil {
		mylogger.Error("logConverter.UnpackLog", "err", err)
		return err
	}

	contractAddress := newBookNFTLog.Address
	tokenId := transferEvent.TokenId
	tokenURI, err := e.evmClient.GetTokenURI(ctx, contractAddress, tokenId)

	if err != nil {
		mylogger.Error("e.evmClient.GetTokenURI", "err", err)
		return err
	}

	metadata := new(model.ERC721Metadata)

	err = tokenURI.Resolve(e.httpClient, metadata)
	if err != nil {
		mylogger.Error("tokenURI.Resolve", "err", err)
		return err
	}

	totalSupply, err := e.evmClient.GetTotalSupply(ctx, contractAddress)

	if err != nil {
		mylogger.Error("e.evmClient.GetTotalSupply", "err", err)
		return err
	}

	owner, err := e.accountRepository.GetOrCreateAccount(ctx, &ent.Account{
		EvmAddress: transferEvent.To.Hex(),
	})

	if err != nil {
		mylogger.Error("e.accountRepository.GetOrCreateAccount", "err", err)
		return err
	}

	nftClass, err := e.nftClassRepository.QueryNFTClassByAddress(ctx, contractAddress.Hex())

	if err != nil {
		mylogger.Error("e.nftClassRepository.QueryNFTClassByAddress", "err", err)
		return err
	}

	_, err = e.nftRepository.GetOrCreate(
		ctx,
		contractAddress.Hex(),
		tokenId,
		string(tokenURI),
		metadata.Image,
		&metadata.ImageData,
		&metadata.ExternalUrl,
		metadata.Description,
		metadata.Name,
		metadata.Attributes,
		&metadata.BackgroundColor,
		&metadata.AnimationUrl,
		&metadata.YoutubeUrl,
		transferEvent.To.Hex(),
		owner,
		nftClass,
	)

	if err != nil {
		mylogger.Error("e.nftRepository.GetOrCreate", "err", err)
		return err
	}

	err = e.nftClassRepository.UpdateTotalSupply(ctx, contractAddress.Hex(), totalSupply)

	if err != nil {
		mylogger.Error("e.nftClassRepository.UpdateTotalSupply", "err", err)
		return err
	}

	err = e.nftRepository.UpdateOwner(
		ctx,
		contractAddress.Hex(),
		tokenId,
		transferEvent.To.Hex(),
		owner,
	)

	if err != nil {
		mylogger.Error("e.nftRepository.UpdateOwner", "err", err)
		return err
	}

	return nil
}

func init() {
	registerEventProcessor(
		"Transfer",
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeTransferProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
				inj.nftRepository,
				inj.accountRepository,
			)
		},
	)
}
