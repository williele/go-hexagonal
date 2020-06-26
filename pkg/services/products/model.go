package products

import (
	categories "demo/pkg/services/categories"
	"time"
)

type Product struct {
	ID          int64                 `json:"id"`
	Title       string                `json:"title"`
	Slug        string                `json:"slug"`
	Price       float64               `json:"price"`
	Description string                `json:"description,omitempty"`
	Published   bool                  `json:"published"`
	Categories  []categories.Category `json:"categories" pg:"many2many:products_categories"`
	CreatedAt   time.Time             `json:"createdAt"`
	UpdatedAt   time.Time             `json:"updatedAt"`
}

type ProductsToCategories struct {
	ProductID  int64
	CategoryID int64
}

type ProductCreateInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"price"`
}

type ProductUpdateInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       *float64 `json:"price" validate:"omitempty,price"`
}
