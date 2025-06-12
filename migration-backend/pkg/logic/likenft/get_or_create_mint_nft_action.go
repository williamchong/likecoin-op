package likenft

import (
	"database/sql"
	"errors"
	"fmt"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type ErrNFTAlreadyMinted struct {
	mintedEvmOwner string
	txHash         string
}

func MakeErrAlreadyMinted(
	mintedEvmOwner string,
	txHash string,
) *ErrNFTAlreadyMinted {
	return &ErrNFTAlreadyMinted{
		mintedEvmOwner,
		txHash,
	}
}

func (e *ErrNFTAlreadyMinted) Error() string {
	return fmt.Sprintf(
		"err already minted, mintedEvmOwner:%s, txHash:%s",
		e.mintedEvmOwner,
		e.txHash,
	)
}

func GetOrCreateMintNFTAction(
	db *sql.DB,
	evmClassId string,
	cosmosNFTId string,
	initialBatchMintOwner string,
	evmOwner string,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	m, err := appdb.QueryLikeNFTMigrationActionMintNFTByEvmClassIDAndCosmosNFTID(
		db,
		evmClassId,
		cosmosNFTId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			m = &model.LikeNFTMigrationActionMintNFT{
				EvmClassId:            evmClassId,
				CosmosNFTId:           cosmosNFTId,
				EvmOwner:              evmOwner,
				Status:                model.LikeNFTMigrationActionMintNFTStatusInit,
				InitialBatchMintOwner: initialBatchMintOwner,
			}
			m, err = appdb.InsertLikeNFTMigrationActionMintNFT(db, m)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	if m.Status == model.LikeNFTMigrationActionMintNFTStatusCompleted && m.EvmOwner != evmOwner {
		return nil, MakeErrAlreadyMinted(
			m.EvmOwner,
			*m.EvmTxHash,
		)
	}
	return m, nil
}
