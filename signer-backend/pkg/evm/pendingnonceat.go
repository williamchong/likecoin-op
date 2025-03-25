package evm

import "context"

func (l *Client) GetPendingNonce(ctx context.Context) (uint64, error) {
	pubkey, err := l.SignerAddress()
	if err != nil {
		return 0, err
	}
	n, err := l.ethClient.PendingNonceAt(ctx, pubkey)
	if err != nil {
		return 0, err
	}
	return n, nil
}
