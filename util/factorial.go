package util

import (
	"math/big"
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
