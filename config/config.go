package config

import "os"

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
	if dsn := os.Getenv("MYSQL_DSN"); dsn == "" {
		cfg.DBType = Memory
	} else {
		cfg.DBType = MySQL
		cfg.DSN = dsn
	}
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}
	return cfg
}
