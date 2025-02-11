package evm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	client *ethclient.Client
}

func NewEvmClient(url string) (*EvmClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EvmClient{
		client: client,
	}, nil
}

func (c *EvmClient) GetNonce(address common.Address) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// The ABI for the LikeProtocol,
// TODO: generate from the contract
const ownerABI = `[{
    "inputs": [],
    "name": "owner",
    "outputs": [{"name": "", "type": "address"}],
    "stateMutability": "view",
    "type": "function"
}]`

func (c *EvmClient) GetLikeProtocolOwner() (ownerAddress common.Address, err error) {
	// TODO: get from env
	contractAddress := common.HexToAddress("0xfF79df388742f248c61A633938710559c61faEF1")

	parsedABI, err := abi.JSON(strings.NewReader(ownerABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse ABI: %v", err)
	}

	data, err := parsedABI.Pack("owner")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to pack data: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	result, err := c.client.CallContract(context.Background(), msg, nil) // nil for latest block
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call contract: %v", err)
	}

	err = parsedABI.UnpackIntoInterface(&ownerAddress, "owner", result)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack result: %v", err)
	}

	return ownerAddress, nil
}
