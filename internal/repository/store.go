package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "simplebank.com/pkg"
)

type Store struct {
	*Repository
	db *sql.DB
}


func NewStore(db *sql.DB) *Store {
    return &Store{
        db:      db,
        Repository: New(db),
    }
}


func (store *Store) execTx(ctx context.Context, fn func(*Repository) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}


type TransferTxResult struct {
    Transfer    models.Transfer `json:"transfer"`
    FromAccount  models.Account  `json:"from_account"`
    ToAccount    models.Account  `json:"to_account"`
    FromEntry    models.Entry    `json:"from_entry"`
    ToEntry      models.Entry    `json:"to_entry"`
}


type TransferTxParams struct {
    FromAccountID int64 `json:"from_account_id"`
    ToAccountID   int64 `json:"to_account_id"`
    Amount        int64 `json:"amount"`
}


var txKey = struct{}{}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

	var result TransferTxResult

	err := store.execTx(ctx, func(r *Repository) error {
		
		var err error

		txName := ctx.Value(txKey).(string)

		fmt.Println(txName, "get account 1")
		account1, err := r.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = r.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 1")
		account2, err := r.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = r.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: account2.Balance + arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entries 1")
		result.FromEntry, err = r.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entries 2")
		result.ToEntry, err = r.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create transfer")
		result.Transfer, err = r.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return TransferTxResult{}, err
	}

	return result, nil
}