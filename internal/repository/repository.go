package repository

import (
	"context"
	"database/sql"
  )

  type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
  }



//Repository is contains a DBTX which is either a *sql.DB or *sql.Tx
type Repository struct {
	db DBTX
}


//New creates a new repository for the API and returns it to the caller
func New(db DBTX) *Repository {
	return &Repository{db: db}
}

/*
WithTx creates a new repository for the API and returns it to the caller
* sql.Tx is a transaction that is used to execute multiple statements
* as a single unit of work
*/

func (r *Repository) WithTx(tx *sql.Tx) *Repository {
	return &Repository{
	  db: tx,
	}
  }


