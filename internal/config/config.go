package config

import "os"

type Config struct {
	HTTPAddress  string
	SqliteDBName string
}

const (
	httpAddr     = "HTTP_ADDRESS"
	sqliteDBName = "SQLITE_DB_NAME"
)

func LoadConfig() *Config {
	// TODO: load .env files
	return &Config{
		HTTPAddress:  os.Getenv(httpAddr),
		SqliteDBName: os.Getenv(sqliteDBName),
	}
}
