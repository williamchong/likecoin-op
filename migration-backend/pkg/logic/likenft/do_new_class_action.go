package likenft

import (
	"database/sql"
	"errors"
	"fmt"
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
	db *sql.DB,
	c *cosmos.LikeNFTCosmosClient,
	n *evm.LikeProtocol,
	a *model.LikeNFTMigrationActionNewClass,
) (*model.LikeNFTMigrationActionNewClass, error) {
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

	newClassId, txHash, err := n.NewClass(like_protocol.MsgNewClass{
		Creator:  initialOwnerAddress,
		Updaters: []common.Address{initialUpdaterAddress},
		Minters:  []common.Address{initialMinterAddress},
		Input: like_protocol.ClassInput{
			Name:     cosmosClass.Name,
			Symbol:   cosmosClass.Symbol,
			Metadata: "{}", // TODO
			Config: like_protocol.ClassConfig{
				MaxSupply: maxSupply,
			},
		},
	})

	if err != nil {
		return nil, doNewClassActionFailed(db, a, err)
	}

	evmClassId := hexutil.Encode(newClassId.Bytes())
	a.EvmClassId = &evmClassId
	evmTxHash := hexutil.Encode(txHash.Bytes())
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
