package pointers

import (
	"fmt"
)

type ErrInsufficientFunds struct {
	Amount  Bitcoin
	Balance Bitcoin
}

func (e ErrInsufficientFunds) Error() string {
	return fmt.Sprintf("could not withdraw %s insufficient funds %s", e.Amount, e.Balance)
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds{amount, w.balance}
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
