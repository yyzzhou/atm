package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to BTW~ \nplease input your AccountNumber:")
	var accountId string
	if _, err := fmt.Scanln(&accountId); err != nil {
		fmt.Println(fmt.Sprintf("scan err:%v", err))
		os.Exit(1)
	}

	accountInfo, err := getAccount(accountId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("welecome %s", accountInfo.CustomerId))

	for {
		fmt.Println("The following options:\n  1. Balance enquiry\n  2. Deposit\n  3. Withdrawal\n  4. Quit")
		var option int64
		fmt.Println("please input your option:")
		if _, err = fmt.Scanln(&option); err != nil {
			fmt.Println(fmt.Sprintf("scan err:%v", err))
			os.Exit(1)
		}
		switch option {
		case balanceEnquiry:
			err = doBalanceEnquiry(accountInfo)
		case deposit:
			err = doDeposit(accountInfo)
		case withdrawal:
			err = doWithdrawal(accountInfo)
		case quit:
			os.Exit(0)
		default:
			fmt.Println("unknown option")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
