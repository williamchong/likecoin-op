package evm

import (
	"context"
	"fmt"
	"math/big"

	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/like_protocol"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jellydator/ttlcache/v3"
)

type EvmClient struct {
	client          *ethclient.Client
	LikeProtocolABI *abi.ABI
	BookNFTABI      *abi.ABI

	chainIdCache     *ttlcache.Cache[string, *big.Int]
	blockNumberCache *ttlcache.Cache[string, uint64]
}

func NewEvmClient(
	url string,
	chainIdCache *ttlcache.Cache[string, *big.Int],
	blockNumberCache *ttlcache.Cache[string, uint64],
) (*EvmClient, error) {
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

	return &EvmClient{
		client:           client,
		LikeProtocolABI:  likeprotocolABI,
		BookNFTABI:       booknftABI,
		chainIdCache:     chainIdCache,
		blockNumberCache: blockNumberCache,
	}, nil
}

func (c *EvmClient) GetNonce(address common.Address) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (c *EvmClient) GetLikeProtocolOwner() (ownerAddress common.Address, err error) {
	// TODO: get from env
	contractAddress := common.HexToAddress("0xfF79df388742f248c61A633938710559c61faEF1")

	parsedABI := c.LikeProtocolABI

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

func (c *EvmClient) GetHeaderByBlockNumber(
	ctx context.Context,
	blockNumber uint64,
) (*types.Header, error) {
	return c.client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
}

func (c *EvmClient) GetHeaderMapByBlockNumbers(
	ctx context.Context,
	blockNumbers []uint64,
) (map[uint64]*types.Header, error) {
	headerMap := make(map[uint64]*types.Header)

	for _, blockNumber := range blockNumbers {
		header, ok := headerMap[blockNumber]
		if ok {
			continue
		}

		header, err := c.GetHeaderByBlockNumber(ctx, blockNumber)
		if err != nil {
			return nil, err
		}
		headerMap[header.Number.Uint64()] = header
	}

	return headerMap, nil
}

func (c *EvmClient) ChainID(ctx context.Context) (*big.Int, error) {
	item := c.chainIdCache.Get("chainId")
	if item != nil {
		return item.Value(), nil
	}
	chainId, err := c.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	c.chainIdCache.Set("chainId", chainId, ttlcache.DefaultTTL)
	return chainId, nil
}

func (c *EvmClient) BlockNumber(ctx context.Context) (uint64, error) {
	item := c.blockNumberCache.Get("blockNumber")
	if item != nil {
		return item.Value(), nil
	}
	blockNumber, err := c.client.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	c.blockNumberCache.Set("blockNumber", blockNumber, ttlcache.DefaultTTL)
	return blockNumber, nil
}

func (c *EvmClient) InvalidateBlockNumberCache() {
	c.blockNumberCache.Delete("blockNumber")
}
