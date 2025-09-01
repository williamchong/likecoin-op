package database

import (
	"context"
	"errors"

	"likecollective-indexer/ent"
)

type AccountRepository interface {
	QueryAccounts(ctx context.Context) (accounts []*ent.Account, count int, nextKey int, err error)
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
) (
	accounts []*ent.Account,
	count int,
	nextKey int,
	err error,
) {
	accounts, err = r.dbService.Client().Account.Query()

	if err != nil {
		return nil, 0, 0, err
	}

	return accounts, len(accounts), 0, nil
}

func (r *accountRepository) QueryAccount(
	ctx context.Context,
	evmAddress string,
) (*ent.Account, error) {
	accounts, err := r.dbService.Client().Account.Query()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.EvmAddress == evmAddress {
			return account, nil
		}
	}

	return nil, errors.New("account not found")
}
