package pkg

import models "simplebank.com/pkg"

 
 type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}



type TransferTxParams struct {
    FromAccountID int64 `json:"from_account_id"`
    ToAccountID   int64 `json:"to_account_id"`
    Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
    Transfer    models.Transfer `json:"transfer"`
    FromAccount  models.Account  `json:"from_account"`
    ToAccount    models.Account  `json:"to_account"`
    FromEntry    models.Entry    `json:"from_entry"`
    ToEntry      models.Entry    `json:"to_entry"`
}