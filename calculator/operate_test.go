package calculator

import "testing"

func TestOperator_Add(t *testing.T) {
	type fields struct {
		Amount float64
	}
	type args struct {
		x float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			fields: fields{Amount: 100.02},
			args:   args{x: 3.1334},
			want:   103.1534,
		},
		{
			fields: fields{Amount: 100.02},
			args:   args{x: -3.1334},
			want:   96.8866,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operator{
				Amount: tt.fields.Amount,
			}
			if got := o.Add(tt.args.x); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
