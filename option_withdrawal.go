package main

import (
	"fmt"

	"atm/model"
)

func doWithdrawal(accountInfo *model.Account) error {
	fmt.Println("please input the amount to withdraw:")
	var amount float64
	if _, err := fmt.Scanln(&amount); err != nil {
		return fmt.Errorf(fmt.Sprintf("invalid amount:%v", err))
	}
	if amount <= 0 {
		return fmt.Errorf(fmt.Sprintf("invalid withdraw amount"))
	}
	if err := accountInfo.Withdraw(amount); err != nil {
		return err
	}
	if err := doBalanceEnquiry(accountInfo); err != nil {
		return err
	}
	return nil
}
