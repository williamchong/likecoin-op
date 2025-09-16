package database

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/account"
	"likecollective-indexer/ent/nftclass"
	"likecollective-indexer/ent/predicate"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/ent/staking"

	"github.com/holiman/uint256"
)

type StakingKey struct {
	AccountEVMAddress string
	BookNFTEVMAddress string
}

func NewStakingKey(
	accountEVMAddress string,
	bookNFTEVMAddress string,
) StakingKey {
	return StakingKey{
		AccountEVMAddress: accountEVMAddress,
		BookNFTEVMAddress: bookNFTEVMAddress,
	}
}

func (k *StakingKey) ToPredicate() predicate.Staking {
	return staking.And(
		staking.HasAccountWith(account.EvmAddressEqualFold(k.AccountEVMAddress)),
		staking.HasNftClassWith(nftclass.AddressEqualFold(k.BookNFTEVMAddress)),
	)
}

type StakingRepository interface {
	QueryStakings(
		ctx context.Context,
		filter QueryStakingsFilter,
		pagination StakingPagination,
	) (stakings []*ent.Staking, count int, nextKey int, err error)

	QueryStakingsByKeys(
		ctx context.Context,
		keys []StakingKey,
	) ([]*ent.Staking, error)

	GetOrCreateStaking(
		ctx context.Context,
		tx *ent.Tx,
		nftClassID int,
		accountID int,
	) (*ent.Staking, error)

	CreateOrUpdateStaking(
		ctx context.Context,
		tx *ent.Tx,
		bookNFTEvmAddress string,
		accountEvmAddress string,
		stakedAmount typeutil.Uint256,
		pendingRewardAmount typeutil.Uint256,
		claimedRewardAmount typeutil.Uint256,
	) (*ent.Staking, error)
}

type stakingRepository struct {
	dbService Service
}

func MakeStakingRepository(
	dbService Service,
) StakingRepository {
	return &stakingRepository{
		dbService: dbService,
	}
}

func (r *stakingRepository) QueryStakings(
	ctx context.Context,
	filter QueryStakingsFilter,
	pagination StakingPagination,
) (
	stakings []*ent.Staking,
	count int,
	nextKey int,
	err error,
) {
	q := r.dbService.Client().Staking.Query().
		WithAccount().
		WithNftClass()
	q = filter.HandleFilter(q)

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	q = pagination.HandlePagination(q)

	stakings, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	nextKey = 0
	if len(stakings) > 0 {
		nextKey = stakings[len(stakings)-1].ID
	}

	return stakings, len(stakings), nextKey, nil
}

func (r *stakingRepository) QueryStakingsByKeys(
	ctx context.Context,
	keys []StakingKey,
) ([]*ent.Staking, error) {
	keysPredicates := make([]predicate.Staking, 0)
	for _, key := range keys {
		keysPredicates = append(
			keysPredicates,
			key.ToPredicate(),
		)
	}
	return r.dbService.Client().Staking.Query().WithAccount().WithNftClass().Where(staking.Or(keysPredicates...)).All(ctx)
}

func (r *stakingRepository) GetOrCreateStaking(
	ctx context.Context,
	tx *ent.Tx,
	nftClassID int,
	accountID int,
) (*ent.Staking, error) {
	staking, err := tx.Staking.Query().
		Where(staking.NftClassID(nftClassID), staking.AccountID(accountID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return tx.Staking.Create().
				SetNftClassID(nftClassID).
				SetAccountID(accountID).
				SetPoolShare("0").
				SetStakedAmount((typeutil.Uint256)(uint256.NewInt(0))).
				SetPendingRewardAmount((typeutil.Uint256)(uint256.NewInt(0))).
				SetClaimedRewardAmount((typeutil.Uint256)(uint256.NewInt(0))).
				Save(ctx)
		}
		return nil, err
	}
	return staking, nil
}

func (r *stakingRepository) CreateOrUpdateStaking(
	ctx context.Context,
	tx *ent.Tx,
	bookNFTEvmAddress string,
	accountEvmAddress string,
	stakedAmount typeutil.Uint256,
	pendingRewardAmount typeutil.Uint256,
	claimedRewardAmount typeutil.Uint256,
) (*ent.Staking, error) {
	account, err := tx.Account.Query().Where(account.EvmAddressEqualFold(accountEvmAddress)).Only(ctx)
	if err != nil {
		return nil, err
	}
	nftClass, err := tx.NFTClass.Query().Where(nftclass.AddressEqualFold(bookNFTEvmAddress)).Only(ctx)
	if err != nil {
		return nil, err
	}
	s, err := tx.Staking.Query().WithAccount().WithNftClass().
		Where(
			staking.AccountIDEQ(account.ID),
			staking.NftClassIDEQ(nftClass.ID),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return tx.Staking.Create().
				SetAccountID(account.ID).
				SetNftClassID(nftClass.ID).
				SetStakedAmount(stakedAmount).
				SetPendingRewardAmount(pendingRewardAmount).
				SetClaimedRewardAmount(claimedRewardAmount).
				Save(ctx)
		}
		return nil, err
	}
	return tx.Staking.UpdateOne(s).
		SetStakedAmount(stakedAmount).
		SetPendingRewardAmount(pendingRewardAmount).
		SetClaimedRewardAmount(claimedRewardAmount).
		Save(ctx)
}
