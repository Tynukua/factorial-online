package test

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator/mathematics"
	"github.com/Tynukua/factorial-online/internal/calculator/mysql"
	"github.com/Tynukua/factorial-online/internal/calculator/service"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCalculatorService(t *testing.T) {
	var a service.AsyncService
	m := mysql.NewMysqlCalculator("root:example@tcp(localhost:3306)/testdb", mathematics.MathCalculator{})
	ctx := context.TODO()
	f := func() {
		log.Println(m.Factorial(ctx, 5555))
	}
	g := func() {
		log.Println(m.Factorial(ctx, 6666))
	}
	fs := []func(){f, g}
	err := a.Do(ctx, fs)

	if err != nil {
		t.Error(err)
	}
}
