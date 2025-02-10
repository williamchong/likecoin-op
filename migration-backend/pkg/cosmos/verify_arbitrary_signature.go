package cosmos

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func VerifyArbitrarySignature(pubKey string, signature string, message string) (bool, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return false, err
	}

	// Create public key object
	pk := secp256k1.PubKey{Key: pubKeyBytes}

	// Decode signature from base64
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	fmt.Printf("sig: %v\n", len(sig))

	// Verify the signature
	// FIXME: construct a cosmos message which is signed from client
	// {
	//   method: 'cos_signMessage',
	//   params: {
	//     chainName,
	//     signer,
	//     message,
	//   },
	// }
	// ref https://github.com/likecoin/likecoin-wallet-connector/blob/1d1cfdb7e3764fa2c94fce9d591620f75fc11f96/src/utils/cosmostation.ts#L72
	// Issue: https://linear.app/oursky/issue/LIK3-74/investigate-cosmos-verify-signature-of-cos-signmessage-message
	messageBytes := []byte(message)
	return pk.VerifySignature(messageBytes, sig), nil
}
