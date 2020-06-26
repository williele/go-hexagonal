package migrate

import "github.com/go-pg/migrations/v8"

func init() {
	const upQuery = `
	CREATE TABLE categories (
		id serial NOT NULL PRIMARY KEY,
		name text NOT NULL UNIQUE,
		slug text NOT NULL UNIQUE
	);

	CREATE TABLE products_categories (
		product_id int REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE,
		category_id int REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE,
		CONSTRAINT products_categories_pkey PRIMARY KEY (product_id, category_id)
	);
	`

	const downQuery = `
	DROP TABLE categories;
	DROP TABLE products_categories;
	`

	migrations.MustRegisterTx(func(tx migrations.DB) error {
		_, err := tx.Exec(upQuery)
		return err
	}, func(tx migrations.DB) error {
		_, err := tx.Exec(downQuery)
		return err
	})
}
