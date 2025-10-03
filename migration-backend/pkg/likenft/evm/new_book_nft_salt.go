package evm

import (
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/common"
)

func ComputeSaltDataFromCandidates(
	msgSender common.Address,
	nonce [2]byte,
	saltData string,
) ([32]byte, error) {
	var data [10]byte
	hash := sha256.Sum256([]byte(saltData))
	copy(data[:], hash[:])

	return newBookNFTSaltLayout(msgSender, nonce, data), nil
}

func newBookNFTSaltLayout(
	msgSender [20]byte,
	nonce [2]byte,
	data [10]byte,
) [32]byte {
	res := [32]byte{}
	copy(res[0:20], msgSender[:])
	copy(res[20:22], nonce[:])
	copy(res[22:32], data[:])
	return res
}
