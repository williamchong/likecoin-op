package evm

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) NewBookNFT(msgNewBookNFT like_protocol.MsgNewBookNFT) (*types.Transaction, error) {
	opts, err := l.transactOpts()
	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}
	opts.NoSend = true

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewBookNFT: %v", err)
	}
	tx, err := instance.NewBookNFT(opts, msgNewBookNFT)
	if err != nil {
		return nil, fmt.Errorf("err instance.NewBookNFT: %v", err)
	}

	err = l.Client.SendTransaction(opts.Context, tx)
	return tx, err
}

func (l *LikeProtocol) GetClassIdFromNewClassTransaction(txReceipt *types.Receipt) (*common.Address, error) {
	filterer, err := like_protocol.NewLikeProtocolFilterer(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikeProtocolFilterer: %v", err)
	}

	logs := txReceipt.Logs

	for _, vLog := range logs {
		newClassEvent, err := filterer.ParseNewBookNFT(*vLog)
		if err == nil {
			return &newClassEvent.BookNFT, nil
		}
	}
	return nil, errors.New("err finding new book nft event from tx log")
}
