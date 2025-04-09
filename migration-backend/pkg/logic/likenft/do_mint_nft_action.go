package likenft

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"regexp"
	"slices"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

var nftIdRegex = regexp.MustCompile("(?P<prefix>.+)-(?P<maybe_num>[0-9]+)")
var numIndex = nftIdRegex.SubexpIndex("maybe_num")

func MakeMatchNFTIdRegex(id string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("^(?P<prefix>[a-zA-Z0-9]+)-(?P<maybe_num>0*%s)$", id))
}

func DoMintNFTAction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	p *evm.LikeProtocol,
	c *evm.BookNFT,
	m *cosmos.LikeNFTCosmosClient,
	a *model.LikeNFTMigrationActionMintNFT,
) (*model.LikeNFTMigrationActionMintNFT, error) {
	mylogger := logger.
		WithGroup("DoMintNFTAction").
		With("mintNFTActionId", a.Id)

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
	initialMemo := "_mint"

	newClassAction, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		EvmClassId: &a.EvmClassId,
	})
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}

	totalSupply, err := c.TotalSupply(evmClassAddress)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}

	matches := nftIdRegex.FindStringSubmatch(a.CosmosNFTId)

	var (
		tx *types.Transaction
	)

	if matches == nil {
		cosmosNFT, err := m.QueryNFT(cosmos.QueryNFTRequest{
			ClassId: newClassAction.CosmosClassId,
			Id:      a.CosmosNFTId,
		})
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
		metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFT(cosmosNFT.NFT))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}

		metadataString := string(metadataBytes)
		tx, _, err = c.MintNFTs(
			ctx,
			mylogger,
			evmClassAddress,
			totalSupply,
			[]common.Address{toOwner},
			[]string{
				"_mint",
			},
			[]string{
				metadataString,
			})
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
	} else {
		nftIdStr := matches[numIndex]
		nftId, err := strconv.ParseUint(nftIdStr, 10, 64)
		if err != nil {
			cosmosNFT, err := m.QueryNFT(cosmos.QueryNFTRequest{
				ClassId: newClassAction.CosmosClassId,
				Id:      a.CosmosNFTId,
			})
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
			metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFT(cosmosNFT.NFT))
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}

			metadataString := string(metadataBytes)
			tx, _, err = c.MintNFTs(
				ctx,
				mylogger,
				evmClassAddress,
				totalSupply,
				[]common.Address{toOwner},
				[]string{
					"_mint",
				},
				[]string{
					metadataString,
				})
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		} else {
			// batch mint
			desireBatchMintAmount := GetDesireBatchMintAmount(totalSupply, nftId)
			if desireBatchMintAmount.Cmp(big.NewInt(0)) == 1 {
				cosmosNFTs, err := m.QueryAllNFTsByClassId(cosmos.QueryAllNFTsByClassIdRequest{
					ClassId: newClassAction.CosmosClassId,
				})

				if err != nil {
					return nil, doMintNFTActionFailed(db, a, err)
				}

				tos := make([]common.Address, 0)
				memos := make([]string, 0)
				metadataList := make([]string, 0)
				for i := big.NewInt(0); i.Cmp(desireBatchMintAmount) == -1; i = i.Add(i, big.NewInt(1)) {
					cosmosNFTIdx := slices.IndexFunc(cosmosNFTs.NFTs, func(n cosmosmodel.NFT) bool {
						return MakeMatchNFTIdRegex(big.NewInt(0).Add(totalSupply, i).String()).MatchString(n.Id)
					})
					metadataStr := "{}"
					if cosmosNFTIdx != -1 {
						cosmosNFT := cosmosNFTs.NFTs[cosmosNFTIdx]
						metadata := evm.ERC721MetadataFromCosmosNFT(&cosmosNFT)
						metadataBytes, err := json.Marshal(metadata)
						if err != nil {
							return nil, doMintNFTActionFailed(db, a, err)
						}
						metadataStr = string(metadataBytes)
					}
					tos = append(tos, initialBatchMintOwnerAddress)
					memos = append(memos, initialMemo)
					metadataList = append(metadataList, metadataStr)
				}
				_, _, err = c.MintNFTs(
					ctx,
					mylogger,
					evmClassAddress,
					totalSupply,
					tos,
					memos,
					metadataList,
				)
				if err != nil {
					return nil, doMintNFTActionFailed(db, a, err)
				}
			}
			tx, _, err = p.TransferNFT(
				ctx,
				mylogger,
				evmClassAddress,
				initialBatchMintOwnerAddress,
				toOwner,
				big.NewInt(int64(nftId)))
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		}
	}

	evmTxHash := hexutil.Encode(tx.Hash().Bytes())
	a.EvmTxHash = &evmTxHash
	a.Status = model.LikeNFTMigrationActionMintNFTStatusCompleted
	err = appdb.UpdateLikeNFTMigrationActionMintNFT(db, a)
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
