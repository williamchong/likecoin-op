package logic

import (
	"context"
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	appdb "github.com/likecoin/likecoin-migration-backend/pkg/db"
	"github.com/likecoin/likecoin-migration-backend/pkg/model"
)

type MigrationRecordStatus string

const (
	MigrationRecordStatusUnknown MigrationRecordStatus = "unknown"
	MigrationRecordStatusPending MigrationRecordStatus = "pending"
	MigrationRecordStatusSuccess MigrationRecordStatus = "success"
	MigrationRecordStatusFailed  MigrationRecordStatus = "failed"
)

func GetLikeCoinMigrationRecord(
	db *sql.DB,
	ethClient *ethclient.Client,
	cosmosTxHash string,
) (*model.MigrationRecord, MigrationRecordStatus, error) {
	migrationRecord, err := appdb.GetMigrationRecordByCosmosTxHash(db, cosmosTxHash)
	if err != nil {
		return nil, MigrationRecordStatusUnknown, err
	}

	_, isPending, err := ethClient.TransactionByHash(
		context.Background(),
		common.HexToHash(migrationRecord.EthTxHash),
	)
	if err != nil {
		return nil, MigrationRecordStatusUnknown, err
	}

	if isPending {
		return migrationRecord, MigrationRecordStatusPending, err
	}

	r, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(migrationRecord.EthTxHash))
	if err != nil {
		return nil, MigrationRecordStatusUnknown, err
	}

	var status MigrationRecordStatus
	if r.Status == 1 {
		status = MigrationRecordStatusSuccess
	} else {
		status = MigrationRecordStatusFailed
	}
	return migrationRecord, status, err
}
