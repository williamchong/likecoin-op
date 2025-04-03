package database

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"
	"likenft-indexer/internal/evm/model"
)

type NFTClassRepository interface {
	QueryAllNFTClasses(ctx context.Context) ([]*ent.NFTClass, error)
	QueryNFTClassByAddress(ctx context.Context, address string) (*ent.NFTClass, error)
	InsertNFTClass(ctx context.Context, nftClass *ent.NFTClass) error
	UpdateMetadata(ctx context.Context, address string, metadata *model.ContractLevelMetadata) error
	UpdateOwner(ctx context.Context, address string, newOwner *ent.Account) error
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

func (r *nftClassRepository) InsertNFTClass(ctx context.Context, nftClass *ent.NFTClass) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		builder := tx.NFTClass.Create().
			SetAddress(nftClass.Address).
			SetBannerImage(nftClass.BannerImage).
			SetDeployedBlockNumber(nftClass.DeployedBlockNumber).
			SetDeployerAddress(nftClass.DeployerAddress).
			SetFeaturedImage(nftClass.FeaturedImage).
			SetMetadata(nftClass.Metadata).
			SetMintedAt(nftClass.MintedAt).
			SetMinterAddresses(nftClass.MinterAddresses).
			SetName(nftClass.Name).
			SetNillableOwnerAddress(nftClass.OwnerAddress).
			SetOwner(nftClass.Edges.Owner).
			SetSymbol(nftClass.Symbol).
			SetTotalSupply(nftClass.TotalSupply).
			SetMaxSupply(nftClass.MaxSupply).
			SetUpdatedAt(nftClass.UpdatedAt)

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
