package evm

import (
	"context"

	"likenft-indexer/internal/evm/book_nft"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *EvmClient) GetBookNFTOwner(
	ctx context.Context,
	contractAddress common.Address,
) (common.Address, error) {
	bookNFTClient, err := book_nft.NewBookNft(contractAddress, c.client)
	if err != nil {
		return common.Address{}, err
	}
	ownerAddress, err := bookNFTClient.Owner(&bind.CallOpts{})
	if err != nil {
		return common.Address{}, err
	}
	return ownerAddress, nil
}
