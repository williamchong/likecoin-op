package likecoin

import (
	"context"
	"database/sql"
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
	"github.com/shopspring/decimal"
)

var ErrAlreadyInProgress = fmt.Errorf("err already in progress")
var ErrEthAddressOnCosmosNotMatch = fmt.Errorf("err eth address on cosmos not match")
var ErrMintingEthPrivateKeyNotMatchMintingEthAddress = fmt.Errorf("err minting eth private key not match minter eth address")
var ErrAmountNotMatch = fmt.Errorf("err amount not match")

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

	txResponse, err := cosmosAPI.QueryTransaction(*a.CosmosTxHash)
	if err != nil {
		mylogger.Error("cosmosAPI.QueryTransaction", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}
	memoString := txResponse.Tx.Body.Memo
	memo, err := ParseMemoData(memoString)
	if err != nil {
		mylogger.Error("likecoin.ParseMemoData", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	cosmosCoinDb, err := types.ParseCoinNormalized(a.Amount)

	if err != nil {
		mylogger.Error("types.ParseCoinNormalized(a.Amount)", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	cosmosCoinMemo := memo.Amount

	if cosmosCoinMemo.Amount != cosmosCoinDb.Amount && cosmosCoinMemo.Denom != cosmosCoinDb.Denom {
		mylogger.Error("coin not matched", "err", ErrAmountNotMatch)
		return nil, doMintLikeCoinFailed(db, a, ErrAmountNotMatch)
	}

	oldDecimals, err := cosmosLikcCoinClient.Decimals()

	if err != nil {
		mylogger.Error("cosmosLikcCoinClient.Decimal", "err", err)
		return nil, doMintLikeCoinFailed(db, a, err)
	}

	evmAmountDecimal := decimal.NewFromBigInt(cosmosCoinMemo.Amount.BigInt(), -int32(oldDecimals))

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
