package likenft

import (
	"database/sql"
	"errors"
	"math/big"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

var nftIdRegex = regexp.MustCompile("(?P<prefix>.+)-(?P<maybe_num>[0-9]+)")
var numIndex = nftIdRegex.SubexpIndex("maybe_num")

func DoMintNFTAction(
	db *sql.DB,
	p *evm.LikeProtocol,
	c *evm.LikeNFTClass,
	a *model.LikeNFTMigrationActionMintNFT,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	if a.Status == model.LikeNFTMigrationActionMintNFTStatusCompleted {
		return a, nil
	}
	if a.Status != model.LikeNFTMigrationActionMintNFTStatusInit &&
		a.Status != model.LikeNFTMigrationActionMintNFTStatusFailed {
		return nil, errors.New("error new class action is not init or failed")
	}

	evmClassAddress := common.HexToAddress(a.EvmClassId)
	toOwner := common.HexToAddress(a.EvmOwner)
	initialBatchMintOwnerAddress := common.HexToAddress(a.InitialBatchMintOwner)

	matches := nftIdRegex.FindStringSubmatch(a.CosmosNFTId)
	var txHash *common.Hash
	if matches == nil {
		_txHash, err := p.MintNFTs(&like_protocol.MsgMintNFTs{
			ClassId: evmClassAddress,
			To:      toOwner,
			Inputs: []like_protocol.NFTData{
				{
					Metadata: "{}", // TODO
				},
			},
		})
		txHash = _txHash
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
	} else {
		nftIdStr := matches[numIndex]
		nftId, err := strconv.ParseUint(nftIdStr, 10, 64)
		if err != nil {
			txHash, err = p.MintNFTs(&like_protocol.MsgMintNFTs{
				ClassId: evmClassAddress,
				To:      toOwner,
				Inputs: []like_protocol.NFTData{
					{
						Metadata: "{}", // TODO
					},
				},
			})
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		} else {
			// batch mint
			totalSupply, err := c.TotalSupply(evmClassAddress)
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
			desireBatchMintAmount := GetDesireBatchMintAmount(totalSupply, nftId)
			if desireBatchMintAmount.Cmp(big.NewInt(0)) == 1 {
				inputs := make([]like_protocol.NFTData, 0)
				for i := big.NewInt(0); i.Cmp(desireBatchMintAmount) == -1; i = i.Add(i, big.NewInt(1)) {
					inputs = append(inputs, like_protocol.NFTData{
						Metadata: "{}", // TODO
					})
				}
				_, err = p.MintNFTs(&like_protocol.MsgMintNFTs{
					ClassId: evmClassAddress,
					To:      initialBatchMintOwnerAddress,
					Inputs:  inputs,
				})
				if err != nil {
					return nil, doMintNFTActionFailed(db, a, err)
				}
			}
			txHash, err = p.TransferNFT(evmClassAddress, initialBatchMintOwnerAddress, toOwner, big.NewInt(int64(nftId)))
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		}
	}

	evmTxHash := hexutil.Encode(txHash.Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeNFTMigrationActionMintNFTStatusCompleted
	err := appdb.UpdateLikeNFTMigrationActionMintNFT(db, a)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	return a, nil
}

// evm token id is 0 based
// so class/token-0001 will be mapped to 1 in evm
// which requires offset 1 from 0
func GetDesireBatchMintAmount(
	totalSupply *big.Int,
	desireAmount uint64,
) *big.Int {
	switch totalSupply.Cmp(big.NewInt(int64(desireAmount))) {
	case 0:
		return big.NewInt(1)
	case 1:
		return big.NewInt(0)
	case -1:
		return big.NewInt(0).Sub(big.NewInt(int64(desireAmount+1)), totalSupply)
	}
	panic("unhandled switch")
}

func doMintNFTActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionMintNFT, err error) error {
	a.Status = model.LikeNFTMigrationActionMintNFTStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionMintNFT(db, a))
}
