package database

import (
	"math/big"
)

type FactorialDatabase interface {
	InitDatabase() error
	SaveFactorial(number int, result *big.Int) error
	GetClosestFactorial(number int) (found int, result *big.Int, err error)
}
