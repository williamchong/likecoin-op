package likenft

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math"
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
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/event"
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
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

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

	newClassAction, err := appdb.QueryLikeNFTMigrationActionNewClass(db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		EvmClassId: &a.EvmClassId,
	})
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}

	totalSupplyBigInt, err := c.TotalSupply(evmClassAddress)
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	totalSupply := totalSupplyBigInt.Uint64()

	cosmosClass, err := m.QueryClassByClassId(cosmos.QueryClassByClassIdRequest{
		ClassId: newClassAction.CosmosClassId,
	})
	if err != nil {
		return nil, doMintNFTActionFailed(db, a, err)
	}
	iscnDataResponse, err := m.GetISCNRecord(
		cosmosClass.Class.Data.Parent.IscnIdPrefix,
		cosmosClass.Class.Data.Parent.IscnVersionAtMint,
	)
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
		metadataOverride, err := m.QueryNFTExternalMetadata(cosmosNFT.NFT)
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}
		metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
			erc721ExternalURLBuilder,
			cosmosNFT.NFT,
			cosmosClass.Class,
			iscnDataResponse,
			metadataOverride,
			a.EvmClassId,
			totalSupply,
		))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}

		events, err := m.QueryAllNFTEvents(m.MakeQueryNFTEventsRequest(newClassAction.CosmosClassId, a.CosmosNFTId))
		if err != nil {
			return nil, doMintNFTActionFailed(db, a, err)
		}

		metadataString := string(metadataBytes)
		tx, _, err = c.MintNFTs(
			ctx,
			mylogger,
			evmClassAddress,
			totalSupplyBigInt,
			[]common.Address{toOwner},
			[]string{
				event.MakeMemoFromEvent(events),
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
			metadataOverride, err := m.QueryNFTExternalMetadata(cosmosNFT.NFT)
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
			metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
				erc721ExternalURLBuilder,
				cosmosNFT.NFT,
				cosmosClass.Class,
				iscnDataResponse,
				metadataOverride,
				a.EvmClassId,
				nftId,
			))
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
			events, err := m.QueryAllNFTEvents(m.MakeQueryNFTEventsRequest(newClassAction.CosmosClassId, a.CosmosNFTId))
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}

			metadataString := string(metadataBytes)
			tx, _, err = c.MintNFTs(
				ctx,
				mylogger,
				evmClassAddress,
				totalSupplyBigInt,
				[]common.Address{toOwner},
				[]string{
					event.MakeMemoFromEvent(events),
				},
				[]string{
					metadataString,
				})
			if err != nil {
				return nil, doMintNFTActionFailed(db, a, err)
			}
		} else {
			// batch mint
			// totalSupply = 10, nftId = 10, contract nftids = [0..9] should mint 1 to [0..10]
			desireSupply := nftId + 1
			desireBatchMintAmount := uint64(math.Max(float64(desireSupply)-float64(totalSupply), 0))
			if desireBatchMintAmount > 0 {
				cosmosNFTs, err := m.QueryAllNFTsByClassId(cosmos.QueryAllNFTsByClassIdRequest{
					ClassId: newClassAction.CosmosClassId,
				})

				if err != nil {
					return nil, doMintNFTActionFailed(db, a, err)
				}

				tos := make([]common.Address, 0)
				memos := make([]string, 0)
				metadataList := make([]string, 0)
				for i := totalSupply; i < desireSupply; i = i + 1 {
					cosmosNFTIdx := slices.IndexFunc(cosmosNFTs.NFTs, func(n cosmosmodel.NFT) bool {
						return MakeMatchNFTIdRegex(strconv.FormatUint(i, 10)).MatchString(n.Id)
					})
					metadataStr := "{}"
					memo := ""
					if cosmosNFTIdx != -1 {
						cosmosNFT := cosmosNFTs.NFTs[cosmosNFTIdx]
						metadataOverride, err := m.QueryNFTExternalMetadata(&cosmosNFT)
						if err != nil {
							return nil, doMintNFTActionFailed(db, a, err)
						}
						metadataBytes, err := json.Marshal(evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
							erc721ExternalURLBuilder,
							&cosmosNFT,
							cosmosClass.Class,
							iscnDataResponse,
							metadataOverride,
							a.EvmClassId,
							nftId,
						))
						if err != nil {
							return nil, doMintNFTActionFailed(db, a, err)
						}
						metadataStr = string(metadataBytes)

						events, err := m.QueryAllNFTEvents(m.MakeQueryNFTEventsRequest(cosmosNFT.ClassId, cosmosNFT.Id))
						if err != nil {
							return nil, doMintNFTActionFailed(db, a, err)
						}
						memo = event.MakeMemoFromEvent(events)
					}
					tos = append(tos, initialBatchMintOwnerAddress)
					memos = append(memos, memo)
					metadataList = append(metadataList, metadataStr)
				}
				_, _, err = c.MintNFTs(
					ctx,
					mylogger,
					evmClassAddress,
					totalSupplyBigInt,
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

func doMintNFTActionFailed(db *sql.DB, a *model.LikeNFTMigrationActionMintNFT, err error) error {
	a.Status = model.LikeNFTMigrationActionMintNFTStatusFailed
	failedReason := err.Error()
	a.FailedReason = &failedReason
	return errors.Join(err, appdb.UpdateLikeNFTMigrationActionMintNFT(db, a))
}
