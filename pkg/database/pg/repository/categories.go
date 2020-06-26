package repository

import (
	"demo/pkg/database/pg"
	"demo/pkg/services"
	. "demo/pkg/services/categories"

	p "github.com/go-pg/pg/v10"
)

type CategoryRepository struct {
	conn *pg.Connection
}

func NewCategoryRepository(conn *pg.Connection) *CategoryRepository {
	return &CategoryRepository{conn}
}

// implement
func (r *CategoryRepository) GetAll(categories *[]Category) error {
	return r.conn.DB.Model(categories).Select()
}

func (r *CategoryRepository) GetByID(category *Category, id int64) error {
	if err := r.conn.DB.Model(category).Where("id = ?", id).Select(); err != nil {
		if err == p.ErrNoRows {
			return services.NewErrNotFound("category")
		} else {
			return err
		}
	}

	return nil
}

func (r *CategoryRepository) GetBySlug(category *Category, slug string) error {
	if err := r.conn.DB.Model(category).Where("slug = ?", slug).Select(); err != nil {
		if err == p.ErrNoRows {
			return services.NewErrNotFound("category")
		} else {
			return err
		}
	}

	return nil
}

func (r *CategoryRepository) CheckSlugExists(slug string) (bool, error) {
	query := r.conn.DB.Model(&Category{}).Where("slug = ?", slug)
	return query.Exists()
}

func (r *CategoryRepository) Create(category *Category) error {
	return r.conn.DB.Insert(category)
}

func (r *CategoryRepository) Update(category *Category) error {
	return r.conn.DB.Update(category)
}

func (r *CategoryRepository) Delete(category *Category) error {
	return r.conn.DB.Delete(category)
}
