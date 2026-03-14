package pointers

import "fmt"

type Bitcoin float64

func (b Bitcoin) Stringer() string {
	return fmt.Sprintf("BTC %g", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) (err error) {
	if deposit <= 0 {
		err = fmt.Errorf("deposit value %g cannot be less than 0.", deposit)
		return
	}
	w.balance = w.balance + deposit
	return
}

func(w *Wallet) Withdraw(withdraw Bitcoin) (err error) {
	if withdraw <= 0 {
		err = fmt.Errorf("invalid withdraw : withdraw value %g cannot be less than 0.", withdraw)
		return
	}

	if withdraw > w.balance {
		err = fmt.Errorf("insufficient funds : withdraw value %g cannot be allowed.", withdraw)
		return
	}
	w.balance = w.balance - withdraw
	return
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
