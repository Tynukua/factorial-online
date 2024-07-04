package main

import (
	"database/sql"
	"log"
	"math/big"
	"os"
)

type FactorialDatabase interface {
	InitDatabase() error
	SaveFactorial(number int, result *big.Int) error
	GetClosestFactorial(number int) (found int, result *big.Int, err error)
}

type MySQLFactorialDatabase struct {
	FactorialDatabase

	db *sql.DB
}

func NewMySQLFactorialDatabase(dsn string) MySQLFactorialDatabase {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	return MySQLFactorialDatabase{
		db: db,
	}
}

func (o MySQLFactorialDatabase) InitDatabase() error {
	query := `CREATE TABLE IF NOT EXISTS factorials (
		number INT NOT NULL,
		result TEXT NOT NULL,
		PRIMARY KEY (number)
	);`
	result, err := o.db.Exec(query)
	log.Print(result, err)
	return err
}

func (o MySQLFactorialDatabase) SaveFactorial(number int, result *big.Int) error {
	query := "INSERT INTO factorials (number, result) VALUES (?, ?)"
	qres, err := o.db.Exec(query, number, result.String())
	log.Print(qres, number, err)

	return err
}

func (o MySQLFactorialDatabase) GetClosestFactorial(number int) (found int, result *big.Int, err error) {
	query := `SELECT number, result 
		FROM factorials 
		WHERE number < ?
		ORDER BY ABS(number - ?) 
		LIMIT 1;`
	err = o.db.QueryRow(query, number, number).Scan(&found, &result)
	if err != nil {
		result = big.NewInt(1)
		found = 1
	}
	log.Print(found, result.String(), err)
	return
}
