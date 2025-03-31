package evmeventprocessor

import (
	"context"
	"log/slog"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/util/logconverter"
)

type contractURIUpdatedProcessor struct {
	httpClient         *http.Client
	evmClient          *evm.EvmClient
	nftClassRepository database.NFTClassRepository
}

func MakeContractURIUpdatedProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
) eventProcessor {
	return &contractURIUpdatedProcessor{
		httpClient:         httpClient,
		evmClient:          evmClient,
		nftClassRepository: nftClassRepository,
	}
}

func (e *contractURIUpdatedProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	mylogger := logger.WithGroup("processContractURIUpdated").
		With("evmEventId", evmEvent.ID)

	logConverter := logconverter.NewLogConverter(e.evmClient.LikeProtocolABI)

	newBookNFTLog := logConverter.ConvertEvmEventToLog(evmEvent)

	contractURI, err := e.evmClient.GetContractURI(ctx, newBookNFTLog.Address)

	if err != nil {
		mylogger.Error("e.evmClient.GetContractURI", "err", err)
		return err
	}

	contractLevelMetadata := make(map[string]any)
	err = contractURI.Resolve(e.httpClient, &contractLevelMetadata)

	if err != nil {
		mylogger.Error("contractURI.Resolve", "err", err)
		return err
	}

	err = e.nftClassRepository.UpdateMetadata(
		ctx,
		newBookNFTLog.Address.Hex(),
		contractLevelMetadata,
	)

	if err != nil {
		mylogger.Error("e.nftClassRepository.UpdateMetadata", "err", err)
		return err
	}

	return nil
}

func init() {
	registerEventProcessor(
		evmeventprocessedblockheight.EventContractURIUpdated.String(),
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeContractURIUpdatedProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
			)
		},
	)
}
