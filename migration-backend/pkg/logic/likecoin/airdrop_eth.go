package likecoin

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/likecoin/like-migration-backend/pkg/ethereum"
)

var ErrSkippedAirdrop = errors.New("skipped airdrop")
var ErrBalanceNotZero = errors.New("balance is not zero")

func DoAirdropEth(
	ctx context.Context,
	logger *slog.Logger,

	ethereumClient ethereum.EthereumClient,

	toAddress common.Address,
	amount *big.Int,
	likecoinMigrationTx *string,
) (*types.Transaction, *types.Receipt, error) {
	mylogger := logger.
		WithGroup("DoAirdropLikeCoin").
		With("toAddress", toAddress.Hex()).
		With("amount", amount.String())
	txLogger := logger.WithGroup("TxLog").
		With("toAddress", toAddress.Hex()).
		With("amount", amount.String())

	if likecoinMigrationTx != nil {
		mylogger = mylogger.With("likecoinMigrationTx", *likecoinMigrationTx)
		txLogger = txLogger.With("likecoinMigrationTx", *likecoinMigrationTx)
	}

	if signerAddress, err := ethereumClient.GetSignerAddress(ctx); err == nil {
		txLogger = txLogger.With("signerAddress", signerAddress.Hex())
		beforeBalance, err := ethereumClient.BalanceOf(ctx, signerAddress)
		if err != nil {
			txLogger = txLogger.With("SignerBeforeBalance", fmt.Sprintf("(err: %+v)", err))
		} else {
			txLogger = txLogger.With("SignerBeforeBalance", fmt.Sprintf("%s ETH", ethereum.ConvertToEther(beforeBalance).String()))
		}
	}

	balance, err := ethereumClient.BalanceOf(ctx, toAddress)
	if err != nil {
		txLogger = txLogger.With("ToAddressBeforeBalance", fmt.Sprintf("(err: %+v)", err))
		txLogger.Info("failed to get to address balance", "error", err)
		mylogger.Error("failed to get to address balance", "error", err)
		return nil, nil, errors.Join(fmt.Errorf("err get to address balance"), err)
	}
	txLogger = txLogger.With("ToAddressBeforeBalance", fmt.Sprintf("%s ETH", ethereum.ConvertToEther(balance).String()))
	if balance.Cmp(big.NewInt(0)) == 1 {
		txLogger.Info("skipped airdrop because to address balance is not zero")
		mylogger.Error("skipped airdrop because to address balance is not zero")
		return nil, nil, errors.Join(ErrSkippedAirdrop, ErrBalanceNotZero)
	}

	tx, txReceipt, err := ethereumClient.TransferToken(ctx, toAddress, amount)
	if err != nil {
		mylogger.Error("ethereumClient.TransferToken", "error", err)
		return nil, nil, errors.Join(fmt.Errorf("err transfer token"), err)
	}

	if signerAddress, err := ethereumClient.GetSignerAddress(ctx); err == nil {
		afterBalance, err := ethereumClient.BalanceOf(ctx, signerAddress)
		if err != nil {
			txLogger = txLogger.With("SignerAfterBalance", fmt.Sprintf("(err: %+v)", err))
		} else {
			txLogger = txLogger.With("SignerAfterBalance", fmt.Sprintf("%s ETH", ethereum.ConvertToEther(afterBalance).String()))
		}
	}

	afterBalance, err := ethereumClient.BalanceOf(ctx, toAddress)
	if err != nil {
		txLogger = txLogger.With("ToAddressAfterBalance", fmt.Sprintf("(err: %+v)", err))
	} else {
		txLogger = txLogger.With("ToAddressAfterBalance", fmt.Sprintf("%s ETH", ethereum.ConvertToEther(afterBalance).String()))
	}

	mylogger.Info("airdrop eth completed", "txHash", tx.Hash().Hex())
	txLogger.Info(
		"airdrop eth completed",
		"txHash", tx.Hash().Hex(),
	)

	return tx, txReceipt, nil
}
