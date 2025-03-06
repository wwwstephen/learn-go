package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// using a pointer as it's less expensive
// but fine to pass the duplicate.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	if w.balance < 0 {
		w.balance += amount
		return ErrInsufficientFunds

	}
	return nil
}
