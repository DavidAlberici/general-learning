package pointers

import (
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
		assertNoError(err, t)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(15)
		wallet := Wallet{balance: startingBalance}

		expectedError := wallet.Withdraw(Bitcoin(20))

		assertBalance(wallet, startingBalance, t)
		assertError(expectedError, ErrorInsufficientFunds, t)
	})
}

func assertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("did not want an error but got one: %q", err)
	}
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

func assertError(actualError, expectedError error, t testing.TB) {
	t.Helper()
	if actualError == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if actualError != expectedError {
		t.Errorf("expected %q but got %q", expectedError, actualError)
	}
}
