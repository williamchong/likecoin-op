package database

import (
	"context"
	"math/big"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/evm/model"
)

type NFTClassRepository interface {
	QueryAllNFTClasses(ctx context.Context) ([]*ent.NFTClass, error)
	QueryNFTClassByAddress(ctx context.Context, address string) (*ent.NFTClass, error)
	QueryNFTClassesByAddressesExact(
		ctx context.Context,
		addresses []string,
	) ([]*ent.NFTClass, error)
	InsertNFTClass(
		ctx context.Context,
		address string,
		name string,
		symbol string,
		ownerAddress *string,
		minterAddresses []string,
		totalSupply *big.Int,
		maxSupply typeutil.Uint64,
		metadata *model.ContractLevelMetadata,
		bannerImage string,
		featuredImage string,
		deployerAddress string,
		deployedBlockNumber typeutil.Uint64,
		latestEventBlockNumber typeutil.Uint64,
		mintedAt time.Time,
		owner *ent.Account,
	) error
	UpdateMetadata(ctx context.Context, address string, metadata *model.ContractLevelMetadata) error
	UpdateOwner(ctx context.Context, address string, newOwner *ent.Account) error
	UpdateTotalSupply(ctx context.Context, address string, newTotalSupply *big.Int) error
}

type nftClassRepository struct {
	dbService Service
}

func MakeNFTClassRepository(
	dbService Service,
) NFTClassRepository {
	return &nftClassRepository{
		dbService: dbService,
	}
}

func (r *nftClassRepository) QueryAllNFTClasses(ctx context.Context) ([]*ent.NFTClass, error) {
	return r.dbService.Client().NFTClass.Query().All(ctx)
}

func (r *nftClassRepository) QueryNFTClassByAddress(ctx context.Context, address string) (*ent.NFTClass, error) {
	return r.dbService.Client().NFTClass.Query().
		Where(nftclass.AddressEqualFold(address)).
		Only(ctx)
}

func (r *nftClassRepository) QueryNFTClassesByAddressesExact(
	ctx context.Context,
	addresses []string,
) ([]*ent.NFTClass, error) {
	res, err := r.dbService.Client().NFTClass.Query().
		Where(nftclass.AddressIn(addresses...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	for _, address := range addresses {
		if !slices.ContainsFunc(res, func(nftClass *ent.NFTClass) bool {
			return nftClass.Address == address
		}) {
			return nil, &ent.NotFoundError{}
		}
	}

	return res, nil
}

func (r *nftClassRepository) InsertNFTClass(
	ctx context.Context,
	address string,
	name string,
	symbol string,
	ownerAddress *string,
	minterAddresses []string,
	totalSupply *big.Int,
	maxSupply typeutil.Uint64,
	metadata *model.ContractLevelMetadata,
	bannerImage string,
	featuredImage string,
	deployerAddress string,
	deployedBlockNumber typeutil.Uint64,
	latestEventBlockNumber typeutil.Uint64,
	mintedAt time.Time,
	owner *ent.Account,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		builder := tx.NFTClass.Create().
			SetAddress(address).
			SetName(name).
			SetSymbol(symbol).
			SetNillableOwnerAddress(ownerAddress).
			SetMinterAddresses(minterAddresses).
			SetTotalSupply(totalSupply).
			SetMaxSupply(maxSupply).
			SetMetadata(metadata).
			SetBannerImage(bannerImage).
			SetFeaturedImage(featuredImage).
			SetDeployerAddress(deployerAddress).
			SetDeployedBlockNumber(deployedBlockNumber).
			SetLatestEventBlockNumber(latestEventBlockNumber).
			SetMintedAt(mintedAt).
			SetOwner(owner).
			SetUpdatedAt(time.Now())

		return builder.Exec(ctx)
	})
}

func (r *nftClassRepository) UpdateMetadata(
	ctx context.Context,
	address string,
	metadata *model.ContractLevelMetadata,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.NFTClass.Query().
			Where(nftclass.AddressEqualFold(address)).
			Only(ctx)
		if err != nil {
			return err
		}
		return r.dbService.Client().NFTClass.Update().
			SetMetadata(metadata).
			Where(nftclass.AddressEqualFold(address)).
			Exec(ctx)
	})
}

func (r *nftClassRepository) UpdateOwner(
	ctx context.Context,
	address string,
	newOwner *ent.Account,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.NFTClass.Query().
			Where(nftclass.AddressEqualFold(address)).
			Only(ctx)
		if err != nil {
			return err
		}
		return r.dbService.Client().NFTClass.Update().
			SetOwner(newOwner).
			Where(nftclass.AddressEqualFold(address)).
			Exec(ctx)
	})
}

func (r *nftClassRepository) UpdateTotalSupply(
	ctx context.Context,
	address string,
	newTotalSupply *big.Int,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		_, err := tx.NFTClass.Query().
			Where(nftclass.AddressEqualFold(address)).
			Only(ctx)
		if err != nil {
			return err
		}
		return r.dbService.Client().NFTClass.Update().
			SetTotalSupply(newTotalSupply).
			Where(nftclass.AddressEqualFold(address)).
			Exec(ctx)
	})
}
