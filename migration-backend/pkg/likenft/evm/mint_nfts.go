package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) MintNFTs(msgMintNFTsFromTokenId *like_protocol.MsgMintNFTsFromTokenId) (*types.Transaction, error) {
	opts, err := l.transactOpts()
	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikenft: %v", err)
	}
	tx, err := instance.SafeMintNFTsWithTokenId(opts, *msgMintNFTsFromTokenId)
	if err != nil {
		return nil, fmt.Errorf("err instance.NewClass: %v", err)
	}

	err = l.Client.SendTransaction(opts.Context, tx)
	return tx, err
}
