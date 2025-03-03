package likenft

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func DoTransferClassAction(
	logger *slog.Logger,

	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	n *evm.LikeNFTClass,
	a *model.LikeNFTMigrationActionTransferClass,
) (*model.LikeNFTMigrationActionTransferClass, error) {
	mylogger := logger.
		WithGroup("DoTransferClassAction").
		With("transferClassActionId", a.Id)

	if a.Status == model.LikeNFTMigrationActionTransferClassStatusCompleted {
		return a, nil
	}
	if a.Status != model.LikeNFTMigrationActionTransferClassStatusInit &&
		a.Status != model.LikeNFTMigrationActionTransferClassStatusFailed {
		return nil, errors.New("error transfer class action is not init or failed")
	}

	a.Status = model.LikeNFTMigrationActionTransferClassStatusInProgress
	err := appdb.UpdateLikeNFTMigrationActionTransferClass(db, a)
	if err != nil {
		return nil, fmt.Errorf(": %v", err)
	}

	// TODO check class owner on cosmos

	newOwnerAddress := common.HexToAddress(a.EvmOwner)
	ethClassAddress := common.HexToAddress(a.EvmClassId)
	tx, err := n.TransferClass(ethClassAddress, newOwnerAddress)

	if err != nil {
		return nil, doTransferClassActionFailed(db, a, err)
	}

	_, err = ethereum.AwaitTx(mylogger, n.Client, tx)

	if err != nil {
		return nil, doTransferClassActionFailed(db, a, err)
	}

	evmTxHash := hexutil.Encode(tx.Hash().Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeNFTMigrationActionTransferClassStatusCompleted

	err = appdb.UpdateLikeNFTMigrationActionTransferClass(db, a)
	if err != nil {
		return nil, doTransferClassActionFailed(db, a, err)
	}
	return a, nil
}

func doTransferClassActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionTransferClass, err error) error {
	a.Status = model.LikeNFTMigrationActionTransferClassStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionTransferClass(db, a))
}
