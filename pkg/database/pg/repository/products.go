package repository

import (
	"demo/pkg/database/pg"
	. "demo/pkg/services/products"
)

type ProductRepository struct {
	conn *pg.Connection
}

func NewProductRepository(conn *pg.Connection) *ProductRepository {
	return &ProductRepository{conn}
}

// implement
func (r *ProductRepository) GetAll(products *[]Product) error {
	return r.conn.DB.Model(products).Select()
}

func (r *ProductRepository) GetByID(product *Product, id int64) error {
	return r.conn.DB.Model(product).Where("id = ?", id).Select()
}

func (r *ProductRepository) GetBySlug(product *Product, slug string) error {
	return r.conn.DB.Model(product).Where("slug = ?", slug).Select()
}

func (r *ProductRepository) CheckSlugExists(slug string) (bool, error) {
	query := r.conn.DB.Model(&Product{}).Where("slug = ?", slug)
	return query.Exists()
}

func (r *ProductRepository) Create(product *Product) error {
	return r.conn.DB.Insert(product)
}

func (r *ProductRepository) Update(product *Product) error {
	return r.conn.DB.Update(product)
}

func (r *ProductRepository) Delete(product *Product) error {
	return r.conn.DB.Delete(product)
}
