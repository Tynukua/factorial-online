package main

import (
	"database/sql"
	"log"
	"math/big"
)

func InitDatabase(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS factorials (
		number INT NOT NULL,
		result TEXT NOT NULL,
		PRIMARY KEY (number)
	);`
	result, err := db.Exec(query)
	log.Print(result, err)
	return err
}

func SaveFactorialToDatabase(db *sql.DB, number int, result *big.Int) error {
	query := "INSERT INTO factorials (number, result) VALUES (?, ?)"
	qres, err := db.Exec(query, number, result.String())
	log.Print(qres, number, err)

	return err
}

func GetClosestFactorial(db *sql.DB, number int) (found int, result *big.Int, err error) {
	query := `SELECT number, result 
		FROM factorials 
		WHERE number < ?
		ORDER BY ABS(number - ?) 
		LIMIT 1;`
	err = db.QueryRow(query, number, number).Scan(&found, &result)
	if err != nil {
		result = big.NewInt(1)
		found = 1
	}
	log.Print(found, result.String(), err)
	return
}
