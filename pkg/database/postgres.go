package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lokicodess/CatalogX/pkg/config"
)

func OpenDB(app *config.Application, dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		app.Logger.Error(err.Error())
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		app.Logger.Error(err.Error())
		return nil, err
	}

	return db, nil
}
