package webservices

import (
	"context"
	"database/sql"
	factorial_online "github.com/Tynukua/factorial-online"
	"github.com/Tynukua/factorial-online/internal/config"
	calculatormathematics "github.com/Tynukua/factorial-online/internal/webcalculator/calculator/mathematics"
	calculatormysql "github.com/Tynukua/factorial-online/internal/webcalculator/calculator/mysql"
	"github.com/Tynukua/factorial-online/internal/webcalculator/service"
	"log"
	"math/big"
)

type FactorialService struct {
	service.AsyncService
	calculator factorial_online.Calculator
}

func NewFactorialService(cfg config.Config) FactorialService {
	fallback := calculatormathematics.New()
	switch expression := cfg.DBType; expression {
	case config.MySQL:
		db, err := sql.Open(`mysql`, cfg.DSN)
		if err != nil {
			log.Fatal(err)
		}
		return FactorialService{calculator: calculatormysql.New(db, fallback)}
	default:
		return FactorialService{calculator: fallback}
	}
}

func (s FactorialService) DoubleFactorial(ctx context.Context, a int, b int) (*big.Int, *big.Int) {
	var aFactorial, bFactorial *big.Int

	aFunc := func() {
		aFactorial = s.calculator.Factorial(ctx, a)
	}
	bFunc := func() {
		bFactorial = s.calculator.Factorial(ctx, b)
	}
	s.Do(ctx, []func(){aFunc, bFunc})
	return aFactorial, bFactorial
}
