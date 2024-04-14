package db

import (
	"context"
	"testing"
	"time"

	"github.com/kapeel-mopkar/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAcct Account, toAcct Account) Transfer {
	arg := CreateTransferParams{
		Amount:        util.RandomPositiveAmount(),
		FromAccountID: fromAcct.ID,
		ToAccountID:   toAcct.ID,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}
func TestCreateTransfer(t *testing.T) {
	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)

	createRandomTransfer(t, fromAcct, toAcct)
}

func TestGetTransfer(t *testing.T) {
	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, fromAcct, toAcct)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestListTransfersByFromAccount(t *testing.T) {

	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAcct, toAcct)
	}

	arg := ListTransfersByFromAccountParams{
		FromAccountID: fromAcct.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfersByFromAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, fromAcct.ID)
		require.Equal(t, transfer.ToAccountID, toAcct.ID)
	}

}

func TestListTransfersByToAccount(t *testing.T) {

	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAcct, toAcct)
	}

	arg := ListTransfersByToAccountParams{
		ToAccountID: toAcct.ID,
		Limit:       5,
		Offset:      0,
	}

	transfers, err := testQueries.ListTransfersByToAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, fromAcct.ID)
		require.Equal(t, transfer.ToAccountID, toAcct.ID)
	}

}

func TestListTransfersByFromAndToAccount(t *testing.T) {

	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAcct, toAcct)
	}

	arg := ListTransfersByFromAndToAccountParams{
		FromAccountID: fromAcct.ID,
		ToAccountID:   toAcct.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfersByFromAndToAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, fromAcct.ID)
		require.Equal(t, transfer.ToAccountID, toAcct.ID)
	}

}
