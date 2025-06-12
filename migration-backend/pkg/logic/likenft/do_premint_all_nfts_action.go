package likenft

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmosnftidclassifier"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

var (
	ErrCosmosNFTIDsNotConsistent = errors.New("cosmos nftids not consistent")
)

// Will do premint all nfts if the new class action
// specified `ShouldPremintAllNFTs` and checked the nft ids are serial
func DoPremintAllNFTsActionIfNeeded(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	likenftCosmosClient *cosmos.LikeNFTCosmosClient,
	likecoinAPI *likecoin_api.LikecoinAPI,
	likeProtocolEvmClient *evm.LikeProtocol,
	bookNFTEvmClient *evm.BookNFT,
	cosmosNFTIDClassifier cosmosnftidclassifier.CosmosNFTIDClassifier,
	erc721ExternalURLBuilder erc721externalurl.ERC721ExternalURLBuilder,

	shouldPremintArbitraryNFTIDs bool,

	batchMintPerPage uint64,
	newClassAction *model.LikeNFTMigrationActionNewClass,
) error {
	mylogger := logger.WithGroup("DoPremintAllNFTsAction").
		With("ShouldPremintAllNFTs", newClassAction.ShouldPremintAllNFTs)

	if newClassAction.ShouldPremintAllNFTs {
		queryAllNFTsByClassIdResponse, err := likenftCosmosClient.
			QueryAllNFTsByClassId(cosmos.QueryAllNFTsByClassIdRequest{
				ClassId: newClassAction.CosmosClassId,
			})
		if err != nil {
			mylogger.Error("c.QueryAllNFTsByClassId", "err", err)
			return err
		}

		nfts := queryAllNFTsByClassIdResponse.NFTs
		if len(nfts) == 0 {
			mylogger.Info("no nfts found")
			return nil
		}

		classifyResult := cosmosNFTIDClassifier.Classify(nfts[0], nfts[1:]...)

		mylogger.Info(
			"classification completed",
			"len(serialNFTIDs)",
			len(classifyResult.SerialNFTIDs),
			"len(arbitraryNFTIds)",
			len(classifyResult.ArbitraryNFTIDs),
		)

		if allSerial, isAllSerial := classifyResult.AllSerial(); isAllSerial {
			// the nft ids are in serial.
			// Will batch mint by expected supply
			// The supply diff will be handled in DoBatchMintNFTsFromCosmosAction
			maxId := allSerial.Max()
			expectedSupply := maxId + 1

			newClassAction.NumberOfCosmosNFTsFound = &expectedSupply
			err = appdb.UpdateLikeNFTMigrationActionNewClass(db, newClassAction)
			if err != nil {
				mylogger.Error("appdb.UpdateLikeNFTMigrationActionNewClass", "err", err)
				return err
			}

			err := DoBatchMintNFTsFromCosmosAction(
				ctx,
				mylogger,
				db,
				likeProtocolEvmClient,
				bookNFTEvmClient,
				likecoinAPI,
				likenftCosmosClient,
				erc721ExternalURLBuilder,
				*newClassAction.EvmClassId,
				expectedSupply,
				batchMintPerPage,
				newClassAction.InitialBatchMintOwner,
			)

			if err != nil {
				return err
			}
		} else if allArbitrary, isAllArbitrary := classifyResult.AllArbitrary(); isAllArbitrary {
			if !shouldPremintArbitraryNFTIDs {
				mylogger.Info("found nft ids all arbitrary but shouldMintWithArbitraryNFTIDs is false. skip")
				return nil
			}
			// the nft ids are not in serial
			// Will mint arbitrarily and reassign a new evm nft id
			// Will skip minting if the cosmosNFTId already minted
			l := uint64(len(allArbitrary))
			newClassAction.NumberOfCosmosNFTsFound = &l
			err = appdb.UpdateLikeNFTMigrationActionNewClass(db, newClassAction)
			if err != nil {
				mylogger.Error("appdb.UpdateLikeNFTMigrationActionNewClass", "err", err)
				return err
			}

			for _, cosmosNFTId := range allArbitrary {
				mintNFTAction, err := GetOrCreateMintNFTAction(
					db,
					*newClassAction.EvmClassId,
					cosmosNFTId,
					newClassAction.InitialBatchMintOwner,
					newClassAction.InitialOwner,
				)
				if err != nil {
					return err
				}
				_, err = DoMintNFTAction(
					ctx,
					mylogger,
					db,
					likeProtocolEvmClient,
					bookNFTEvmClient,
					likenftCosmosClient,
					erc721ExternalURLBuilder,
					mintNFTAction,
				)
				if err != nil {
					return err
				}
			}
		} else {
			return ErrCosmosNFTIDsNotConsistent
		}
	} else {
		mylogger.Info("skipped")
	}

	return nil
}
