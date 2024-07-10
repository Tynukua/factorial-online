package service

import (
	"context"
	"sync"
)

type AsyncService struct {
}

func New() *AsyncService {
	return &AsyncService{}
}

func (a *AsyncService) Do(ctx context.Context, fs []func()) error {
	wg := sync.WaitGroup{}
	wg.Add(len(fs))
	for _, f := range fs {
		go func() { f(); wg.Done() }()
	}
	wg.Wait()
	return nil
}
