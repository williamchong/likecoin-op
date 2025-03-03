package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) MintNFTs(msgMintNFTs *like_protocol.MsgMintNFTs) (*types.Transaction, error) {
	opts, err := l.transactOpts()

	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikenft: %v", err)
	}
	tx, err := instance.MintNFTs(opts, *msgMintNFTs)
	if err != nil {
		return nil, fmt.Errorf("err instance.NewClass: %v", err)
	}

	return tx, nil
}
