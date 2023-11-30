package pricerepository

import (
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/adapter/sqllite"
	"github.com/rafaelsouzaribeiro/client-serve-api-go/server/core/domain"
)

type Repository struct {
	db sqllite.DataBaseInterface
}

func CreateTable(db sqllite.DataBaseInterface) error {
	createTable := `CREATE TABLE IF NOT EXISTS price(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"code" TEXT,
		"code_in" TEXT,
		"name" TEXT,
		"high" TEXT,
		"low" TEXT,
		"var_bid" TEXT,
		"pct_change" TEXT,
		"bid" TEXT,
		"ask" TEXT,
		"timestamp" TEXT,
		"create_date" TEXT
	)`

	stmt, err := db.Prepare(createTable)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	return nil
}

func New(db sqllite.DataBaseInterface) domain.PriceRepository {
	err := CreateTable(db)

	if err != nil {
		panic(err)
	}

	return &Repository{
		db: db,
	}
}
