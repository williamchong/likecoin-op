package database

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmeventprocessedblockheight"
)

type EVMEventProcessedBlockHeightRepository interface {
	GetEVMEventProcessedBlockHeight(
		ctx context.Context,

		contractType evmeventprocessedblockheight.ContractType,
		contractAddress string,
		event evmeventprocessedblockheight.Event,
	) (uint64, error)

	UpdateEVMEventProcessedBlockHeight(
		ctx context.Context,

		contractType evmeventprocessedblockheight.ContractType,
		contractAddress string,
		event evmeventprocessedblockheight.Event,
		blockHeight uint64,
	) error
}

type evmEventProcessedBlockHeightRepository struct {
	dbService Service
}

var _ EVMEventProcessedBlockHeightRepository = &evmEventProcessedBlockHeightRepository{}

func MakeEVMEventProcessedBlockHeightRepository(
	dbService Service,
) *evmEventProcessedBlockHeightRepository {
	return &evmEventProcessedBlockHeightRepository{
		dbService: dbService,
	}
}

func (s *evmEventProcessedBlockHeightRepository) GetEVMEventProcessedBlockHeight(
	ctx context.Context,

	contractType evmeventprocessedblockheight.ContractType,
	contractAddress string,
	event evmeventprocessedblockheight.Event,
) (uint64, error) {
	processedBlockHeight, err := s.dbService.Client().EVMEventProcessedBlockHeight.
		Query().
		Where(
			evmeventprocessedblockheight.ContractAddressEQ(contractAddress),
			evmeventprocessedblockheight.EventEQ(event),
			evmeventprocessedblockheight.ContractTypeEQ(contractType),
		).
		Only(ctx)

	if err != nil {
		return 0, err
	}

	return processedBlockHeight.BlockHeight, nil
}

func (s *evmEventProcessedBlockHeightRepository) UpdateEVMEventProcessedBlockHeight(
	ctx context.Context,

	contractType evmeventprocessedblockheight.ContractType,
	contractAddress string,
	event evmeventprocessedblockheight.Event,
	blockHeight uint64,
) error {
	return WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		processedBlockHeight, err := tx.EVMEventProcessedBlockHeight.
			Query().
			Where(
				evmeventprocessedblockheight.ContractAddressEQ(contractAddress),
				evmeventprocessedblockheight.EventEQ(event),
				evmeventprocessedblockheight.ContractTypeEQ(contractType),
			).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				return tx.EVMEventProcessedBlockHeight.
					Create().
					SetContractType(contractType).
					SetContractAddress(contractAddress).
					SetEvent(event).
					SetBlockHeight(blockHeight).
					Exec(ctx)
			}
			return err
		}

		return processedBlockHeight.
			Update().
			SetContractAddress(contractAddress).
			SetEvent(event).
			SetBlockHeight(blockHeight).
			Exec(ctx)
	})
}
