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

func updateBalance(cus *customer, trns transaction) error {
	switch trns.transactionType {
	case transactionDeposit:
		cus.balance += trns.amount
		return nil
	case transactionWithdrawal:
		if cus.balance < trns.amount {
			return errors.New("insufficient funds")
		}
		cus.balance -= trns.amount
		return nil
	default:
		return errors.New("unknown transaction type")
	}

}
