package mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online/internal/calculator"
	"math/big"
)

type MysqlCalculator struct {
	calculator.Calculator
	db       *sql.DB
	fallback calculator.Calculator
}

func NewMysqlCalculator(dsn string, fallback calculator.Calculator) MysqlCalculator {
	db, _ := sql.Open("mysql", dsn)
	return MysqlCalculator{db: db, fallback: fallback}
}

func (m MysqlCalculator) Factorial(ctx context.Context, sn int) *big.Int {

	return m.fallback.Factorial(ctx, sn)
}
