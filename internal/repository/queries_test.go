package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	pkg "simplebank.com/pkg/params"

	"context"
	"simplebank.com/internal/utils"
	models "simplebank.com/pkg"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)


const (
    dbDriver = "postgres"
    dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)


// test repository is a global variable of type *Repository that will be used in all tests
var testRepository *Repository



func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testRepository = New(conn)

	os.Exit(m.Run())
}



func createRandomAccount(t *testing.T) models.Account {
	arg := pkg.CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testRepository.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}


func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testRepository.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}



func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := pkg.UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testRepository.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}


func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testRepository.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testRepository.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}


func TestListAccounts(t *testing.T) {
	for i:=0; i<10; i++ {
		createRandomAccount(t)
	}

	arg := pkg.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testRepository.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

