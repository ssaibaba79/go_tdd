package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		deposit := Bitcoin(10.0)
		err := wallet.Deposit(deposit)

		if err != nil {
			t.Errorf("deposit failed with error, %v", err)
		}
		balance := wallet.Balance()

		if deposit != balance {
			t.Errorf("%#v got:%v, want:%v", wallet, balance, deposit)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(10.0)
		wallet := Wallet{startingBalance}
		withdraw := Bitcoin(5.0)
		err := wallet.Withdraw(Bitcoin(withdraw))
		if err != nil {
			t.Errorf("withdraw failed with error, %v", err)
			return
		}
		balance := wallet.Balance()

		if balance != (startingBalance - withdraw) {
			t.Errorf("%#v withdrawn :%v, balance:%v", wallet, withdraw, balance)
		}
	})

	t.Run("Withdraw With Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10.0)
		wallet := Wallet{startingBalance}
		withdraw := Bitcoin(15.0)
		err := wallet.Withdraw(Bitcoin(withdraw))
		if err == nil {
			t.Errorf("withdraw did not fail %v", err)
			return
		}
		balance := wallet.Balance()

		if balance != startingBalance {
			t.Errorf("%#v withdraw > balance was allowed :%v, balance:%v", wallet, startingBalance, balance)
		}
	})

}
