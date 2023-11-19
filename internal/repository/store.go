package repository

import (
	"context"
	"database/sql"
	"fmt"

	pkg "simplebank.com/pkg/params"
)

type SQLStore struct {
	*Repository
	db *sql.DB
}


type Store interface {
	Querier
	TransferTx(ctx context.Context, arg pkg.TransferTxParams) (pkg.TransferTxResult, error)
}


func NewStore(db *sql.DB) Store {
    return &SQLStore{
        db:      db,
        Repository: New(db),
    }
}


func (store *SQLStore) execTx(ctx context.Context, fn func(*Repository) error) error {
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






var txKey = struct{}{}

func (store *SQLStore) TransferTx(ctx context.Context, arg pkg.TransferTxParams) (pkg.TransferTxResult, error) {

	var result pkg.TransferTxResult

	err := store.execTx(ctx, func(r *Repository) error {
		
		var err error

		txName := ctx.Value(txKey).(string)

		fmt.Println(txName, "get account 1")
		account1, err := r.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = r.UpdateAccount(ctx, pkg.UpdateAccountParams{
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
		result.ToAccount, err = r.UpdateAccount(ctx, pkg.UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: account2.Balance + arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entries 1")
		result.FromEntry, err = r.CreateEntry(ctx, pkg.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entries 2")
		result.ToEntry, err = r.CreateEntry(ctx, pkg.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create transfer")
		result.Transfer, err = r.CreateTransfer(ctx, pkg.CreateTransferParams{
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
		return pkg.TransferTxResult{}, err
	}

	return result, nil
}