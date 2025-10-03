package likenft

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/likenftindexer"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func DoNewClassAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
	likenftIndexer likenftindexer.LikeNFTIndexerClient,
	n *evm.LikeProtocol,
	a *model.LikeNFTMigrationActionNewClass,
) (*model.LikeNFTMigrationActionNewClass, error) {
	mylogger := logger.
		WithGroup("DoNewClassAction").
		With("newClassActionId", a.Id)

	if a.Status == model.LikeNFTMigrationActionNewClassStatusCompleted {
		return a, nil
	}
	if a.Status != model.LikeNFTMigrationActionNewClassStatusInit &&
		a.Status != model.LikeNFTMigrationActionNewClassStatusFailed {
		return nil, errors.New("error new class action is not init or failed")
	}

	a.Status = model.LikeNFTMigrationActionNewClassStatusInProgress
	err := appdb.UpdateLikeNFTMigrationActionNewClass(db, a)
	if err != nil {
		return nil, fmt.Errorf(": %v", err)
	}
	cosmosClassResponse, err := c.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
		ClassId: a.CosmosClassId,
	})
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	cosmosClass := cosmosClassResponse.Class

	iscnDataResponse, err := c.GetISCNRecord(
		cosmosClass.Data.Parent.IscnIdPrefix,
		cosmosClass.Data.Parent.IscnVersionAtMint,
	)

	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	royaltyConfigsResponse, err := c.QueryRoyaltyConfigsByClassId(cosmos.QueryRoyaltyConfigsByClassIdRequest{
		ClassId: cosmosClass.Id,
	})

	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	var royaltyConfig *cosmosmodel.RoyaltyConfig = nil
	if royaltyConfigsResponse != nil {
		royaltyConfig = royaltyConfigsResponse.RoyaltyConfig
	}

	initialMinterAddresses := make([]common.Address, len(a.InitialMintersStr.ToSlice()))
	for i, initialMinterStr := range a.InitialMintersStr.ToSlice() {
		initialMinterAddresses[i] = common.HexToAddress(initialMinterStr)
	}
	initialUpdaterAddresses := make([]common.Address, len(a.InitialUpdatersStr.ToSlice()))
	for i, initialMinterStr := range a.InitialUpdatersStr.ToSlice() {
		initialUpdaterAddresses[i] = common.HexToAddress(initialMinterStr)
	}

	initialOwnerAddress := common.HexToAddress(a.InitialOwner)

	maxSupply := ^uint64(0)
	if cosmosClass.Data.Config.MaxSupply != "" {
		maxSupply, err = strconv.ParseUint(cosmosClass.Data.Config.MaxSupply, 10, 64)
		if err != nil {
			return nil, doNewClassActionFailed(db, a, err)
		}
		if maxSupply == 0 {
			maxSupply = ^uint64(0)
		}
	}

	contractLevelMetadata := evm.ContractLevelMetadataFromCosmosClassAndISCN(
		cosmosClass,
		iscnDataResponse,
		royaltyConfig,
	)

	metadataBytes, err := json.Marshal(contractLevelMetadata)
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	msgSenderStr, err := n.Signer.GetSignerAddress()
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}
	msgSender := common.HexToAddress(*msgSenderStr)

	var saltData string = ""
	if contractLevelMetadata.PotentialAction != nil &&
		len(contractLevelMetadata.PotentialAction.Target) > 0 {
		saltData = contractLevelMetadata.PotentialAction.Target[0].Url
	}

	salt, err := evm.ComputeSaltDataFromCandidates(msgSender, [2]byte{0, 0}, saltData)
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	tx, txReceipt, err := n.NewBookNFT(
		ctx,
		mylogger,
		salt,
		like_protocol.MsgNewBookNFT{
			Creator:  initialOwnerAddress,
			Updaters: initialUpdaterAddresses,
			Minters:  initialMinterAddresses,
			Config: like_protocol.BookConfig{
				Name:      cosmosClass.Name,
				Symbol:    cosmosClass.Symbol,
				Metadata:  string(metadataBytes),
				MaxSupply: maxSupply,
			},
		}, a.DefaultRoyaltyFraction)

	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	newClassId, err := n.GetClassIdFromNewClassTransaction(txReceipt)

	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	evmClassId := hexutil.Encode(newClassId.Bytes())
	a.EvmClassId = &evmClassId
	evmTxHash := hexutil.Encode(tx.Hash().Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeNFTMigrationActionNewClassStatusCompleted
	err = appdb.UpdateLikeNFTMigrationActionNewClass(db, a)
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	_, err = likenftIndexer.IndexLikeProtocol(ctx)
	if err != nil {
		mylogger.Error("failed to index like protocol", "error", err)
	}

	go SubmitEvmBookMigrated(
		context.Background(),
		mylogger,
		db,
		likecoinAPI,
		a.CosmosClassId,
		evmClassId,
	)

	return a, nil
}

func doNewClassActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionNewClass, err error) error {
	a.Status = model.LikeNFTMigrationActionNewClassStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionNewClass(db, a))
}
