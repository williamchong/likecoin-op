package db

import (
	"github.com/likecoin/likecoin-migration-backend/pkg/model"
)

func GetMigrationRecordByCosmosTxHash(
	tx TxLike,
	cosmosTxHash string,
) (*model.MigrationRecord, error) {
	row := tx.QueryRow(
		`SELECT cosmos_tx_hash, eth_tx_hash, cosmos_address, eth_address FROM migration_record WHERE cosmos_tx_hash = $1`,
		cosmosTxHash,
	)

	migrationRecord := &model.MigrationRecord{}

	err := row.Scan(
		&migrationRecord.CosmosTxHash,
		&migrationRecord.EthTxHash,
		&migrationRecord.CosmosAddress,
		&migrationRecord.EthAddress,
	)

	if err != nil {
		return nil, err
	}

	return migrationRecord, nil
}

func InsertMigrationRecord(
	tx TxLike,
	migrationRecord *model.MigrationRecord,
) error {
	_, err := tx.Exec(
		"INSERT INTO migration_record (cosmos_tx_hash, eth_tx_hash, cosmos_address, eth_address) VALUES ($1, $2, $3, $4)",
		migrationRecord.CosmosTxHash,
		migrationRecord.EthTxHash,
		migrationRecord.CosmosAddress,
		migrationRecord.EthAddress,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateMigrationRecord(
	tx TxLike,
	migrationRecord *model.MigrationRecord,
) error {
	_, err := tx.Exec(
		"UPDATE migration_record SET eth_tx_hash=$2, cosmos_address=$3, eth_address=$4 where cosmos_tx_hash=$1",
		migrationRecord.CosmosTxHash,
		migrationRecord.EthTxHash,
		migrationRecord.CosmosAddress,
		migrationRecord.EthAddress,
	)

	if err != nil {
		return err
	}

	return nil
}
