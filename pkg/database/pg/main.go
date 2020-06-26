package pg

import (
	"github.com/go-pg/pg/v10"
)

type Connection struct {
	DB *pg.DB
}

func NewConnection() (*Connection, error) {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})

	return &Connection{DB: db}, nil
}

func NewTestConnection() (*Connection, error) {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "tests",
	})

	return &Connection{DB: db}, nil
}
