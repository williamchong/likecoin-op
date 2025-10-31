package evmeventprocessor

import (
	"context"
	"log/slog"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/model"
	"likenft-indexer/internal/evm/util/logconverter"
)

type transferWithMemoProcessor struct {
	httpClient                *http.Client
	evmClient                 *evm.EvmClient
	nftClassRepository        database.NFTClassRepository
	nftRepository             database.NFTRepository
	transactionMemoRepository database.TransactionMemoRepository
	accountRepository         database.AccountRepository
}

func MakeTransferWithMemoProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
	nftRepository database.NFTRepository,
	transactionMemoRepository database.TransactionMemoRepository,
	accountRepository database.AccountRepository,
) *transferWithMemoProcessor {
	return &transferWithMemoProcessor{
		httpClient:                httpClient,
		evmClient:                 evmClient,
		nftClassRepository:        nftClassRepository,
		nftRepository:             nftRepository,
		transactionMemoRepository: transactionMemoRepository,
		accountRepository:         accountRepository,
	}
}

func (e *transferWithMemoProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	mylogger := logger.WithGroup("processTransferWithMemo").
		With("evmEventId", evmEvent.ID)

	logConverter := logconverter.NewLogConverter(e.evmClient.BookNFTABI)

	newBookNFTLog := logConverter.ConvertEvmEventToLog(evmEvent)

	transferWithMemoEvent := new(book_nft.BookNftTransferWithMemo)
	err := logConverter.UnpackLog(newBookNFTLog, transferWithMemoEvent)
	if err != nil {
		mylogger.Error("logConverter.UnpackLog", "err", err)
		return err
	}

	contractAddress := newBookNFTLog.Address
	tokenId := transferWithMemoEvent.TokenId
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
		EvmAddress: transferWithMemoEvent.To.Hex(),
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
		transferWithMemoEvent.To.Hex(),
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

	err = e.transactionMemoRepository.InsertTransactionMemo(ctx, &ent.TransactionMemo{
		TransactionHash: newBookNFTLog.TxHash.Hex(),
		BookNftID:       newBookNFTLog.Address.Hex(),
		From:            transferWithMemoEvent.From.Hex(),
		To:              transferWithMemoEvent.To.Hex(),
		TokenID:         typeutil.Uint64(transferWithMemoEvent.TokenId.Uint64()),
		Memo:            transferWithMemoEvent.Memo,
		BlockNumber:     typeutil.Uint64(newBookNFTLog.BlockNumber),
	})

	if err != nil {
		mylogger.Error("e.transactionMemoRepository.InsertTransactionMemo", "err", err)
		return err
	}

	err = e.nftRepository.UpdateOwner(
		ctx,
		contractAddress.Hex(),
		tokenId,
		transferWithMemoEvent.To.Hex(),
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
		"TransferWithMemo",
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeTransferWithMemoProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
				inj.nftRepository,
				inj.transactionMemoRepository,
				inj.accountRepository,
			)
		},
	)
}
