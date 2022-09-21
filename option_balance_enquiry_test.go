package main

import (
	"testing"

	"awesomeProject1/model"
)

func Test_doBalanceEnquiry(t *testing.T) {
	type args struct {
		accountInfo *model.Account
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				accountInfo: &model.Account{
					Balance: 3.1415926,
				},
			},
		},
		{
			args: args{
				accountInfo: &model.Account{
					Balance: 1000000000000000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := doBalanceEnquiry(tt.args.accountInfo); (err != nil) != tt.wantErr {
				t.Errorf("doBalanceEnquiry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_beautifyAmount(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{amount: 100.12},
			want: "100.12",
		},
		{
			args: args{amount: 100000000.12},
			want: "100,000,000.12",
		},
		{
			args: args{amount: 999900000000.12},
			want: "999,900,000,000.12",
		},
		{
			args: args{amount: 999991111111111.12},
			want: "999,991,111,111,111.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifyAmount(tt.args.amount); got != tt.want {
				t.Errorf("beautifyAmount1() = %v, want %v", got, tt.want)
			}
		})
	}
}
