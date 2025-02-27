package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) NewClass(msgNewClass like_protocol.MsgNewClass) (*types.Transaction, error) {
	opts, err := l.transactOpts()

	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikenft: %v", err)
	}
	tx, err := instance.NewClass(opts, msgNewClass)
	if err != nil {
		return nil, fmt.Errorf("err instance.NewClass: %v", err)
	}
	return tx, nil
}
