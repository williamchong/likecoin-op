package db

import (
	"github.com/likecoin/like-migration-backend/pkg/model"
)

func InsertLikecoinEvmMigrateBookSubmission(
	tx TxLike,
	submission *model.LikecoinEvmMigrateBookSubmission,
) (*model.LikecoinEvmMigrateBookSubmission, error) {
	row := tx.QueryRow(
		`INSERT INTO likecoin_evm_migrate_book_submission (
			like_class_id,
			evm_class_id,
			status,
			failed_reason
	) VALUES ($1, $2, $3, $4)
	RETURNING id`,
		submission.LikeClassID,
		submission.EvmClassID,
		submission.Status,
		submission.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	submission.Id = uint64(lastInsertId)

	return submission, nil
}
