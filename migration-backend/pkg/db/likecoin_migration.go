package db

import (
	"database/sql"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryNonEndedLikeCoinMigration(
	tx TxLike,
	cosmosAddress string,
) (*model.LikeCoinMigration, error) {
	row := tx.QueryRow(`SELECT
    id,
    created_at,
    user_cosmos_address,
    burning_cosmos_address,
    minting_eth_address,
    user_eth_address,
    amount,
    evm_signature,
    evm_signature_message,
    status,
    cosmos_tx_hash,
    evm_tx_hash,
    failed_reason
FROM likecoin_migration
WHERE user_cosmos_address = $1 AND status NOT IN ($2, $3)
`,
		cosmosAddress,
		model.LikeCoinMigrationStatusCompleted,
		model.LikeCoinMigrationStatusFailed)

	m := &model.LikeCoinMigration{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.UserCosmosAddress,
		&m.BurningCosmosAddress,
		&m.MintingEthAddress,
		&m.UserEthAddress,
		&m.Amount,
		&m.EvmSignature,
		&m.EvmSignatureMessage,
		&m.Status,
		&m.CosmosTxHash,
		&m.EvmTxHash,
		&m.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func QueryPaginatedLikeCoinMigration(
	tx TxLike,
	limit int,
	offset int,
	status *model.LikeCoinMigrationStatus,
	keyword string,
) ([]*model.LikeCoinMigration, error) {
	rows, err := tx.Query(`SELECT
    id,
    created_at,
    user_cosmos_address,
    burning_cosmos_address,
    minting_eth_address,
    user_eth_address,
    amount,
    evm_signature,
    evm_signature_message,
    status,
    cosmos_tx_hash,
    evm_tx_hash,
    failed_reason
FROM likecoin_migration
WHERE ($3::text IS NULL OR status = $3) AND
(
	$4::text = '' OR 
	failed_reason ILIKE '%' || $4 || '%' OR 
	user_cosmos_address ILIKE '%' || $4 || '%' OR 
	burning_cosmos_address ILIKE '%' || $4 || '%' OR 
	minting_eth_address ILIKE '%' || $4 || '%' OR 
	user_eth_address ILIKE '%' || $4 || '%' OR
	evm_signature ILIKE '%' || $4 || '%' OR 
	evm_signature_message ILIKE '%' || $4 || '%' OR 
	cosmos_tx_hash ILIKE '%' || $4 || '%' OR 
	evm_tx_hash ILIKE '%' || $4 || '%'
)
ORDER BY created_at DESC
LIMIT $1
OFFSET $2

`, limit, offset, status, keyword)

	if err != nil {
		return nil, err
	}

	migrations := []*model.LikeCoinMigration{}

	for rows.Next() {
		m := &model.LikeCoinMigration{}
		err := rows.Scan(
			&m.Id,
			&m.CreatedAt,
			&m.UserCosmosAddress,
			&m.BurningCosmosAddress,
			&m.MintingEthAddress,
			&m.UserEthAddress,
			&m.Amount,
			&m.EvmSignature,
			&m.EvmSignatureMessage,
			&m.Status,
			&m.CosmosTxHash,
			&m.EvmTxHash,
			&m.FailedReason,
		)

		if err != nil {
			return nil, err
		}

		migrations = append(migrations, m)
	}

	return migrations, nil
}

func QueryLatestLikeCoinMigration(
	tx TxLike,
	cosmosAddress string,
) (*model.LikeCoinMigration, error) {
	row := tx.QueryRow(`SELECT
    id,
    created_at,
    user_cosmos_address,
    burning_cosmos_address,
    minting_eth_address,
    user_eth_address,
    amount,
    evm_signature,
    evm_signature_message,
    status,
    cosmos_tx_hash,
    evm_tx_hash,
    failed_reason
FROM likecoin_migration
WHERE user_cosmos_address = $1
ORDER BY created_at DESC
LIMIT 1
`, cosmosAddress)

	m := &model.LikeCoinMigration{}

	err := row.Scan(
		&m.Id,
		&m.CreatedAt,
		&m.UserCosmosAddress,
		&m.BurningCosmosAddress,
		&m.MintingEthAddress,
		&m.UserEthAddress,
		&m.Amount,
		&m.EvmSignature,
		&m.EvmSignatureMessage,
		&m.Status,
		&m.CosmosTxHash,
		&m.EvmTxHash,
		&m.FailedReason,
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func InsertLikeCoinMigration(
	tx TxLike,
	m *model.LikeCoinMigration,
) (*model.LikeCoinMigration, error) {
	row := tx.QueryRow(
		`INSERT INTO likecoin_migration (
    user_cosmos_address,
    burning_cosmos_address,
    minting_eth_address,
    user_eth_address,
    amount,
    evm_signature,
    evm_signature_message,
    status,
    cosmos_tx_hash,
    evm_tx_hash,
    failed_reason
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id`,
		m.UserCosmosAddress,
		m.BurningCosmosAddress,
		m.MintingEthAddress,
		m.UserEthAddress,
		m.Amount,
		m.EvmSignature,
		m.EvmSignatureMessage,
		m.Status,
		m.CosmosTxHash,
		m.EvmTxHash,
		m.FailedReason,
	)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	m.Id = uint64(lastInsertId)

	return m, nil
}

func UpdateLikeCoinMigration(
	tx TxLike,
	m *model.LikeCoinMigration,
) error {
	result, err := tx.Exec(
		`UPDATE likecoin_migration SET
    user_cosmos_address = $1,
    burning_cosmos_address = $2,
    minting_eth_address = $3,
    user_eth_address = $4,
    amount = $5,
    evm_signature = $6,
    evm_signature_message = $7,
    status = $8,
    cosmos_tx_hash = $9,
    evm_tx_hash = $10,
    failed_reason = $11
WHERE id = $12`,
		m.UserCosmosAddress,
		m.BurningCosmosAddress,
		m.MintingEthAddress,
		m.UserEthAddress,
		m.Amount,
		m.EvmSignature,
		m.EvmSignatureMessage,
		m.Status,
		m.CosmosTxHash,
		m.EvmTxHash,
		m.FailedReason,
		m.Id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
