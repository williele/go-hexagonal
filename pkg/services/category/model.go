package categories

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryCreateInput struct {
	Name string `json:"title" validate:"required"`
}

type CategoryUpdateInput struct {
	Name string `json:"title"`
}
