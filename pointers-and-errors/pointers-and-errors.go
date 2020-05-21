package main

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	fmt.Printf("address of balance in Deposit is %T \n", &w.balance)
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	fmt.Printf("address of balance in Deposit is %#v \n", &w.balance)
	fmt.Printf("address of balance in Deposit is %% \n")
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}


var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}