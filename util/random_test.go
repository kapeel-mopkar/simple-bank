package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.NotEmpty(t, owner)
	require.Equal(t, 6, len(owner))
}

func TestRandomBalance(t *testing.T) {
	balance := RandomBalance()
	require.NotEmpty(t, balance)
	require.GreaterOrEqual(t, balance, int64(1000))
	require.LessOrEqual(t, balance, int64(10000000))
}

func TestRandomAmount(t *testing.T) {
	amount := RandomAmount()
	require.NotEmpty(t, amount)
	require.GreaterOrEqual(t, amount, int64(-100000))
	require.LessOrEqual(t, amount, int64(100000))
}

func TestRandomPositiveAmount(t *testing.T) {
	amount := RandomPositiveAmount()
	require.NotEmpty(t, amount)
	require.GreaterOrEqual(t, amount, int64(1))
	require.LessOrEqual(t, amount, int64(100000))
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.NotEmpty(t, currency)
	require.Equal(t, 3, len(currency))
	require.Contains(t, []string{"INR", "USD", "EUR", "AED"}, currency)
}
