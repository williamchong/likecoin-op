package evm

import (
	"context"
	"math/big"

	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/util/jsondatauri"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *EvmClient) GetTokenURI(
	ctx context.Context,
	contractAddress common.Address,
	tokenId *big.Int,
) (jsondatauri.JSONDataUri, error) {
	bookNFTClient, err := book_nft.NewBookNft(contractAddress, c.client)
	if err != nil {
		return "", err
	}
	tokenURI, err := bookNFTClient.TokenURI(&bind.CallOpts{}, tokenId)
	if err != nil {
		return "", err
	}
	return jsondatauri.JSONDataUri(tokenURI), nil
}
