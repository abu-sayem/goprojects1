package repository

import (
	"context"

	models "simplebank.com/pkg"
	pkg "simplebank.com/pkg/params"
)

const (
	createAccount = `-- name: CreateAccount :one
	INSERT INTO accounts (
	  owner,
	  balance,
	  currency
	) VALUES (
	  $1, $2, $3
	) RETURNING id, owner, balance, currency, created_at
	`

	deleteAccount = `-- name: DeleteAccount :exec
	DELETE FROM accounts
	WHERE id = $1
	`

	getAccount = `-- name: GetAccount :one
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
	`

	listAccounts = `-- name: ListAccounts :many
	SELECT id, owner, balance, currency, created_at FROM accounts
	ORDER BY id
	LIMIT $1
	OFFSET $2
	`

	updateAccount = `-- name: UpdateAccount :one
	UPDATE accounts
	SET balance = $2
	WHERE id = $1
	RETURNING id, owner, balance, currency, created_at
	`

	createEntry = `-- name: CreateEntry :one
	INSERT INTO entries (
	  account_id,
	  amount
	) VALUES (
	  $1, $2
	) RETURNING id, account_id, amount, created_at
	`

	createTransfer = `-- name: CreateTransfer :one
	INSERT INTO transfers (
	  from_account_id,
	  to_account_id,
	  amount
	) VALUES (
	  $1, $2, $3
	) RETURNING id, from_account_id, to_account_id, amount, created_at
	`

	getTransfer = `-- name: GetTransfer :one
	SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
	WHERE id = $1 LIMIT 1
	`

	getAccountForUpdate = `-- name: GetAccountForUpdate :one
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
	FOR UPDATE
	`
)

func (q *Repository) CreateAccount(ctx context.Context, arg pkg.CreateAccountParams) (models.Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}



func (q *Repository) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.QueryContext(ctx, deleteAccount, id)
	return err
}

func (q *Repository) GetAccount(ctx context.Context, id int64) (models.Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}



func (q *Repository) ListAccounts(ctx context.Context, arg pkg.ListAccountsParams) ([]models.Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.Account{}
	for rows.Next() {
		var i models.Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}



func (q *Repository) UpdateAccount(ctx context.Context, arg pkg.UpdateAccountParams) (models.Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)
	var i models.Account
	err := row.Scan(
	  &i.ID,
	  &i.Owner,
	  &i.Balance,
	  &i.Currency,
	  &i.CreatedAt,
	)
	return i, err
  }





func (q *Repository) CreateEntry(ctx context.Context, arg pkg.CreateEntryParams) (models.Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i models.Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}


func (q *Repository) CreateTransfer(ctx context.Context, arg pkg.CreateTransferParams) (models.Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i models.Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

func (r *Repository) GetTransfer(ctx context.Context, id int64) (models.Transfer, error) {
	row := r.db.QueryRowContext(ctx, getTransfer, id)
	var i models.Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}


func (r *Repository) GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error) {
	row := r.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}


