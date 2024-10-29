package pointers

import (
	asserts "hello-world"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(wallet, expected, t)
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(2))
		expected := Bitcoin(8)

		assertBalance(wallet, expected, t)
		asserts.AssertNoError(err, t)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(15)
		wallet := Wallet{balance: startingBalance}

		expectedError := wallet.Withdraw(Bitcoin(20))

		assertBalance(wallet, startingBalance, t)
		asserts.AssertErrorEquals(expectedError, ErrorInsufficientFunds, t)
	})
}

func assertBalance(wallet Wallet, expected Bitcoin, t testing.TB) {
	t.Helper()
	assertBitcoinEquals(wallet.Balance(), expected, t)
}

func assertBitcoinEquals(actual, expected Bitcoin, t testing.TB) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
