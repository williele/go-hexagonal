package products

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
