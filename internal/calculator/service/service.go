package service

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator"
	"sync"
)

type AsyncService struct {
	calculator.AsyncService
}

func (a AsyncService) Do(ctx context.Context, fs []func()) error {
	wg := ctx.Value("waitgroup").(*sync.WaitGroup)
	wg.Add(len(fs))
	for _, f := range fs {
		go f()
	}
	wg.Wait()
	return nil
}
