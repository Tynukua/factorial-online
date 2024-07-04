package main

import (
	"math/big"
	"runtime"
	"sync"
)

func MulRange(a, b int) *big.Int {
	product := big.NewInt(1)
	for i := a; i <= b; i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}
	return product
}

func MulRangeParallel(a int, b int, numWorkers int) *big.Int {
	if a > b {
		return big.NewInt(1)
	}

	var wg sync.WaitGroup
	results := make(chan *big.Int, numWorkers)
	step := (b - a + 1) / numWorkers

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		start := a + i*step
		end := start + step - 1
		if i == numWorkers-1 {
			end = b
		}

		go func(start, end int) {
			defer wg.Done()
			results <- MulRange(start, end)
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	product := big.NewInt(1)
	for result := range results {
		product.Mul(product, result)
	}

	return product
}

func DoubleFactorial(o FactorialDatabase, a int, b int) (*big.Int, *big.Int) {
	var swapped bool
	if a > b {
		a, b = b, a
		swapped = true
	}
	var af, bf *big.Int
	var ac, bc int
	var acf, bcf *big.Int
	ac, acf, _ = o.GetClosestFactorial(a)
	bc, bcf, _ = o.GetClosestFactorial(b)
	af = big.NewInt(1)
	bf = big.NewInt(1)

	af.Mul(acf, MulRangeParallel(ac+1, a, runtime.NumCPU()))
	if a > bc {
		bc = a
		bcf = af
	}
	bf.Mul(bcf, MulRangeParallel(bc+1, b, runtime.NumCPU()))

	o.SaveFactorial(a, af)
	o.SaveFactorial(b, bf)
	if swapped {
		af, bf = bf, af
	}
	return af, bf
}
