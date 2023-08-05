package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"io.defassio/simplebank/util"
	"testing"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        account1.Balance,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, account1.Balance)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer.ID)
	require.Equal(t, transfer2.Amount, transfer.Amount)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer2.CreatedAt, transfer.CreatedAt)

}
func TestUpdateTransfer(t *testing.T) {

	transfer := CreateRandomTransfer(t)
	money := util.RandomMoney()
	args := UpdateTransferParams{ID: transfer.ID, Amount: money}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, money, transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestGetTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := GetTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.GetTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, account := range transfers {
		require.NotEmpty(t, account)
	}
}
