package ethereum

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyStringToAddress(privateKeyString string) (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)

	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &addr, nil
}
