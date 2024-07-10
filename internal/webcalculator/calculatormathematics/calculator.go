package calculatormathematics

import (
	"context"
	"math/big"
)

type Calculator struct {
}

func New() *Calculator {
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
