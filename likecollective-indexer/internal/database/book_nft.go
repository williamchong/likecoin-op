package database

import (
	"context"
	"errors"

	"likecollective-indexer/ent"
)

type BookNFTRepository interface {
	QueryBookNFTs(
		ctx context.Context,
		filter QueryBookNFTsFilter,
	) (
		bookNFTs []*ent.BookNFT,
		count int,
		nextKey int,
		err error,
	)
	QueryBookNFT(
		ctx context.Context,
		evmAddress string,
	) (*ent.BookNFT, error)
}

type bookNFTRepository struct {
	dbService Service
}

func MakeBookNFTRepository(dbService Service) BookNFTRepository {
	return &bookNFTRepository{dbService: dbService}
}

func (r *bookNFTRepository) QueryBookNFTs(
	ctx context.Context,
	filter QueryBookNFTsFilter,
) (
	bookNFTs []*ent.BookNFT,
	count int,
	nextKey int,
	err error,
) {
	bookNFTs, err = r.dbService.Client().BookNFT.Query()
	if err != nil {
		return nil, 0, 0, err
	}
	bookNFTs = filter.HandleFilter(bookNFTs)

	return bookNFTs, len(bookNFTs), 0, nil
}

func (r *bookNFTRepository) QueryBookNFT(
	ctx context.Context,
	evmAddress string,
) (*ent.BookNFT, error) {
	bookNFTs, err := r.dbService.Client().BookNFT.Query()
	if err != nil {
		return nil, err
	}

	for _, bookNFT := range bookNFTs {
		if bookNFT.EvmAddress == evmAddress {
			return bookNFT, nil
		}
	}

	return nil, errors.New("book nft not found")
}
