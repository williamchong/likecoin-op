package database

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/account"
)

type AccountRepository interface {
	QueryAccounts(ctx context.Context, pagination AccountPagination, filter QueryAccountsFilter) (accounts []*ent.Account, count int, nextKey int, err error)
	QueryAccount(ctx context.Context, evmAddress string) (*ent.Account, error)
}

type accountRepository struct {
	dbService Service
}

func MakeAccountRepository(dbService Service) AccountRepository {
	return &accountRepository{dbService: dbService}
}

func (r *accountRepository) QueryAccounts(
	ctx context.Context,
	pagination AccountPagination,
	filter QueryAccountsFilter,
) (
	accounts []*ent.Account,
	count int,
	nextKey int,
	err error,
) {
	q := r.dbService.Client().Account.Query()
	q = filter.HandleFilter(q)

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

	return accounts, len(accounts), 0, nil
}

func (r *accountRepository) QueryAccount(
	ctx context.Context,
	evmAddress string,
) (*ent.Account, error) {
	account, err := r.dbService.Client().Account.Query().Where(
		account.EvmAddressEqualFold(evmAddress),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return account, nil
}
