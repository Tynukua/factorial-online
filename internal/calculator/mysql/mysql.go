package mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online/internal/calculator"
	"log"
	"math/big"
)

type MysqlCalculator struct {
	calculator.Calculator
	db       *sql.DB
	fallback calculator.Calculator
}

func NewMysqlCalculator(dsn string, fallback calculator.Calculator) MysqlCalculator {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	query := `CREATE TABLE IF NOT EXISTS factorials (
		number INT NOT NULL,
		result TEXT NOT NULL,
		PRIMARY KEY (number)
	);`
	result, err := db.Exec(query)
	log.Print(result, err)
	return MysqlCalculator{db: db, fallback: fallback}
}

func (m MysqlCalculator) Factorial(ctx context.Context, sn int) *big.Int {
	query := `SELECT number, result 
		FROM factorials 
		WHERE number = ?
		LIMIT 1;`
	var result_s string
	var found int
	err := m.db.QueryRow(query, sn).Scan(&found, &result_s)
	result := big.NewInt(1)
	if err != nil {
		result = m.fallback.Factorial(ctx, sn)
		query := "INSERT INTO factorials (number, result) VALUES (?, ?)"
		m.db.Exec(query, sn, result.String())
		return result
	} else {
		result.SetString(result_s, 10)
		return result
	}

}
