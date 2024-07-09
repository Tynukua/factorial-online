package mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online/internal/calculator"
	"github.com/Tynukua/factorial-online/internal/calculator/mathematics"
	"math/big"
)

type MysqlCalculator struct {
	calculator.Calculator
	db *sql.DB
}

func NewMysqlCalculator(dsn string) MysqlCalculator {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return MysqlCalculator{db: db}
}

func (m MysqlCalculator) Factorial(ctx context.Context, sn int) *big.Int {

	return mathematics.MathCalculator{}.Factorial(ctx, sn)
}
