package config

import (
	"database/sql"
	"log/slog"
)

const Version = "1.0.0"

type Config struct {
	Port      int
	Env       string
	DB_Config struct {
		Dsn string
	}
}

type Application struct {
	Config Config
	Logger *slog.Logger
	DB     *sql.DB
}
