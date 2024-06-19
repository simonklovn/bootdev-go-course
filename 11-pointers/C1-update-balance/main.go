package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
	switch t.transactionType {
		case transactionDeposit:
			c.balance += t.amount
			return nil
		case transactionWithdrawal:
			if c.balance < t.amount {
				return errors.New("insufficient funds")
			}
			c.balance -= t.amount
			return nil
		default:
			return errors.New("unknown transaction type")
	}
}
