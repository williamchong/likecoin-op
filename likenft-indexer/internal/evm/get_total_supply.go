package evm

import (
	"context"
	"math/big"

	"likenft-indexer/internal/evm/book_nft"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *EvmClient) GetTotalSupply(
	ctx context.Context,
	contractAddress common.Address,
) (*big.Int, error) {
	bookNFTClient, err := book_nft.NewBookNft(contractAddress, c.client)
	if err != nil {
		return nil, err
	}
	totalSupply, err := bookNFTClient.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return totalSupply, nil
}
