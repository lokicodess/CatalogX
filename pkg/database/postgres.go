package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lokicodess/CatalogX/pkg/config"
)

func OpenDB(app *config.Application, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		app.Logger.Error(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		app.Logger.Error(err.Error())
		return nil, err
	}

	return db, nil
}
