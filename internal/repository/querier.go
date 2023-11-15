package repository

import (
	"context"

	models "simplebank.com/pkg"
)



type Querier interface {
    CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error)
    CreateEntry(ctx context.Context, arg CreateEntryParams) (models.Entry, error)
    CreateTransfer(ctx context.Context, arg CreateTransferParams) (models.Transfer, error)
    DeleteAccount(ctx context.Context, id int64) error
    GetAccount(ctx context.Context, id int64) (models.Account, error)
    ListAccounts(ctx context.Context, arg ListAccountsParams) ([]models.Account, error)
    GetTransfer(ctx context.Context, id int64) (models.Transfer, error)
    UpdateAccount(ctx context.Context, arg UpdateAccountParams) (models.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error)
}


var _ Querier = (*Repository)(nil)