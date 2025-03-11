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
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func DoNewClassAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
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
	initialOwnerAddress := common.HexToAddress(a.InitialOwner)
	initialMinterAddress := common.HexToAddress(a.InitialMinter)
	initialUpdaterAddress := common.HexToAddress(a.InitialUpdater)

	maxSupply := uint64(0)
	if cosmosClass.Data.Config.MaxSupply != "" {
		maxSupply, err = strconv.ParseUint(cosmosClass.Data.Config.MaxSupply, 10, 64)
		if err != nil {
			return nil, doNewClassActionFailed(db, a, err)
		}
	}

	metadataBytes, err := json.Marshal(evm.ContractLevelMetadataFromCosmosClass(cosmosClass))
	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}
	tx, txReceipt, err := n.NewBookNFT(ctx, mylogger, like_protocol.MsgNewBookNFT{
		Creator:  initialOwnerAddress,
		Updaters: []common.Address{initialUpdaterAddress},
		Minters:  []common.Address{initialMinterAddress},
		Config: like_protocol.BookConfig{
			Name:      cosmosClass.Name,
			Symbol:    cosmosClass.Symbol,
			Metadata:  string(metadataBytes),
			MaxSupply: maxSupply,
		},
	})

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
	return a, nil
}

func doNewClassActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionNewClass, err error) error {
	a.Status = model.LikeNFTMigrationActionNewClassStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionNewClass(db, a))
}
