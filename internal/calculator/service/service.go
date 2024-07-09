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
	wg := sync.WaitGroup{}
	wg.Add(len(fs))
	for _, f := range fs {
		go func() { f(); wg.Done() }()
	}
	wg.Wait()
	return nil
}
