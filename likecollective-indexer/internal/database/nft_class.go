package database

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/nftclass"
)

type NFTClassRepository interface {
	QueryNFTClasses(
		ctx context.Context,
		filter QueryNFTClassesFilter,
		pagination NFTClassPagination,
	) (
		bookNFTs []*ent.NFTClass,
		count int,
		nextKey int,
		err error,
	)
	QueryNFTClass(
		ctx context.Context,
		address string,
	) (*ent.NFTClass, error)
}

type nftClassRepository struct {
	dbService Service
}

func MakeNFTClassRepository(dbService Service) NFTClassRepository {
	return &nftClassRepository{dbService: dbService}
}

func (r *nftClassRepository) QueryNFTClasses(
	ctx context.Context,
	filter QueryNFTClassesFilter,
	pagination NFTClassPagination,
) (
	nftClasses []*ent.NFTClass,
	count int,
	nextKey int,
	err error,
) {
	q := r.dbService.Client().NFTClass.Query()
	q = filter.HandleFilter(q)

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	q = pagination.HandlePagination(q)

	nftClasses, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	nextKey = 0
	if len(nftClasses) > 0 {
		nextKey = nftClasses[len(nftClasses)-1].ID
	}

	return nftClasses, count, nextKey, nil
}

func (r *nftClassRepository) QueryNFTClass(
	ctx context.Context,
	address string,
) (*ent.NFTClass, error) {
	nftClass, err := r.dbService.Client().NFTClass.Query().Where(
		nftclass.AddressEqualFold(address),
	).Only(ctx)
	if err != nil {
		return nil, err
	}

	return nftClass, nil
}
