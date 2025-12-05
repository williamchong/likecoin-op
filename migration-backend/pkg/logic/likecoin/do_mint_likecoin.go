package likecoin

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/util"
	"github.com/likecoin/like-migration-backend/pkg/model"
	cosmosModel "github.com/likecoin/like-migration-backend/pkg/model/cosmos"
	"github.com/shopspring/decimal"
)

var ErrAlreadyInProgress = fmt.Errorf("err already in progress")
var ErrEthAddressOnCosmosNotMatch = fmt.Errorf("err eth address on cosmos not match")
var ErrMintingEthPrivateKeyNotMatchMintingEthAddress = fmt.Errorf("err minting eth private key not match minter eth address")
var ErrAmountNotMatch = fmt.Errorf("err amount not match")
var ErrNoMessagesInTransaction = fmt.Errorf("err no messages found in transaction")
var ErrNoMsgSendInTransaction = fmt.Errorf("err no MsgSend found in transaction")
var ErrNoAmountInMsgSend = fmt.Errorf("err no amount found in MsgSend")
var ErrTransactionRecipientMismatch = fmt.Errorf("err transaction recipient does not match burning address")
var ErrTransactionSenderMismatch = fmt.Errorf("err transaction sender does not match user address")
var ErrTransactionAmountMismatch = fmt.Errorf("err transaction amount does not match expected amount")

func verifyCosmosTransaction(
	logger *slog.Logger,
	txResponse *cosmosModel.TxResponse,
	expectedFromAddress string,
	expectedToAddress string,
	expectedAmount string,
	memo *MemoData,
) (types.Coin, error) {
	if len(txResponse.Tx.Body.Messages) == 0 {
		logger.Error("no messages in transaction", "err", ErrNoMessagesInTransaction)
		return types.Coin{}, ErrNoMessagesInTransaction
	}

	var msgSend *cosmosModel.MsgSend
	for _, rawMsg := range txResponse.Tx.Body.Messages {
		var tempMsg cosmosModel.MsgSend
		if err := json.Unmarshal(rawMsg, &tempMsg); err != nil {
			continue
		}
		if tempMsg.TypeURL == "/cosmos.bank.v1beta1.MsgSend" {
			msgSend = &tempMsg
			break
		}
	}

	if msgSend == nil {
		logger.Error("no MsgSend found", "err", ErrNoMsgSendInTransaction)
		return types.Coin{}, ErrNoMsgSendInTransaction
	}

	if !strings.EqualFold(msgSend.To, expectedToAddress) {
		err := fmt.Errorf("%w: got %s, expected %s", ErrTransactionRecipientMismatch, msgSend.To, expectedToAddress)
		logger.Error("recipient mismatch", "err", err, "got", msgSend.To, "expected", expectedToAddress)
		return types.Coin{}, err
	}

	if !strings.EqualFold(msgSend.From, expectedFromAddress) {
		err := fmt.Errorf("%w: got %s, expected %s", ErrTransactionSenderMismatch, msgSend.From, expectedFromAddress)
		logger.Error("sender mismatch", "err", err, "got", msgSend.From, "expected", expectedFromAddress)
		return types.Coin{}, err
	}

	if len(msgSend.Amount) == 0 {
		logger.Error("no amount in MsgSend", "err", ErrNoAmountInMsgSend)
		return types.Coin{}, ErrNoAmountInMsgSend
	}

	actualTransferredCoin := msgSend.Amount[0]
	cosmosCoinTx, err := types.ParseCoinNormalized(actualTransferredCoin.Amount + actualTransferredCoin.Denom)
	if err != nil {
		logger.Error("failed to parse transaction amount", "err", err)
		return types.Coin{}, fmt.Errorf("failed to parse transaction amount: %w", err)
	}

	cosmosCoinDb, err := types.ParseCoinNormalized(expectedAmount)
	if err != nil {
		logger.Error("failed to parse database amount", "err", err)
		return types.Coin{}, fmt.Errorf("failed to parse database amount: %w", err)
	}

	cosmosCoinMemo := memo.Amount

	if !cosmosCoinTx.Amount.Equal(cosmosCoinDb.Amount) || cosmosCoinTx.Denom != cosmosCoinDb.Denom {
		err := fmt.Errorf("%w: transaction %s != database %s", ErrTransactionAmountMismatch, cosmosCoinTx.String(), cosmosCoinDb.String())
		logger.Error("amount verification failed", "err", err, "txAmount", cosmosCoinTx.String(), "dbAmount", cosmosCoinDb.String())
		return types.Coin{}, err
	}

	if !cosmosCoinTx.Amount.Equal(cosmosCoinMemo.Amount) || cosmosCoinTx.Denom != cosmosCoinMemo.Denom {
		err := fmt.Errorf("%w: transaction %s != memo %s", ErrTransactionAmountMismatch, cosmosCoinTx.String(), cosmosCoinMemo.String())
		logger.Error("amount verification failed", "err", err, "txAmount", cosmosCoinTx.String(), "memoAmount", cosmosCoinMemo.String())
		return types.Coin{}, err
	}

	return cosmosCoinTx, nil
}

func DoMintLikeCoinByCosmosAddress(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	ethClient *ethclient.Client,
	cosmosAPI *api.CosmosAPI,
	evmLikeCoinClient *evm.LikeCoin,
	cosmosLikcCoinClient *cosmos.LikeCoin,

	cosmosAddress string,
) (*model.LikeCoinMigration, error) {
	mylogger := logger.
		WithGroup("DoMintLikeCoinByCosmosAddress")

	a, err := appdb.QueryLatestLikeCoinMigration(db, cosmosAddress)

	if err != nil {
		logger.Error("appdb.QueryLatestLikeCoinMigration", "err", err)
		return nil, err
	}

	if a.Status == model.LikeCoinMigrationStatusCompleted {
		logger.Info("already completed")
		return a, nil
	}

	if a.Status == model.LikeCoinMigrationStatusFailed {
		logger.Info("already failed", "err", a.FailedReason)
		return nil, errors.New(*a.FailedReason)
	}

	if a.Status != model.LikeCoinMigrationStatusVerifyingCosmosTx {
		logger.Error("already failed", "err", ErrAlreadyInProgress)
		return nil, ErrAlreadyInProgress
	}

	return DoMintLikeCoin(
		ctx,
		mylogger,
		db,
		ethClient,
		cosmosAPI,
		evmLikeCoinClient,
		cosmosLikcCoinClient,
		a,
	)
}

func DoMintLikeCoinByLikeCoinMigrationId(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	ethClient *ethclient.Client,
	cosmosAPI *api.CosmosAPI,
	evmLikeCoinClient *evm.LikeCoin,
	cosmosLikcCoinClient *cosmos.LikeCoin,

	likecoinMigrationId uint64,
) (*model.LikeCoinMigration, error) {
	mylogger := logger.
		WithGroup("DoMintLikeCoinByCosmosAddress")

	a, err := appdb.QueryLikeCoinMigrationById(db, likecoinMigrationId)

	if err != nil {
		logger.Error("appdb.QueryLatestLikeCoinMigration", "err", err)
		return nil, err
	}

	if a.Status == model.LikeCoinMigrationStatusCompleted {
		logger.Info("already completed")
		return a, nil
	}

	if a.Status == model.LikeCoinMigrationStatusFailed {
		logger.Info("already failed", "err", a.FailedReason)
		return nil, errors.New(*a.FailedReason)
	}

	if a.Status != model.LikeCoinMigrationStatusVerifyingCosmosTx {
		logger.Error("already failed", "err", ErrAlreadyInProgress)
		return nil, ErrAlreadyInProgress
	}

	return DoMintLikeCoin(
		ctx,
		mylogger,
		db,
		ethClient,
		cosmosAPI,
		evmLikeCoinClient,
		cosmosLikcCoinClient,
		a,
	)
}

func DoMintLikeCoin(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	ethClient *ethclient.Client,
	cosmosAPI *api.CosmosAPI,
	evmLikeCoinClient *evm.LikeCoin,
	cosmosLikcCoinClient *cosmos.LikeCoin,

	a *model.LikeCoinMigration,
) (*model.LikeCoinMigration, error) {
	mylogger := logger.
		WithGroup("DoMintLikeCoin")
	txLogger := logger.WithGroup("TxLog")

	txResponse, err := cosmosAPI.QueryTransactionWithRetry(*a.CosmosTxHash, 1)
	if err != nil {
		mylogger.Error("cosmosAPI.QueryTransactionWithRetry", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	if txResponse.Code != 0 {
		reason := txResponse.RawLog
		if reason == "" {
			reason = fmt.Sprintf("cosmos tx %s failed with code %d", *a.CosmosTxHash, txResponse.Code)
		}
		mylogger.Error("cosmosAPI.QueryTransactionWithRetry", "err", reason)
		return nil, doMintLikeCoinFailed(db, a, errors.New(reason))
	}

	memoString := txResponse.Tx.Body.Memo
	memo, err := ParseMemoData(memoString)
	if err != nil {
		mylogger.Error("likecoin.ParseMemoData", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	cosmosCoinDb, err := verifyCosmosTransaction(
		mylogger,
		txResponse,
		a.UserCosmosAddress,
		a.BurningCosmosAddress,
		a.Amount,
		memo,
	)
	if err != nil {
		mylogger.Error("verifyCosmosTransaction", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	oldDecimals, err := cosmosLikcCoinClient.Decimals()

	if err != nil {
		mylogger.Error("cosmosLikcCoinClient.Decimal", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	// Use the verified amount from the actual transaction
	evmAmountDecimal := decimal.NewFromBigInt(cosmosCoinDb.Amount.BigInt(), -int32(oldDecimals))

	message := GetEthSigningMessage(evmAmountDecimal)

	recoveredAddr, err := ethereum.RecoverAddress(memo.Signature, []byte(message))
	if err != nil {
		mylogger.Error("ethereum.RecoverAddress", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	recoveredAddrStr := hexutil.Encode(recoveredAddr.Bytes())
	if !strings.EqualFold(recoveredAddrStr, a.UserEthAddress) {
		mylogger.Error("strings.EqualFold(recoveredAddrStr, a.UserCosmosAddress)", "err", ErrEthAddressOnCosmosNotMatch, "recoveredAddrStr", recoveredAddrStr, "a.UserCosmosAddress", a.UserCosmosAddress)
		return nil, doMintLikeCoinFailed(db, a, ErrEthAddressOnCosmosNotMatch)
	}

	mintingEthAddressStr, err := evmLikeCoinClient.Signer.GetSignerAddress()

	if err != nil {
		mylogger.Error("ethereum.PrivateKeyStringToAddress", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	if !strings.EqualFold(*mintingEthAddressStr, a.MintingEthAddress) {
		mylogger.Error("minting eth address not match", "err", ErrMintingEthPrivateKeyNotMatchMintingEthAddress)
		return nil, doMintLikeCoinFailed(db, a, ErrMintingEthPrivateKeyNotMatchMintingEthAddress)
	}

	txLogger = txLogger.With(
		"FromSignerAddress", *mintingEthAddressStr,
		"ToUserAddress", a.UserEthAddress,
	)

	beforeBalance, err := evmLikeCoinClient.BalanceOf(ctx, common.HexToAddress(*mintingEthAddressStr))
	if err != nil {
		txLogger = txLogger.With("BeforeBalance", fmt.Sprintf("(err: %+v)", err))
	} else {
		txLogger = txLogger.With("BeforeBalance", beforeBalance.String())
	}

	a.Status = model.LikeCoinMigrationStatusEvmMinting
	err = appdb.UpdateLikeCoinMigration(db, a)

	if err != nil {
		mylogger.Error("appdb.UpdateLikeCoinMigration", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	newDecimals, err := evmLikeCoinClient.Decimals()

	if err != nil {
		mylogger.Error("evmLikeCoinClient.LikeCoin.Decimals", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	evmAmount, err := util.ConvertAmountByDecimals(
		cosmosCoinDb.Amount, oldDecimals, newDecimals)

	if err != nil {
		mylogger.Error("util.ConvertAmountByDecimals", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	txLogger = txLogger.With(
		"cosmosCoin", cosmosCoinDb.String(),
		"oldDecimals", oldDecimals,
		"newDecimals", newDecimals,
		"evmAmount", evmAmount.String(),
	)

	mylogger.Info(
		"Coin conversion info",
		"cosmosCoin", cosmosCoinDb.String(),
		"oldDecimals", oldDecimals,
		"newDecimals", newDecimals,
		"evmAmount", evmAmount.String(),
	)

	tx, _, err := evmLikeCoinClient.TransferTo(
		ctx, mylogger, *recoveredAddr, evmAmount,
	)

	if err != nil {
		mylogger.Error("ethereum.TransferToken", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	txLogger = txLogger.With("EvmTxHash", tx.Hash().Hex())

	afterBalance, err := evmLikeCoinClient.BalanceOf(ctx, common.HexToAddress(*mintingEthAddressStr))
	if err != nil {
		txLogger = txLogger.With("AfterBalance", fmt.Sprintf("(err: %+v)", err))
	} else {
		txLogger = txLogger.With("AfterBalance", afterBalance.String())
	}

	evmTxHash := hexutil.Encode(tx.Hash().Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeCoinMigrationStatusCompleted
	err = appdb.UpdateLikeCoinMigration(db, a)

	if err != nil {
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	txLogger.Info("likecoin transfer completed")

	return a, nil
}

func doMintLikeCoinFailed(db *sql.DB, a *model.LikeCoinMigration, err error) error {
	a.Status = model.LikeCoinMigrationStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeCoinMigration(db, a))
}
