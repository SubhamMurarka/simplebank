package db

import (
	"context"
	"testing"

	"github.com/SubhamMurarka/simplebank/util"
	"github.com/stretchr/testify/require"
)

// Testing CreateEntry function by updating an account_id = 10 balance
func TestCreateEntry(t *testing.T) {
	accountid := int64(10)
	account, err := testQueries.GetAccount(context.Background(), accountid)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	arg := CreateEntryParams{
		AccountID: int64(10),
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, account.ID)
}

// func TestGetEntry(t *testing.T) {
// 	accountid := int64(10)
// 	entry, err := testQueries.GetEntry(context.Background(), accountid)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, entry)
// 	require.Equal(t, entry.AccountID.Int64, accountid)
// }
