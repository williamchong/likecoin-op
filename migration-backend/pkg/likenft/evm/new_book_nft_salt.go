package evm

import (
	"crypto/sha256"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/model"
)

func ComputeNewBookNFTSalt(
	msgSender common.Address,
	nonce [2]byte,
	metadata *model.ContractLevelMetadata,
) ([32]byte, error) {
	dataCandidates := getDataCandidatesFromContractLevelMetadata(metadata)
	return ComputeSaltDataFromCandidates(msgSender, nonce, dataCandidates...)
}

func ComputeSaltDataFromCandidates(
	msgSender common.Address,
	nonce [2]byte,
	dataCandidates ...*string,
) ([32]byte, error) {
	for _, dataCandidate := range dataCandidates {
		if dataCandidate == nil || *dataCandidate == "" {
			continue
		}

		var data [10]byte
		hash := sha256.Sum256([]byte(*dataCandidate))
		copy(data[:], hash[:])

		return newBookNFTSaltLayout(msgSender, nonce, data), nil
	}

	return [32]byte{}, errors.New("err cannot determine salt")
}

func getDataCandidatesFromContractLevelMetadata(
	metadata *model.ContractLevelMetadata,
) []*string {
	dataCandidates := make([]*string, 0)

	if metadata.PotentialAction != nil &&
		metadata.PotentialAction.Target != nil &&
		len(metadata.PotentialAction.Target) > 0 {
		dataCandidates = append(dataCandidates, &metadata.PotentialAction.Target[0].Url)
	}

	if metadata.Name != "" {
		dataCandidates = append(dataCandidates, &metadata.Name)
	}

	return dataCandidates
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
