package calculator

import (
	"context"
	"math/big"
)

type Calculator interface {
	Factorial(ctx context.Context, sn int) *big.Int
}
type AsyncService interface {
	Do(ctx context.Context, fs []func()) error
}
