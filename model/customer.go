package model

type (
	// Customer customerInfo struct
	Customer struct {
		Id       string
		Name     string
		Accounts []*Account
	}
)
