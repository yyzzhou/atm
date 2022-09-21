package calculator

import "github.com/shopspring/decimal"

type (
	// Operator 计算器
	Operator struct {
		Amount float64
	}
)

// NewOperator new operator with amount
func NewOperator(amount float64) *Operator {
	return &Operator{
		Amount: amount,
	}
}

// Add return operator.Amount + x
func (o *Operator) Add(x float64) float64 {
	newX := decimal.NewFromFloat(x)
	newAmt := decimal.NewFromFloat(o.Amount)
	return newAmt.Add(newX).InexactFloat64()
}

// Sub return operator.Amount - x
func (o *Operator) Sub(x float64) float64 {
	newX := decimal.NewFromFloat(x)
	newAmt := decimal.NewFromFloat(o.Amount)
	return newAmt.Sub(newX).InexactFloat64()
}
