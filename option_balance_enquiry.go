package main

import (
	"fmt"
	"strconv"
	"strings"

	"atm/model"
)

func doBalanceEnquiry(accountInfo *model.Account) error {
	amountStr := beautifyAmount(accountInfo.GetBalance())
	fmt.Println(fmt.Sprintf("your balance is: %s", amountStr))
	return nil
}

// 1111111.00 -> 1,111,111.00
func beautifyAmount(amount float64) string {
	amountStr := fmt.Sprintf("%.f", amount)
	original := fmt.Sprintf("%.2f", amount)
	i, _ := strconv.ParseInt(amountStr, 10, 64)
	d := i
	for {
		if d < 1000 {
			break
		}
		d = i / 1000
		tmpPrefix := fmt.Sprintf("%d", d)
		tmpSuffix := strings.TrimPrefix(original, tmpPrefix)
		tmp := tmpPrefix + "," + tmpSuffix
		original = tmp
		i = d
	}
	return original
}
