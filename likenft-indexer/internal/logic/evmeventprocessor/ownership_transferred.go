package evmeventprocessor

import (
	"context"
	"log/slog"

	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/util/logconverter"
)

type bookNFTOwnershipTransferredProcessor struct {
	evmClient          *evm.EvmClient
	nftClassRepository database.NFTClassRepository
	accountRepository  database.AccountRepository
}

func MakeBookNFTOwnershipTransferredProcessor(
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
	accountRepository database.AccountRepository,
) eventProcessor {
	return &bookNFTOwnershipTransferredProcessor{
		evmClient:          evmClient,
		nftClassRepository: nftClassRepository,
		accountRepository:  accountRepository,
	}
}

func (e *bookNFTOwnershipTransferredProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	mylogger := logger.WithGroup("processBookNFTOwnershipTransferred").
		With("evmEventId", evmEvent.ID)

	logConverter := logconverter.NewLogConverter(e.evmClient.LikeProtocolABI)

	bookNFTOwnershipTransferredLog := logConverter.ConvertEvmEventToLog(evmEvent)

	bookNFTOwnershipTransferredEvent := new(book_nft.BookNftOwnershipTransferred)

	err := logConverter.UnpackLog(bookNFTOwnershipTransferredLog, bookNFTOwnershipTransferredEvent)

	if err != nil {
		mylogger.Error("logConverter.UnpackLog", "err", err)
		return err
	}

	account, err := e.accountRepository.GetOrCreateAccount(ctx, &ent.Account{
		CosmosAddress: nil,
		EvmAddress:    bookNFTOwnershipTransferredEvent.NewOwner.Hex(),
		Likeid:        nil,
	})
	if err != nil {
		mylogger.Error("e.accountRepository.GetOrCreateAccount", "err", err)
		return err
	}

	err = e.nftClassRepository.UpdateOwner(ctx, bookNFTOwnershipTransferredLog.Address.Hex(), account)

	if err != nil {
		mylogger.Error("e.nftClassRepository.UpdateOwner", "err", err)
		return err
	}

	return nil
}

func init() {
	registerEventProcessor(
		"OwnershipTransferred",
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeBookNFTOwnershipTransferredProcessor(
				inj.evmClient,
				inj.nftClassRepository,
				inj.accountRepository,
			)
		},
	)
}
