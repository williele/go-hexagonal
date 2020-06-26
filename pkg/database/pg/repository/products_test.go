package repository

import (
	"demo/pkg/database/pg"
	"demo/pkg/database/pg/migrate"
	. "demo/pkg/services/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {
	conn, _ := pg.NewTestConnection()

	migrate.Migrate(conn, []string{"init"})
	migrate.Migrate(conn, []string{"up"})

	repo := NewProductRepository(conn)

	products := &[]Product{{
		Title: "foo product",
		Slug:  "foo-product",
		Price: 32,
	}, {
		Title:       "bar pro",
		Slug:        "bar-pro",
		Price:       10,
		Description: "Bar Production",
	}}

	// mocking data
	assert.NoError(t, conn.DB.Insert(products))

	// read
	products = &[]Product{}
	assert.NoError(t, repo.GetAll(products), "get all without error")
	assert.Equal(t, len(*products), 2, "get all correctly")

	product := &Product{}
	assert.NoError(t, repo.GetByID(product, 1), "get by id without error")
	assert.Equal(t, product.Slug, "foo-product", "get by id correctly")

	product = &Product{}
	assert.NoError(t, repo.GetBySlug(product, "bar-pro"), "get by slug without error")
	assert.Equal(t, product.Description, "Bar Production")

	product = &Product{}
	exist, err := repo.CheckSlugExists("bar-pro")
	assert.Equal(t, exist, true, "check slug exists correctly")
	assert.NoError(t, err)

	// shutdown
	migrate.Migrate(conn, []string{"down"})
}
