package database

import (
	"database/sql"
	"log"
	"math/big"
	"os"
)

type MySQLFactorialDatabase struct {
	FactorialDatabase

	db *sql.DB
}

func NewMySQLFactorialDatabase(dsn string) MySQLFactorialDatabase {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	query := `CREATE TABLE IF NOT EXISTS factorials (
		number INT NOT NULL,
		result TEXT NOT NULL,
		PRIMARY KEY (number)
	);`
	result, err := db.Exec(query)
	log.Print(result, err)
	return MySQLFactorialDatabase{
		db: db,
	}
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
	var result_s string
	err = o.db.QueryRow(query, number, number).Scan(&found, &result_s)
	result = big.NewInt(1)
	if err != nil {
		found = 1
	} else {
		result.SetString(result_s, 10)
	}
	log.Print(found, result.String(), err)
	return
}
