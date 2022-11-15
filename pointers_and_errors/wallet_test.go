package pointersanderrors

import "testing"

// TestWallet tests different functionalities of Bitcoin Wallet
func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("Wanted an error but didn't give one")
		}
	})
}

/* * Assertion utilities * */

// assertBalance asserts if the balance in the wallet is equal to
// the 'want' value
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// assertErrors asserts if there are error & error of the right
// type returned
func assertError(t *testing.T, err error, want error) {
	t.Helper()

	if err == nil {
		t.Error("wanted an error but didn't give one")
	}

	if err != want {
		t.Errorf("got %s want %s", err.Error(), want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Error("got an error but didn't want one")
	}
}
