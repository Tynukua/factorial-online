package config

import (
	"log"
	"os"
)

// enum db type
type DBType string

const (
	MySQL  DBType = "mysql"
	Memory DBType = "memory"
)

type Config struct {
	Port   string
	DBType DBType
	DSN    string
}

func NewConfig() Config {
	var cfg Config
	cfg.Port = "8989"
	cfg.DBType = Memory
	if dsn := os.Getenv("MYSQL_DSN"); dsn != "" {
		cfg.DBType = MySQL
		cfg.DSN = dsn
	}
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}
	log.Println("Database type:", cfg.DBType)
	return cfg
}