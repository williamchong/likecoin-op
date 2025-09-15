package likenft

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
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
	"github.com/likecoin/like-migration-backend/pkg/likenftindexer"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func DoMintNFTAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	p *evm.LikeProtocol,
	c *evm.BookNFT,
	m *cosmos.LikeNFTCosmosClient,
	likenftIndexer likenftindexer.LikeNFTIndexerClient,
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
		metadataOverride, err := cosmos.ProcessQueryNFTExternalMetadataErrors(
			m.QueryNFTExternalMetadata(cosmosNFT.NFT),
		)
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
		metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNDataArbitrary(
			erc721ExternalURLBuilder,
			cosmosNFT.NFT,
			cosmosClass.Class,
			iscnDataResponse,
			metadataOverride,
			a.EvmClassId,
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
		// serial nft id
		// Assume the desire supply is prepared *by initial batch mint owner (signer)* before minting the nftid
		// Otherwise the call will return no token error
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
	_, err = likenftIndexer.IndexBookNFT(ctx, a.EvmClassId)
	if err != nil {
		mylogger.Error("failed to index book nft", "error", err)
	}
	return a, nil
}

func doMintNFTActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionMintNFT, err error) error {
	a.Status = model.LikeNFTMigrationActionMintNFTStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionMintNFT(db, a))
}
