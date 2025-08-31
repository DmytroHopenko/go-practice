package main

import (
	"fmt"
)

type Account struct {
	id      int
	balance float64
}

type Bank interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited %.2f, new balance: %.2f\n", amount, a.balance)
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		fmt.Println("Insufficient funds")
		return
	}
	a.balance -= amount
	fmt.Printf("Withdrew %.2f, new balance: %.2f\n", amount, a.balance)
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func main() {
	account := &Account{
		id:      1,
		balance: 0,
	}

	var bank Bank = account

	bank.Deposit(100)
	bank.Withdraw(50)
	bank.Withdraw(100)
	fmt.Println("Final balance:", bank.GetBalance())
}
