package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, got Bitcoin, want Bitcoin) {
		t.Helper()
		assert.Equal(t, want, got)
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(5))
		assert.NoError(t, err)

		assertBalance(t, wallet.Balance(), Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startedBalance := Bitcoin(20)
		wallet := Wallet{balance: startedBalance}

		err := wallet.Withdraw(Bitcoin(200))

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "could not withdraw 200 BTC insufficient funds 20 BTC")
		assertBalance(t, wallet.Balance(), startedBalance)
	})
}
