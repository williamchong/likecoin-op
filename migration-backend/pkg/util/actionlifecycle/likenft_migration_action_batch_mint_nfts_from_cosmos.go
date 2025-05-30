package actionlifecycle

import (
	"context"

	"github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp struct {
	FromID    uint64
	ToID      uint64
	EvmTxHash string
}

type LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle struct {
	batchmintNFTsFromCosmosRepository db.LikeNFTMigrationActionBatchMintNFTsFromCosmosRepository

	evmClassId            string
	currentSupply         uint64
	expectedSupply        uint64
	batchMintSize         uint64
	initialBatchMintOwner string

	model *model.LikeNFTMigrationActionBatchMintNFTsFromCosmos
}

func MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle(
	batchmintNFTsFromCosmosRepository db.LikeNFTMigrationActionBatchMintNFTsFromCosmosRepository,
	evmClassId string,
	currentSupply uint64,
	expectedSupply uint64,
	batchMintSize uint64,
	initialBatchMintOwner string,
) ActionLifecycle[LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp, model.LikeNFTMigrationActionBatchMintNFTsFromCosmos] {
	return &LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle{
		batchmintNFTsFromCosmosRepository: batchmintNFTsFromCosmosRepository,
		evmClassId:                        evmClassId,
		currentSupply:                     currentSupply,
		expectedSupply:                    expectedSupply,
		batchMintSize:                     batchMintSize,
		initialBatchMintOwner:             initialBatchMintOwner,
	}
}

func (l *LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle) Begin(
	ctx context.Context,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	m, err := l.batchmintNFTsFromCosmosRepository.Create(
		ctx,
		l.evmClassId,
		l.currentSupply,
		l.expectedSupply,
		l.batchMintSize,
		l.initialBatchMintOwner,
		model.LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusInProgress,
	)
	if err != nil {
		return nil, err
	}
	l.model = m
	return m, nil
}

func (l *LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle) Success(
	ctx context.Context,
	sucResp *LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycleSucResp,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	m, err := l.batchmintNFTsFromCosmosRepository.Success(
		ctx,
		l.model.Id,
		sucResp.FromID,
		sucResp.ToID,
		sucResp.EvmTxHash,
	)
	if err != nil {
		return nil, err
	}
	l.model = m
	return m, nil
}

func (l *LikeNFTMigrationActionBatchMintNFTsFromCosmosActionLifecycle) Failed(
	ctx context.Context,
	err error,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	m, err := l.batchmintNFTsFromCosmosRepository.Failed(ctx, l.model.Id, err)
	if err != nil {
		return nil, err
	}
	l.model = m
	return m, nil
}
