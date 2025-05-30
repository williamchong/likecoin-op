package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

type LikeNFTMigrationActionBatchMintNFTsFromCosmosRepository interface {
	Create(
		ctx context.Context,
		evmClassId string,
		currentSupply uint64,
		expectedSupply uint64,
		mintAmount uint64,
		initialBatchMintOwner string,
		status model.LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus,
	) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error)

	Success(
		ctx context.Context,
		id uint64,
		fromID uint64,
		toID uint64,
		evmTxHash string,
	) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error)

	Failed(
		ctx context.Context,
		id uint64,
		err error,
	) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error)
}

type likeNFTMigrationActionBatchMintNFTsFromCosmosRepository struct {
	db *sql.DB
}

func MakeLikeNFTMigrationActionBatchMintNFTsFromCosmosRepository(
	db *sql.DB,
) LikeNFTMigrationActionBatchMintNFTsFromCosmosRepository {
	return &likeNFTMigrationActionBatchMintNFTsFromCosmosRepository{
		db: db,
	}
}

func (r *likeNFTMigrationActionBatchMintNFTsFromCosmosRepository) Create(
	ctx context.Context,
	evmClassId string,
	currentSupply uint64,
	expectedSupply uint64,
	mintAmount uint64,
	initialBatchMintOwner string,
	status model.LikeNFTMigrationActionBatchMintNFTsFromCosmosStatus,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	mChan := make(chan *model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, 1)
	err := WithTx(ctx, r.db, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx,
			`INSERT INTO likenft_migration_action_batch_mint_nfts (
	evm_class_id,
	current_supply,
	expected_supply,
	batch_mint_amount,
	initial_batch_mint_owner,
	status
) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING
		id,
		created_at,
		evm_class_id,
		current_supply,
		expected_supply,
		batch_mint_amount,
		initial_batch_mint_owner,
		status,
		from_id,
		to_id,
		evm_tx_hash,
		failed_reason
`,
			evmClassId,
			currentSupply,
			expectedSupply,
			mintAmount,
			initialBatchMintOwner,
			status,
		)

		m := &model.LikeNFTMigrationActionBatchMintNFTsFromCosmos{}

		err := row.Scan(
			&m.Id,
			&m.CreatedAt,
			&m.EvmClassId,
			&m.CurrentSupply,
			&m.ExpectedSupply,
			&m.BatchMintAmount,
			&m.InitialBatchMintOwner,
			&m.Status,
			&m.FromID,
			&m.ToID,
			&m.EvmTxHash,
			&m.FailedReason,
		)

		if err != nil {
			return err
		}

		mChan <- m
		return nil
	})

	if err != nil {
		close(mChan)
		return nil, err
	}

	return <-mChan, nil
}

func (r *likeNFTMigrationActionBatchMintNFTsFromCosmosRepository) Success(
	ctx context.Context,

	id uint64,
	fromID uint64,
	toID uint64,
	evmTxHash string,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	mChan := make(chan *model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, 1)
	err := WithTx(ctx, r.db, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx,
			`UPDATE likenft_migration_action_batch_mint_nfts SET
	status = $2,
	from_id = $3,
	to_id = $4,
	evm_tx_hash = $5
WHERE id = $1
RETURNING
	id,
	created_at,
	evm_class_id,
	current_supply,
	expected_supply,
	batch_mint_amount,
	initial_batch_mint_owner,
	status,
	from_id,
	to_id,
	evm_tx_hash,
	failed_reason
`,
			id,
			model.LikeNFTMigrationActionBatchMintNFTsFromCosmosStatusCompleted,
			fromID,
			toID,
			evmTxHash,
		)

		m := &model.LikeNFTMigrationActionBatchMintNFTsFromCosmos{}

		err := row.Scan(
			&m.Id,
			&m.CreatedAt,
			&m.EvmClassId,
			&m.CurrentSupply,
			&m.ExpectedSupply,
			&m.BatchMintAmount,
			&m.InitialBatchMintOwner,
			&m.Status,
			&m.FromID,
			&m.ToID,
			&m.EvmTxHash,
			&m.FailedReason,
		)

		if err != nil {
			return err
		}

		mChan <- m
		return nil
	})
	if err != nil {
		close(mChan)
		return nil, err
	}
	return <-mChan, nil
}

func (r *likeNFTMigrationActionBatchMintNFTsFromCosmosRepository) Failed(
	ctx context.Context,
	id uint64,
	err error,
) (*model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, error) {
	mChan := make(chan *model.LikeNFTMigrationActionBatchMintNFTsFromCosmos, 1)
	txErr := WithTx(ctx, r.db, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx,
			`UPDATE likenft_migration_action_batch_mint_nfts SET
	status = $2,
	failed_reason = $3
WHERE id = $1
RETURNING
	id,
	created_at,
	evm_class_id,
	current_supply,
	expected_supply,
	batch_mint_amount,
	initial_batch_mint_owner,
	status,
	from_id,
	to_id,
	evm_tx_hash,
	failed_reason
`,
			id,
			model.LikeCoinMigrationStatusFailed,
			err.Error(),
		)

		m := &model.LikeNFTMigrationActionBatchMintNFTsFromCosmos{}

		err := row.Scan(
			&m.Id,
			&m.CreatedAt,
			&m.EvmClassId,
			&m.CurrentSupply,
			&m.ExpectedSupply,
			&m.BatchMintAmount,
			&m.InitialBatchMintOwner,
			&m.Status,
			&m.FromID,
			&m.ToID,
			&m.EvmTxHash,
			&m.FailedReason,
		)

		if err != nil {
			return err
		}
		mChan <- m
		return nil
	})

	if txErr != nil {
		close(mChan)
		return nil, errors.Join(err, txErr)
	}

	return <-mChan, err
}
