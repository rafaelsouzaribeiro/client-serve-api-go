package sqllite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DataBaseInterface interface {
	Close() error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
}

func Connect() *sql.DB {

	db, err := sql.Open("sqlite3", "price.db")

	if err != nil {
		panic(err)
	}

	return db

}
