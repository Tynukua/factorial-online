package database

import (
	"math/big"
)

type MemoryFactorialDatabase struct {
	FactorialDatabase
	factorials map[int]*big.Int
}

func NewMemoryFactorialDatabase() MemoryFactorialDatabase {
	return MemoryFactorialDatabase{
		factorials: make(map[int]*big.Int),
	}
}

func (o MemoryFactorialDatabase) SaveFactorial(number int, result *big.Int) error {
	o.factorials[number] = result
	return nil
}

func (o MemoryFactorialDatabase) GetClosestFactorial(number int) (found int, result *big.Int, err error) {
	for i := number; i > 0; i-- {
		if o.factorials[i] != nil {
			found = i
			result = o.factorials[i]
			return
		}
	}
	return 1, big.NewInt(1), nil
}
