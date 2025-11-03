package evm

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrNoCode = errors.New("no code")

type Client struct {
	db            *sql.DB
	ethClient     *ethclient.Client
	nonceProvider NonceProvider
	privateKeyStr string
}

func NewClient(
	db *sql.DB,
	ethClient *ethclient.Client,
	nonceProvider NonceProvider,
	privateKeyStr string,
) *Client {
	return &Client{
		db:            db,
		ethClient:     ethClient,
		nonceProvider: nonceProvider,
		privateKeyStr: privateKeyStr,
	}
}

func (c *Client) Client() *ethclient.Client {
	return c.ethClient
}

func (c *Client) privateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(c.privateKeyStr)
}

func (c *Client) SignerAddress() (common.Address, error) {
	privateKey, err := c.privateKey()
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey), nil
}

func (c *Client) estimateGasLimit(
	ctx context.Context,
	signerAddress common.Address,
	contract *common.Address,
	input []byte,
	gasPrice, gasTipCap, gasFeeCap, value *big.Int,
) (uint64, error) {
	if contract != nil {
		// Gas estimation cannot succeed without code for method invocations.
		if code, err := c.ethClient.PendingCodeAt(ctx, *contract); err != nil {
			return 0, err
		} else if len(code) == 0 {
			return 0, ErrNoCode
		}
	}
	msg := ethereum.CallMsg{
		From:      signerAddress,
		To:        contract,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      input,
	}
	return c.ethClient.EstimateGas(ctx, msg)
}

func (c *Client) estimateGasLimitForTransfer(
	ctx context.Context,
	signerAddress common.Address,
	toAddress common.Address,
	input []byte,
	gasPrice, gasTipCap, gasFeeCap, value *big.Int,
) (uint64, error) {
	msg := ethereum.CallMsg{
		From:      signerAddress,
		To:        &toAddress,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      input,
	}
	return c.ethClient.EstimateGas(ctx, msg)
}
