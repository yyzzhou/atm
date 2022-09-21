package model

type (
	// Company companyInfo struct
	Company struct {
		Id       string
		Name     string
		Accounts []*Account
	}
)
