package database

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/ent/likeprotocol"
	"likenft-indexer/ent/schema/typeutil"
)

type LikeProtocolRepository interface {
	GetLikeProtocol(
		ctx context.Context,
		contractAddress string,
	) (*ent.LikeProtocol, error)
	CreateOrUpdateLatestEventBlockHeight(
		ctx context.Context,
		contractAddress string,
		latestEventBlockHeight typeutil.Uint64,
	) error
}

type likeProtocolRepository struct {
	dbService Service
}

func MakeLikeProtocolRepository(
	dbService Service,
) LikeProtocolRepository {
	return &likeProtocolRepository{
		dbService: dbService,
	}
}

func (r *likeProtocolRepository) GetLikeProtocol(
	ctx context.Context,
	contractAddress string,
) (*ent.LikeProtocol, error) {
	likeProtocol, err := r.dbService.Client().LikeProtocol.Query().
		Where(likeprotocol.AddressEqualFold(contractAddress)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return likeProtocol, nil
}

func (r *likeProtocolRepository) CreateOrUpdateLatestEventBlockHeight(
	ctx context.Context,
	contractAddress string,
	latestEventBlockHeight typeutil.Uint64,
) error {
	return WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		affectedCount, err := tx.LikeProtocol.Update().
			Where(likeprotocol.AddressEqualFold(contractAddress)).
			SetLatestEventBlockNumber(latestEventBlockHeight).
			Save(ctx)
		if err != nil {
			return err
		}
		if affectedCount != 0 {
			return nil
		}
		return tx.LikeProtocol.Create().
			SetAddress(contractAddress).
			SetLatestEventBlockNumber(latestEventBlockHeight).
			Exec(ctx)
	})
}
