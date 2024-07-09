package test

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator/mysql"
	"github.com/Tynukua/factorial-online/internal/calculator/service"
	"log"
	"sync"
	"testing"
)

func TestCalculatorService(t *testing.T) {
	var a service.AsyncService
	var m mysql.MysqlCalculator
	ctx := context.TODO()
	wg := &sync.WaitGroup{}
	ctx = context.WithValue(ctx, "waitgroup", wg)
	f := func() {
		log.Println(m.Factorial(ctx, 5555))
		wg := ctx.Value("waitgroup").(*sync.WaitGroup)
		defer wg.Done()
	}
	g := func() {
		log.Println(m.Factorial(ctx, 6666))
		wg := ctx.Value("waitgroup").(*sync.WaitGroup)
		defer wg.Done()
	}
	err := a.Do(ctx, []func(){f, g})
	if err != nil {
		t.Error(err)
	}
}
