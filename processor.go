package main

import (
	"fmt"
	"strings"

	"atm/data"
	"atm/model"

	"github.com/howeyc/gopass"
)

const (
	balanceEnquiry = 1
	deposit        = 2
	withdrawal     = 3
	quit           = 4
)

func isPSNId(accountId string) bool {
	return strings.HasPrefix(accountId, model.PSN.Name())
}

func isBIZId(accountId string) bool {
	return strings.HasPrefix(accountId, model.BIZ.Name())
}

func validPSNAccountPin(accountId, pin string) (*model.Account, error) {
	curAccount := new(model.Account)
	for _, acc := range data.PersonAccounts {
		if accountId != acc.Id {
			continue
		}
		curAccount.Id = acc.Id
		if pin != acc.Pin {
			return nil, fmt.Errorf("invalid pin")
		}
		curAccount = acc
		break
	}
	if curAccount.Id == "" {
		return nil, fmt.Errorf("unknown account")
	}
	return curAccount, nil
}

func validBIZAccountPin(accountId, pin string) (*model.Account, error) {
	curAccount := new(model.Account)
	for _, acc := range data.CompanyAccounts {
		if accountId != acc.Id {
			continue
		}
		curAccount.Id = acc.Id
		if pin != acc.Pin {
			return nil, fmt.Errorf("invalid pin")
		}
		curAccount = acc
		break
	}
	if curAccount.Id == "" {
		return nil, fmt.Errorf("unknown account")
	}
	return curAccount, nil
}

func getAccount(accountId string) (*model.Account, error) {
	var (
		pin         string
		accountInfo *model.Account
		err         error
	)

	if isPSNId(accountId) {
		fmt.Println("please input your PIN:")
		pin, err = GetPin()
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("scan err:%v", err))
		}
		accountInfo, err = validPSNAccountPin(accountId, pin)
		if err != nil {
			return nil, err
		}
		return accountInfo, nil
	}

	if isBIZId(accountId) {
		fmt.Println("please input your CompanyCode:")
		var companyCode string
		if _, err = fmt.Scanln(&companyCode); err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("scan err:%v", err))
		}

		fmt.Println("please input your EmployeeId:")
		var employeeId string
		if _, err = fmt.Scanln(&employeeId); err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("scan err:%v", err))
		}

		fmt.Println("please input your PIN:")
		pin, err = GetPin()
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("scan err:%v", err))
		}

		accountInfo, err = validBIZAccountPin(accountId, pin)
		if err != nil {
			return nil, err
		}
		return accountInfo, nil
	}

	return nil, fmt.Errorf("unknown account")
}

func GetPin() (string, error) {
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		return "", err
	}
	return string(pass), nil
}
