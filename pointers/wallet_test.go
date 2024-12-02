package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		if actual != expected {
			t.Errorf("expected %s actual %s", expected, actual)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(100)
		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		_ = wallet.Withdraw(90)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{100}
		err := wallet.Withdraw(Bitcoin(110))
		assertError(t, err, errorInsufficientBalance)
		assertBalance(t, wallet, Bitcoin(100))
	})
}

func assertError(t *testing.T, actual error, expected error) {
	if !errors.Is(actual, expected) {
		t.Errorf("epected an error but didnt received")
	}
}
