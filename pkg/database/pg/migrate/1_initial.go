package migrate

import "github.com/go-pg/migrations/v8"

func init() {
	const upQuery = `
	CREATE TABLE products (
		id serial NOT NULL PRIMARY KEY,
		title text NOT NULL UNIQUE,
		slug text NOT NULL UNIQUE,
		price float NOT NULL,
		description text,
		published bool DEFAULT false,
		created_at timestamp with time zone DEFAULT current_timestamp,
		updated_at timestamp with time zone DEFAULT current_timestamp
	);
	`

	const downQuery = `
	DROP TABLE products;
	`

	migrations.MustRegisterTx(func(tx migrations.DB) error {
		_, err := tx.Exec(upQuery)
		return err
	}, func(tx migrations.DB) error {
		_, err := tx.Exec(downQuery)
		return err
	})
}
