package likenft

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/event"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func DoMintNFTAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	p *evm.LikeProtocol,
	c *evm.BookNFT,
	m *cosmos.LikeNFTCosmosClient,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	a *model.LikeNFTMigrationActionMintNFT,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	mylogger := logger.
		WithGroup("DoMintNFTAction").
		With("mintNFTActionId", a.Id)

	if a.Status == model.LikeNFTMigrationActionMintNFTStatusCompleted {
		return a, nil
	}
	if a.Status != model.LikeNFTMigrationActionMintNFTStatusInit &&
		a.Status != model.LikeNFTMigrationActionMintNFTStatusFailed {
		return nil, errors.New("error mint nft action is not init or failed")
	}

	nftIDMatcher := nftidmatcher.MakeNFTIDMatcher()

	evmClassAddress := common.HexToAddress(a.EvmClassId)
	toOwner := common.HexToAddress(a.EvmOwner)
	initialBatchMintOwnerAddress := common.HexToAddress(a.InitialBatchMintOwner)

	newClassAction, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		EvmClassId: &a.EvmClassId,
	})
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}

	totalSupplyBigInt, err := c.TotalSupply(evmClassAddress)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	totalSupply := totalSupplyBigInt.Uint64()

	cosmosClass, err := m.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
		ClassId: newClassAction.CosmosClassId,
	})
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	iscnDataResponse, err := m.GetISCNRecord(
		cosmosClass.Class.Data.Parent.IscnIdPrefix,
		cosmosClass.Class.Data.Parent.IscnVersionAtMint,
	)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}

	nftId, ok := nftIDMatcher.ExtractSerialID(a.CosmosNFTId)

	var (
		tx *types.Transaction
	)

	if !ok {
		// arbitrary id: wnft
		// Given the mint nft action is indexed by classid and nftid, the action should be simply mint.
		//
		// Dup check (another evm owner wants to mint the same token again) is checked
		// when the record is retrieved by get or create
		cosmosNFT, err := m.QueryNFT(cosmos.QueryNFTRequest{
			ClassId: newClassAction.CosmosClassId,
			Id:      a.CosmosNFTId,
		})
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
		metadataOverride, err := m.QueryNFTExternalMetadata(cosmosNFT.NFT)
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
		metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
			erc721ExternalURLBuilder,
			cosmosNFT.NFT,
			cosmosClass.Class,
			iscnDataResponse,
			metadataOverride,
			a.EvmClassId,
			totalSupply,
		))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}

		events, err := m.QueryAllNFTEvents(m.MakeQueryNFTEventsRequest(newClassAction.CosmosClassId, a.CosmosNFTId))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}

		metadataString := string(metadataBytes)
		tx, _, err = c.MintNFTs(
			ctx,
			mylogger,
			evmClassAddress,
			totalSupplyBigInt,
			[]common.Address{toOwner},
			[]string{
				event.MakeMemoFromEvent(events),
			},
			[]string{
				metadataString,
			})
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
	} else {

		desireSupply := nftId + 1
		desireBatchMintAmount := uint64(math.Max(float64(desireSupply)-float64(totalSupply), 0))
		if desireBatchMintAmount > 0 {
			cosmosNFTs, err := m.QueryAllNFTsByClassId(cosmos.QueryAllNFTsByClassIdRequest{
				ClassId: newClassAction.CosmosClassId,
			})

			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}

			tos := make([]common.Address, 0)
			memos := make([]string, 0)
			metadataList := make([]string, 0)
			for i := totalSupply; i < desireSupply; i = i + 1 {
				cosmosNFT, ok := nftIDMatcher.FindCosmosNFTBySerial(cosmosNFTs.NFTs, nftId)
				metadataStr := "{}"
				memo := ""
				if ok {
					metadataOverride, err := m.QueryNFTExternalMetadata(cosmosNFT)
					if err != nil {
						return nil, doMintNFTActionFailed(db, a, err)
					}
					metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
						erc721ExternalURLBuilder,
						cosmosNFT,
						cosmosClass.Class,
						iscnDataResponse,
						metadataOverride,
						a.EvmClassId,
						nftId,
					))
					if err != nil {
						return nil, doMintNFTActionFailed(db, a, err)
					}
					metadataStr = string(metadataBytes)

					events, err := m.QueryAllNFTEvents(m.MakeQueryNFTEventsRequest(cosmosNFT.ClassId, cosmosNFT.Id))
					if err != nil {
						return nil, doMintNFTActionFailed(db, a, err)
					}
					memo = event.MakeMemoFromEvent(events)
				}
				tos = append(tos, initialBatchMintOwnerAddress)
				memos = append(memos, memo)
				metadataList = append(metadataList, metadataStr)
			}
			_, _, err = c.MintNFTs(
				ctx,
				mylogger,
				evmClassAddress,
				totalSupplyBigInt,
				tos,
				memos,
				metadataList,
			)
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		}
		tx, _, err = p.TransferNFT(
			ctx,
			mylogger,
			evmClassAddress,
			initialBatchMintOwnerAddress,
			toOwner,
			big.NewInt(int64(nftId)))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
	}

	evmTxHash := hexutil.Encode(tx.Hash().Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeNFTMigrationActionMintNFTStatusCompleted
	err = appdb.UpdateLikeNFTMigrationActionMintNFT(db, a)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	return a, nil
}

func doMintNFTActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionMintNFT, err error) error {
	a.Status = model.LikeNFTMigrationActionMintNFTStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionMintNFT(db, a))
}
