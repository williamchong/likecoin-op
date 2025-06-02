package database

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/ent/account"
	"likenft-indexer/ent/nft"
	"likenft-indexer/ent/nftclass"
)

type AccountRepository interface {
	GetAccountByEvmAddress(ctx context.Context, evmAddress string) (*ent.Account, error)
	GetOrCreateAccount(ctx context.Context, acct *ent.Account) (*ent.Account, error)

	GetTokenAccountsByBookNFT(
		ctx context.Context,
		bookNFTId string,
		pagination AccountPagination,
	) (accounts []*ent.Account, count int, nextKey int, err error)
}

type accountRepository struct {
	dbService Service
}

func MakeAccountRepository(
	dbService Service,
) AccountRepository {
	return &accountRepository{
		dbService: dbService,
	}
}

func (r *accountRepository) GetAccountByEvmAddress(ctx context.Context, evmAddress string) (*ent.Account, error) {
	return r.dbService.Client().Account.Query().
		Where(account.EvmAddressEqualFold(evmAddress)).
		Only(ctx)
}

func (r *accountRepository) GetOrCreateAccount(ctx context.Context, acct *ent.Account) (*ent.Account, error) {
	resChan := make(chan *ent.Account, 1)

	err := WithTx(ctx, r.dbService.Client(), func(tx *ent.Tx) error {
		dbAccount, err := r.dbService.Client().Account.Query().
			Where(account.EvmAddressEqualFold(acct.EvmAddress)).
			Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		} else {
			resChan <- dbAccount
			return nil
		}

		err = r.dbService.Client().Account.Create().
			SetNillableCosmosAddress(acct.CosmosAddress).
			SetEvmAddress(acct.EvmAddress).
			SetNillableLikeid(acct.Likeid).Exec(ctx)
		if err != nil {
			return err
		}

		dbAccount, err = r.dbService.Client().Account.Query().
			Where(account.EvmAddressEqualFold(acct.EvmAddress)).
			Only(ctx)
		if err != nil {
			return err
		}

		resChan <- dbAccount
		return nil
	})

	if err != nil {
		return nil, err
	}

	return <-resChan, nil
}

func (r *accountRepository) GetTokenAccountsByBookNFT(
	ctx context.Context,
	bookNFTId string,
	pagination AccountPagination,
) (accounts []*ent.Account, count int, nextKey int, err error) {
	q := r.dbService.Client().Account.Query().
		Where(
			account.HasNftsWith(
				nft.HasClassWith(nftclass.AddressEqualFold(bookNFTId)),
			),
		)

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	q = pagination.HandlePagination(q)

	accounts, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	nextKey = 0
	if len(accounts) > 0 {
		nextKey = accounts[len(accounts)-1].ID
	}

	return accounts, count, nextKey, nil
}
