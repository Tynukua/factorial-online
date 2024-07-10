package calculator_mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online"
	"log"
	"math/big"
)

type Calculator struct {
	db       *sql.DB
	fallback factorial_online.Calculator
}

func New(db *sql.DB, fallback factorial_online.Calculator) *Calculator {
	query := `CREATE TABLE IF NOT EXISTS factorials (
		number INT NOT NULL,
		result TEXT NOT NULL,
		PRIMARY KEY (number)
	);`
	result, err := db.Exec(query)
	log.Print(result, err)
	return &Calculator{db: db, fallback: fallback}
}

func (m *Calculator) Factorial(ctx context.Context, sn int) *big.Int {
	query := `SELECT number, result 
		FROM factorials 
		WHERE number = ?
		LIMIT 1;`
	var storedValue string
	var foundNumber int
	err := m.db.QueryRow(query, sn).Scan(&foundNumber, &storedValue)
	result := big.NewInt(1)
	if err != nil {
		result = m.fallback.Factorial(ctx, sn)
		query := "INSERT INTO factorials (number, result) VALUES (?, ?)"
		_, err = m.db.Exec(query, sn, result.String())
		if err != nil {
			log.Println(err)
		}
		return result
	} else {
		result.SetString(storedValue, 10)
		return result
	}

}
