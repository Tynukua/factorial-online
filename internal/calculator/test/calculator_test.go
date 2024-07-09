package test

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator/mathematics"
	"github.com/Tynukua/factorial-online/internal/calculator/mysql"
	"github.com/Tynukua/factorial-online/internal/calculator/service"
	"log"
	"sync"
	"testing"
)

func TestCalculatorService(t *testing.T) {
	var a service.AsyncService
	m := mysql.NewMysqlCalculator("shrug", mathematics.MathCalculator{})
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	f := func() {
		log.Println(m.Factorial(ctx, 5555))
		wg.Done()
	}
	g := func() {
		log.Println(m.Factorial(ctx, 6666))
		wg.Done()
	}
	fs := []func(){f, g}
	wg.Add(len(fs))
	err := a.Do(ctx, fs)
	wg.Wait()
	if err != nil {
		t.Error(err)
	}
}
