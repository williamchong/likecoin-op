package evm

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *Client) Sign(
	ctx context.Context,
	logger *slog.Logger,

	to common.Address,
	data []byte,
	value *big.Int,
) (*types.Transaction, error) {
	mylogger := logger.
		WithGroup("Sign").
		With("to", to).
		With("value", value)

	chainID, err := c.ethClient.ChainID(ctx)
	if err != nil {
		mylogger.Error("failed to get chain ID", "error", err)
		return nil, err
	}
	privateKey, err := c.privateKey()
	if err != nil {
		mylogger.Error("failed to get private key", "error", err)
		return nil, err
	}
	mylogger = mylogger.With("chain_id", chainID)
	fromAddress, err := c.SignerAddress()
	if err != nil {
		mylogger.Error("failed to get signer address", "error", err)
		return nil, err
	}
	mylogger = mylogger.With("from_address", fromAddress)
	nonce, err := c.awaitAvailableNonce(ctx, mylogger)
	if err != nil {
		mylogger.Error("failed to get nonce", "error", err)
		return nil, err
	}
	mylogger = mylogger.With("nonce", nonce)
	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		mylogger.Error("failed to get gas price", "error", err)
		return nil, err
	}
	mylogger = mylogger.With("gas_price", gasPrice)
	gasLimit, err := c.estimateGasLimit(
		ctx,
		fromAddress,
		&to,
		data,
		gasPrice,
		nil,
		nil,
		value,
	)
	if err != nil {
		mylogger.Error("failed to estimate gas limit", "error", err)
		return nil, err
	}
	mylogger = mylogger.With("gas_limit", gasLimit)
	tx := types.NewTransaction(
		nonce,
		to,
		value,
		gasLimit,
		gasPrice,
		data,
	)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		mylogger.Error("failed to sign transaction", "error", err)
		return nil, err
	}
	return tx, nil
}
