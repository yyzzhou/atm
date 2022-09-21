package model

import (
	"fmt"
	"sync"

	"atm/calculator"
)

type (
	// Account accountInfo struct
	Account struct {
		Id string
		// 1-personal account 2-business account
		AccType    int64
		Balance    float64
		Pin        string
		CustomerId string
		mu         sync.Mutex
	}

	AccountType int64
)

const (
	// PSN personal account
	PSN AccountType = 1
	// BIZ business account
	BIZ AccountType = 2
)

// IsBusinessAccount return true if business account
func (a *Account) IsBusinessAccount() bool {
	return a.AccType == int64(BIZ)
}

// IsPersonAccount return true if person account
func (a *Account) IsPersonAccount() bool {
	return a.AccType == int64(PSN)
}

// Name get name of account type
func (a AccountType) Name() string {
	if a == PSN {
		return "PSN"
	}
	if a == BIZ {
		return "BIZ"
	}
	return ""
}

// GetBalance get balance
func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

// Deposit deposit amount
func (a *Account) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	cal := calculator.NewOperator(a.Balance)
	ret := cal.Add(amount)
	a.Balance = ret
}

// Withdraw do withdraw amount
func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	cal := calculator.NewOperator(a.Balance)
	ret := cal.Sub(amount)
	if ret < 0 {
		return fmt.Errorf("unenough balance")
	}
	a.Balance = ret
	return nil
}
