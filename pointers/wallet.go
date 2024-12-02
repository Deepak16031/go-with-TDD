package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

var errorInsufficientBalance = errors.New("insufficient funds for withdrawal")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = w.balance + amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance < b {
		return errorInsufficientBalance
	}
	w.balance -= b
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
