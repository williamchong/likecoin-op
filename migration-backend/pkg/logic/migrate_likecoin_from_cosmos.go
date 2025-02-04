package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

type Memo struct {
	Signature  string `json:"signature"`
	EthAddress string `json:"ethAddress"`
	Amount     string `json:"amount"`
	Denom      string `json:"denom"`
}

func getMessage(amount string, denom string) string {
	return fmt.Sprintf(`You are going to deposit %v %v to migration program.

This sign make sure the address is correct.`, amount, denom)
}

func MigrateLikeCoinFromCosmos(
	db *sql.DB,
	ethClient *ethclient.Client,
	cosmosAPI *api.CosmosAPI,
	ethNetworkPublicRPCURL string,
	ethWalletPrivateKey string,
	ethTokenAddress string,
	cosmosTxHash string,
) (*types.Transaction, error) {
	txResponse, err := cosmosAPI.QueryTransaction(cosmosTxHash)
	if err != nil {
		return nil, err
	}

	migrationRecord, err := appdb.GetMigrationRecordByCosmosTxHash(db, cosmosTxHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			memoString := txResponse.Tx.Body.Memo
			var memo Memo
			err = json.Unmarshal([]byte(memoString), &memo)
			if err != nil {
				return nil, err
			}
			message := getMessage(memo.Amount, memo.Denom)

			recoveredAddr, err := ethereum.RecoverAddress(memo.Signature, []byte(message))
			if err != nil {
				return nil, err
			}

			migrationRecord = &model.MigrationRecord{
				CosmosTxHash: cosmosTxHash,
				EthAddress:   recoveredAddr.Hex(),
			}

			// TODO: Use worker
			tx, err := ethereum.TransferToken(
				ethNetworkPublicRPCURL,
				ethWalletPrivateKey,
				*recoveredAddr,
				common.HexToAddress(ethTokenAddress),
				memo.Amount)
			if err != nil {
				return nil, err
			}

			migrationRecord.EthTxHash = tx.Hash().Hex()

			err = appdb.InsertMigrationRecord(db, migrationRecord)

			if err != nil {
				return nil, err
			}

			return tx, nil
		}

		return nil, err
	}

	transaction, _, err := ethClient.TransactionByHash(
		context.Background(),
		common.HexToHash(migrationRecord.EthTxHash),
	)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
