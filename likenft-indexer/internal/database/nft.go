package database

import (
	"context"
	"log/slog"
	"math/big"
	"strconv"
	"strings"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/ent/nft"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/evm/model"
)

type NFTRepository interface {
	QueryNFTsByEvmAddress(
		ctx context.Context,
		accountEvmAddress string,
		contractLevelMetadataEQ ContractLevelMetadataFilterEquatable,
		contractLevelMetadataNEQ ContractLevelMetadataFilterEquatable,
		pagination NFTPagination,
	) (nfts []*ent.NFT, count int, nextKey int, err error)
	QueryNFTsByBookNFTAndEvmAddress(
		ctx context.Context,
		bookNFTId string,
		accountEvmAddress string,
		pagination NFTPagination,
	) (nfts []*ent.NFT, count int, nextKey int, err error)
	GetOrCreate(
		ctx context.Context,
		ContractAddress string,
		TokenID *big.Int,
		TokenURI string,
		Image string,
		ImageData *string,
		ExternalURL *string,
		Description string,
		Name string,
		Attributes []model.ERC721MetadataAttribute,
		BackgroundColor *string,
		AnimationURL *string,
		YoutubeURL *string,
		OwnerAddress string,
		Owner *ent.Account,
		nftClass *ent.NFTClass,
	) (*ent.NFT, error)

	UpdateOwner(
		ctx context.Context,
		ContractAddress string,
		TokenID *big.Int,
		NewOwnerAddress string,
		NewOwner *ent.Account,
		UpdatedAt time.Time,
	) error

	BackfillUpdatedAtFromTransferEvents(
		ctx context.Context,
		logger *slog.Logger,
	) (updatedCount int, err error)
}

type nftRepository struct {
	dbService Service
}

func MakeNFTRepository(
	dbService Service,
) NFTRepository {
	return &nftRepository{
		dbService: dbService,
	}
}

func (r *nftRepository) QueryNFTsByEvmAddress(
	ctx context.Context,
	accountEvmAddress string,
	contractLevelMetadataEQ ContractLevelMetadataFilterEquatable,
	contractLevelMetadataNEQ ContractLevelMetadataFilterEquatable,
	pagination NFTPagination,
) (nfts []*ent.NFT, count int, nextKey int, err error) {
	bookNFTQ := r.dbService.Client().NFTClass.Query()

	bookNFTQ = contractLevelMetadataEQ.ApplyEQ(bookNFTQ)
	bookNFTQ = contractLevelMetadataNEQ.ApplyNEQ(bookNFTQ)

	q := bookNFTQ.QueryNfts().Where(
		nft.OwnerAddressEqualFold(accountEvmAddress),
	)

	q = pagination.HandlePagination(q)

	nfts, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	count = len(nfts)
	nextKey = 0
	if len(nfts) > 0 {
		nextKey = nfts[len(nfts)-1].ID
	}

	return nfts, count, nextKey, nil
}

func (r *nftRepository) QueryNFTsByBookNFTAndEvmAddress(
	ctx context.Context,
	bookNFTId string,
	accountEvmAddress string,
	pagination NFTPagination,
) (nfts []*ent.NFT, count int, nextKey int, err error) {
	q := r.dbService.Client().NFT.Query().
		Where(
			nft.HasClassWith(nftclass.AddressEqualFold(bookNFTId)),
			nft.OwnerAddressEqualFold(accountEvmAddress),
		)

	q = pagination.HandlePagination(q)

	nfts, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	count = len(nfts)
	nextKey = 0
	if len(nfts) > 0 {
		nextKey = nfts[len(nfts)-1].ID
	}

	return nfts, count, nextKey, nil
}

func (r *nftRepository) GetOrCreate(
	ctx context.Context,
	contractAddress string,
	tokenID *big.Int,
	tokenURI string,
	image string,
	imageData *string,
	externalURL *string,
	description string,
	name string,
	attributes []model.ERC721MetadataAttribute,
	backgroundColor *string,
	animationURL *string,
	youtubeURL *string,
	ownerAddress string,
	owner *ent.Account,
	nftClass *ent.NFTClass,
) (*ent.NFT, error) {
	resChan := make(chan *ent.NFT, 1)

	err := WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		n, err := tx.NFT.Query().
			Where(
				nft.ContractAddressEqualFold(contractAddress),
				nft.TokenIDEQ(typeutil.Uint64(tokenID.Uint64()))).
			Only(ctx)

		if err == nil {
			resChan <- n
			return nil
		}

		if !ent.IsNotFound(err) {
			return err
		}

		err = tx.NFT.Create().
			SetContractAddress(contractAddress).
			SetTokenID(typeutil.Uint64(tokenID.Uint64())).
			SetTokenURI(tokenURI).
			SetImage(image).
			SetNillableImageData(imageData).
			SetNillableExternalURL(externalURL).
			SetDescription(description).
			SetName(name).
			SetAttributes(attributes).
			SetNillableBackgroundColor(backgroundColor).
			SetNillableAnimationURL(animationURL).
			SetNillableYoutubeURL(youtubeURL).
			SetOwnerAddress(ownerAddress).
			SetOwner(owner).
			SetMintedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SetClass(nftClass).Exec(ctx)

		if err != nil {
			return err
		}

		n, err = tx.NFT.Query().
			Where(
				nft.ContractAddressEqualFold(contractAddress),
				nft.TokenIDEQ(typeutil.Uint64(tokenID.Uint64()))).
			Only(ctx)

		if err != nil {
			return err
		}

		resChan <- n
		return nil

	})

	if err != nil {
		return nil, err
	}
	return <-resChan, nil
}

func (r *nftRepository) UpdateOwner(
	ctx context.Context,
	contractAddress string,
	tokenID *big.Int,
	newOwnerAddress string,
	newOwner *ent.Account,
	updatedAt time.Time,
) error {
	err := WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.NFT.Query().
			Where(
				nft.ContractAddressEqualFold(contractAddress),
				nft.TokenIDEQ(typeutil.Uint64(tokenID.Uint64())),
			).
			Only(ctx)

		if err != nil {
			return err
		}

		return tx.NFT.Update().
			SetOwnerAddress(newOwnerAddress).
			SetOwner(newOwner).
			SetUpdatedAt(updatedAt).
			Where(
				nft.ContractAddressEqualFold(contractAddress),
				nft.TokenIDEQ(typeutil.Uint64(tokenID.Uint64())),
			).Exec(ctx)
	})

	if err != nil {
		return err
	}
	return nil
}

// Backfills nfts.updated_at from the block time of each token's latest
// Transfer/TransferWithMemo event. Historical rows only carried the indexing
// time (and transfers never bumped updated_at before this was recorded), so
// this makes updated_at mean "when the current owner acquired the token".
// Idempotent — safe to re-run.
func (r *nftRepository) BackfillUpdatedAtFromTransferEvents(
	ctx context.Context,
	logger *slog.Logger,
) (updatedCount int, err error) {
	// topic3 holds the tokenId of Transfer/TransferWithMemo as a decimal string
	// (see logconverter.ConvertLogToEvmEvent).
	var rows []struct {
		Address   string    `json:"address"`
		Topic3    string    `json:"topic3"`
		Timestamp time.Time `json:"max_timestamp"`
	}
	err = r.dbService.Client().EVMEvent.Query().
		Where(
			evmevent.NameIn("Transfer", "TransferWithMemo"),
			evmevent.Removed(false),
			// A failed event's ownership change never reached the nfts row,
			// so its timestamp must not be stamped either.
			evmevent.StatusEQ(evmevent.StatusProcessed),
			evmevent.Topic3NotNil(),
		).
		GroupBy(evmevent.FieldAddress, evmevent.FieldTopic3).
		Aggregate(ent.As(ent.Max(evmevent.FieldTimestamp), "max_timestamp")).
		Scan(ctx, &rows)
	if err != nil {
		return 0, err
	}

	// Update by primary key: matching on contract_address would compile
	// EqualFold to ILIKE, which cannot use the btree index and would turn
	// every update into a full table scan.
	nfts, err := r.dbService.Client().NFT.Query().
		Select(nft.FieldID, nft.FieldContractAddress, nft.FieldTokenID, nft.FieldUpdatedAt).
		All(ctx)
	if err != nil {
		return 0, err
	}
	nftByKey := make(map[string]*ent.NFT, len(nfts))
	for _, n := range nfts {
		nftByKey[makeNFTBackfillKey(n.ContractAddress, uint64(n.TokenID))] = n
	}

	for i, row := range rows {
		tokenID, err := strconv.ParseUint(row.Topic3, 10, 64)
		if err != nil {
			logger.Warn("skipping unparsable topic3", "address", row.Address, "topic3", row.Topic3)
			continue
		}
		n, exists := nftByKey[makeNFTBackfillKey(row.Address, tokenID)]
		if !exists || n.UpdatedAt.Equal(row.Timestamp) {
			continue
		}
		err = r.dbService.Client().NFT.UpdateOneID(n.ID).
			SetUpdatedAt(row.Timestamp).
			Exec(ctx)
		if err != nil {
			return updatedCount, err
		}
		updatedCount++
		if (i+1)%10000 == 0 {
			logger.Info("backfill progress", "processed", i+1, "total", len(rows), "updated", updatedCount)
		}
	}
	return updatedCount, nil
}

func makeNFTBackfillKey(contractAddress string, tokenID uint64) string {
	return strings.ToLower(contractAddress) + "/" + strconv.FormatUint(tokenID, 10)
}
