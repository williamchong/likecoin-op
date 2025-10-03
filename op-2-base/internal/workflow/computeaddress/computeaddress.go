package computeaddress

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/preparebooknfts"
)

type Input struct {
	preparebooknfts.Output
}

type Output struct {
	OldAddress string `json:"old_address"`
	Salt       string `json:"salt"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	NewAddress string `json:"new_address"`
}

type ComputeAddress interface {
	Compute(input *Input) (*Output, error)
}

type computeAddress struct {
	creationCode    creationcode.CreationCode
	protocolAddress common.Address
	signerAddress   common.Address
}

func NewComputeAddress(
	creationCode creationcode.CreationCode,
	protocolAddress common.Address,
	signerAddress common.Address,
) ComputeAddress {
	return &computeAddress{
		creationCode,
		protocolAddress,
		signerAddress,
	}
}

func (c *computeAddress) Compute(input *Input) (*Output, error) {
	saltData := input.Salt
	if saltData == nil || *saltData == "" {
		saltData = input.Salt2
	}

	salt, err := evm.ComputeSaltDataFromCandidates(
		c.signerAddress,
		[2]byte{0, 0},
		*saltData,
	)
	if err != nil {
		return nil, fmt.Errorf("evm.ComputeNewBookNFTSalt: %v", err)
	}

	initHash, err := c.creationCode.MakeInitCodeHash(
		c.protocolAddress,
		input.Metadata.Name,
		input.Metadata.Symbol,
	)
	if err != nil {
		return nil, fmt.Errorf("creationCode.MakeInitCodeHash: %v", err)
	}

	newAddress := crypto.CreateAddress2(c.protocolAddress, salt, initHash)

	return &Output{
		OldAddress: input.OpAddress,
		Salt:       "0x" + hex.EncodeToString(salt[:]),
		Name:       input.Metadata.Name,
		Symbol:     input.Metadata.Symbol,
		NewAddress: newAddress.Hex(),
	}, nil
}
