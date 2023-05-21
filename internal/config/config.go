package config

import "os"

type Config struct {
	HTTPAddress string
}

const httpAddr = "HTTP_ADDRESS"

func LoadConfig() *Config {
	// TODO: load .env files
	return &Config{
		HTTPAddress: os.Getenv(httpAddr),
	}
}
