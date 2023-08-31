package common

import "os"

type Config struct {
	Dsn       string
	SecretKey string
}

func NewConfig() *Config {
	return &Config{
		Dsn:       os.Getenv("MYSQL_CONN_STRING"),
		SecretKey: os.Getenv("SYSTEM_SECRET"),
	}
}
