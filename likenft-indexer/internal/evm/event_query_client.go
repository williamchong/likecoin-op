package evm

import (
	"context"

	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/like_protocol"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EVMQueryClient interface {
	GetLikeProtocolABI() *abi.ABI
	GetBookNFTABI() *abi.ABI

	QueryBookNFTOwnershipTransferred(
		ctx context.Context,
		contractAddress common.Address,
		startBlock uint64,
	) ([]types.Log, error)
	QueryContractURIUpdated(
		ctx context.Context,
		contractAddress common.Address,
		startBlock uint64,
	) ([]types.Log, error)
	QueryNewBookNFT(
		ctx context.Context,
		contractAddress common.Address,
		startBlock uint64,
	) ([]types.Log, error)
	QueryTransferWithMemo(
		ctx context.Context,
		contractAddress common.Address,
		startBlock uint64,
	) ([]types.Log, error)
	QueryTransfer(
		ctx context.Context,
		contractAddress common.Address,
		startBlock uint64,
	) ([]types.Log, error)
}

type evmQueryClient struct {
	client          *ethclient.Client
	LikeProtocolABI *abi.ABI
	BookNFTABI      *abi.ABI
}

func NewEvmQueryClient(url string) (EVMQueryClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	likeprotocolABI, err := like_protocol.LikeProtocolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	booknftABI, err := book_nft.BookNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &evmQueryClient{
		client:          client,
		LikeProtocolABI: likeprotocolABI,
		BookNFTABI:      booknftABI,
	}, nil
}

func (c *evmQueryClient) GetLikeProtocolABI() *abi.ABI {
	return c.LikeProtocolABI
}

func (c *evmQueryClient) GetBookNFTABI() *abi.ABI {
	return c.BookNFTABI

}
