package database

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/ent/transactionmemo"
)

type TransactionMemoRepository interface {
	InsertTransactionMemo(ctx context.Context, m *ent.TransactionMemo) error
}

type transactionMemoRepository struct {
	dbService Service
}

func MakeTransactionMemoRepository(
	dbService Service,
) TransactionMemoRepository {
	return &transactionMemoRepository{
		dbService: dbService,
	}
}

func (r *transactionMemoRepository) InsertTransactionMemo(ctx context.Context, m *ent.TransactionMemo) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.TransactionMemo.Query().Where(
			transactionmemo.TransactionHashEqualFold(m.TransactionHash),
			transactionmemo.BookNftIDEqualFold(m.BookNftID),
			transactionmemo.TokenIDEQ(m.TokenID),
		).Only(ctx)

		if err == nil {
			// Should already be inserted
			return nil
		}

		if !ent.IsNotFound(err) {
			return err
		}

		return tx.TransactionMemo.Create().
			SetBlockNumber(m.BlockNumber).
			SetBookNftID(m.BookNftID).
			SetFrom(m.From).
			SetMemo(m.Memo).
			SetTo(m.To).
			SetTokenID(m.TokenID).
			SetTransactionHash(m.TransactionHash).Exec(ctx)
	})
}
