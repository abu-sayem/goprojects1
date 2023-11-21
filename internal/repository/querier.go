package repository

import (
	"context"

	models "simplebank.com/pkg"
	pkg "simplebank.com/pkg/params"
)



type Querier interface {
    GetAccount(ctx context.Context, id int64) (models.Account, error)
    CreateAccount(ctx context.Context, arg pkg.CreateAccountParams) (models.Account, error)
    CreateEntry(ctx context.Context, arg pkg.CreateEntryParams) (models.Entry, error)
    CreateTransfer(ctx context.Context, arg pkg.CreateTransferParams) (models.Transfer, error)
    DeleteAccount(ctx context.Context, id int64) error
    ListAccounts(ctx context.Context, arg pkg.ListAccountsParams) ([]models.Account, error)
    GetTransfer(ctx context.Context, id int64) (models.Transfer, error)
    UpdateAccount(ctx context.Context, arg pkg.UpdateAccountParams) (models.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error)
}


var _ Querier = (*Repository)(nil)