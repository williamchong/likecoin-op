package evm

import (
	"context"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	appdb "github.com/likecoin/like-signer-backend/pkg/db"
)

type NonceAvailability string

var (
	NonceAvailabilityAvailable NonceAvailability = "available"
	NonceAvailabilityUsing     NonceAvailability = "using"
)

func getNonceAvailability(
	logger *slog.Logger,
	db appdb.TxLike,
	ethAddress common.Address,
	remoteNonce uint64,
) (*NonceAvailability, error) {
	mylogger := logger.WithGroup("getNonceAvailability").
		With("eth_address", ethAddress.Hex()).
		With("remote_nonce", remoteNonce)

	latestNonce, exists, err := appdb.QueryLatestNonce(db, ethAddress.Hex())
	if err != nil {
		mylogger.Error("failed to query latest nonce", "error", err)
		return nil, err
	}

	if !exists {
		mylogger.Info("no latest nonce found, assume nonce is available")
		return &NonceAvailabilityAvailable, nil
	}

	mylogger = mylogger.With("latest_nonce", latestNonce)

	if remoteNonce > latestNonce {
		mylogger.Info("nonce is available")
		return &NonceAvailabilityAvailable, nil
	}

	mylogger.Info("nonce is using")
	return &NonceAvailabilityUsing, nil
}

func (c *Client) awaitAvailableNonce(
	ctx context.Context,
	logger *slog.Logger,
) (uint64, error) {
	mylogger := logger.WithGroup("awaitAvailableNonce")

	fromAddress, err := c.SignerAddress()
	if err != nil {
		mylogger.Error("failed to get signer address", "error", err)
		return 0, err
	}
	mylogger = mylogger.With("from_address", fromAddress)

	availableNonceChan := make(chan uint64)
	errorChan := make(chan error)

	go func() {
		for {
			select {
			case <-ctx.Done():
				errorChan <- ctx.Err()
				return
			case <-time.After(1 * time.Second):
			}

			remoteNonce, err := c.GetPendingNonce(ctx)

			if err != nil {
				errorChan <- err
				return
			}

			nonceAvailability, err := getNonceAvailability(mylogger, c.db, fromAddress, remoteNonce)

			if err != nil {
				errorChan <- err
				return
			}

			if *nonceAvailability == NonceAvailabilityAvailable {
				availableNonceChan <- remoteNonce
				return
			}
		}
	}()

	select {
	case err := <-errorChan:
		mylogger.Error("failed to await available nonce", "error", err)
		return 0, err
	case nonce := <-availableNonceChan:
		mylogger.Info("available nonce found", "nonce", nonce)
		return nonce, nil
	}
}
