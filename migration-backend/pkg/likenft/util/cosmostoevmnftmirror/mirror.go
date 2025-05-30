package cosmostoevmnftmirror

import (
	"context"
	"encoding/json"
	"log/slog"
	"math"

	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/event"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
)

type CosmosToEVMNFTMirrorResult struct {
	FromID    uint64
	ToID      uint64
	EvmTxHash string
}

type CosmosToEVMNFTMirror interface {
	Mirror(
		ctx context.Context,
		cosmosClassId string,
		evmClassId string,
		expectedSupply uint64,
	) ([]CosmosToEVMNFTMirrorResult, error)
}

type cosmosToEVMNFTMirror struct {
	logger *slog.Logger

	cosmosNftIDMatcher nftidmatcher.CosmosNFTIDMatcher

	likeNFTCosmosClient *cosmos.LikeNFTCosmosClient
	bookNFTEvmClient    *evm.BookNFT

	initialBatchMintOwnerAddress string
	mintPageSize                 uint64
}

func MakeCosmosToEVMNFTMirror(
	logger *slog.Logger,

	cosmosNftIDMatcher nftidmatcher.CosmosNFTIDMatcher,

	likeNFTCosmosClient *cosmos.LikeNFTCosmosClient,
	bookNFTEvmClient *evm.BookNFT,

	initialBatchMintOwnerAddress string,
	mintPageSize uint64,
) CosmosToEVMNFTMirror {
	return &cosmosToEVMNFTMirror{
		logger:                       logger,
		cosmosNftIDMatcher:           cosmosNftIDMatcher,
		likeNFTCosmosClient:          likeNFTCosmosClient,
		bookNFTEvmClient:             bookNFTEvmClient,
		initialBatchMintOwnerAddress: initialBatchMintOwnerAddress,
		mintPageSize:                 mintPageSize,
	}
}

func (m *cosmosToEVMNFTMirror) Mirror(
	ctx context.Context,
	cosmosClassId string,
	evmClassId string,
	expectedSupply uint64,
) ([]CosmosToEVMNFTMirrorResult, error) {
	mylogger := m.logger.WithGroup("Mirror").
		With("cosmosClassId", cosmosClassId).
		With("evmClassId", evmClassId).
		With("expectedSupply", expectedSupply)

	evmClassAddress := common.HexToAddress(evmClassId)

	// First check the total supply to estimate the number of pages to process
	currentTotalSupplyBigInt, err := m.bookNFTEvmClient.TotalSupply(evmClassAddress)
	if err != nil {
		return nil, err
	}
	currentTotalSupply := currentTotalSupplyBigInt.Uint64()
	mylogger = mylogger.With("currentTotalSupply", currentTotalSupply)

	if currentTotalSupply >= expectedSupply {
		mylogger.Info("Current Total Supply suffice. No need to mirror.")
		return nil, nil
	}

	cosmosClassResponse, err := m.likeNFTCosmosClient.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
		ClassId: cosmosClassId,
	})
	if err != nil {
		return nil, err
	}
	cosmosClass := cosmosClassResponse.Class
	mylogger = mylogger.
		With("iscnIdPrefix", cosmosClass.Data.Parent.IscnIdPrefix).
		With("iscnVersionAtMint", cosmosClass.Data.Parent.IscnVersionAtMint)

	iscnDataResponse, err := m.likeNFTCosmosClient.GetISCNRecord(
		cosmosClass.Data.Parent.IscnIdPrefix,
		cosmosClass.Data.Parent.IscnVersionAtMint,
	)
	if err != nil {
		return nil, err
	}

	cosmosNFTsResponse, err := m.likeNFTCosmosClient.QueryAllNFTsByClassId(cosmos.QueryAllNFTsByClassIdRequest{
		ClassId: cosmosClassId,
	})
	if err != nil {
		return nil, err
	}
	cosmosNFTs := cosmosNFTsResponse.NFTs

	numberOfPageToCall := uint64(math.Ceil((float64(expectedSupply) - float64(currentTotalSupply)) / float64(m.mintPageSize)))
	mylogger = mylogger.With("numberOfPageToCall", numberOfPageToCall)

	results := make([]CosmosToEVMNFTMirrorResult, 0)
	for p := uint64(0); p < numberOfPageToCall; p++ {
		// As the total supply should be updated for any reason, revise the total supply.
		totalSupplyBigInt, err := m.bookNFTEvmClient.TotalSupply(evmClassAddress)
		if err != nil {
			return nil, err
		}
		totalSupply := totalSupplyBigInt.Uint64()
		mylogger := mylogger.With("totalSupply", totalSupply)

		if totalSupply >= expectedSupply {
			break
		}

		tos := make([]common.Address, 0)
		memos := make([]string, 0)
		metadataList := make([]string, 0)
		mylogger = mylogger.With("callPage", p)

		fromID := totalSupply
		toID := fromID
		for i := totalSupply; i < totalSupply+m.mintPageSize; i++ {
			if i >= expectedSupply {
				break
			}
			mylogger := mylogger.With("i", i)
			cosmosNFT, found := m.cosmosNftIDMatcher.FindCosmosNFTBySerial(cosmosNFTs, i)
			metadataStr := "{}"
			memo := ""
			if found {
				metadataOverride, err := m.likeNFTCosmosClient.QueryNFTExternalMetadata(cosmosNFT)
				if err != nil {
					return nil, err
				}
				metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
					cosmosNFT,
					cosmosClass,
					iscnDataResponse,
					metadataOverride,
				))
				if err != nil {
					return nil, err
				}
				metadataStr = string(metadataBytes)

				events, err := m.likeNFTCosmosClient.QueryAllNFTEvents(
					m.likeNFTCosmosClient.MakeQueryNFTEventsRequest(cosmosNFT.ClassId, cosmosNFT.Id),
				)
				if err != nil {
					return nil, err
				}
				memo = event.MakeMemoFromEvent(events)
			}
			mylogger = mylogger.With("metadataStr", metadataStr).With("memo", memo)
			tos = append(tos, common.HexToAddress(m.initialBatchMintOwnerAddress))
			memos = append(memos, memo)
			metadataList = append(metadataList, metadataStr)
			mylogger.Info("Data Preparation Done")
			toID = i
		}

		if len(tos) == 0 {
			// No more minting
			break
		}

		_, txReceipt, err := m.bookNFTEvmClient.MintNFTs(
			ctx,
			mylogger,
			evmClassAddress,
			totalSupplyBigInt,
			tos,
			memos,
			metadataList,
		)

		if err != nil {
			return nil, err
		}

		mylogger.With("txHash", txReceipt.TxHash).Info("Page Completed")
		results = append(results, CosmosToEVMNFTMirrorResult{
			FromID:    fromID,
			ToID:      toID,
			EvmTxHash: txReceipt.TxHash.Hex(),
		})
	}

	return results, nil
}
