package data

import "atm/model"

var (
	// PersonAccounts person account list
	PersonAccounts = []*model.Account{
		{
			Id:         "PSN-1",
			AccType:    int64(model.PSN),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "a1",
		},
		{
			Id:         "PSN-2",
			AccType:    int64(model.PSN),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "a2",
		},
		{
			Id:         "PSN-3",
			AccType:    int64(model.PSN),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "a3",
		},
	}

	// CompanyAccounts company account list
	CompanyAccounts = []*model.Account{
		{
			Id:         "BIZ-1",
			AccType:    int64(model.BIZ),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "b1",
		},
		{
			Id:         "BIZ-2",
			AccType:    int64(model.BIZ),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "b2",
		},
		{
			Id:         "BIZ-3",
			AccType:    int64(model.BIZ),
			Balance:    100.00,
			Pin:        "123",
			CustomerId: "b3",
		},
	}
)
