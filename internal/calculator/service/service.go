package service

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator"
)

type AsyncService struct {
	calculator.AsyncService
}

func (a AsyncService) Do(ctx context.Context, fs []func()) error {

	for _, f := range fs {
		go f()
	}
	return nil
}
