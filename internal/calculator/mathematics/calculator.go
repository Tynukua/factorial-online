package mathematics

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator"
	"math/big"
)

type Calculator struct {
	calculator.Calculator
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Factorial(ctx context.Context, n int) *big.Int {
	result := big.NewInt(1)
	for n > 1 {
		result.Mul(result, big.NewInt(int64(n)))
		n -= 1
	}
	return result
}
