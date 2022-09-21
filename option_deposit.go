package main

import (
	"fmt"

	"atm/model"
)

func doDeposit(accountInfo *model.Account) error {
	fmt.Println("please input the amount to deposit:")
	var amount float64
	if _, err := fmt.Scanln(&amount); err != nil {
		return fmt.Errorf(fmt.Sprintf("invalid amount:%v", err))
	}
	if amount <= 0 {
		return fmt.Errorf(fmt.Sprintf("invalid deposit amount"))
	}
	accountInfo.Deposit(amount)
	if err := doBalanceEnquiry(accountInfo); err != nil {
		return err
	}
	return nil
}
