package migrate

import (
	"demo/pkg/database/pg"
	"log"

	"github.com/go-pg/migrations/v8"
)

func Migrate(conn *pg.Connection, args []string) {
	old, new, err := migrations.Run(conn.DB, args...)
	if err != nil {
		log.Fatal(err)
	}

	if old != new {
		log.Printf("Migrate from version %v to %v", old, new)
	} else {
		log.Printf("Current version %v", old)
	}
}
