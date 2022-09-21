package main

import "testing"

func Test_isPSNId(t *testing.T) {
	type args struct {
		accountId string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				accountId: "xxx",
			},
			want: false,
		},
		{
			args: args{
				accountId: "PSN-1",
			},
			want: true,
		},
		{
			args: args{
				accountId: "psn-1",
			},
			want: false,
		},
		{
			args: args{
				accountId: "Pxx-1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPSNId(tt.args.accountId); got != tt.want {
				t.Errorf("isPSNId() = %v, want %v", got, tt.want)
			}
		})
	}
}
