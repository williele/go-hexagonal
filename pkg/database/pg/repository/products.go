package repository

import (
	"demo/pkg/database/pg"
	"demo/pkg/services"
	. "demo/pkg/services/products"

	p "github.com/go-pg/pg/v10"
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
	if err := r.conn.DB.Model(product).Where("id = ?", id).Select(); err != nil {
		if err == p.ErrNoRows {
			return services.NewErrNotFound("product")
		} else {
			return err
		}
	}

	return nil
}

func (r *ProductRepository) GetBySlug(product *Product, slug string) error {
	if err := r.conn.DB.Model(product).Where("slug = ?", slug).Select(); err != nil {
		if err == p.ErrNoRows {
			return services.NewErrNotFound("product")
		} else {
			return err
		}
	}

	return nil
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
