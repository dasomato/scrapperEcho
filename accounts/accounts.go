package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoreBalance = errors.New("Can't Change Withdraw")

//NewAccount creates accounts
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit add balance of accounts
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance get balance of account
func (a Account) Balance() int {
	return a.balance
}

//ChangeWithdraw minus balance
func (a *Account) ChangeWithdraw(balance int) error {
	if a.balance < balance {
		return errNoMoreBalance
	}
	a.balance -= balance
	return nil
}

//ChangeOwner is to change owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nhas amount ", a.Balance())
}
