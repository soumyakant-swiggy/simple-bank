package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T, randomAccount Account) Entry {

	arg := CreateEntryParams{
		AccountID: sql.NullInt64{Int64: randomAccount.ID, Valid: true},
		Amount:    10,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	randomAccount := createRandomAccount(t)
	CreateRandomEntry(t, randomAccount)
}

func TestGetEntry(t *testing.T) {
	randomAccount := createRandomAccount(t)
	randomEntry := CreateRandomEntry(t, randomAccount)

	entry, err := testQueries.GetEntry(context.Background(), randomEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, randomEntry.ID, entry.ID)
	require.Equal(t, randomEntry.AccountID, entry.AccountID)
	require.Equal(t, randomEntry.Amount, entry.Amount)
	require.WithinDuration(t, randomEntry.CreatedAt.Time, entry.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {

	randomAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, randomAccount)
	}

	arg := ListEntriesParams{
		AccountID: sql.NullInt64{Int64: randomAccount.ID, Valid: true},
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}

}
