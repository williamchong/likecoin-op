package creationcode

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreationCode interface {
	MakeInitCodeHash(protocolAddress common.Address, name string, symbol string) ([]byte, error)
}

type creationCode struct {
	byteCode []byte
}

func NewCreationCode(byteCode []byte) CreationCode {
	return &creationCode{byteCode}
}

func (c *creationCode) MakeInitCodeHash(protocolAddress common.Address, name string, symbol string) ([]byte, error) {
	// Bytecode file is expected to be a hex string like 0x.... ; decode to raw bytes
	hexStr := strings.TrimSpace(string(c.byteCode))[2:]
	creationCode, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to hex-decode creation code: %v\n", err)
	}

	creationCode, err = hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to hex-decode creation code: %v\n", err)
	}

	parsedAbi, _ := abi.JSON(strings.NewReader(`[
		{
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "symbol",
          "type": "string"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }]`))

	initData, err := parsedAbi.Pack("initialize", name, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to pack initialize data: %v\n", err)
	}

	// Encode constructor args (address beacon, bytes data)
	addressType, _ := abi.NewType("address", "address", nil)
	bytesType, _ := abi.NewType("bytes", "bytes", nil)
	constructorArgs := abi.Arguments{
		{Type: addressType},
		{Type: bytesType},
	}
	encodedArgs, err := constructorArgs.Pack(protocolAddress, initData)
	if err != nil {
		return nil, fmt.Errorf("failed to pack constructor args: %v\n", err)
	}
	proxyCreationCode := append(creationCode, encodedArgs...)
	initCodeHash := crypto.Keccak256(proxyCreationCode)
	return initCodeHash, nil
}
